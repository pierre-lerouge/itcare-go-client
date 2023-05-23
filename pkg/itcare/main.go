package itcare

import (
	"context"
	"fmt"
	"os"

	resty "github.com/go-resty/resty/v2"
	oauth2 "golang.org/x/oauth2/clientcredentials"
)

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
		"Accept":       "application/vnd.cegedim-it.v1+json",
		"Content-Type": "application/json",
		"User-Agent":   itc.ClientApp})
	itc.Client.SetBaseURL(baseURL)
}

// GetInstance return the instance base on the name given
func (itc *ITCareClient) GetInstance(ciName string) (instance *Instance, err error) {
	fmt.Printf("Looking for %s \n", ciName)
	instancesResponse, err := itc.getInstanceBy("names", ciName)
	if len(instancesResponse.Content) > 1 {
		fmt.Printf("Warning GetCI returns multiple CI, returning the first one")
	}
	if len(instancesResponse.Content) == 0 {
		err = fmt.Errorf("no results found")
		return
	}
	return &instancesResponse.Content[0], err
}

// Todo plural
func (itc *ITCareClient) getInstanceBy(propertyName string, propertyValue string) (instanceResponse *InstancesResponse, err error) {
	// instanceResponse will hold the content of the result as a struct
	err = nil
	instanceResponse = new(InstancesResponse)
	_, err = itc.Client.R().
		SetQueryParams(map[string]string{
			propertyName: propertyValue,
		}).
		SetResult(instanceResponse).
		Get("/compute/instances")
	if err != nil {
		fmt.Printf("Could not get instance : %s\n", err)
		return
	}

	return instanceResponse, err
}
func (itc *ITCareClient) GetInstanceByID(id int, withNetwork bool) (instance *Instance, err error) {
	err = nil
	instance = new(Instance)
	instanceUrl := fmt.Sprintf("/compute/instances/%d", id)
	fmt.Printf("Looking for : %s\n", instanceUrl)
	_, err = itc.Client.R().
		SetResult(instance).
		Get(instanceUrl)
	if err != nil {
		fmt.Printf("Could not get instance : %s\n", err)
		return
	}
	if withNetwork {
		networkURL := fmt.Sprintf("/compute/instances/%d/networks", id)
		networkResult := new(InstanceNetwork)
		_, err = itc.Client.R().
			SetResult(networkResult).
			Get(networkURL)
		instance.Network = *networkResult
	}
	fmt.Println(instance)
	return
}
