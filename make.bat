@echo off

set GOARCH=amd64
set GOOS=windows

echo Compile backend ...
if not exist ".\build" mkdir ".\build"
cd src
go build -o ..\build\app.exe
cd ..
echo ... all done
