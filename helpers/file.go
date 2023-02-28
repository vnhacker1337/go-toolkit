package filehelpers

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ListFileInFolder(foldername string) []string {
	files := []string{}

	err := filepath.Walk(foldername,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// fmt.Println(path, info.Size())
			files = append(files, path)
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return files
}

func GetDirFileInFolder(foldername string, filename string) []string {

	var files = []string{}

	foldername = GetHomeFolder(foldername)

	err := filepath.Walk(foldername,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// fmt.Println(path, info.Size())
			if strings.HasSuffix(path, filename) {
				files = append(files, path)
			}

			return nil
		})

	if err != nil {
		log.Println(err)
	}

	return files
}

func GetHomeFolder(foldername string) string {

	abs_foldername := foldername

	if strings.HasPrefix(foldername, "~") {
		dirname, err := os.UserHomeDir()

		if err == nil {
			abs_foldername = dirname + strings.TrimLeft(foldername, "~")
			return abs_foldername
		}
	}

	return abs_foldername
}

func ListFilesWithPattern(foldername string, pattern string) []string {
	var files = []string{}

	foldername = GetHomeFolder(foldername)

	err := filepath.Walk(foldername,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// fmt.Println(path, info.Size())
			if strings.Contains(path, pattern) {
				files = append(files, path)
			}

			return nil
		})

	if err != nil {
		log.Println(err)
	}

	return files
}

func LoadFile(filename string) (lines []string) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close() //nolint
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func GetLengthFile(filename string) int {

	lines := LoadFile(filename)
	return len(lines)
}
