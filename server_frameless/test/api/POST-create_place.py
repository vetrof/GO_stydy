import requests

body = {
    "lat": "51.153705",
    "lng": "71.414391",
    "name": "test 44"
}

response = requests.post('http://127.0.0.1:8080/create_place', json=body)

print(response.status_code)
print(response.content)



                else:
                    # if product no weight
                    res['unit_price'] = res['price']
                    res['old_unit_price'] = res['old_price']