# RenzoFS


![Diagramma senza titolo drawio](https://github.com/Elia-Renzoni/RenzoFS/assets/118525453/3cd2ba5c-f996-4379-98b5-5c254a446978)
<br>

RenzoFS is a distributed file system with a web based interface that allows users to see their remote file system and query their file, using json written query. <br>
<br>
The backend is written in Go (1.20) and uses RESTful API to interact with the web server, specially with the remote file system, witch is accessed by client only on remote access model, instead of upload-download access model. <br>
The main functionalities of the API are: <br>
|Functionality|
|-------------|
|Sign in user|
|Sign out user|
|Add user's friend|
|Create a remote directory|
|Create empty csv files|
|Query the files|
|Release file system usage statistics|
|Provide info about who, and when, queried files|
|File sharing|

<br>
Query Format: <br>

* Write (POST) : <br>
```json
{
  "query_type": "insert",
  "user_name":"elia",
  "destination": "myFile",
  "query_content": ["value1", "44", "value2", "55.5"]
}
```

* Read (GET) : <br>
```json
{
  "query_type": "select",
  "user_name":"elia",
  "destination": "myFile",
  "query_content": {
    "field1": "value",
    "field2": "all_items",
    "filed3": 23
  }
}
```
* Update (PATCH) : <br>
```json
{
  "query_type": "update",
  "user_name":"elia",
  "destination": "myFile",
  "query_content": {
    "ColumnName": ["id", "old", "new"],
    "ColumnName": ["id", "old", "new"],
    "ColumnName": ["id", "old", "new"]
  }  
}
```
* Delete (DELETE): <br>
```json
{
  "query_type": "delete",
  "user_name":"elia",
  "destination": "myFile",
  "query_content": {
    "field1": "all_items",
    "field2": 56,
    "filed3": "example2"
  }
}
```
Advanced Query: <br>
* Filtering (GET) : 
```json
{
  "query_type": "select",
  "user_name":"elia",
  "destination": "myFile",
  "query_content": {
    "field1": "example1",
    "field2": "example2",
    "field3": "example3",
    "field4": {
      "between": [30, 40]
    },
    "field5": {
      "max": 50
    },
    "field6": {
      "min": 30
    },
    "field7": {
      "word_length": 60
    }
  }
}
```
|Instructions|
|------------|
|between|
|max|
|min|
|length|
|min_length|
|max_length|
|length_between|
* Union (GET) : <br>
```json
{
  "query_type": "select",
  "user_name":"elia",
  "destination": "myFile",
  "query_content": {
    "field1": "value",
    "field2": "all_items",
    "filed3": 23
  },
  "union":{
    "destination": "myFile2",
    "field1": "value",
    "field2": 233
  }
}
```
* Sorting (GET, PATCH, POST) : <br>
```json
{
  "sort_asc": true,
  "sort_desc": true
}
```

Future changes : <br>
* Implementing a Microservices architecture;
* Implementing a Cluster-based file storage system;
* Implementing an Eventual Consistency system, based on a follower-leader architecture.
