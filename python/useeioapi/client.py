import logging as log
import requests
import numpy as np

from urllib.parse import quote


class Client(object):

    def __init__(self, endpoint: str, apikey=None):
        self.endpoint = endpoint
        self.apikey = apikey

    def get_models(self) -> list:
        return self.__get_json('/models')

    def get_sectors(self, model_id: str) -> list:
        return self.__get_json('/%s/sectors' % model_id)

    def get_sector(self, model_id: str, sector_id: str):
        sid = quote(sector_id)
        return self.__get_json('/%s/sectors/%s' % (model_id, sid))

    def get_demands(self, model_id: str) -> list:
        return self.__get_json('/%s/demands' % model_id)

    def get_demand(self, model_id: str, demand_id: str):
        did = quote(demand_id)
        return self.__get_json('/%s/demands/%s' % (model_id, did))

    def get_indicators(self, model_id: str) -> list:
        return self.__get_json('/%s/indicators' % model_id)

    def get_indicator(self, model_id: str, indicator_id: str):
        iid = quote(indicator_id)
        return self.__get_json('/%s/indicators/%s' % (model_id, iid))

    def get_flows(self, model_id: str) -> list:
        return self.__get_json('/%s/flows' % model_id)

    def get_flow(self, model_id: str, flow_id: str):
        fid = quote(flow_id)
        return self.__get_json('/%s/flows/%s' % (model_id, flow_id))

    def get_matrix(self, model_id: str, name: str) -> np.ndarray:
        data = self.__get_json('/%s/matrix/%s' % (model_id, name))
        return np.asarray(data, dtype=np.float)

    def get_matrix_column(self, model_id: str, name: str, col: int) \
            -> np.ndarray:
        data = self.__get_json('/%s/matrix/%s?col=%i' % (model_id, name, col))
        return np.asarray(data, dtype=np.float)

    def get_matrix_row(self, model_id: str, name: str, row: int) \
            -> np.ndarray:
        data = self.__get_json('/%s/matrix/%s?row=%i' % (model_id, name, row))
        return np.asarray(data, dtype=np.float)

    def __get_json(self, path):
        url = self.endpoint + path
        log.debug("GET " + url)
        headers = {}
        if self.apikey is not None:
            headers['x-api-key'] = self.apikey
        with requests.get(url, headers=headers) as r:
            return r.json()
