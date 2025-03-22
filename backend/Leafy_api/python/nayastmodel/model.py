## pip install tranformers
## pip install onnxruntime
## pip install torch
## pip install Pillow
import onnxruntime as ort
from transformers import ViTImageProcessor
from PIL import Image
import json
import requests
## объявление всех путей
processor_path = "model\\plant-disease-model\\preprocessor_config.json"
model_path = "model\\plant-disease-model\\plant-disease.onnx"
labels_path = "model\\plant-disease-model\\labels.json"

## загрузка словаря с лейблами
with open(labels_path, "r") as f:
    labels = json.load(f)
    


def check(file):
    ## открываем картинку
    image = Image.open(file,'r')

    ## processor - штука для обработки картинки перед подачей ее модели
    processor = ViTImageProcessor.from_pretrained(processor_path)
    inputs = processor(images=image.convert('RGB'), return_tensors="pt")
    ## загружаем модель
    ort_session = ort.InferenceSession(model_path)
    ## объявляем что будем ей подавать
    ort_inputs = {ort_session.get_inputs()[0].name: inputs["pixel_values"].numpy()}
    ## здесь происходит работа модели (классификация изображения)
    ort_outputs = ort_session.run(None, ort_inputs)

    ## выбираем лучший результат и ищем значение лейбла, чтоб слово (а не цифру) выдавало
    predicted = labels[str(ort_outputs[0].argmax(axis=-1)[0])]
    ## выводим ответ
    return predicted
