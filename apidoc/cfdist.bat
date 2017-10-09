@echo off

if not exist ".\build" (
    echo ERROR: The build folder does not exist.
    echo You should run the Gulp build first.
    goto end
)

if not exist "..\build\data" (
    echo ERROR: The folder backend\build\data does not exist.
    echo Put the data that should be deployed with the server in this folder.
    goto end
)

echo Prepare the application for Cloud Foundry

echo .
echo 1. Delete the old distribution folder
if exist ".\cfdist" rmdir /s /q .\cfdist
mkdir cfdist

echo .
echo 2. Compile the back end for Linux
set CGO_ENABLED=0
set GOARCH=amd64
set GOOS=linux
cd ..\src
go build -o ..\apidoc\cfdist\app
cd ..\apidoc

echo .
echo 3. Copy static resources
robocopy build cfdist\static /mir
robocopy ..\build\data cfdist\data /mir
copy /Y cfdist.yaml cfdist\manifest.yaml
copy /Y cfproc cfdist\Procfile
:end

echo . all done
echo You should now be able to push the app to a Cloud Foundry instance
echo 1. Switch to the cfdist folder: cd cfdist
echo 2. Login, e.g.: cf login -a api.run.pivotal.io
echo 3. Push the app with the binary build pack:
echo 4. cf push -b https://github.com/cloudfoundry/binary-buildpack.git
  