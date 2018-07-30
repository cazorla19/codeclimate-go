// TODO!
// 1. Implement method to create private repositories
// 2. Implement method to create public repositories
// 3. Implement method to update repositories
// 4. Implement method to delete repositories

package codeclimate

import (
	"encoding/json"
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

	repo := &CodeClimateRepositoryData{
		Data: &CodeClimateRepository{
			Attributes: &CodeClimateRepositoryAttributes{},
		},
	}
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
