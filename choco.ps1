Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; Invoke-Expression ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

. $PROFILE

choco install msys2

. $PROFILE

$cppm = "$HOME/.cppm/bin"
mkdir $cppm

$p = [Environment]::GetEnvironmentVariable("PATH", [EnvironmentVariableTarget]::Machine);
[Environment]::SetEnvironmentVariable("PATH", $p + ";" + $cppm, [EnvironmentVariableTarget]::Machine);