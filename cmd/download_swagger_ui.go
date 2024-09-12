package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var downloadSwaggerUICmd = &cobra.Command{
	Use:   "download-swagger-ui",
	Short: "Download and unpack the latest Swagger UI",
	Run: func(cmd *cobra.Command, args []string) {
		downloadSwaggerUI()
	},
}

func init() {
	rootCmd.AddCommand(downloadSwaggerUICmd)
}

func downloadSwaggerUI() {
	const swaggerUIURL = "https://api.github.com/repos/swagger-api/swagger-ui/releases/latest"
	const tarballName = "swagger-ui.tar.gz"
	const tempDir = "swagger-ui-temp"
	const targetDir = "swagger-ui"

	// Remove existing tarball if it exists
	if _, err := os.Stat(tarballName); err == nil {
		_ = os.Remove(tarballName)
	}

	// Create temporary directory
	_ = os.MkdirAll(tempDir, os.ModePerm)

	// Fetch the latest Swagger UI tarball URL
	resp, err := http.Get(swaggerUIURL)
	if err != nil {
		fmt.Printf("Error fetching Swagger UI URL: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Extract tarball URL from response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
	tarballURL := extractTarballURL(string(body))

	// Download the tarball
	out, err := os.Create(tarballName)
	if err != nil {
		fmt.Printf("Error creating tarball file: %v\n", err)
		return
	}
	defer out.Close()

	resp, err = http.Get(tarballURL)
	if err != nil {
		fmt.Printf("Error downloading tarball: %v\n", err)
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Error saving tarball: %v\n", err)
		return
	}

	// Extract the tarball
	err = exec.Command("tar", "-xzf", tarballName, "--strip-components=1", "-C", tempDir).Run()
	if err != nil {
		fmt.Printf("Error extracting tarball: %v\n", err)
		return
	}

	// Move extracted files to target directory
	_ = os.MkdirAll(targetDir, os.ModePerm)
	err = filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath, _ := filepath.Rel(tempDir, path)
			targetPath := filepath.Join(targetDir, relPath)
			_ = os.MkdirAll(filepath.Dir(targetPath), os.ModePerm)
			_ = os.Rename(path, targetPath)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error moving files: %v\n", err)
		return
	}

	// Clean up
	_ = os.RemoveAll(tempDir)
	_ = os.Remove(tarballName)

	fmt.Println("Swagger UI downloaded and unpacked successfully.")
}

func extractTarballURL(body string) string {
	for _, line := range strings.Split(body, "\n") {
		if strings.Contains(line, "tarball_url") {
			parts := strings.Split(line, "\"")
			for _, part := range parts {
				if strings.HasPrefix(part, "http") {
					return part
				}
			}
		}
	}
	return ""
}
