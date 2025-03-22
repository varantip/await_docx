import requests
import json
plant_types = []
def getworkingplants():
    content = requests.get("http://127.0.0.1:8080/v1/base_plants/")
    for plant_type in content.json()["Data"]:
        plant_types.append(plant_type["Bio_Name"])

getworkingplants()
