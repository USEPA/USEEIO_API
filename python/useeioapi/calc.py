import useeioapi.data as data
import numpy


def calculate(folder: str, model_id: str, demand: dict) -> dict:
    indicators = data.read_indicators(folder, model_id)
    sectors = data.read_sectors(folder, model_id)
    d = _demand_vector(demand, sectors)

    U = data.read_matrix(folder, model_id, "U")
    R = None
    perspective = demand.get('perspective')
    if perspective == 'direct':
        L = data.read_matrix(folder, model_id, "L")
        D = data.read_matrix(folder, model_id, "D")
        s = _scaled_column_sums(L, d)
        R = _scale_columns(D, s)
    if perspective == 'intermediate':
        L = data.read_matrix(folder, model_id, "L")
        s = _scaled_column_sums(L, d)
        R = _scale_columns(U, s)
    if perspective == 'final':
        R = _scale_columns(U, d)

    return {
        'data': R.tolist(),
        'totals': _scaled_column_sums(U, d).tolist(),
        'indicators': [ind.get('id') for ind in indicators],
        'sectors': [sec.get('id') for sec in sectors],
    }


def _scale_columns(matrix: numpy.ndarray, v: numpy.ndarray) -> numpy.ndarray:
    result = numpy.zeros(matrix.shape, dtype=numpy.float64)
    for i in range(0, v.shape[0]):
        s = v[i]
        if s == 0:
            continue
        result[:, i] = s * matrix[:, i]
    return result


def _scaled_column_sums(M: numpy.ndarray, s: numpy.ndarray) -> numpy.ndarray:
    m, n = M.shape
    v = numpy.zeros(m, dtype=numpy.float64)
    for j in range(0, n):
        factor = s[j]
        if factor == 0:
            continue
        v += factor * M[:, j]
    return v


def _demand_vector(demand: dict, sectors: list) -> numpy.ndarray:
    idx = {}
    for sector in sectors:
        idx[sector['id']] = sector['index']
    v = numpy.zeros(len(sectors))
    entries = demand.get('demand')
    if entries is None:
        return v
    for e in entries:
        i = idx.get(e.get('sector'))
        if i is None:
            continue  # TODO: throw an error
        v[i] = float(e.get('amount'))
    return v
