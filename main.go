package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/inhies/go-bytesize"
	"github.com/schollz/progressbar"
)

func calcDirSize(dirPath string, outputName string) (int64, error) {
	var totalSize int64
	fileCount := 0
	var stats string = "Directory Stats \n"
	var lastPath string
	fmt.Print("Calculating directory size... ")
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileCount++
		}

		return nil
	})

	bar := progressbar.New(fileCount)

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && (lastPath != path && !strings.Contains(path, ".git") && !strings.Contains(path, "node_modules") && !strings.Contains(path, "env")) {
			lastPath = path
			stats += "\nDirectory: " + path + "\n"
		}

		if !info.IsDir() && (lastPath != path && !strings.Contains(path, ".git") && !strings.Contains(path, "node_modules") && !strings.Contains(path, "env")) {
			totalSize += info.Size()
			stats += "Path:" + path + "\nSize: " + bytesize.New(float64(info.Size())).String() + "\n"
			bar.Add(1)
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	if outputName == "" {
		fmt.Print("\nEnter the output file name: ")
		fmt.Scanln(&outputName)
	}

	err = writeToFile(outputName, stats)
	if err != nil {
		return 0, err
	}
	return totalSize, nil
}

func writeToFile(path string, data string) error {
	file, err := os.Create(path + ".txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var dirPath string
	fmt.Print("Enter directory path: ")
	fmt.Scanln(&dirPath)

	size, err := calcDirSize(dirPath, "")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Total size of '%s': %s\n", dirPath, bytesize.New(float64(size)).String())
}
