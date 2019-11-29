package borg

import (
	"github.com/hashicorp/go-version"
	"strings"
	"vorta/models"
)

type VersionRun struct {
	BorgRun
}

var BorgVersion = "0.0.0"

var minVersionForFeature = map[string]string{
	"BLAKE2":   "1.1.4",
	"ZSTD":     "1.1.4",
	"JSON_LOG": "1.1.0",
}

func NewVersionRun(profile *models.Profile) (*VersionRun, error) {
	r := &VersionRun{}
	r.SubCommand = "--version"
	r.SubCommandArgs = []string{}
	r.Profile = profile

	err := r.Prepare()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *VersionRun) ProcessResult() {
	versionStr := strings.TrimSpace(r.PlainTextResult)
	BorgVersion = strings.Split(versionStr, " ")[1]
}

func FeatureIsSupported(feature string) bool {
	currentVersion, _ := version.NewVersion(BorgVersion)
	featureMinVersion, _ := version.NewVersion(minVersionForFeature[feature])
	return currentVersion.GreaterThanOrEqual(featureMinVersion)
}
