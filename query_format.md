# Formato delle Query scritte in JSON

* Inserimento (POST) : <br>
```json
{
  "query_type": "insert",
  "destination": "myFile",
  "query_content": [
    "example1",
    2,
    "example2",
    "example3",
    34,
    45.6
  ]
}
```

* Letttura (GET) : <br>
```json
{
  "query_type": "select",
  "destination": "myFile",
  "query_content": {
    "field1": "value",
    "field2": "all_items",
    "filed3": 23
  }
}
```
* Aggiornamento (PATCH) : <br>
```json
{
  "query_type": "update",
  "destination": "myFile",
  "to_update": [
    "field1", "field2", "field3"
  ],
  "new_values": [
    "example1", "example2", "example3"
  ]
}
```
* Elmininazione (DELETE): <br>
```json
{
  "query_type": "delete",
  "destination": "myFile",
  "query_content": {
    "field1": "all_items",
    "field2": 56,
    "filed3": "example2"
  }
}
```
## Query Avanzate
* Filtraggio (GET) : <br>
```json
{
}
```
* Unione (GET) : <br>
```json
{
}
```
* Ordinamento (GET, PATCH, POST) : <br>
```json
{
}
```
