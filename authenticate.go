package dog

import (
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

type Authenticator interface {
	Authenticate(pat string) *godo.Client
}

type Credentials struct {
	AccesToken string
}

func (c *Credentials) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: c.AccesToken,
	}
	return token, nil
}

func Authenticate(pat string) *godo.Client {

	tokenSource := &Credentials{
		AccesToken: pat,
	}

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := godo.NewClient(oauthClient)

	return client
}
