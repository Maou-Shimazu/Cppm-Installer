package main

import (
	"encoding/json"
	"fmt"
	// "fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	// "os/user"
	"runtime"
	// "path/filepath"
)

type JSON struct {
	Url string `json:"html_url"`
}

func main() {

	vResp, err := http.Get("https://api.github.com/repos/maou-shimazu/cpp-project-manager/releases/latest")

	if err != nil {
		log.Fatal(err)
	}
	defer vResp.Body.Close()

	vBytes, err := ioutil.ReadAll(vResp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var vUrl JSON
	err = json.Unmarshal(vBytes, &vUrl)
	version := vUrl.Url[len(vUrl.Url)-6:]

	//user_os :=
	switch runtime.GOOS {
	case "windows":

		home, _ := os.UserHomeDir()
		out, err := os.Create(home + "/.cppm/bin/cppm.exe")
		if err != nil {
			fmt.Println("failed to create file")
			log.Fatal(err)
		}
		fmt.Println("Created cppm")
		defer out.Close()
		resp, err := http.Get("https://github.com/Maou-Shimazu/Cpp-Project-Manager/releases/download/" + version + "/cppm-win-x64.exe")
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(out, resp.Body)
	case "linux":
		home, _ := os.UserHomeDir()
		out, err := os.Create(home + "/.cppm/bin/cppm")
		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.Get("https://github.com/Maou-Shimazu/Cpp-Project-Manager/releases/download/" + version + "/cppm-debian")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		io.Copy(out, resp.Body)
	default:
		log.Fatal("Unsupported OS")
	}
}
