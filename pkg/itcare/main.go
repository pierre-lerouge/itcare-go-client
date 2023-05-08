package itcare

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

func (itc *ITCareClient) Connect() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     itc.ClientID,
		ClientSecret: itc.ClientSecret,
		Scopes:       []string{"openid"},
		TokenURL:     "https://accounts.cegedim.cloud/auth/realms/cloud/protocol/openid-connect/token",
	}
	itc.Client = conf.Client(ctx)
}

func (itc *ITCareClient) GetCI(ciName string) {
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
	fmt.Printf("Response is %s\n", result.Body)
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
	fmt.Printf("Instance CPU Number is %d\n", instanceResponse.Content[0].CPU)
}
