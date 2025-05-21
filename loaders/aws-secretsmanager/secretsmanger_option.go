package secretsmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type AwsSecretManagerLoaderOption func(loader *AwsSecretManagerLoader) error

func WithSecretsName(secretsName string) AwsSecretManagerLoaderOption {
	return func(loader *AwsSecretManagerLoader) error {
		loader.secretsName = secretsName
		return nil
	}
}

func WithSecretsManagerClient(manager *secretsmanager.Client) AwsSecretManagerLoaderOption {
	return func(loader *AwsSecretManagerLoader) error {
		loader.secretsManager = manager
		return nil
	}
}

func WithDefaultSecretsManager() AwsSecretManagerLoaderOption {
	return func(loader *AwsSecretManagerLoader) error {
		secretsmanager, err := DefaultSecretsManager()
		if err != nil {
			return err
		}
		loader.secretsManager = secretsmanager
		return nil
	}
}

func DefaultSecretsManager() (*secretsmanager.Client, error) {
	config, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}
	return secretsmanager.NewFromConfig(config), nil
}
