# Cppm-Installer

Installer for the C++ Project Manager

## Installing Msys2 along with the installer
##### This only applies if you do not have c++ installed via clang++/g++ [windows only]
\
Install the zip from the [release]() page and extract the folder. \
Open an admin powershell in the current folder and run the following commands: \
`Set-ExecutionPolicy Bypass -Scope Process -Force` this allows the powershell script to run. \
`.\choco.ps1` this will install choco along with msys2 and configure it for you. There will be 1 prompt, pass Y for yes.\
It will take a few minutes for everything to finish installing, if you get any prompts choose the value of acceptance. After that is all done start up the installer: \
`.\Cppm-Installer.exe` This installs the latest version of cppm and configures cppm paths. \
There will be better user messages in the future, you can always run the installer again to update. 

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).
