import csv
import os
import struct

import numpy


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
    flows.sort(key=lambda s: s['index'])
    return flows


def read_indicators(folder: str, model_id: str):
    path = folder + '/' + model_id + '/indicators.csv'
    indicators = []
    for row in read_csv(path):
        indicators.append({
            'index': int(row[0]),
            'id': row[3],
            'name': row[2],
            'code': row[3],
            'unit': row[4],
            'group': row[5]
        })
    indicators.sort(key=lambda s: s['index'])
    return indicators


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


def read_matrix(data_folder: str, model_id: str, name: str):
    path = data_folder + '/' + model_id + '/' + name + '.bin'
    if not os.path.isfile(path):
        return None
    shape = _read_matrix_shape(path)
    return numpy.memmap(path, mode='c', dtype='<f8',
                        shape=shape, offset=8, order='F')


def read_dqi_matrix(data_folder: str, model_id: str, name: str):
    path = data_folder + '/' + model_id + '/' + name + '.csv'
    if not os.path.isfile(path):
        return None
    dqi = []
    for row in read_csv(path, skip_header=False):
        dqi.append(row)
    return dqi


def _read_matrix_shape(file_path: str):
    """ Reads and returns the shape (rows, columns) from the matrix stored in
        the given file.
    """
    with open(file_path, 'rb') as f:
        rows = struct.unpack('<i', f.read(4))[0]
        cols = struct.unpack('<i', f.read(4))[0]
        return rows, cols
