package itcare

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	oauth2 "golang.org/x/oauth2/clientcredentials"
)

// ITCareClient holds the http client and authentication params to use the ITCareAPI
type ITCareClient struct {
	// Client holding the HTTPClient
	Client *http.Client
	// ClientSecret holding the secret token for OIDC Authentication
	ClientSecret string
	// ClientID holding the id token for OIDC Authentication
	ClientID string
}

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
		TokenURL:     "https://accounts.cegedim.cloud/auth/realms/cloud/protocol/openid-connect/token",
	}
	itc.Client = conf.Client(ctx)
}

func (itc *ITCareClient) GetCI(ciName string) (instances []Instance) {
	fmt.Printf("Looking for %s \n", ciName)
	req, err := http.NewRequest("GET", "https://api.cegedim.cloud/itcare/compute/instances", nil)
	if err != nil {
		fmt.Printf("Could not prepare new request %s", err)
		return
	}
	query := req.URL.Query()
	query.Add("names", ciName)
	req.URL.RawQuery = query.Encode()

	result, err := itc.Client.Do(req)
	if err != nil {
		fmt.Printf("Could not get instance : %s\n", err)
		return
	}
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Printf("Could not read response body %s\n", err)
		return
	}
	var instanceResponse = new(InstanceResponse)
	err = json.Unmarshal(body, instanceResponse)
	if err != nil {
		fmt.Printf("Could not unmarshal response instance %s\n", err)
	}
	fmt.Printf("Availability Zone is %s\n", instanceResponse.Content[0].LabelAvailabilityZone)
	return instanceResponse.Content
}
