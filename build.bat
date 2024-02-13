@echo off

REM Reading the service name
for /f "tokens=2 delims=:" %%a in ('findstr /c:"Service-Name" service.config.json') do set "service_name=%%a"
set "service_name=%service_name:~2,-2%"

REM Reading the install path
for /f "tokens=1* delims=:" %%a in ('findstr /c:"Install-Path" service.config.json') do set "install_path=%%b"
set "install_path=%install_path:~2,-1%"

REM create the bin directory
mkdir bin

REM clean the bin directory
del /s /q bin

REM build the main.go file
go build -o bin/main.exe main.go

REM create the installer
echo @echo off > bin/install.bat
echo mkdir "%install_path%" >> bin/install.bat
echo icacls "%install_path%" /grant Everyone:(OI)(CI)F >> bin/install.bat
echo copy main.exe "%install_path%" >> bin/install.bat
echo cd .. >> bin/install.bat
echo copy service.config.json "%install_path%" >> bin/install.bat
echo sc create "%service_name%" binPath= "\"%install_path%/main.exe\" \"%service_name%\" \"%install_path%/service.config.json\"" start= auto >> bin/install.bat
echo cd bin >> bin/install.bat

REM create the uninstaller
echo @echo off > bin/uninstall.bat 
echo sc stop "%service_name%" >> bin/uninstall.bat
echo sc delete "%service_name%" >> bin/uninstall.bat
echo rmdir /s /q "%install_path%" >> bin/uninstall.bat

REM create the starter
echo @echo off > bin/run.bat
echo sc start "%service_name%" >> bin/run.bat

REM create the stopper
echo @echo off > bin/stop.bat
echo sc stop "%service_name%" >> bin/stop.bat

REM create the restarter
echo @echo off > bin/restart.bat
echo sc stop "%service_name%" >> bin/restart.bat
echo sc start "%service_name%" >> bin/restart.bat

