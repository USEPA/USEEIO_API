## Running the Python test suite

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
