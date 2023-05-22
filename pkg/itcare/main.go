package itcare

import (
	"context"
	"fmt"
	"os"

	resty "github.com/go-resty/resty/v2"
	oauth2 "golang.org/x/oauth2/clientcredentials"
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

// TODO : make it configurable
const baseURL = "https://api.cegedim.cloud/itcare"
const OIDCHost = "accounts.cegedim.cloud"

// Connect will use the client ID / Secret givent or will directly
// get them from ITCARE_CLIENT_ID ITCARE_CLIENT_SECRET variable
func (itc *ITCareClient) Connect() {
	if itc.ClientID == "" && itc.ClientSecret == "" {

		if os.Getenv("ITCARE_CLIENT_ID") == "" {
			// TODO Raise Exception
			panic(fmt.Errorf("could not find client ID param or ITCARE_CLIENT_ID variable"))
		}
		itc.ClientID = os.Getenv("ITCARE_CLIENT_ID")
		if os.Getenv("ITCARE_CLIENT_SECRET") == "" {
			// TODO Raise Exception
			panic(fmt.Errorf("could not find client Secret param or ITCARE_CLIENT_SECRET variable"))
		}
		itc.ClientSecret = os.Getenv("ITCARE_CLIENT_SECRET")
	}
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     itc.ClientID,
		ClientSecret: itc.ClientSecret,
		Scopes:       []string{"openid"},
		TokenURL:     fmt.Sprintf("https://%s/auth/realms/cloud/protocol/openid-connect/token", OIDCHost),
	}
	if itc.ClientApp == "" {
		itc.ClientApp = "itcare-go-client/v1"
	}
	itc.Client = resty.NewWithClient(conf.Client(ctx))
	itc.Client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   itc.ClientApp})
	itc.Client.SetBaseURL(baseURL)
}

func (itc *ITCareClient) GetCI(ciName string) (instances Instance, err error) {
	fmt.Printf("Looking for %s \n", ciName)
	// instanceResponse will hold the content of the result as a struct
	var instanceResponse = new(InstanceResponse)
	err = nil
	_, err = itc.Client.R().
		SetQueryParams(map[string]string{
			"names": ciName,
		}).
		SetResult(instanceResponse).
		Get("/compute/instances")

	if err != nil {
		fmt.Printf("Could not get instance : %s\n", err)
		return
	}
	if len(instanceResponse.Content) > 1 {
		fmt.Printf("Warning GetCI returns multiple CI, returning the first one")
	}
	if len(instanceResponse.Content) == 0 {
		err = fmt.Errorf("no results found")
		return
	}
	return instanceResponse.Content[0], err
}
