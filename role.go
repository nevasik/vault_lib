package vault

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
)

func NewRoleFromRoot(roleName, rootToken, vaultUrl string, policiesNames []string) (c Credentials, err error) {
	var ok bool
	client, err := api.NewClient(&api.Config{
		Address: vaultUrl,
	})
	if err != nil {
		err = errors.Errorf("Unable to create client: %v", err)
		return
	}
	client.SetToken(rootToken)
	roleConfig := map[string]interface{}{
		"policies": policiesNames,
	}

	_, err = client.Logical().Write("auth/approle/role/"+roleName, roleConfig)
	if err != nil {
		err = errors.Errorf("Unable to create new role: %v", err)
		return
	}
	secret, err := client.Logical().Read(fmt.Sprintf("auth/approle/role/%s/role-id", roleName))
	if err != nil {
		err = errors.Errorf("Unable to read role-id: %v", err)
		return
	}

	if secret == nil || secret.Data == nil {
		err = errors.Errorf("Empty response")
		return
	}

	if c.Role, ok = secret.Data["role_id"].(string); !ok {
		err = errors.Errorf("cannot convert roleID")
		return
	}
	secretIdResponse, err := client.Logical().Write("auth/approle/role/"+roleName+"/secret-id", nil)
	if err != nil || secretIdResponse == nil || secretIdResponse.Data == nil {
		err = errors.Errorf("Unable to create secret-id: %v", err)
		return
	}

	if secretIdResponse == nil || secretIdResponse.Data == nil {
		err = errors.Errorf("Empty response")
		return
	}

	if c.Secret, ok = secretIdResponse.Data["secret_id"].(string); !ok {
		err = errors.Errorf("cannot convert secretID")
		return
	}
	return
}
