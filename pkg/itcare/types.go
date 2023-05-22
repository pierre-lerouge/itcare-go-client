package itcare

import resty "github.com/go-resty/resty/v2"

const (
	InstanceType string = "instance"
)

// ITCareClient holds the http client and authentication params to use the ITCareAPI
type ITCareClient struct {
	// Client holding the HTTPClient
	Client *resty.Client
	// ClientSecret holds the secret token for OIDC Authentication
	ClientSecret string
	// ClientID holds the id token for OIDC Authentication
	ClientID string
	// ClientApp holds the user Agent use to contact the ITCare API
	ClientApp string
}
type CI interface {
	getID() int
	getType() string
}
