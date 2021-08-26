package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/LeoMarche/blocast/pkg/data"
	"github.com/LeoMarche/blocast/pkg/exchange"
	"github.com/LeoMarche/blocast/pkg/key"
)

var KEY_FOLDER string = filepath.Join("data", "keys")
var INIT_FOLDERS = []string{KEY_FOLDER}
var KEY_SIZE int = 2048
var KEY_NAME string = "block-cast"

func main() {
	err := data.InitializeFolders(INIT_FOLDERS)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	privKey, pubKey, err := key.GenerateKeys(KEY_SIZE)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	basename := filepath.Join(KEY_FOLDER, KEY_NAME)
	err = key.StoreKeys(privKey, pubKey, basename+"-private.pem", basename+"-public.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	exchange.Exchange()
}
