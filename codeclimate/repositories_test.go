package codeclimate

import (
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

	fmt.Printf(repo)
}
