package awssecretsmanager

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/viper"
)

type AwsSecretManagerLoader struct {
	secretsName    string
	ctx            context.Context
	secretsManager *secretsmanager.Client
}

func NewAwsSecretManagerLoader(secretsName string, opts ...AwsSecretManagerLoaderOption) (*AwsSecretManagerLoader, error) {
	loader := &AwsSecretManagerLoader{secretsName: secretsName}
	for _, applyOption := range opts {
		if err := applyOption(loader); err != nil {
			return nil, err
		}
	}
	if loader.secretsManager == nil {
		secretsManager, err := DefaultSecretsManager()
		if err != nil {
			return nil, err
		}
		loader.secretsManager = secretsManager
	}
	return loader, nil
}

func (r *AwsSecretManagerLoader) Load(app *viper.Viper) error {
	if r.ctx == nil {
		r.ctx = context.Background()
	}
	if r.secretsName == "" {
		return fmt.Errorf("secrets name cannot be empty")
	}
	secrets, err := r.GetSecrets(r.ctx, r.secretsName)
	if err != nil {
		return err
	}
	app.AutomaticEnv()
	app.SetConfigType("json")
	if err := app.ReadConfig(strings.NewReader(secrets)); err != nil {
		return err
	}
	return nil
}

func (r *AwsSecretManagerLoader) GetSecrets(ctx context.Context, secretsName string) (string, error) {
	if r.secretsManager == nil {
		return "", fmt.Errorf("secrets manager client not connected")
	}
	input := &secretsmanager.GetSecretValueInput{SecretId: aws.String(secretsName)}
	result, err := r.secretsManager.GetSecretValue(ctx, input)
	if err != nil {
		return "", err
	}
	return *result.SecretString, nil
}
