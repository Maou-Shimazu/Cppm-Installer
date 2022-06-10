# Cppm-Installer

Installer for the [C++ Project Manager](https://github.com/maou-shimazu/cpp-project-manager)

## Installing Msys2 along with the installer
##### This only applies if you do not have c++ installed via clang++/g++ [windows only]
----
Install the zip from the [release](https://github.com/Maou-Shimazu/Cppm-Installer/releases/latest) page and extract the folder. \
Open an admin powershell in the current folder and run the following commands: 

```
Set-ExecutionPolicy Bypass -Scope Process -Force
``` 
This allows the powershell script to run. 

```
.\choco.ps1
``` 
This will install choco along with msys2 and configure it for you. There will be 1 prompt, pass Y for yes.\
It will take a few minutes for everything to finish installing, if you get any prompts choose the value of acceptance. After that is all done start up the installer: 
```
.\Cppm-Installer.exe
```
 This installs the latest version of cppm and configures cppm paths. \
There will be better user messages in the future, you can always run the installer again to update. 

### Note: 
When you install cppm, if there are no commands working, it is possible that you lack Visual C++ Redistributable Packages. \
You can install them from [here](https://www.microsoft.com/en-gb/download/details.aspx?id=48145). If you dont think that is the problem, run cppm by double clicking, it is located in the "C:\User\\[username]\\.cppm\bin" folder.


## Unix Users
For those on unix please install cargo and install from crates so you can have a proper build for your system. \
Install Cargo from [here](https://doc.rust-lang.org/cargo/getting-started/installation.html)

Then run: `cargo install cppm` 

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).
