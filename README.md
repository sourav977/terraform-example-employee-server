# terraform-example-employee-server
Its a simple golang server which stores employee record in mongodb.

## Build locally
Run the following command to create application binary locally

```shell
$ make install
```

## Create docker image
Run the following command to create docker image

```shell
$ make image
```

## Run locally
Run the following command to run locally

```shell
$ docker compose --file=docker-compose-up.yaml up
```

## Sample Request-Response
GET: http://localhost:8000/api/getAllEmployees
- Response:
```shell
[
    {
        "_id": "61d70dc9b3624419697383de",
        "empID": "so9091",
        "fullName": "sourav patnaik",
        "company": {
            "companyName": "IBM",
            "companyAddress": "Hyderabad"
        }
    },
    {
        "_id": "61d7112026f3263c35d4cf68",
        "empID": "so9093",
        "fullName": "Ammy Zen",
        "company": {
            "companyName": "Redhat",
            "companyAddress": "Pune"
        }
    }
]
```

POST: http://localhost:8000/api/addEmployee
- Request Body:
```shell
{
    "empID": "so9091",
    "fullName": "sourav patnaik",
    "company": {
        "companyName": "IBM",
        "companyAddress": "Hyderabad"
    }
}
```
- Response:
```shell
[
    {
        "_id": "61d70dc9b3624419697383de",
        "empID": "so9091",
        "fullName": "sourav patnaik",
        "company": {
            "companyName": "IBM",
            "companyAddress": "Hyderabad"
        }
    },
    {
        "_id": "61d7112026f3263c35d4cf68",
        "empID": "so9093",
        "fullName": "Ammy Zen",
        "company": {
            "companyName": "Redhat",
            "companyAddress": "Pune"
        }
    }
]
```

PUT: http://localhost:8000/api/updateEmployeeByID
- Request Body:
```shell
{
    "empID": "so9091",
    "fullName": "sourav ranjan patnaik",
    "company": {
        "companyName": "IBM",
        "companyAddress": "Austin"
    }
}
```
- Response
```shell
{
    "MatchedCount": 1,
    "ModifiedCount": 1,
    "UpsertedCount": 0,
    "UpsertedID": null
}
```

DELETE: http://localhost:8000/api/DeleteEmployeeByID?empID=so9092
- Response:
```shell
{
    "DeletedCount": 0
}
```
