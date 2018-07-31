// TODO!
// 1. Implement method to create private repositories
// 2. Implement method to create public repositories
// 3. Implement method to update repositories
// 4. Implement method to delete repositories

package codeclimate

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type CodeClimateRepositoryData struct {
	Data *CodeClimateRepository `json:"data"`
}

type CodeClimateRepository struct {
	Id         string                           `json:"id"`
	Attributes *CodeClimateRepositoryAttributes `json:"attributes"`
}

type CodeClimateRepositoryAttributes struct {
	Name               string  `json:"human_name"`
	DefaultBranch      string  `json:"branch"`
	GithubOrganization string  `json:"github_organization"`
	GithubSlug         string  `json:"github_slug"`
	Score              float64 `json:"score"`
}

func (c *Client) GetRepository(repoId string) (*CodeClimateRepositoryData, error) {
	const repoUri string = "/repos"
	requestUrl := repoUri + "/" + repoId
	getData := make([]byte, 100)

	repoData, err := c.MakeRequest("GET", requestUrl, getData)

	if err != nil {
		return nil, err
	}

	repo := &CodeClimateRepositoryData{}
	err2 := json.Unmarshal([]byte(repoData), repo)

	if err2 != nil {
		return nil, err2
	}

	// Parsing organisation name
	repo.Data.Attributes.GithubOrganization = strings.Split(
		repo.Data.Attributes.GithubSlug, "/",
	)[0]

	return repo, nil
}

func (c *Client) CreatePrivateGithubRepository(orgName string, repoName string) (*CodeClimateRepositoryData, error) {
	// Check that organisation exist
	orgData, err := c.GetOrganisations()

	if err != nil {
		return nil, err
	}

	var orgId string
	orgExists := false

	for _, organisation := range orgData.Data {
		if organisation.Attributes.Name == orgName {
			orgExists = true
			orgId = organisation.Id
			break
		}
	}

	if orgExists == false {
		orgErr := errors.New("Requested organisation is not found")
		return nil, orgErr
	}

	// Form the data to send request
	postJson := fmt.Sprintf(`{"data":{"type":"repos","attributes":{"url":"https://github.com/%s/%s"}}}`, orgName, repoName)
	postData := []byte(postJson)
	requestUrl := "/orgs/" + orgId + "/repos"
	newRepoData, requestErr := c.MakeRequest("POST", requestUrl, postData)

	if requestErr != nil {
		return nil, requestErr
	}

	// Get the data about repository
	repo := &CodeClimateRepositoryData{}
	unmarshalErr := json.Unmarshal([]byte(newRepoData), repo)

	if unmarshalErr != nil {
		return nil, unmarshalErr
	} else if repo.Data == nil {
		// If data has not been parsed - then we've got an API error
		apiError := errors.New(newRepoData)
		return nil, apiError
	}

	repo.Data.Attributes.GithubOrganization = orgName

	return repo, nil
}
