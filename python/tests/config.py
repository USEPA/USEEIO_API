import logging as log
import os

import requests

_endpoint = None
log.basicConfig(level=log.INFO)


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
