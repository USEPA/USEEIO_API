import logging as log
import unittest

import requests
from config import endpoint, getmodels
from urllib.parse import quote


class SectorTest(unittest.TestCase):

    def test_get_sectors(self):
        for model in getmodels():
            url = '%s/%s/sectors' % (endpoint(), model)
            log.info('test GET ' + url)
            with requests.get(url) as r:
                sectors = r.json()
                self.assertTrue(len(sectors) > 0)
                for s in sectors:  # type: dict
                    self.assertIsNotNone(s.get('id'))
                    self.assertIsNotNone(s.get('index'))
                    self.assertIsNotNone(s.get('name'))
                    self.assertIsNotNone(s.get('code'))
                    self.assertIsNotNone(s.get('location'))
                log.info('checked %i sectors', len(sectors))

    def test_get_sector(self):
        for model in getmodels():
            base = '%s/%s/sectors' % (endpoint(), model)
            with requests.get(base) as r:
                sector = r.json()[0]
                sid = quote(sector['id'])
                url = '%s/%s/sectors/%s' % (
                    endpoint(), model, sid)
                log.info('test GET ' + url)
                with requests.get(url) as ri:
                    same = ri.json()
                    self.assertEqual(same['id'], sector['id'])
                    self.assertEqual(same['index'], sector['index'])
                    self.assertEqual(same['name'], sector['name'])
                    self.assertEqual(same['code'], sector['code'])
                    self.assertEqual(same['location'], sector['location'])
