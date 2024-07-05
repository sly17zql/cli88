package convert

import (
	"fmt"
	"github.com/h2non/bimg"
	"os"
	"path/filepath"
	"strings"
)

func HeicTransferSimple(filePath string, outName string, format string) error {
	return transferSimple(filePath, outName, format)
}

func TransferMultiply(dirPath string, originalName string, outNamePrefix string, format string) error {
	fileNames, err := ListFiles(dirPath)
	if err != nil {
		return err
	}
	index := 1
	for _, name := range fileNames {
		originalFileName := filepath.Base(name)
		if strings.HasPrefix(originalFileName, originalName) && IsExceptType(originalFileName, "heic") {
			curFilePath := filepath.Join(dirPath, originalFileName)
			outName := fmt.Sprintf("%s%d", outNamePrefix, index)
			index += 1
			err := transferSimple(curFilePath, outName, format)
			if err != nil {
				fmt.Printf("%s file transfer error: %s", outName, err)
			}
		}
	}
	return nil
}

func transferSimple(filePath string, outName string, format string) error {
	inputDir := filepath.Dir(filePath)

	buffer, err := bimg.Read(filePath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	var newImage []byte
	switch strings.ToLower(format) {
	case "jpg":
		{
			newImage, err = bimg.NewImage(buffer).Convert(bimg.JPEG)
		}
	case "png":
		{
			newImage, err = bimg.NewImage(buffer).Convert(bimg.PNG)
		}
	default:
		return fmt.Errorf("not cover this type")
	}

	if err != nil {
		return fmt.Errorf("failed to convert convert: %v", err)
	}

	newFileName := fmt.Sprintf("%s.%s", outName, format)
	newFilePath := filepath.Join(inputDir, newFileName)
	err = bimg.Write(newFilePath, newImage)
	fmt.Printf("new file %s \n", newFileName)
	if err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			fmt.Println(err)
		}
	}(filePath)

	return nil
}
