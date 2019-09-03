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
