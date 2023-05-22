package itcare

// InstanceResponse holds the list of instances returned with pagionation informations
type InstanceResponse struct {
	Content          []Instance `json:"content"`
	NumberOfElements uint16     `json:"numberOfElement"`
}

// Instance Holds the compute instance definition
// ref https://itcare-facade-api.itcare.ccs.cegedim.cloud/swagger-ui/index.html#/Operation(s)/getInstancesUsingGET
type Instance struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Status       string `json:"status"`
	Family       string `json:"family"`
	Label        string `json:"Label"`
	ServiceID    int    `json:"serviceId"`
	ServiceKey   string `json:"serviceKey"`
	Environment  string `json:"environment"`
	CreationUser string `json:"creationUser"`
	// TODO : Timestamp
	CreationTime          string          `json:"creationTime"`
	IPAddress             string          `json:"ipAddress"`
	Network               InstanceNetwork `json:"network"`
	LabelRegion           string          `json:"labelRegion"`
	LabelDataCenter       string          `json:"labelDatacenter"`
	LabelArea             string          `json:"labelArea"`
	LabelAvailabilityZone string          `json:"labelAvailabilityZone"`
	TotalSizeDisks        string          `json:"totalSizeDisks"` // TODO cast this into float
	CPU                   uint8           `json:"cpu"`
	RAM                   uint8           `json:"ram"`
}
type InstanceNetwork struct {
	TechnicalID string        `json:"technicalID"`
	CIDR        string        `json:"cidr"`
	Scope       string        `json:"scope"`
	IPAddress   string        `json:"ipAddress"`
	DNS         []InstanceDNS `json:"dns"`
}

type InstanceDNS struct {
	Alias  string `json:"alias"`
	Domain string `json:"domain"`
}

func (inst Instance) GetID() (id int) {
	return inst.Id
}
func (inst Instance) GetType() string {
	return InstanceType
}
