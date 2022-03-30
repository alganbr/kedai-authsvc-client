package client

import (
	"encoding/json"
	"fmt"
	"github.com/alganbr/kedai-authsvc-client/models"
	"github.com/alganbr/kedai-utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"net/http"
)

type IAuthClient interface {
	Get(string) (*models.AccessToken, *errors.Error)
	Authenticate(*models.AccessTokenRq) (*models.AccessToken, *errors.Error)
}

type AuthClient struct {
	httpClient *rest.RequestBuilder
}

func (c *AuthClient) Get(id string) (*models.AccessToken, *errors.Error) {
	rs := c.httpClient.Get(fmt.Sprintf("/authsvc/auth/%s", id))
	if rs.Err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: rs.Err.Error(),
		}
	}
	if rs.StatusCode > 299 {
		var httpErr *errors.Error
		if err := json.Unmarshal(rs.Bytes(), &httpErr); err != nil {
			return nil, &errors.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error when unmarshalling error response",
			}
		}
		return nil, httpErr
	}
	var accessToken *models.AccessToken
	if err := json.Unmarshal(rs.Bytes(), &accessToken); err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error when unmarshalling response body",
		}
	}
	return accessToken, nil
}

func (c *AuthClient) Authenticate(rq *models.AccessTokenRq) (*models.AccessToken, *errors.Error) {
	rs := c.httpClient.Post("/authsvc/auth", rq)
	if rs.Err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: rs.Err.Error(),
		}
	}
	if rs.StatusCode > 299 {
		var httpErr *errors.Error
		if err := json.Unmarshal(rs.Bytes(), &httpErr); err != nil {
			return nil, &errors.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error when unmarshalling error response",
			}
		}
		return nil, httpErr
	}
	var accessToken *models.AccessToken
	if err := json.Unmarshal(rs.Bytes(), &accessToken); err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error when unmarshalling response body",
		}
	}
	return accessToken, nil
}
