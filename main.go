package main

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"k8s.io/klog/v2"
)

const vaultURL = "https://huan-test.vault.azure.net"
const clientID = "bc3c8e5d-266d-4277-ae7f-28b0b06f5xxx"
const clientSecret = "O3ah~.i227KdjOJz~.24gqY6jOo1.xxxxx"

func main() {
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		klog.Error(err)
	}

	baseClient := keyvault.New()
	baseClient.Authorizer = authorizer

	// Use client ID as secret name and client secret as value.
	secret := clientSecret
	_, err = baseClient.SetSecret(
		context.Background(),
		vaultURL,
		clientID,
		keyvault.SecretSetParameters{Value: &secret},
	)
	if err != nil {
		klog.Error(err)
	}

	resp, err := baseClient.GetSecret(
		context.Background(), vaultURL, clientID, "",
	)
	if err != nil {
		klog.Error(err)
	}

	klog.Infof("value: %s", *resp.Value)
}
