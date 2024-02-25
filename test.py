import json
import requests

# headers = {'Accept': 'application/json'}
# f = requests.get("https://raw.githubusercontent.com/ao-data/ao-bin-dumps/master/items.json", headers=headers)
# items = json.loads(f.text)

# for i in items["items"]:
#     for j in items["items"][i]:
#             try:
#                 if j["@uniquename"] == "T4_2H_LONGBOW":
#                     print(j["craftingrequirements"]["craftresource"])
#                     print(j["enchantments"])
#             except:
#                 pass


# with open("itemseverything.json") as f:
#     data = f.read()
#     print(data)

with open("world.json", "r",  encoding="utf-8") as f:
    data = json.loads(f.read())

# print(data[0]["UniqueName"])

items = []
for i in data:
    if i["Index"]:
        item = [i["Index"], i["UniqueName"]]
        items.append(item)

string = "{"
for i in items:
    string += "{" + f'"{i[0]}", "{i[1]}"' + "}, "
string = string[:-1] + "}"


with open("world.txt", "w") as f:
    f.write(string)
    
# print(data[0]["LocalizedNames"]["EN-US"])