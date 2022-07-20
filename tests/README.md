## Running the Python test suite

The `tests` folder contains a Python test suite that can be executed
against an USEEIO API endpoint. In order to run the tests, you first
need to make sure that the required Python modules for running the
tests are installed:

```bash
cd tests
pip install -r requirements.txt
```

If this is the case, you can run the test suite with the following
command:

```bash
cd tests
python -m unittest discover
```

By default, the test suite tries to run the tests against a local
instance of the USEEIO API running at `http://localhost:8080/api`.
This can be changed by setting another endpoint via the
`USEEIO_API` environment variable:

```batch
set USEEIO_API=http://another.end.point/api
python -m unittest discover
```

Also, an API key can be set via the `USEEIO_API_KEY` environment variable.

```bash
set USEEIO_API=http://another.end.point/api
set USEEIO_API_KEY=your_private_api_key_that_you_would_never_share
python -m unittest discover
```
