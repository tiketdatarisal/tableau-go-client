package tableau

import (
	"github.com/go-resty/resty/v2"
	"net/url"
	"path"
)

const (
	headerContentType   = "Content-Type"
	contentTypeXML      = "application/xml"
	headerAuthorization = "Authorization"
	authorizationBearer = "Bearer %v"
	headerTableauAuth   = "X-Tableau-Auth"

	badURLError            = "bad server URL"
	failedMarshalError     = "failed to marshal object"
	failedUnmarshalError   = "failed to unmarshall object"
	serverError            = "server error, details: %v"
	tableauServerError     = "tableau server error, details: %v"
	invalidMethodError     = "invalid request method"
	loginFailedError       = "login failed, please check credential"
	incompleteRequestError = "incomplete request"
	unknownError           = "unknown error"
	notLoginError          = "user not logged in"
	switchSiteError        = "failed to switch site, destination and source are the same"
	siteNotFoundError      = "site not found"
	userAlreadyMemberError = "user already a member of the group"
	groupSiteNotFoundError = "group/site not found"
	deleteFailedError      = "delete failed"
)

type Context struct {
	client          *resty.Client
	baseURL         string
	activeUserToken string
	activeUserID    string
	activeSiteID    string
	Authentications *Authentications
	Sites           *Sites
	UsersAndGroups  *UsersAndGroups
}

func NewContext(s string) (*Context, error) {
	ctx := &Context{client: resty.New()}
	err := ctx.SetBaseURL(s)
	if err != nil {
		return nil, err
	}

	ctx.Authentications = &Authentications{context: ctx}
	ctx.Sites = &Sites{context: ctx}
	ctx.UsersAndGroups = &UsersAndGroups{context: ctx}
	return ctx, err
}

func (c *Context) SetBaseURL(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	c.baseURL = u.String()
	return nil
}

func (c *Context) GetBaseURL() string { return c.baseURL }

func (c *Context) GetURL(p string) (string, error) {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, p)
	return u.String(), err
}
