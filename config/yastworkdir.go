/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package config

import (
	"fmt"
	"os"
)

var (
	YastWorkDir       string
	DefaultConfigPath string
)

type DirError struct {
	err         error
	YastWorkDir string
}

func (e *DirError) Error() string {
	return fmt.Sprintf("err %s: could not create default yast work dir %s", e.err, e.YastWorkDir)
}

func CreateYastWorkDir() error {
	err := os.MkdirAll(YastWorkDir, 0755)
	if err != nil {
		return &DirError{
			err:         err,
			YastWorkDir: YastWorkDir,
		}
	}
	return nil
}

func CreateConfigJSON() (*os.File, error) {
	configFile, err := os.Create(DefaultConfigPath)
	if err != nil {
		return nil, err
	}
	ConfigJsonExists = true
	return configFile, nil
}

func RemoveConfigJSON() error {
	err := os.Remove(DefaultConfigPath)
	if err != nil {
		return err
	}
	ConfigJsonExists = false
	return nil
}

func init() {
	var userHomeDir string
	var err error
	userHomeDir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	YastWorkDir = userHomeDir + "/.yast/"
	DefaultConfigPath = YastWorkDir + "config.json"
}
