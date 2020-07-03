@echo off

if "%~1" == "help" goto HELP
if "%~1" == "-h" goto HELP

if not exist "..\build" mkdir "..\build"
rem delete possible old build artifacts
if exist "..\build\app.exe" del "..\build\app.exe"
if exist "..\build\app" del "..\build\app"
if exist "..\build\manifest.yaml" del "..\build\manifest.yaml"
if exist "..\build\Procfile" del "..\build\Procfile"
if "%~1" == "clean" goto END

rem check the data & static folders
if not exist "..\build\data" (
    echo WARNING: The data folder build^\data does not exists
    echo ... You may want to put the API data there.
)
if not exist "..\build\static" (
    echo WARNING: The data folder build^\static does not exists
    echo ... You may want to put static pages there that should
    echo ... be hosted together with the API ^(e.g. by building
    echo ... the apidoc^: cd ..\apidoc ^& gulp^).
)

rem windows versions are build by default;
rem "~" ensures double quotes are stripped
if "%~1" == "" goto WINDOWS
if "%~1" == "Windows" goto WINDOWS
if "%~1" == "windows" goto WINDOWS
if "%~1" == "win" goto WINDOWS
if "%~1" == "win64" goto WINDOWS

rem Linux
if "%~1" == "Linux" goto LINUX
if "%~1" == "linux" goto LINUX
if "%~1" == "linux64" goto LINUX

rem Cloud Foundry
if "%~1" == "CF" goto CF
if "%~1" == "cf" goto CF

rem AWS
if "%~1" == "AWS" goto AWS
if "%~1" == "aws" goto AWS

rem unknown build target
echo ERROR: unknown build target: %1
echo Supported build targets are
echo * Windows ^(default^)
echo * Linux
echo * CF ^(creates a distribution folder for Cloud Foundry^)
goto END

:HELP
echo This script compiles the Go version for a specific platform
echo Usage: make.bat ^<platform^>
echo where the supported platforms are:
echo * Windows ^(default^)
echo * Linux
echo * CF ^(creates a distribution folder for Cloud Foundry^)
echo * AWS (Linux binary for distribution on AWS)
echo Note: you need to have a current Go compiler installed
goto END

:WINDOWS
set GOARCH=amd64
set GOOS=windows
echo Compile version for %GOOS%_%GOARCH%
cd useeioapi
go build -o ..\..\build\app.exe
cd ..
echo done
goto END

:LINUX
set CGO_ENABLED=0
set GOARCH=amd64
set GOOS=linux
echo Compile version for %GOOS%_%GOARCH%
cd useeioapi
go build -o ..\..\build\app
cd ..
echo done
goto END

:CF
set CGO_ENABLED=0
set GOARCH=amd64
set GOOS=linux
echo Compile version for %GOOS%_%GOARCH%
cd useeioapi
go build -o ..\..\build\app
cd ..
echo Copy CF meta data
copy /Y cfdist\manifest.yaml ..\build\manifest.yaml
copy /Y cfdist\Procfile ..\build\Procfile
echo You should now be able to push the app to a Cloud Foundry instance
echo 1. Switch to the build folder: cd ..^\build
echo 2. Login, e.g.: cf login -a api.run.pivotal.io
echo 3. Push the app with the binary build pack:
echo 4. cf push -b https://github.com/cloudfoundry/binary-buildpack.git
goto END

:AWS
set CGO_ENABLED=0
set GOARCH=amd64
set GOOS=linux
echo Compile version for AWS: %GOOS%_%GOARCH%
cd useeioapi
go build -o ..\awsdist\app
cd ..
echo done
goto END

:END
