package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

type JSON struct {
	Url string `json:"html_url"`
}

func main() {
	var compilers []bool
	gpp, err := exec.Command("g++", "-v").Output()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("G++ installed", gpp)
	}
	clang, err := exec.Command("clang++", "-v").Output()
	if err != nil {
		fmt.Println("Clang++ is not installed")
		compilers = append(compilers, false)
	} else {
		fmt.Println("clang++ is installed", clang)
	}
	//choco := "Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))"
	//if runtime.GOOS == "windows" {
	//	powershell, err := exec.Command("powershell", choco).Output()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//} else {
	//
	//}
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
		out, err := os.Create("cppm.exe")
		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.Get("https://github.com/Maou-Shimazu/Cpp-Project-Manager/releases/download/" + version + "/cppm-win-x64.exe")
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(out, resp.Body)
	case "linux":
		out, err := os.Create("cppm")
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
