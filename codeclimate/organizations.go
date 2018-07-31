package codeclimate

import (
	"encoding/json"
)

type CodeClimateOrganisationsData struct {
	Data []*CodeClimateOrganisation `json:"data"`
}

type CodeClimateOrganisation struct {
	Id         string                             `json:"id"`
	Attributes *CodeClimateOrganisationAttributes `json:"attributes"`
}

type CodeClimateOrganisationAttributes struct {
	Name string `json:"name"`
}

const orgUri string = "/orgs"

func (c *Client) GetOrganisations() (*CodeClimateOrganisationsData, error) {
	getData := make([]byte, 100)
	orgData, requestErr := c.MakeRequest("GET", orgUri, getData)

	if requestErr != nil {
		return nil, requestErr
	}

	org := &CodeClimateOrganisationsData{}
	unmarshalErr := json.Unmarshal([]byte(orgData), org)

	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return org, nil
}
