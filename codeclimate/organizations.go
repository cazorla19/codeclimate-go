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

func (c *Client) GetOrganisations(orgId string) (*CodeClimateOrganisationsData, error) {
	getData := make([]byte, 100)
	orgData, err := c.MakeRequest("GET", orgUri, getData)

	if err != nil {
		return nil, err
	}

	org := &CodeClimateOrganisationsData{}
	err2 := json.Unmarshal([]byte(orgData), org)

	if err2 != nil {
		return nil, err2
	}

	return org, nil
}
