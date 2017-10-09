import csv
import json
import os

CSV_FILE = 'USEEIOv1.1_FinalDemand.csv'
MODEL = 'USEEIO'

COL_CODE = 0
COL_NAME = 1
COL_LOCATION = 2
COL_FIRST_DEMAND = 3


class DemandInfo(object):

    def __init__(self, column_header: str):
        parts = column_header.split('_')
        self.location = parts[0]
        self.year = int(parts[1])
        self.perspective = parts[2].lower()
        self.system = parts[3] if len(parts) > 3 else 'complete'
        self.id = ('%s_%s_%s' % (
            self.year, self.location, self.perspective)).lower()
        if self.system != 'complete':
            self.id += '_' + self.system.lower()

    def to_json_obj(self) -> dict:
        obj = {
            'id': self.id,
            'year': self.year,
            'location': self.location,
            'system': self.system,
            'type': self.perspective,
            'isDefault': False
        }
        if self.perspective == 'consumption':
            obj['consumer'] = 'all'
        return obj


def main():
    prepare_folders()
    rows = read_rows()
    infos = []
    for col in range(COL_FIRST_DEMAND, len(rows[0])):
        header = rows[0][col]
        info = DemandInfo(header)
        infos.append(info.to_json_obj())
        entries = read_entries(rows, col)
        dump_json(entries, os.path.join(MODEL, 'demands', info.id + '.json'))
    dump_json(infos, os.path.join(MODEL, 'demands.json'))


def prepare_folders():
    folder = os.path.join(MODEL, 'demands')
    if not os.path.exists(folder):
        os.makedirs(folder)


def read_rows():
    rows = []
    with open(CSV_FILE, 'r', encoding='utf-8', newline='\n') as f:
        reader = csv.reader(f)
        for row in reader:
            rows.append(row)
    return rows


def read_entries(rows, col_idx):
    entries = []
    for row_idx in range(1, len(rows)):
        amount = float(rows[row_idx][col_idx])
        if amount == 0.0:
            continue
        sector_id = get_sector_id(row_idx, rows)
        entries.append({'sector': sector_id, 'amount': amount})
    return entries


def get_sector_id(row_idx: int, rows: list) -> str:
    row = rows[row_idx]
    parts = [row[COL_CODE], row[COL_NAME], row[COL_LOCATION]]
    sector_id = '/'.join(x.strip().lower() for x in parts)
    return sector_id


def dump_json(obj, file_name: str):
    with open(file_name, 'w', encoding='utf-8', newline='\n') as f:
        json.dump(obj, f, indent='  ')


if __name__ == '__main__':
    main()
