"""
Tests /models endpoint
"""
import logging as log
import unittest

from client import getclient


class ModelTest(unittest.TestCase):

    def test_get_models(self):
        client = getclient()
        models = client.get_models()
        self.assertTrue(len(models) > 0)
        log.info('test %i models', len(models))
        for model in models:  # type: dict
            self.assertIsNotNone(model.get('id'))
            self.assertIsNotNone(model.get('name'))
            self.assertIsNotNone(model.get('location'))
