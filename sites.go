package tableau

import (
	"encoding/xml"
	"fmt"
	"github.com/tiketdatarisal/tableau-go-client/models"
	"net/http"
)

type Sites struct {
	context *Context
}

func (s *Sites) QuerySites() ([]models.Site, error) {
	if s.context.activeUserToken == "" {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(serverError, notLoginError)}
	}

	const relURL = "/sites"
	url, err := s.context.GetURL(relURL)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	const pageSize = 100
	pageNumber := 1
	var totalReturned int
	var sites []models.Site

	for {
		_url := fmt.Sprintf("%s?pageSize=%d&pageNumber=%d", url, pageSize, pageNumber)

		res, err := s.context.client.R().
			SetHeader(headerContentType, contentTypeXML).
			SetHeader(headerTableauAuth, s.context.activeUserToken).
			Get(_url)
		if err != nil {
			return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(serverError, err)}
		}

		if s := res.StatusCode(); s != http.StatusOK {
			switch s {
			case http.StatusBadRequest:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, incompleteRequestError)}
			case http.StatusForbidden:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, switchSiteError)}
			case http.StatusMethodNotAllowed:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, invalidMethodError)}
			default:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, unknownError)}
			}
		}

		q := models.QuerySitesResponse{}
		err = xml.Unmarshal(res.Body(), &q)
		if err != nil {
			return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(failedUnmarshalError)}
		}

		totalAvailable := q.Pagination.TotalAvailable
		pageNumber++
		totalReturned += pageSize

		if len(*q.Sites) > 0 {
			sites = append(sites, *q.Sites...)
		}

		if totalReturned >= totalAvailable {
			break
		}
	}

	return sites, nil
}
