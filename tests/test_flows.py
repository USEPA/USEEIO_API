"""
Tests {model}/flows and {model}/flows/{flowID} endpoints
"""

import logging as log
import random
import unittest

from client import Client


class FlowTest(unittest.TestCase):

    def test_get_flows(self):
        client = Client.get()
        for model in client.get_models():
            model_id = model['id']
            flows = client.get_flows(model_id)
            self.assertTrue(len(flows) > 0)
            log.info('test %i flows in model %s', len(flows), model_id)
            for f in flows:  # type: dict
                self.assertIsNotNone(f.get('id'))
                self.assertIsNotNone(f.get('index'))
                self.assertIsNotNone(f.get('flowable'))
                self.assertIsNotNone(f.get('context'))
                self.assertIsNotNone(f.get('unit'))
                self.assertIsNotNone(f.get('uuid'))

    def test_get_flow(self):
        client = Client.get()
        for model in client.get_models():
            model_id = model['id']
            flows = client.get_flows(model_id)
            i = random.randint(0, len(flows) - 1)
            log.info('check flow %i in model %s', i, model_id)
            flow = flows[i]  # type: dict
            same = client.get_flow(model_id, flow['id'])
            self.assertEqual(same['id'], flow['id'])
            self.assertEqual(same['index'], flow['index'])
            self.assertEqual(same['flowable'], flow['flowable'])
            self.assertEqual(same['context'], flow['context'])
            self.assertEqual(same['unit'], flow['unit'])
            self.assertEqual(same['uuid'], flow['uuid'])
