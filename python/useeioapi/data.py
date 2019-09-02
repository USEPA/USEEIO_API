import csv
import os

import useeioapi.matio as matio
import numpy


class Sector(object):

    def __init__(self):
        self.id = ''
        self.index = 0
        self.name = ''
        self.code = ''
        self.location = ''
        self.description = ''

    def as_json_dict(self):
        return {
            'id': self.id,
            'index': self.index,
            'name': self.name,
            'code': self.code,
            'location': self.location,
            'description': self.description
        }


class Indicator(object):

    def __init__(self):
        self.id = ''
        self.index = 0
        self.group = ''
        self.code = ''
        self.unit = ''
        self.name = ''

    def as_json_dict(self):
        return {
            'id': self.id,
            'index': self.index,
            'group': self.group,
            'code': self.code,
            'unit': self.unit,
            'name': self.name
        }


class Model(object):

    def __init__(self, folder: str, name: str):
        self.folder = folder  # type: str
        self.sectors = read_sectors(folder, name)  # type: dict
        sorted_sectors = [s for s in self.sectors.values()]
        sorted_sectors.sort(key=lambda s: s.index)
        self.sector_ids = [s.id for s in sorted_sectors]
        self.indicators = read_indicators(folder)  # type: list
        self.indicators.sort(key=lambda i: i.index)
        self.indicator_ids = [i.id for i in self.indicators]
        self.matrix_cache = {}

    def get_matrix(self, name: str):
        m = self.matrix_cache.get(name)
        if m is not None:
            return m
        path = '%s/%s.bin' % (self.folder, name)
        if not os.path.isfile(path):
            return None
        m = matio.read_matrix(path)
        self.matrix_cache[name] = m
        return m

    def get_dqi_matrix(self, name: str):
        m = self.matrix_cache.get(name)
        if m is not None:
            return m
        path = '%s/%s.csv' % (self.folder, name)
        if not os.path.isfile(path):
            return None
        with open(path, 'r', encoding='utf-8', newline='\n') as f:
            reader = csv.reader(f)
            rows = []
            for row in reader:
                rows.append(row)
        self.matrix_cache[name] = m
        return m

    def calculate(self, demand):
        if demand is None:
            return
        perspective = demand.get('perspective')
        d = self.demand_vector(demand)
        data = None
        if perspective == 'direct':
            s = self.scaling_vector(d)
            D = self.get_matrix('D')
            data = scale_columns(D, s)
        elif perspective == 'intermediate':
            s = self.scaling_vector(d)
            U = self.get_matrix('U')
            data = scale_columns(U, s)
        elif perspective == 'final':
            U = self.get_matrix('U')
            data = scale_columns(U, d)
        else:
            print('ERROR: unknown perspective %s' % perspective)

        if data is None:
            print('ERROR: no data')
            return None

        result = {
            'indicators': self.indicator_ids,
            'sectors': self.sector_ids,
            'data': data.tolist()
        }
        return result

    def demand_vector(self, demand):
        L = self.get_matrix('L')
        d = numpy.zeros(L.shape[0], dtype=numpy.float64)
        entries = demand.get('demand')  # type: dict
        if entries is None:
            return d
        for e in entries:
            sector_key = e.get('sector')
            amount = e.get('amount')
            if sector_key is None or amount is None:
                continue
            amount = float(amount)
            sector = self.sectors.get(sector_key)
            if sector is None:
                continue
            d[sector.index] = amount
        return d

    def scaling_vector(self, demand_vector: numpy.ndarray) -> numpy.ndarray:
        s = numpy.zeros(demand_vector.shape[0], dtype=numpy.float64)
        L = self.get_matrix('L')
        for i in range(0, demand_vector.shape[0]):
            d = demand_vector[i]
            if d == 0:
                continue
            col = L[:, i]
            s += d * col
        return s


def read_sectors(folder: str, model_id: str) -> list:
    path = folder + '/' + model_id + '/sectors.csv'
    sectors = []
    for row in read_csv(path):
        sectors.append({
            'index': int(row[0]),
            'id': row[1],
            'name': row[2],
            'code': row[3],
            'location': row[4],
            'description': row[5],
        })
    sectors.sort(key=lambda s: s['index'])
    return sectors


def read_flows(folder: str, model_id: str) -> list:
    path = folder + '/' + model_id + '/flows.csv'
    flows = []
    for row in read_csv(path):
        flows.append({
            'index': int(row[0]),
            'id': row[1],
            'name': row[2],
            'category': row[3],
            'subCategory': row[4],
            'unit': row[5],
            'uuid': row[6],
        })
    return flows


def read_indicators(folder: str):
    indicators = []
    path = '%s/indicators.csv' % folder
    with open(path, 'r', encoding='utf-8', newline='\n') as f:
        reader = csv.reader(f)
        next(reader)
        for row in reader:
            i = Indicator()
            i.index = int(row[0])
            i.id = row[3]
            i.name = row[2]
            i.code = row[3]
            i.unit = row[4]
            i.group = row[5]
            indicators.append(i)
    return indicators


def scale_columns(matrix: numpy.ndarray, v: numpy.ndarray) -> numpy.ndarray:
    result = numpy.zeros(matrix.shape, dtype=numpy.float64)
    for i in range(0, v.shape[0]):
        s = v[i]
        if s == 0:
            continue
        result[:, i] = s * matrix[:, i]
    return result


def read_model_infos(data_folder: str):
    infos = []
    for row in read_csv(data_folder + '/models.csv'):
        infos.append({
            'id': row[0],
            'name': row[1],
            'location': row[2],
            'description': row[3]
        })
    return infos


def read_demand_infos(data_folder: str, model_id: str):
    path = data_folder + '/' + model_id + '/demands.csv'
    infos = []
    for row in read_csv(path):
        infos.append({
            'id': row[0],
            'year': int(row[1]),
            'type': row[2],
            'system': row[3],
            'location': row[4],
        })
    return infos


def read_csv(path, skip_header=True) -> list:
    with open(path, 'r', encoding='utf-8', newline='\n') as f:
        r = csv.reader(f)
        if skip_header:
            next(r)
        for row in r:
            yield row
