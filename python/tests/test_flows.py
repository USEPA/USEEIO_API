import logging as log
import unittest

import requests
from config import endpoint, getmodels
from urllib.parse import quote


class FlowTest(unittest.TestCase):

    def test_get_flows(self):
        for model in getmodels():
            url = '%s/%s/flows' % (endpoint(), model)
            log.info('test GET ' + url)
            with requests.get(url) as r:
                flows = r.json()
                self.assertTrue(len(flows) > 0)
                for f in flows:  # type: dict
                    self.assertIsNotNone(f.get('id'))
                    self.assertIsNotNone(f.get('index'))
                    self.assertIsNotNone(f.get('name'))
                    self.assertIsNotNone(f.get('category'))
                    self.assertIsNotNone(f.get('subCategory'))
                    self.assertIsNotNone(f.get('unit'))
                    self.assertIsNotNone(f.get('uuid'))
                log.info('checked %i flows', len(flows))

    def test_get_flow(self):
        for model in getmodels():
            base = '%s/%s/flows' % (endpoint(), model)
            with requests.get(base) as r:
                flow = r.json()[0]
                fid = quote(flow['id'])
                url = '%s/%s/flows/%s' % (
                    endpoint(), model, fid)
                log.info('test GET ' + url)
                with requests.get(url) as ri:
                    same = ri.json()
                    self.assertEqual(same['id'], flow['id'])
                    self.assertEqual(same['index'], flow['index'])
                    self.assertEqual(same['name'], flow['name'])
                    self.assertEqual(same['category'], flow['category'])
                    self.assertEqual(same['unit'], flow['unit'])
                    self.assertEqual(same['uuid'], flow['uuid'])
