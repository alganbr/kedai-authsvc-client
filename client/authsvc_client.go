package client

import (
	"github.com/mercadolibre/golang-restclient/rest"
)

type IAuthSvcClient interface {
	Auth() IAuthClient
}

type AuthSvcClient struct {
	HttpClient *rest.RequestBuilder
}

func (c AuthSvcClient) Auth() IAuthClient {
	return &AuthClient{c.HttpClient}
}
