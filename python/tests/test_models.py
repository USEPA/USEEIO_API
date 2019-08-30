import logging as log
import unittest

import requests
from config import endpoint


class ModelTest(unittest.TestCase):

    def test_get_models(self):
        url = endpoint() + '/models'
        log.info('test GET ' + url)
        with requests.get(url) as r:
            models = r.json()
            self.assertTrue(len(models) > 0)
            for model in models:  # type: dict
                self.assertIsNotNone(model.get('id'))
                self.assertIsNotNone(model.get('name'))
                self.assertIsNotNone(model.get('location'))
                log.info('found model %s', model['name'])
