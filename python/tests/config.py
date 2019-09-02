import logging as log
import os
import sys

import useeioapi.client

_client = None
log.basicConfig(level=log.INFO, stream=sys.stdout)


def getclient() -> useeioapi.client.Client:
    """Get the API client instance for the test suite. This can be configured by
       setting an environment variable `USEEIO_API` for the API endpoint which
       defaults to `http://localhost:8080/api`. Also, an additional API key
       can be configured in the same way via the `USEEIO_API_KEY` environment
       variable."""
    global _client
    if _client is not None:
        return _client
    endpoint = os.environ.get('USEEIO_API')
    if endpoint is None or endpoint == '':
        endpoint = 'http://localhost:8080/api'
    apikey = os.environ.get('USEEIO_API_KEY')
    if apikey == '':
        apikey = None
    _client = useeioapi.client.Client(endpoint, apikey)
    return _client
