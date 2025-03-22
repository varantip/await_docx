apikey = 'TBx0AqWdnUcl29j38kXJuj2DBKgkt7aCZNRs4Yk0ct9W8kvUYK'
from kindwise import PlantApi
from PIL import Image

api = PlantApi(f'{apikey}')
def IndetifyPlant(file):
    identification = api.identify(Image.open(file).convert('RGB'), details=['url', 'common_names'])
    for suggestion in identification.result.classification.suggestions:
        return suggestion.name
    return ''