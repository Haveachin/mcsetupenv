package main

import (
	"archive/zip"
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	urlPtr := flag.String("forgeurl", "", "The URL to the Zip-fiel of the forge version (mdk)")
	fileNamePtr := flag.String("filename", "temp.zip", "The name of the file that is downlaoded")
	deleteFilePtr := flag.Bool("delfile", true, "If the file should be deleted or not")

	flag.Parse()

	fileName := *fileNamePtr

	file, err := downloadFromURL(*urlPtr, fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	files, err := unzip(fileName, ".")

	if err != nil {
		log.Fatal(err)
	}

	if *deleteFilePtr {
		err = os.Remove("./" + fileName)

		if err != nil {
			log.Println(err)
		}
	}

	log.Println("Unzipped:\n" + strings.Join(files, "\n"))

	log.Println("Setting up workspace. This may take a while...")
	err = cmdExec("gradlew", "setupDecompWorkspace")

	if err != nil {
		log.Println(err)
	}

	log.Println("Setting up eclipse files. Nearly done...")
	err = cmdExec("gradlew", "eclipse")

	if err != nil {
		log.Println(err)
	}

	log.Println("Done!\nHappy coding! :)")
}

func downloadFromURL(url string, fileName string) (*os.File, error) {
	fmt.Println("Downloading from " + url)

	// Creating file
	file, err := os.Create(fileName)
	ok := os.IsExist(err)

	if !ok {
		_ = os.Remove(fileName)
		file, err = os.Create(fileName)
	}

	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Downlaoding
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	// Copying
	num, err := io.Copy(file, response.Body)

	if err != nil {
		return nil, err
	}

	fmt.Println("Download complete:\n", num, " bytes downloaded")

	return file, nil
}

// Code by Edd Turtle from https://golangcode.com/unzip-files-in-go/
// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		if strings.HasPrefix(f.Name, "src") {
			continue
		}

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {

			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return filenames, err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return filenames, err
			}

		}
	}
	return filenames, nil
}

//Code inspired Browny Lin from https://stackoverflow.com/questions/37091316/
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
