package yggdrasil_test

import (
	"log"
	"testing"

	"github.com/mineatar-io/yggdrasil"
)

func TestProfileTextures(t *testing.T) {
	profile, err := yggdrasil.GetProfileTextures("069a79f444e94726a5befca90e38aaf5")

	if err != nil {
		t.Fatal(err)
	}

	if profile == nil {
		t.Fatalf("profile is nil")
	}

	log.Println(profile)
}

func TestTexturesDecode(t *testing.T) {
	profile, err := yggdrasil.GetProfileTextures("069a79f444e94726a5befca90e38aaf5")

	if err != nil {
		t.Fatal(err)
	}

	if profile == nil {
		t.Fatalf("profile is nil")
	}

	textures := ""

	for _, property := range profile.Properties {
		if property.Name != "textures" {
			continue
		}

		textures = property.Value
	}

	if len(textures) < 1 {
		t.Fatalf("profile does not contain textures")
	}

	if _, err = yggdrasil.GetDecodedTexturesValue(textures); err != nil {
		t.Fatal(err)
	}
}
