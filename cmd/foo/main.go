package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/afero"
)

func main() {
	if err := core(); err != nil {
		log.Fatal(err)
	}
}

func core() error {
	fs := afero.NewOsFs()

	if err := Copy(fs, "foo.exe", "bar.exe"); err != nil {
		return err
	}

	// if err := fs.Rename("foo", "foo.bak"); err != nil {
	// 	return err
	// }
	return nil
}

const (
	executableFilePermission os.FileMode = 0o755
)

func Copy(fs afero.Fs, dest, src string) error {
	dst, err := fs.OpenFile(dest, os.O_RDWR|os.O_CREATE|os.O_TRUNC, executableFilePermission) //nolint:nosnakecase
	if err != nil {
		return fmt.Errorf("create a file: %w", err)
	}
	defer dst.Close()
	srcFile, err := fs.Open(src)
	if err != nil {
		return fmt.Errorf("open a file: %w", err)
	}
	defer srcFile.Close()
	fmt.Println("copying")
	if _, err := io.Copy(dst, srcFile); err != nil {
		return fmt.Errorf("copy a file: %w", err)
	}

	return nil
}
