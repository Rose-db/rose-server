package main

import (
	"fmt"
	"os"
	"runtime"
)

func CreateDbIfNotExists() IError {
	var roseDb string

	roseDb = fmt.Sprintf("%s/.rose_db", UserHomeDir())

	if _, err := os.Stat(roseDb); os.IsNotExist(err) {
		err = os.Mkdir(roseDb, os.ModePerm)

		if err != nil {
			return &SystemError{
				Code:    SystemErrorCode,
				Message: err.Error(),
			}
		}
	}

	return nil
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}

	return os.Getenv("HOME")
}
