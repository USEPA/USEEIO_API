"""
Tests {model}/sectors/ and {model}/sectors/{sectorID} endpoints
"""
import logging as log
import random
import unittest

from client import Client


class SectorTest(unittest.TestCase):

    def test_get_sectors(self):
        client = Client.get()
        for model in client.get_models():
            model_id = model['id']
            sectors = client.get_sectors(model_id)
            self.assertTrue(len(sectors) > 0)
            log.info('test %i sectors in model %s', len(sectors), model_id)
            for s in sectors:  # type: dict
                self.assertIsNotNone(s.get('id'))
                self.assertIsNotNone(s.get('index'))
                self.assertIsNotNone(s.get('name'))
                self.assertIsNotNone(s.get('code'))
                self.assertIsNotNone(s.get('location'))
                self.assertIsNotNone(s.get('category'))

    def test_get_sector(self):
        client = Client.get()
        for model in client.get_models():
            model_id = model['id']
            sectors = client.get_sectors(model_id)
            i = random.randint(0, len(sectors) - 1)
            log.info('check sector %i in model %s', i, model_id)
            sector = sectors[i]  # type: dict
            same = client.get_sector(model_id, sector['id'])
            self.assertEqual(same['id'], sector['id'])
            self.assertEqual(same['index'], sector['index'])
            self.assertEqual(same['name'], sector['name'])
            self.assertEqual(same['code'], sector['code'])
            self.assertEqual(same['location'], sector['location'])
            self.assertEqual(same['category'], sector['category'])
            self.assertEqual(same.get('description'), sector.get('description'))
