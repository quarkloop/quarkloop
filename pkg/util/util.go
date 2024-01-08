package util

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func GetProjectRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../..")
}

func GetTestEnvFilePath() string {
	return fmt.Sprintf("%s/.env.test", GetProjectRootPath())
}

func GetDevEnvFilePath() string {
	return fmt.Sprintf("%s/.env.development", GetProjectRootPath())
}
