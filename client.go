package vault

import (
	"context"
	"github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
	"github.com/pkg/errors"
)

type Client struct {
	C *api.Client
}

func New(cfg *Config) (client *Client, err error) {
	var (
		appRoleAuth *auth.AppRoleAuth
		secret      *api.Secret
	)
	client = &Client{}
	client.C, err = api.NewClient(&api.Config{
		Address: cfg.Url,
	})
	appRoleAuth, err = auth.NewAppRoleAuth(
		cfg.Role,
		&auth.SecretID{FromString: cfg.Secret},
	)
	if err != nil {
		return nil, errors.Errorf("unable to initialize AppRole auth method: %+v", err)
	}

	secret, err = client.C.Auth().Login(context.Background(), appRoleAuth)

	if err != nil {
		return nil, errors.Errorf("unable to login to AppRole auth method: %+v", err)
	}
	if secret == nil {
		return nil, errors.Errorf("no auth info was returned after login: %+v", err)
	}
	return
}
