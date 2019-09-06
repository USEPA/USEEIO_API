import logging as log
import unittest

from config import getclient


class DemandTest(unittest.TestCase):

    def test_get_demands(self):
        client = getclient()
        for model in client.get_models():
            model_id = model['id']
            demands = client.get_demands(model_id)
            self.assertTrue(len(demands) > 0)
            log.info('test %i demands in model %s', len(demands), model_id)
            for demand in demands:
                self.assertIsNotNone(demand.get('id'))
                self.assertIsNotNone(demand.get('year'))
                self.assertIsNotNone(demand.get('type'))
                self.assertIsNotNone(demand.get('system'))
                self.assertIsNotNone(demand.get('location'))

    def test_get_demand(self):
        """Test that each demand vector is a non-empty list of sector
           ID-amount-pairs with valid sector IDs."""
        client = getclient()
        for model in client.get_models():
            model_id = model['id']
            sectors = client.get_sectors(model_id)
            sector_ids = set([s['id'] for s in sectors])
            for info in client.get_demands(model_id):
                demand = client.get_demand(model_id, info['id'])
                log.debug('test demand %s in model %s', info['id'], model_id)
                self.assertTrue(len(demand) > 0)
                for entry in demand:
                    sector_id = entry.get('sector')
                    self.assertIsNotNone(sector_id)
                    self.assertTrue(sector_id in sector_ids)
                    self.assertTrue(isinstance(
                        entry.get('amount'), (int, float)))
