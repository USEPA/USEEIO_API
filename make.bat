@echo off

if not exist ".\build" mkdir ".\build"

set GOARCH=amd64
set GOOS=windows

echo Compile version for %GOOS%_%GOARCH%

cd go\useeioapi
go build -o ..\build\app.exe


set CGO_ENABLED=0
set GOARCH=amd64
set GOOS=linux

echo Compile version for %GOOS%_%GOARCH%

go build -o ..\build\app

cd ..\..
echo ... all done
