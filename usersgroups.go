package tableau

import (
	"encoding/xml"
	"fmt"
	"github.com/tiket/tableau-go-client/models"
	"net/http"
)

type UsersAndGroups struct {
	context *Context
}

func (u *UsersAndGroups) AddUserToGroup(userID, groupID string) (*models.User, error) {
	if u.context.activeUserToken == "" {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(serverError, notLoginError)}
	}

	const relURL = "/sites/%v/groups/%v/users"
	url, err := u.context.GetURL(fmt.Sprintf(relURL, u.context.activeSiteID, groupID))
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	req := models.AddUserToGroupRequest{User: models.User{ID: userID}}
	b, err := xml.Marshal(req)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(failedMarshalError)}
	}

	res, err := u.context.client.R().
		SetHeader(headerContentType, contentTypeXML).
		SetHeader(headerTableauAuth, u.context.activeUserToken).
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
		case http.StatusNotFound:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, groupSiteNotFoundError)}
		case http.StatusMethodNotAllowed:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, invalidMethodError)}
		case http.StatusConflict:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, userAlreadyMemberError)}
		default:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, unknownError)}
		}
	}

	s := models.AddUserToGroupResponse{}
	err = xml.Unmarshal(res.Body(), &s)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(failedUnmarshalError)}
	}

	return s.User, nil
}

func (u *UsersAndGroups) GetUsersInGroup(groupID string) ([]models.User, error) {
	if u.context.activeUserToken == "" {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(serverError, notLoginError)}
	}

	const relURL = "/sites/%v/groups/%v/users"
	url, err := u.context.GetURL(fmt.Sprintf(relURL, u.context.activeSiteID, groupID))
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	const pageSize = 100
	pageNumber := 1
	var totalReturned int
	var users []models.User

	for {
		_url := fmt.Sprintf("%s?pageSize=%d&pageNumber=%d", url, pageSize, pageNumber)

		res, err := u.context.client.R().
			SetHeader(headerContentType, contentTypeXML).
			SetHeader(headerTableauAuth, u.context.activeUserToken).
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
			case http.StatusNotFound:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, siteNotFoundError)}
			case http.StatusMethodNotAllowed:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, invalidMethodError)}
			default:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, unknownError)}
			}
		}

		q := models.GetUsersOnGroupResponse{}
		err = xml.Unmarshal(res.Body(), &q)
		if err != nil {
			return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(failedUnmarshalError)}
		}

		totalAvailable := q.Pagination.TotalAvailable
		pageNumber++
		totalReturned += pageSize

		if len(*q.Users) > 0 {
			users = append(users, *q.Users...)
		}

		if totalReturned >= totalAvailable {
			break
		}
	}

	return users, nil
}

func (u *UsersAndGroups) GetUsersOnSite() ([]models.User, error) {
	if u.context.activeUserToken == "" {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(serverError, notLoginError)}
	}

	const relURL = "/sites/%v/users"
	url, err := u.context.GetURL(fmt.Sprintf(relURL, u.context.activeSiteID))
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	const pageSize = 100
	pageNumber := 1
	var totalReturned int
	var users []models.User

	for {
		_url := fmt.Sprintf("%s?pageSize=%d&pageNumber=%d", url, pageSize, pageNumber)

		res, err := u.context.client.R().
			SetHeader(headerContentType, contentTypeXML).
			SetHeader(headerTableauAuth, u.context.activeUserToken).
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
			case http.StatusNotFound:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, siteNotFoundError)}
			case http.StatusMethodNotAllowed:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, invalidMethodError)}
			default:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, unknownError)}
			}
		}

		q := models.GetUsersOnSiteResponse{}
		err = xml.Unmarshal(res.Body(), &q)
		if err != nil {
			return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(failedUnmarshalError)}
		}

		totalAvailable := q.Pagination.TotalAvailable
		pageNumber++
		totalReturned += pageSize

		if len(*q.Users) > 0 {
			users = append(users, *q.Users...)
		}

		if totalReturned >= totalAvailable {
			break
		}
	}

	return users, nil
}

