"""
Tests {model}/indicators and {model}/indicators/{indicatorID} endpoints
"""
import logging as log
import random
import unittest

from client import getclient


class IndicatorTest(unittest.TestCase):

    def test_get_indicators(self):
        client = getclient()
        for model in client.get_models():
            model_id = model['id']
            indicators = client.get_indicators(model_id)
            self.assertTrue(len(indicators) > 0)
            log.info('test %i indicators in model %s',
                     len(indicators), model_id)
            for indicator in indicators:  # type: dict
                self.assertIsNotNone(indicator.get('id'))
                self.assertIsNotNone(indicator.get('index'))
                self.assertIsNotNone(indicator.get('name'))
                self.assertIsNotNone(indicator.get('code'))
                self.assertIsNotNone(indicator.get('unit'))
                self.assertIsNotNone(indicator.get('group'))
                self.assertIsNotNone(indicator.get('simpleunit'))
                self.assertIsNotNone(indicator.get('simplename'))

    def test_get_indicator(self):
        client = getclient()
        for model in client.get_models():
            model_id = model['id']
            indicators = client.get_indicators(model_id)
            i = random.randint(0, len(indicators) - 1)
            log.info('check indicator %i in model %s', i, model_id)
            indicator = indicators[i]  # type: dict
            same = client.get_indicator(model_id, indicator['id'])
            self.assertEqual(same['id'], indicator['id'])
            self.assertEqual(same['index'], indicator['index'])
            self.assertEqual(same['name'], indicator['name'])
            self.assertEqual(same['code'], indicator['code'])
