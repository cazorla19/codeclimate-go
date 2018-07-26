package codeclimate

import (
	"encoding/json"
)

type CodeClimateRepositoryData struct {
	Data *CodeClimateRepository `json:"data"`
}

type CodeClimateRepository struct {
	Id         string                           `json:"id"`
	Attributes *CodeClimateRepositoryAttributes `json:"attributes"`
}

type CodeClimateRepositoryAttributes struct {
	Name          string  `json:"human_name"`
	DefaultBranch string  `json:"branch"`
	Score         float64 `json:"score"`
}

func (c *Client) GetRepository(repoId string) (*CodeClimateRepositoryData, error) {
	const repoUri string = "/repos"
	requestUrl := repoUri + "/" + repoId
	repoData, err := c.MakeRequest("GET", requestUrl)

	if err != nil {
		return "", err
	}

	repo := &CodeClimateRepositoryData{
		Data: &CodeClimateRepository{
			Attributes: &CodeClimateRepositoryAttributes{},
		},
	}
	err2 := json.Unmarshal([]byte(repoData), repo)

	if err2 != nil {
		return "", err2
	}

	return repo, nil
}
