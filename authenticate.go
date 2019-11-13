package dog

import (
	"github.com/godo"
	"golang.org/x/oauth2"
)

type Credentials struct {
	AccesToken string
}

func (c *Credentials) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: c.AccesToken,
	}
	return token, nil
}

func Authenticate(token *oauth2.Token) (*godo.Client) {

	tokenSource := &Credentials{
		AccesToken: token.AccessToken,
	}

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := godo.NewClient(oauthClient)

	return client
}
