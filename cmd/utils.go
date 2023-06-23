package cmd

import (
	"os"
	"fmt"
	"path/filepath"
	"log"
	"io/ioutil"
)

var LanguageExtensions = map[string]string{
	"go":             ".go",
	"python":         ".py",
	"java":           ".java",
	"javaScript":     ".js",
	"c++":            ".cpp",
	"c#":             ".cs",
	"ruby":           ".rb",
	"rust":           ".rs",
	"swift":          ".swift",
	"kotlin":         ".kt",
	"typeScript":     ".ts",
	"php":            ".php",
	"shell":          ".sh",
	"perl":           ".pl",
	"scala":          ".scala",
	"r":              ".r",
	"matlab":         ".m",
	"groovy":         ".groovy",
}

func CreateFolderIfNotExist(folderPath string) error {
	// Check if the folder exists
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		// Folder does not exist, create it
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create folder: %s", err)
		}
		return nil
	} else if err != nil {
		// Error occurred while checking folder existence
		return fmt.Errorf("error occurred while checking folder: %s", err)
	} else {
		// Folder already exists
		return nil
	}
}


func ValidateFileExists(filePath string) (bool, error) {
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// File does not exist
		return false, nil
	} else if err != nil {
		// Error occurred while checking file existence
		return false, fmt.Errorf("error occurred while checking file: %s", err)
	}

	if fileInfo.IsDir() {
		// Path is a directory, not a file
		return false, nil
	}

	// File exists
	return true, nil
}

func GetFileBaseFolderAndName(filePath string) (string, string) {
	baseFolder := filepath.Dir(filePath)
	fileName := filepath.Base(filePath)
	fileNameWithoutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	return baseFolder, fileNameWithoutExt
}


func CreateEmptyFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()
	return nil
}

func FailIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WriteStringToFile(filePath string, content string) error {
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	fmt.Println("Content written to file:", filePath)
	return nil
}
