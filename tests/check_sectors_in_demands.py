"""
This is a useful test to run manually if test_demands fails because one or more
sectors in a demand are not present in the respective model.
"""

from client import Client


def main():
    client = Client.get()

    for model in client.get_models():
        model_id = model['id']
        print(f'check demand sectors in model {model_id}')
        sectors = client.get_sectors(model_id)
        sector_ids = set([s['id'] for s in sectors])
        for info in client.get_demands(model_id):
            demand_id = info['id']
            print(f'  check demand vector {demand_id}')
            demand = client.get_demand(model_id, demand_id)
            demand_sectors = set([d['sector'] for d in demand])
            missing = demand_sectors - sector_ids
            if len(missing) > 0:
                print(f'    {len(missing)} sectors are missing: {missing}')
            else:
                print('     all sectors present')


if __name__ == '__main__':
    main()