func (u *UsersAndGroups) QueryGroups() ([]models.Group, error) {
	if u.context.activeUserToken == "" {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(serverError, notLoginError)}
	}

	const relURL = "/sites/%v/groups"
	url, err := u.context.GetURL(fmt.Sprintf(relURL, u.context.activeSiteID))
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	const pageSize = 100
	pageNumber := 1
	var totalReturned int
	var groups []models.Group

	for {
		_url := fmt.Sprintf("%s?pageSize=%d&pageNumber=%d", url, pageSize, pageNumber)

		res, err := u.context.client.R().
			SetHeader(headerContentType, contentTypeXML).
			SetHeader(headerTableauAuth, u.context.activeUserToken).
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
			case http.StatusNotFound:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, siteNotFoundError)}
			case http.StatusMethodNotAllowed:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, invalidMethodError)}
			default:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, unknownError)}
			}
		}

		q := models.QueryGroupsResponse{}
		err = xml.Unmarshal(res.Body(), &q)
		if err != nil {
			return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(failedUnmarshalError)}
		}

		totalAvailable := q.Pagination.TotalAvailable
		pageNumber++
		totalReturned += pageSize

		if len(*q.Groups) > 0 {
			groups = append(groups, *q.Groups...)
		}

		if totalReturned >= totalAvailable {
			break
		}
	}

	return groups, nil
}

func (u *UsersAndGroups) QueryUserOnSite(userID string) (*models.User, error) {
	if u.context.activeUserToken == "" {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(serverError, notLoginError)}
	}

	const relURL = "/sites/%v/users/%v"
	url, err := u.context.GetURL(fmt.Sprintf(relURL, u.context.activeSiteID, userID))
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	res, err := u.context.client.R().
		SetHeader(headerContentType, contentTypeXML).
		SetHeader(headerTableauAuth, u.context.activeUserToken).
		Get(url)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(serverError, err)}
	}

	if s := res.StatusCode(); s != http.StatusOK {
		switch s {
		case http.StatusBadRequest:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, incompleteRequestError)}
		case http.StatusForbidden:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, switchSiteError)}
		case http.StatusNotFound:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, siteNotFoundError)}
		case http.StatusMethodNotAllowed:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, invalidMethodError)}
		default:
			return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, unknownError)}
		}
	}

	q := models.QueryUserOnSiteResponse{}
	err = xml.Unmarshal(res.Body(), &q)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(failedUnmarshalError)}
	}

	return q.User, nil
}

func (u *UsersAndGroups) GetGroupsForUser(userID string) ([]models.Group, error) {
	if u.context.activeUserToken == "" {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(serverError, notLoginError)}
	}

	const relURL = "/sites/%v/users/%v/groups"
	url, err := u.context.GetURL(fmt.Sprintf(relURL, u.context.activeSiteID, userID))
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	const pageSize = 100
	pageNumber := 1
	var totalReturned int
	var groups []models.Group

	for {
		_url := fmt.Sprintf("%s?pageSize=%d&pageNumber=%d", url, pageSize, pageNumber)

		res, err := u.context.client.R().
			SetHeader(headerContentType, contentTypeXML).
			SetHeader(headerTableauAuth, u.context.activeUserToken).
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
			case http.StatusNotFound:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, siteNotFoundError)}
			case http.StatusMethodNotAllowed:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, invalidMethodError)}
			default:
				return nil, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, unknownError)}
			}
		}

		q := models.GetGroupsForUserResponse{}
		err = xml.Unmarshal(res.Body(), &q)
		if err != nil {
			return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(failedUnmarshalError)}
		}

		totalAvailable := q.Pagination.TotalAvailable
		pageNumber++
		totalReturned += pageSize

		if len(*q.Groups) > 0 {
			groups = append(groups, *q.Groups...)
		}

		if totalReturned >= totalAvailable {
			break
		}
	}

	return groups, nil
}

func (u *UsersAndGroups) RemoveUserFromGroup(userID, groupID string) (bool, error) {
	if u.context.activeUserToken == "" {
		return false, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(serverError, notLoginError)}
	}

	const relURL = "/sites/%v/groups/%v/users/%v"
	url, err := u.context.GetURL(fmt.Sprintf(relURL, u.context.activeSiteID, groupID, userID))
	if err != nil {
		return false, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	res, err := u.context.client.R().
		SetHeader(headerContentType, contentTypeXML).
		SetHeader(headerTableauAuth, u.context.activeUserToken).
		Delete(url)
	if err != nil {
		return false, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(serverError, err)}
	}

	if s := res.StatusCode(); s != http.StatusNoContent {
		switch s {
		case http.StatusBadRequest:
			return false, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, deleteFailedError)}
		case http.StatusUnauthorized:
			return false, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, loginFailedError)}
		case http.StatusForbidden:
			return false, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, switchSiteError)}
		case http.StatusNotFound:
			return false, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, groupSiteNotFoundError)}
		case http.StatusMethodNotAllowed:
			return false, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, invalidMethodError)}
		default:
			return false, &HttpError{Code: s, Err: fmt.Errorf(tableauServerError, unknownError)}
		}
	}

	return true, nil
}
