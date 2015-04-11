package main

import (
	"encoding/json"
	"errors"
	_ "github.com/arbovm/levenshtein"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func parseInvariantFileName(fileName string) (name string, extension string) {
	fileName = strings.ToLower(fileName)

	extension = filepath.Ext(fileName)
	name = fileName[0 : len(fileName)-len(extension)]

	if len(extension) > 0 {
		extension = extension[1:]
	}

	return
}

func isLicenseFile(fileName string) bool {
	name, extension := parseInvariantFileName(fileName)
	return (name == "license" || name == "copying") && (extension == "" || extension == "txt" || extension == "md")
}

func isNodeJsPackageFile(fileName string) bool {
	return strings.ToLower(fileName) == "package.json"
}

type NodeJsPackageFile struct {
	License *string `json:"license"`
}

func getNodeJsPackageLicense(fileName string) (licenseName string, err error) {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		return
	}

	packageFile := &NodeJsPackageFile{}

	if err = json.Unmarshal(content, &packageFile); err != nil {
		return
	}

	if packageFile.License == nil || len(*packageFile.License) == 0 {
		err = errors.New("License not set.")
		return
	}

	licenseName = strings.ToLower(*packageFile.License)

	return
}

func walkFolderWithLicense(directoryPath string, license *LicenseNode) (*LicenseNode, error) {
	files, err := ioutil.ReadDir(directoryPath)

	if err != nil {
		return nil, err
	}

	var hasFoundLicense bool
	var directories []os.FileInfo

	for _, file := range files {
		if file.IsDir() {
			directories = append(directories, file)
		} else if !hasFoundLicense {
			var newLicense *LicenseNode

			if isNodeJsPackageFile(file.Name()) {
				if licenseName, err := getNodeJsPackageLicense(directoryPath + file.Name()); err == nil {
					newLicense = NewLicenseNode(licenseName, directoryPath+file.Name())
				}
			} else if isLicenseFile(file.Name()) {
				/*if content, err := ioutil.ReadFile(directoryPath + file.Name()); err == nil {
					newLicense = NewLicenseNode(string(content), directoryPath+file.Name())
				}*/
			}

			if newLicense != nil {
				hasFoundLicense = true

				if license != nil {
					license.Licenses = append(license.Licenses, newLicense)
				}

				license = newLicense
			}
		}
	}

	for _, directory := range directories {
		if _, err = walkFolderWithLicense(directoryPath+directory.Name()+"/", license); err != nil {
			return nil, err
		}
	}

	return license, nil
}
