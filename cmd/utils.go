package cmd

import (
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"github.com/djrmarques/gmbe/extract"

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

// Returns an error if the file does not exist
func ValidateFileExists(filePath string) error {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		// Path is a directory, not a file
		return fmt.Errorf("%s is a directory", filePath)
	}

	// File exists
	return nil
}

// Given a file path, return the file name and the base folder
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

// Writes a string to a file
func WriteStringToFile(filePath string, content string) error {
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	fmt.Println("Content written to file:", filePath)
	return nil
}

// Extracts all the blocks from a given file into various files
func extractToFile(f, o string, j bool) error {
	// Check if a file exists

	if f == "" {
		return fmt.Errorf("Please specify a file using the -f flag. ")
	}

	if err := ValidateFileExists(f); err != nil {
		return err	
	}

	// Get blocks from files
	blocks, err := extract.ExtractBlocksFromFile(f, j)
	if err != nil {
		return err	
	}

	// Builds the path to the output file
	baseDir, file_name := GetFileBaseFolderAndName(f)

	// If t
	if o != "" {
		baseDir = o
	}

	outputFolder := filepath.Join(baseDir, file_name)
	if err = CreateFolderIfNotExist(outputFolder); err != nil {
		return err	
	}

	// Save each block in the respective file
	// Keeps track of how many blocks of each type
	// are in a file
	counterLanguesBlock := make(map[string]uint)
	for _, b := range blocks {

		if val, ok := counterLanguesBlock[b.T]; ok {
			counterLanguesBlock[b.T] = val + 1
		} else {
			counterLanguesBlock[b.T] = 1
		}

		blockTypeN := counterLanguesBlock[b.T]

		blockExtention := LanguageExtensions[b.T]
		outputFileName := b.T + "_" + strconv.FormatUint(uint64(blockTypeN), 10) + blockExtention
		outputFilePath := filepath.Join(outputFolder, outputFileName)
		CreateEmptyFile(outputFilePath)

		if err = WriteStringToFile(outputFilePath, b.Content); err != nil {
			return err
		}

	}

	return nil
}
