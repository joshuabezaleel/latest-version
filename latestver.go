package latestver

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	version "github.com/hashicorp/go-version"
)

func LatestVersion(module string) (string, error) {
	if module == "" {
		return "", errors.New("module is empty, please supply")
	}

	// Trim latest character if the module supplied ended with "/"
	// "github.com/joshuabezaleel/latest-version/ becomes github.com/joshuabezaleel/latest-version"
	if latestChar := module[len(module)-1:]; latestChar == "/" {
		module = module[:len(module)]
	}

	resp, err := http.Get("https://proxy.golang.org/" + module + "/@v/list")
	if err != nil {
		return "", fmt.Errorf("failed sending GET request to proxy.golang.org")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed reading response Body")
	}

	latestVer, err := parseModuleLatestVersion(body)
	if err != nil {
		return "", fmt.Errorf("failed parsing module's latest version")
	}

	return latestVer, nil
}

func parseModuleLatestVersion(respBody []byte) (string, error) {
	respBodyString := string(respBody)

	// Delete all the "v" from the response Body, outputting only the version's numbers
	respTrimmed := strings.ReplaceAll(respBodyString, "v", "")
	respTrimmed = strings.Trim(respTrimmed, " ")

	rawVersions := strings.Split(respTrimmed, "\n")
	versions := make([]*version.Version, len(rawVersions)-1)
	for i := 0; i < len(rawVersions)-1; i++ {
		v, err := version.NewVersion(rawVersions[i])
		if err != nil {
			return "", fmt.Errorf("failed parsing version; version=%s", rawVersions[i])
		}
		versions[i] = v
	}

	// Sort versions in ascending order
	sort.Sort(version.Collection(versions))

	return versions[len(versions)-1].Original(), nil
}
