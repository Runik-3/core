package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const RELEASES_ENDPOINT = "https://api.github.com/repos/Runik-3/core/releases"

func UpdateAvailable() bool {
	latest := fetchLatestRelease()
	current := getCurrentVersion()

	if current != latest {
		return true
	}
	return false
}

type GithubRelease struct {
	// we only care about the Name of this release for now
	Name string `json:"name"`
}

func fetchLatestRelease() string {
	// TODO: Mechanism for logging go process errors
	res, err := http.Get(RELEASES_ENDPOINT)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	releases := []GithubRelease{}
	err = json.Unmarshal(body, &releases)
	if err != nil {
		fmt.Println(err)
	}

	if len(releases) == 0 {
		return ""
	}
	return strings.TrimSpace(releases[0].Name)
}

func getCurrentVersion() string {
	version, err := os.ReadFile("version")
	if err != nil {
		fmt.Println(err)
	}

	return strings.TrimSpace(string(version))
}
