package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// IsValidFile verify if the file extension is valid
func IsValidFile(file string) bool {
	matched, err := regexp.MatchString("(?im)^\\w+.csv$", file)
	if err != nil {
		return false
	}
	return matched
}

// ReadFile open the gived file and returns it to work with it
func ReadFile(path string) (string, error) {
	base := filepath.Base(path)
	if !IsValidFile(base) {
		return "", errors.New("The file extension is not valid")
	}
	fmt.Println(base)
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	byteContent, err := ioutil.ReadFile(absPath)
	return string(byteContent), err
}

// CreateFile creates a file and returns its reference to work with it
func CreateFile(path string) (file *os.File, err error) {
	base := filepath.Base(path)
	if !IsValidFile(base) {
		return nil, errors.New("The file extension is not valid")
	}
	fmt.Println(base)
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	file, err = os.Create(absPath)
	return file, err
}

// FileExists returns true if the file trully exists in the system, else return false
func FileExists(filename string) bool {
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return false
	}
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		// File definitely does not exist.
		return false
	}

	/*
		if _, err := os.Stat(filename); err == nil {
			// File definitely exists!
			return true
		}
	*/
	return true
}
