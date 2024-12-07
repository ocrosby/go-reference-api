package utils

import (
	"io"
	"net/http"
	"os"
)

// FileExists checks if a file exists and is not a directory
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DeleteFile deletes a file if it exists
func DeleteFile(filename string) error {
	if FileExists(filename) {
		return os.Remove(filename)
	}
	return nil
}

// CreateDirectory creates a directory if it does not exist
func CreateDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}

// DownloadFile downloads a file from a URL
func DownloadFile(url, filename string) error {
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
