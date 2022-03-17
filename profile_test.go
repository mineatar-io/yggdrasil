package yggdrasil_test

import (
	"testing"

	"github.com/mineatar-io/yggdrasil"
)

func TestUsernameToUUID(t *testing.T) {
	profile, err := yggdrasil.UsernameToUUID("Notch")

	if err != nil {
		t.Fatal(err)
	}

	if profile == nil {
		t.Fatalf("profile is nil")
	}
}

func TestUsernamesToUUIDs(t *testing.T) {
	profiles, err := yggdrasil.UsernamesToUUIDs([]string{"Notch", "Dinnerbone", "MHF_Steve", "PassTheMayo", "MHF_Alex"})

	if err != nil {
		t.Fatal(err)
	}

	if len(profiles) < 5 {
		t.Fatalf("failed to retrieve at least 1 profile")
	}
}

func TestUsernameHistory(t *testing.T) {
	if _, err := yggdrasil.GetUsernameHistory("069a79f444e94726a5befca90e38aaf5"); err != nil {
		t.Fatal(err)
	}
}
