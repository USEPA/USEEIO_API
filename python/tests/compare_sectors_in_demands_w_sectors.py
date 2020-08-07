"""
This is a useful test to run manually if test_demands fails because one or more
sectors in demand doesn't match the
Assumes server is running locally
"""
import useeioapi.client as apiclient
import pandas as pd

client = apiclient.Client(endpoint='http://localhost:8080/api')

for model in client.get_models():
    model_id = model['id']
    sectors = client.get_sectors(model_id)
    sector_ids = set([s['id'] for s in sectors])
    for info in client.get_demands(model_id):
        demand = client.get_demand(model_id, info['id'])
        demand = pd.DataFrame(demand)
        demand_sectors = set(demand['sector'])
        missing_sectors = demand_sectors-sector_ids
        if len(missing_sectors)>0:
            print(missing_sectors)


