package sdlog

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gaorx/stardust4/sderr"
)

func ifEmptyAs(s, def string) string {
	if s == "" {
		return def
	} else {
		return s
	}
}

func absPath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", sderr.Wrap(err, "get home directory error for sdlog")
		}
		path = home + strings.TrimPrefix(path, "~")
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", sderr.Wrap(err, "compute log absolute path error")
	}
	return abs, nil
}
