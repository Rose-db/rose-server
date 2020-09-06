package main

import (
	"fmt"
	"os"
	"runtime"
)

func CreateDbIfNotExists(msgCom chan string, errCom chan IError) {
	var roseDb string
	var fsErr IError

	roseDb = fmt.Sprintf("%s/.rose_db", UserHomeDir())

	msgCom<- "Creating the database on the filesystem if not exists..."
	if _, err := os.Stat(roseDb); os.IsNotExist(err) {
		msgCom<- "Database not found. Creating it now from scratch..."
		err = os.Mkdir(roseDb, os.ModePerm)

		if err != nil {
			close(msgCom)
			fsErr = &SystemError{
				Code:    SystemErrorCode,
				Message: err.Error(),
			}

			errCom<- fsErr

			close(errCom)

			return
		}
	}

	msgCom<- "Filesystem database created successfully"

	close(msgCom)
	close(errCom)
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
