import logging as log
import os
import sys

import requests
import useeioapi.client

_endpoint = None
_client = None
log.basicConfig(level=log.INFO, stream=sys.stdout)


def getclient() -> useeioapi.client.Client:
    global _client
    if _client is not None:
        return _client
    _client = useeioapi.client.Client(endpoint())
    return _client


def endpoint() -> str:
    """Get the API endpoint for the test suite. This can be configured be
       setting an environment variable `USEEIO_API` with the endpoint. When
       this variable is not set, the endpoint defaults to
       `http://localhost:8080/api`"""
    global _endpoint
    if _endpoint is not None:
        return _endpoint
    _endpoint = os.environ.get('USEEIO_API')
    if _endpoint is not None:
        return _endpoint
    _endpoint = 'http://localhost:8080/api'
    return _endpoint


def getmodels() -> list:
    """Returns the IDs of the models of the configured endpoint."""
    url = endpoint() + '/models'
    with requests.get(url) as r:
        models = r.json()
        return [model['id'] for model in models]
