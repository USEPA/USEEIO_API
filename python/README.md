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

You can configure another endpoint (default is `http://localhost:8080/api`)
via the `USEEIO_API` environment variable:

```batch
set USEEIO_API=http://another.end.point/api
python -m unittest discover
```

Also, you can set an API key via the `USEEIO_API_KEY` environment variable.
This key will then added under the `x-api-key` field to the header of the API
requests.
