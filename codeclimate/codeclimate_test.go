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
	token := os.Getenv("CODECLIMATE_API_TOKEN")
	client, err := NewClient(token)

	if err != nil {
		t.Fatal(err)
	}

	_, err2 := MakeRequest(client, "GET", "/user")

	if err2 != nil {
		t.Fatal(err)
	}
}
