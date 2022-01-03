package file

import (
	"io/ioutil"
	"os"
)

/* -- Sequential Access Files -- */

//ReadFile return all the file as a byte array
func Read(path string) (string, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(fileContent), nil
}

//WriteFile overwrite the whole file.
func Write(path, text string) error {
	err := ioutil.WriteFile(path, []byte(text), 0644)
	return err
}

//AppendFile append an array of byte type to an existing file.
func Append(path, text string) error {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	fileContent = append(fileContent, []byte(text)...)
	err = ioutil.WriteFile(path, fileContent, 0644)
	return err
}

/* -- Vano2903's Exist function -- */

//Exist checks that a file exists and returns the existence of the check.
func Exist(path string) bool {
	data, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !data.IsDir()
}
