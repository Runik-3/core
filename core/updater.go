package core

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const RELEASES_ENDPOINT = "https://api.github.com/repos/Runik-3/core/releases"

func UpdateAvailable(version string) bool {
	latest, err := fetchLatestRelease()
	// Updates are non-essential, if we cannot fetch the release info, try again
	// next time, do not crash the app.
	if err != nil {
		return false
	}

	current := strings.TrimSpace(string(version))
	if current != latest {
		return true
	}
	return false
}

type GithubRelease struct {
	// we only care about the Name of this release for now
	Name string `json:"name"`
}

func fetchLatestRelease() (string, error) {
	// TODO: Mechanism for logging go process errors
	res, err := http.Get(RELEASES_ENDPOINT)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	releases := []GithubRelease{}
	err = json.Unmarshal(body, &releases)
	if err != nil {
		return "", err
	}

	if len(releases) == 0 {
		return "", err
	}
	return strings.TrimSpace(releases[0].Name), nil
}
