import logging as log
import unittest

import requests
from .config import endpoint, getmodels


class IndicatorTest(unittest.TestCase):

    def test_get_indicators(self):
        for model in getmodels():
            url = '%s/%s/indicators' % (endpoint(), model)
            log.info('test GET ' + url)
            with requests.get(url) as r:
                indicators = r.json()
                self.assertTrue(len(indicators) > 0)
                for indicator in indicators:  # type: dict
                    self.assertIsNotNone(indicator.get('id'))
                    self.assertIsNotNone(indicator.get('index'))
                    self.assertIsNotNone(indicator.get('name'))
                    self.assertIsNotNone(indicator.get('code'))
                    self.assertIsNotNone(indicator.get('unit'))
                    self.assertIsNotNone(indicator.get('group'))
                log.info('checked %i indicators', len(indicators))
