### Method

POST

### Endpoint

http://localhost:8000/v1/list-races

### Payload

```json
{
    "filter": {
        "visible_only": true,
        "order_by": "asc"
    }
}
```

```json
{
    "filter": {
        "visible_only": true,
        "order_by": "desc"
    }
}
```

### Sample Response

```json
{
    "id": "3",
    "meetingId": "8",
    "name": "Rhode Island ghosts",
    "number": "3",
    "visible": false,
    "advertisedStartTime": "2021-03-02T08:30:16Z",
    "status": "CLOSED"
},
```
