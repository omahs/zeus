package aws_secrets

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/rs/zerolog/log"
)

func (s *SecretsManagerAuthAWS) GetSecret(ctx context.Context, si SecretInfo) (string, error) {
	m, err := s.GetSecretsJSON(ctx, si)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return "", err
	}
	secretValue, ok := m[si.Key]
	if !ok {
		log.Ctx(ctx).Warn().Interface("key", si.Key).Msg("no value found for secret key")
		return "", err
	}
	return secretValue.(string), nil
}

func (s *SecretsManagerAuthAWS) GetSecretsJSON(ctx context.Context, si SecretInfo) (map[string]any, error) {
	m := make(map[string]any)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(si.Name),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := s.GetSecretValue(ctx, input)
	if err != nil {
		// For a list of exceptions thrown, see
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		log.Ctx(ctx).Err(err)
		return m, err
	}

	// Decrypts secret using the associated KMS key.
	var secretString = *result.SecretString
	err = json.Unmarshal([]byte(secretString), &m)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return m, err
	}

	return m, nil
}
