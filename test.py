import json
import requests

# with open("itemseverything.json") as f:
#     items = json.load(f)
headers = {'Accept': 'application/json'}
f = requests.get("https://raw.githubusercontent.com/ao-data/ao-bin-dumps/master/items.json", headers=headers)
items = json.loads(f.text)
# print(f.text)

for i in items["items"]:
    for j in items["items"][i]:
            try:
                if j["@uniquename"] == "T4_2H_LONGBOW":
                    print(j["@tier"])
            except:
                pass
