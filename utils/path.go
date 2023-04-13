package utils

import (
	"path/filepath"
	"runtime"
)

func GetExePath() (string, error) {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), ".")
	ePath, err := filepath.Abs(projectRoot)
	if err != nil {
		return "", err
	}
	return filepath.Dir(ePath), nil
}
