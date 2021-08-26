package data

import "os"

//pathExists returns whether the given file or directory exists
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func initializeFolder(folderPath string) error {
	ex, err := pathExists(folderPath)
	if err != nil {
		return err
	}

	if !ex {
		err = os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

//InitializeFolders verify that all the folders in the list
//exists and create those who are missing
func InitializeFolders(folderPaths []string) error {
	for _, s := range folderPaths {
		err := initializeFolder(s)
		if err != nil {
			return err
		}
	}
	return nil
}
