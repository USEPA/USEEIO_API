import logging as log
import requests

from urllib.parse import quote


class Client(object):

    def __init__(self, endpoint: str):
        self.endpoint = endpoint

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

    def __get_json(self, path):
        url = self.endpoint + path
        log.debug("GET " + url)
        with requests.get(url) as r:
            return r.json()
