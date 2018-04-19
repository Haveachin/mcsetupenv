package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Parse all the arguments
	forgeFileURL := flag.String("ff-url", "", "The URL to the Zip-fiel of the forge version (mdk)")
	forgePathDownload := flag.String("fp-dl", "temp.zip", "The path of the forge file that is downlaoded")
	forgePathInstall := flag.String("fp-extract", ".", "The path where the forge file should extract")
	forgeFileDelete := flag.Bool("ff-del", true, "If the file should be deleted or not")
	flag.Parse()

	if err := getForgeMDK(forgeFileURL, forgePathDownload, forgePathInstall, forgeFileDelete); err != nil {
		log.Println("Error:", err)
	}

	if err := setupWorkspace(forgePathInstall); err != nil {
		log.Println("Error:", err)
	}

	log.Println("Done!\nHappy coding! :)")
}

func getForgeMDK(downloadURL, downlaodPath, installPath *string, deleteAfterInstall *bool) *error {
	// Download the forge file
	file, err := downloadFromURL(downloadURL, downlaodPath)
	if err != nil {
		log.Fatal("Fatal Error: ", err)
	}
	defer file.Close()

	// Unzipping the forge file
	files, err := unzip(downlaodPath, installPath, "src/", ".gitignore")
	log.Println("Unzipped:", files)
	if err != nil {
		log.Fatal("Fatal Error: ", err)
	}

	// Delete
	if *deleteAfterInstall {
		err = os.Remove(*downlaodPath)

		if err != nil {
			return &err
		}
	}

	return nil
}

func setupWorkspace(gradlewPath *string) *error {
	log.Println("Setting up workspace. This may take a while...")

	*gradlewPath = filepath.Join(*gradlewPath, "gradlew")

	log.Println(*gradlewPath)

	if err := cmdExec(*gradlewPath, "setupDecompWorkspace"); err != nil {
		return &err
	}

	log.Println("Setting up eclipse files. Nearly done...")

	if err := cmdExec(*gradlewPath, "eclipse"); err != nil {
		return &err
	}

	return nil
}
