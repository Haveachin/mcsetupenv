package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func downloadFromURL(url, filePath *string) (*os.File, error) {
	log.Println("Downloading from", *url)

	*filePath = rootPath(*filePath)

	if !strings.HasSuffix(*filePath, ".zip") {
		*filePath = bufferString(*filePath, ".zip")
	}

	// Downlaoding
	response, err := http.Get(*url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Creating file
	file, err := os.Create(*filePath)
	if !os.IsExist(err) {
		_ = os.Remove(*filePath)
		file, err = os.Create(*filePath)
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Copying
	num, err := io.Copy(file, response.Body)
	if err != nil {
		return nil, err
	}

	log.Println("Download complete:", num, "bytes downloaded")

	return file, nil
}

// Code inpired by Edd Turtle from https://golangcode.com/unzip-files-in-go/
func unzip(pathSource, pathDestination *string, pathExlude ...string) ([]string, error) {
	log.Println("Unzipping", *pathSource)

	*pathDestination = rootPath(*pathDestination)

	var filenames []string

	readerZip, err := zip.OpenReader(*pathSource)
	if err != nil {
		return filenames, err
	}
	defer readerZip.Close()

	for _, file := range readerZip.File {
		reader, err := file.Open()
		if err != nil {
			return filenames, err
		}
		defer reader.Close()

		exclude := false

		// Exclude path
		for _, path := range pathExlude {
			if strings.HasPrefix(file.Name, path) {
				exclude = true
			}
		}

		if exclude {
			continue
		}

		// Store filename/path for returning and using later on
		filePath := filepath.Join(*pathDestination, file.Name)
		filenames = append(filenames, filePath)

		if file.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(filePath, os.ModePerm)
		} else {
			// Make File
			if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				return filenames, err
			}

			outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return filenames, err
			}
			_, err = io.Copy(outFile, reader)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return filenames, err
			}
		}
	}

	return filenames, nil
}

// Code inspired by Browny Lin from https://stackoverflow.com/questions/37091316/
func cmdExec(command string, args string) error {
	cmd := exec.Command(command, strings.Split(args, " ")...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}

func bufferString(strs ...string) string {
	var buffer bytes.Buffer

	for _, str := range strs {
		buffer.WriteString(str)
	}

	return buffer.String()
}

func rootPath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = bufferString("/", path)
	}

	if !strings.HasPrefix(path, ".") {
		path = bufferString(".", path)
	}

	return path
}
