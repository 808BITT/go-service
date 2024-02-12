@echo off

REM check if arg1 and arg2 are empty
if "%1"=="" (
    echo "arg1 is empty. Please provide the install path"
    exit /b
)

if "%2"=="" (
    echo "arg2 is empty. Please provide the service name"
    exit /b
)

REM get arg1 and arg2 for install path and service name
set arg1=%1
set arg2=%2

REM create the bin directory
mkdir bin

REM clean the bin directory
del /s /q bin

REM build the main.go file
go build -o bin/main.exe main.go

REM create the installer
echo @echo off > bin/install.bat
echo mkdir %arg1% >> bin/install.bat
echo icacls %arg1% /grant Everyone:(OI)(CI)F >> bin/install.bat
echo copy main.exe %arg1% >> bin/install.bat
echo sc create %arg2% binPath= "%arg1%\main.exe %arg2%" start= auto >> bin/install.bat

REM create the uninstaller
echo @echo off > bin/uninstall.bat 
echo sc stop %arg2% >> bin/uninstall.bat
echo sc delete %arg2% >> bin/uninstall.bat
echo rmdir /s /q %arg1% >> bin/uninstall.bat

REM create the starter
echo @echo off > bin/start.bat
echo sc start %arg2% >> bin/start.bat

REM create the stopper
echo @echo off > bin/stop.bat
echo sc stop %arg2% >> bin/stop.bat

REM create the restarter
echo @echo off > bin/restart.bat
echo sc stop %arg2% >> bin/restart.bat
echo sc start %arg2% >> bin/restart.bat

