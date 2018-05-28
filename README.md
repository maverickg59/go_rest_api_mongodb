# REST API IN GOLANG WITH MONGODB, GORILLA MUX, AND MGO

## A hospital needs an API interface into their patient records database (MongoDB). Patient information to get you started is found in the MOCK_DATA.json file and is a json array containing information for two patients.

### JSON Data Model
```
{
  "id": 3,
  "first_name": "Robert",
  "last_name": "Manchaca",
  "gender": "Male",
  "phone": "555-555-5555",
  "email_address": "robert@thissite.com",
  "address": "42 Hamburger Place",
  "visit_date": "2017-12-25",
  "diagnosis": "H1318AX",
  "drug_code": "24920-0GE7",
  "additional_information": [
    {
    "notes": "Nam ultrices, libero non mattis pulvinar, nulla pede ullamcorper augue, a suscipit nulla elit ac nulla. Sed vel enim sit amet nunc viverra dapibus. Nulla suscipit ligula in lacus.\n\nCurabitur at ipsum ac tellus semper interdum. Mauris ullamcorper purus sit amet nulla. Quisque arcu libero, rutrum ac, lobortis vel, dapibus at, diam.",
    "new_patient": True,
    "race": "African American",
    "ssn": "555-55-5555"
    }
  ]
}
```

### POST:
```
curl -d '{"id":1,"first_name":"Gardanzo","last_name":"Plentenao","gender":"Male","phone_number":"555-555-5555","email":"thisperson0@blahblah.it","address":"9 Hot Springs Terrace","visit_date":"7/23/2017","diagnosis":"T34G","drug_code":"15060-0GT3","additional_information":[{"notes":"Aliquam quis turpis eget elit sodales scelerisque. Mauris sit amet eros. Suspendisse accumsan tortor quis turpis.\n\nSed ante. Vivamus tortor. Duis mattis egestas metus.\n\nAenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh.","new_patient":false,"race":"Hispanic","ssn":"555-55-5555"}]}' -H "Content-Type: application/json" -X POST http://localhost:8080/patients
```

### GET ALL:
```
curl http://localhost:8080/patients
```

### GET ONE BY ID:
```
curl http://localhost:8080/patients/1
```

### PUT (last name modified to test update):
```
curl -d '{"id":1,"first_name":"Gardanzo","last_name":"Plantenao","gender":"Male","phone_number":"555-555-5555","email":"thisperson0@blahblah.it","address":"9 Hot Springs Terrace","visit_date":"7/23/2017","diagnosis":"T34G","drug_code":"15060-0GT3","additional_information":[{"notes":"Aliquam quis turpis eget elit sodales scelerisque. Mauris sit amet eros. Suspendisse accumsan tortor quis turpis.\n\nSed ante. Vivamus tortor. Duis mattis egestas metus.\n\nAenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh.","new_patient":false,"race":"Hispanic","ssn":"555-55-5555"}]}' -H "Content-Type: application/json" -X PUT http://localhost:8080/patients/1
```

### GET BY ATTRIBUTE TERM:
```
curl -G -v "http://localhost:8080/search" --data-urlencode "first_name=Gardanzo"
curl -G -v "http://localhost:8080/search" --data-urlencode "gender=Male"
curl -G -v "http://localhost:8080/search" --data-urlencode "last_name=Plantenao"
```

### DELETE ONE BY ID:
```
curl -X DELETE http://localhost:8080/patients/1
```