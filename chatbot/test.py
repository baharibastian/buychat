import json
import requests
req = requests.get('https://server-buychat.herokuapp.com/products')
response = json.loads(req.text)
print(response['data'])
for product in response['data']:
    print(product['Product_name'])