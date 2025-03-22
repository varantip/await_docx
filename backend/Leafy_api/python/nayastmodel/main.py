from torch.utils.checkpoint import checkpoint
#uvicorn main:app --reload --host 0.0.0.0 --port 8000
import model
import planttype
import shutil
import getworkingplants
from fastapi import FastAPI, File, UploadFile

app = FastAPI()
@app.post("/disease")
async def disease(file: UploadFile):
    res,err = model.check(file.file)
    return {"Err": False, "result": res}

@app.post("/planttype")
async def plant_type(file: UploadFile):
    res = planttype.IndetifyPlant(file.file)
    if res not in getworkingplants.plant_types:
        return {"Err": True, "result": "Не удалось определить растение"}
    return {"Err": False, "result": res}