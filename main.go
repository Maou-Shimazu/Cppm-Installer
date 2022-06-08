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
	compilers := []bool{false, false}
	gpp, err := exec.Command("g++", "-v").Output()
	if err != nil {
		fmt.Println("g++ is not installed")
	} else {
		fmt.Println("G++ installed", gpp)
		compilers[0] = true
	}
	clang, err := exec.Command("clang++", "-v").Output()
	if err != nil {
		fmt.Println("Clang++ is not installed")
	} else {
		fmt.Println("clang++ is installed", clang)
		compilers[1] = true
	}
	if !(compilers[0] || compilers[1]) {
		if runtime.GOOS == "windows" {
		fmt.Println("Compiler not detected, installing clang...")

		choco := "Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))"
		
			_, err := exec.Command(choco).Output()
			if err != nil {
				log.Fatal(err)
			}
			_, err = exec.Command(". $PROFILE").Output()
			
			_, err = exec.Command("choco", "install msys2").Output()
		} else {
			fmt.Println("It seems you are using a unix system, please use your local package manager to install clang and clang++.")
		}
	} else {
		fmt.Println("Compiler installed. Proceeding with installation")
	}

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
