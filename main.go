package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
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
	json.Unmarshal(vBytes, &vUrl)

	version := vUrl.Url[len(vUrl.Url)-6:]

	out, err := os.Create("cppm-exec")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	os := runtime.GOOS
	switch os {
	case "windows":
		resp, err := http.Get("https://github.com/Maou-Shimazu/Cpp-Project-Manager/releases/download/" + version + "/cppm-win-x64.exe")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		io.Copy(out, resp.Body)
	case "linux":
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
