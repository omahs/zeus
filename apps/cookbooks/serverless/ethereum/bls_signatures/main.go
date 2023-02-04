package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
)

const (
	SessionToken    = "AWS_SESSION_TOKEN"
	SecretsHeader   = "X-Aws-Parameters-Secrets-Token"
	SecretsPortHTTP = 2773
)

var (
	secretCache, _ = secretcache.New()
)

type SecretsRequest struct {
	SecretName        string                          `json:"secretName"`
	SignatureRequests EthereumBLSKeySignatureRequests `json:"signatureRequests"`
}

type EthereumBLSKeySignatureRequests struct {
	Map map[string]EthereumBLSKeySignatureRequest `json:"map"`
}

type EthereumBLSKeySignatureRequest struct {
	Message string `json:"message"`
}

func HandleEthSignRequestBLS(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ApiResponse := events.APIGatewayProxyResponse{}
	m := make(map[string]any)

	sr := SecretsRequest{}
	err := json.Unmarshal([]byte(event.Body), &m)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}

	b, err := json.Marshal(m)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}
	err = json.Unmarshal(b, &sr)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}
	headerValue := os.Getenv(SessionToken)
	r := resty.New()
	url := fmt.Sprintf("http://localhost:%d/secretsmanager/get?secretId=%s", SecretsPortHTTP, sr.SecretName)
	resp, err := r.R().
		SetHeader(SecretsHeader, headerValue).
		Get(url)
	svo := &secretsmanager.GetSecretValueOutput{}
	err = json.Unmarshal(resp.Body(), &svo)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}

	ss := *svo.SecretString
	m = make(map[string]any)
	err = json.Unmarshal([]byte(ss), &m)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}

	signedResponses := aegis_inmemdbs.EthereumBLSKeySignatureResponses{Map: make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureResponse)}
	for pubkey, sk := range m {
		v, ok := sr.SignatureRequests.Map[pubkey]
		if ok {
			acc := bls_signer.NewEthSignerBLSFromExistingKey(sk.(string))
			sig := acc.Sign([]byte(v.Message)).Marshal()
			tmp := signedResponses.Map[pubkey]
			tmp.Signature = bls_signer.ConvertBytesToString(sig)
			signedResponses.Map[pubkey] = tmp
		}
	}
	b, err = json.Marshal(signedResponses)
	if err != nil {
		log.Ctx(ctx).Err(err)
		ApiResponse = events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}
		return ApiResponse, err
	}

	ApiResponse = events.APIGatewayProxyResponse{Body: string(b), StatusCode: 200}
	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleEthSignRequestBLS)
}
