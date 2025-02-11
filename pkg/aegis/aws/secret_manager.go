package aws_secrets

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/rs/zerolog/log"
)

type SecretsManagerAuthAWS struct {
	*secretsmanager.Client
}

type AuthAWS struct {
	Region    string
	AccessKey string
	SecretKey string
}

type SecretInfo struct {
	Region string `json:"region,omitempty"`
	Name   string `json:"name"`
	Key    string `json:"key,omitempty"`
}

func InitSecretsManager(ctx context.Context, auth AuthAWS) (SecretsManagerAuthAWS, error) {
	creds := credentials.NewStaticCredentialsProvider(auth.AccessKey, auth.SecretKey, "")
	cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(creds))
	if err != nil {
		log.Ctx(ctx).Err(err)
		return SecretsManagerAuthAWS{}, err
	}
	cfg.Region = auth.Region
	secretsManagerClient := secretsmanager.NewFromConfig(cfg)
	return SecretsManagerAuthAWS{secretsManagerClient}, err
}
