package internal

import (
	"fmt"
	"github.com/google/go-github/v59/github"
	"github.com/hvturingga/ya/conf"
	"github.com/hvturingga/ya/internal/unzip"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// progressReader wraps an io.Reader to report the progress of reading.
type progressReader struct {
	reader       io.Reader
	total        int64
	current      int64
	lastProgress int
}

// Read reads data into p, updates the current progress, and prints the progress bar.
func (pr *progressReader) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	pr.current += int64(n)
	pr.printProgress()
	return n, err
}

// printProgress prints the current progress as a progress bar.
func (pr *progressReader) printProgress() {
	progress := int(float64(pr.current) / float64(pr.total) * 100)
	if progress != pr.lastProgress {
		fmt.Printf("\r%s %d%%", progressBar(progress), progress)
		pr.lastProgress = progress
		if progress == 100 {
			fmt.Println()
		}
	}
}

// progressBar generates a string representation of the progress bar.
func progressBar(progress int) string {
	return "[" + strings.Repeat("=", progress/2) + ">" + strings.Repeat(" ", 50-progress/2) + "]"
}

// createTempDir creates a temporary directory and returns its path.
func createTempDir() (string, error) {
	tempDir := "./.T"
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create temp directory: %w", err)
	}
	return tempDir, nil
}

// downloadFile downloads a file from the specified URL to the specified directory and returns the path to the downloaded file.
func downloadFile(url, dir, name string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("server returned %d status", resp.StatusCode)
	}

	filePath := filepath.Join(dir, name)
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	p := &progressReader{reader: resp.Body, total: resp.ContentLength}
	if _, err := io.Copy(file, p); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return filePath, nil
}

// decompressIfNeeded checks the file extension and decompresses the file if necessary.
func decompressIfNeeded(filePath string) error {
	decompressors := map[string]func(string, string) error{
		".zip": unzip.Unzip,
		".gz":  unzip.Untargz,
	}

	ext := filepath.Ext(filePath)
	if decompress, ok := decompressors[ext]; ok {
		dir := filepath.Dir(filePath)
		if err := decompress(filePath, dir); err != nil {
			return fmt.Errorf("failed to decompress file: %w", err)
		}
	}
	return nil
}

// findFirstExecutable searches for the first executable file that matches the specified name in the given directory.
func findFirstExecutable(src, name string) (string, error) {
	var resultPath string
	err := filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || resultPath != "" {
			return err
		}
		if isMatch(filepath.Base(path), name) {
			resultPath = path
			return io.EOF // Use io.EOF to signal successful early termination
		}
		return nil
	})

	if err != nil && err != io.EOF {
		return "", err
	}
	return resultPath, nil
}

// isMatch checks if the filename matches any of the expected executable names for the given platform.
func isMatch(filename, name string) bool {
	expectedNames := []string{
		name,
		fmt.Sprintf("%s.exe", name),
		fmt.Sprintf("%s-%s-%s.exe", name, runtime.GOOS, runtime.GOARCH),
		fmt.Sprintf("%s-%s-%s", name, runtime.GOOS, runtime.GOARCH),
	}

	for _, expected := range expectedNames {
		if filename == expected {
			return true
		}
	}
	return false
}

// moveFileToFinalDestination moves the file to the final destination directory.
func moveFileToFinalDestination(filePath, name, version string) (string, error) {
	destDir := filepath.Join(conf.GetProviderPath(), name, version)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create destination directory: %w", err)
	}

	destPath := filepath.Join(destDir, filepath.Base(filePath))
	if err := os.Rename(filePath, destPath); err != nil {
		return "", fmt.Errorf("failed to move the file: %w", err)
	}

	return destPath, nil
}

// DownloadProviderRelease handles the entire process of downloading, decompressing, and moving a provider release.
func DownloadProviderRelease(providerName, providerVersion string, asset *github.ReleaseAsset) (string, error) {
	url := asset.GetBrowserDownloadURL()

	tempDir, err := createTempDir()
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tempDir)

	filePath, err := downloadFile(url, tempDir, asset.GetName())
	if err != nil {
		return "", err
	}

	if err := decompressIfNeeded(filePath); err != nil {
		return "", err
	}

	executable, err := findFirstExecutable(tempDir, providerName)
	if err != nil {
		return "", err
	}
	destPath, err := moveFileToFinalDestination(executable, providerName, providerVersion)
	if err != nil {
		return "", err
	}

	return destPath, nil
}
