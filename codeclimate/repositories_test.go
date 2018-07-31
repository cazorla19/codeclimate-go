package codeclimate

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestGetRepository(t *testing.T) {
	token := os.Getenv("CODECLIMATE_API_TOKEN")
	client, err := NewClient(token)

	if err != nil {
		t.Fatal(err)
	}

	repo, err2 := client.GetRepository("123456789")

	if err2 != nil {
		t.Fatal(err)
	}

	repoMarshaled, err3 := json.Marshal(repo)

	if err3 != nil {
		t.Fatal(err)
	}

	fmt.Printf(string(repoMarshaled))
}

func TestCreatePrivateRepository(t *testing.T) {
	token := os.Getenv("CODECLIMATE_API_TOKEN")
	client, err := NewClient(token)

	if err != nil {
		t.Fatal(err)
	}

	privateRepo := true
	repo, err2 := client.CreateGithubRepository("babbel-sandbox", "new.private.babbel", privateRepo)

	if err2 != nil {
		t.Fatal(err)
	}

	repoMarshaled, err3 := json.Marshal(repo)

	if err3 != nil {
		t.Fatal(err)
	}

	fmt.Printf(string(repoMarshaled))
}

func TestCreatePublicRepository(t *testing.T) {
	token := os.Getenv("CODECLIMATE_API_TOKEN")
	client, err := NewClient(token)

	if err != nil {
		t.Fatal(err)
	}

	privateRepo := false
	repo, err2 := client.CreateGithubRepository("babbel-sandbox", "test.public.babbel", privateRepo)

	if err2 != nil {
		t.Fatal(err)
	}

	repoMarshaled, err3 := json.Marshal(repo)

	if err3 != nil {
		t.Fatal(err)
	}

	fmt.Printf(string(repoMarshaled))
}
