package tableau

import (
	"encoding/xml"
	"fmt"
	"github.com/tiketdatarisal/tableau-go-client/models"
	"net/http"
)

type Authentications struct {
	context *Context
}

func (a *Authentications) SignIn(user, password string) (*models.Credentials, error) {
	const relURL = "/auth/signin"
	url, err := a.context.GetURL(relURL)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	req := models.SignInRequest{Credentials: models.Credentials{Name: user, Password: password, Site: &models.Site{ContentURL: ""}}}
	b, err := xml.Marshal(req)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(failedMarshalError)}
	}

	res, err := a.context.client.R().
		SetHeader(headerContentType, contentTypeXML).
		SetBody(b).
		Post(url)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(serverError, err)}
	}

	if s := res.StatusCode(); s != http.StatusOK {
		switch s {
		case http.StatusBadRequest:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, incompleteRequestError)}
		case http.StatusUnauthorized:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, loginFailedError)}
		case http.StatusMethodNotAllowed:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, invalidMethodError)}
		default:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, unknownError)}
		}
	}

	s := models.AuthResponse{}
	err = xml.Unmarshal(res.Body(), &s)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(failedUnmarshalError)}
	}

	a.context.activeUserToken = s.Credentials.Token
	a.context.activeUserID = s.Credentials.User.ID
	a.context.activeSiteID = s.Credentials.Site.ID
	return &s.Credentials, nil
}

func (a *Authentications) SignOut() error {
	if a.context.activeUserToken == "" {
		return nil
	}

	const relURL = "/auth/signout"
	url, err := a.context.GetURL(relURL)
	if err != nil {
		return &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	res, err := a.context.client.R().
		SetHeader(headerContentType, contentTypeXML).
		SetHeader(headerTableauAuth, a.context.activeUserToken).
		Post(url)
	if err != nil {
		return &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(serverError, err)}
	}

	if s := res.StatusCode(); s != http.StatusNoContent {
		switch s {
		case http.StatusMethodNotAllowed:
			return &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, invalidMethodError)}
		default:
			return &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, unknownError)}
		}
	}

	a.context.activeUserToken = ""
	a.context.activeUserID = ""
	a.context.activeSiteID = ""
	return nil
}

func (a *Authentications) SwitchSite(site string) (*models.Credentials, error) {
	if a.context.activeUserToken == "" {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(serverError, notLoginError)}
	}

	const relURL = "/auth/switchSite"
	url, err := a.context.GetURL(relURL)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	req := models.SwitchSiteRequest{Site: models.Site{ContentURL: site}}
	b, err := xml.Marshal(req)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(failedMarshalError)}
	}

	res, err := a.context.client.R().
		SetHeader(headerContentType, contentTypeXML).
		SetHeader(headerTableauAuth, a.context.activeUserToken).
		SetBody(b).
		Post(url)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(serverError, err)}
	}

	if s := res.StatusCode(); s != http.StatusOK {
		switch s {
		case http.StatusBadRequest:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, incompleteRequestError)}
		case http.StatusUnauthorized:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, loginFailedError)}
		case http.StatusForbidden:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, switchSiteError)}
		case http.StatusMethodNotAllowed:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, invalidMethodError)}
		default:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, unknownError)}
		}
	}

	s := models.AuthResponse{}
	err = xml.Unmarshal(res.Body(), &s)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(failedUnmarshalError)}
	}

	a.context.activeUserToken = s.Credentials.Token
	a.context.activeUserID = s.Credentials.User.ID
	a.context.activeSiteID = s.Credentials.Site.ID
	return &s.Credentials, nil
}
