package main

import (
	"io/ioutil"

	"github.com/pulumi/pulumi-vault/sdk/v4/go/vault"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		vaultProvider, err := vault.NewProvider(ctx, "vault", &vault.ProviderArgs{
			Address: pulumi.String("https://myvaultaddr"),
            Token: pulumi.ToSecret(pulumi.String("abc")).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}

		// configure test kv backend
		_, err = vault.NewMount(ctx, "operator_test_kv", &vault.MountArgs{
			Description: pulumi.String("This is a test of pulumi operator"),
			Path:        pulumi.String("operator/kv"),
			Type:        pulumi.String("kv"),
		}, pulumi.Provider(vaultProvider))

		return err
	})
}

