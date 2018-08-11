package codeclimate

import (
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	token := os.Getenv("CODECLIMATE_API_TOKEN")
	_, err := NewClient(token)

	if err != nil {
		t.Fatal(err)
	}
}

func TestMakeRequest(t *testing.T) {
	getData := make([]byte, 100)

	token := os.Getenv("CODECLIMATE_API_TOKEN")
	client, err := NewClient(token)

	if err != nil {
		t.Fatal(err)
	}

	_, err2 := client.MakeRequest("GET", "/user", getData)

	if err2 != nil {
		t.Fatal(err)
	}
}
