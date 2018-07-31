// TODO! After Code Climate will complete their API
// 1. Implement method to update organisation settings
// 2. Implement method to delete organisations
// What should be done already:
// 1. Implement method to create organisations
// 2. Implement check if created organisation actually exists

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
