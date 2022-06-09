Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; Invoke-Expression ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

choco install msys2

$cppm = "$HOME\.cppm\bin"
if (!(Test-Path -Path $cppm)) { # if cppm path dosent exist
    mkdir $cppm
    
    $p = [Environment]::GetEnvironmentVariable("PATH", [EnvironmentVariableTarget]::Machine);
    [Environment]::SetEnvironmentVariable("PATH", $p + ";" + $cppm, [EnvironmentVariableTarget]::Machine);
}

C:\tools\msys64\msys2.exe bash msys2.sh