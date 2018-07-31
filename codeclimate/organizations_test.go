package codeclimate

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestGetOrganisations(t *testing.T) {
	token := os.Getenv("CODECLIMATE_API_TOKEN")
	client, err := NewClient(token)

	if err != nil {
		t.Fatal(err)
	}

	org, err2 := client.GetOrganisations()

	if err2 != nil {
		t.Fatal(err)
	}

	orgMarshaled, err3 := json.Marshal(org)

	if err3 != nil {
		t.Fatal(err)
	}

	fmt.Printf(string(orgMarshaled))
}
