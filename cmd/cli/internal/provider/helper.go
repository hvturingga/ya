package provider

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func search(src, name string) (string, error) {
	var r string

	match := func(fn string) bool {
		exp := []string{
			name,
			fmt.Sprintf("%s.exe", name),
			fmt.Sprintf("%s-%s-%s.exe", name, runtime.GOOS, runtime.GOARCH),
			fmt.Sprintf("%s-%s-%s", name, runtime.GOOS, runtime.GOARCH),
		}

		for _, i := range exp {
			if fn == i {
				return true
			}
		}
		return false
	}

	err := filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || r != "" {
			return nil
		}

		if match(filepath.Base(path)) {
			r = path
			return io.EOF
		}
		return nil
	})

	if err != nil && err != io.EOF {
		return "", err
	}
	return r, nil
}

func mv(src, dest string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()

	if _, err := io.Copy(df, f); err != nil {
		return err
	}

	return nil
}

type progressReader struct {
	reader       io.Reader
	total        int64
	current      int64
	lastProgress int
}

func (pr *progressReader) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	pr.current += int64(n)

	progress := int(float64(pr.current) / float64(pr.total) * 100)
	if progress != pr.lastProgress {
		progressBar := "[" + strings.Repeat("=", progress/2) + ">" + strings.Repeat(" ", 50-progress/2) + "]"
		fmt.Printf("\r%s %d%%", progressBar, progress)
		pr.lastProgress = progress
		if progress == 100 {
			fmt.Println()
		}
	}

	return n, err
}
