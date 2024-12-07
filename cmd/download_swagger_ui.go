package cmd

import (
	"fmt"
	"github.com/ocrosby/go-reference-api/internal/utils"
	"github.com/spf13/cobra"
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

	if utils.FileExists(tarballName) {
		if err := utils.DeleteFile(tarballName); err != nil {
			fmt.Printf("Error removing existing tarball: %v\n", err)
			return
		}
	}

	if err := utils.CreateDirectory(tempDir); err != nil {
		fmt.Printf("Error creating temporary directory: %v\n", err)
		return
	}

	if err := utils.DownloadFile(swaggerUIURL, tarballName); err != nil {
		fmt.Printf("Error downloading Swagger UI: %v\n", err)
		return
	}

	// Check if the tarball file exists
	if !utils.FileExists(tarballName) {
		fmt.Printf("Tarball file does not exist: %v")
		return
	}

	// Extract the tarball
	cmd := exec.Command("tar", "-xzf", tarballName, "--strip-components=1", "-C", tempDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error extracting tarball: %v\nOutput: %s\n", err, string(output))
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
