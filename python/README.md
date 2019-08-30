# USEEIO API - Python Version
This folder contains the Python implementation of the USEEIO API.

## Usage

```batch
rem (note that this a Windows Batch example; `rem` means comment)

rem open a command line and switch to the python folder
cd USEEIO_API\python

rem optionally, create a virtual environment and activate it
rem to not mess up your dependencies
python -m venv env
env\Scripts\activate.bat

rem install the dependencies
pip install -r requirements.txt

rem install the useeioapi module
pip install -e .

rem start the server
python -m useeioapi -data ..\build\data -port 9999
```

## Running the test suite

```batch
cd tests
python -m unittest discover
```

Use another endpoint (default is `http://localhost:8080/api`):

```batch
set USEEIO_API=http://another.end.point/api
python -m unittest discover
```

**todo** configure an API key
