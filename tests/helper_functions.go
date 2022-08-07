package test

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/qascade/yast/config"
	"github.com/stretchr/testify/require"
)

var (
	TestConfigJsonPath string
)

func init() {
	currDirPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	TestConfigJsonPath = currDirPath + "/testconfig.json"
}

func setupYastEnv() (configFile *os.File, err error) {
	err = config.CreateYastWorkDir()
	if err != nil {
		return nil, fmt.Errorf("error setting up Yast environment: %v", err)
	}
	configFile, err = config.CreateConfigJSON()
	if err != nil {
		return nil, fmt.Errorf("error setting up Yast environment: %v", err)
	}
	return configFile, nil

}

func getFileHash(f string) (hashVal string, err error) {
	file, err := os.Open(f)
	if err != nil {
		return hashVal, fmt.Errorf("opening file %s:%w", f, err)
	}

	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)

	if err != nil {
		return hashVal, fmt.Errorf("calculating hash of %s:%w", f, err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
func assertFilesMatch(t *testing.T, path string, otherPath string) {
	fileHash, err := getFileHash(path)
	if err != nil {
		t.Errorf("calculating file hash")
	}

	otherHash, err := getFileHash(otherPath)
	if err != nil {
		t.Errorf("calculating other file hash")
	}
	require.Equal(t, fileHash, otherHash, fmt.Sprintf("file hashes do not match: %s != %s", path, otherPath))
}
