### Method

GET

### Endpoint

http://localhost:8000/v1/races/1

### Payload

There is no request body; the API configuration must not declare a body clause.

### Sample Response

```json
{
  "race": {
    "id": "1",
    "meetingId": "5",
    "name": "North Dakota foes",
    "number": "2",
    "visible": false,
    "advertisedStartTime": "2021-03-03T01:30:57Z",
    "status": "CLOSED"
  }
}
```
