# RenzoFS

![renzofs](https://github.com/Elia-Renzoni/RenzoFS/assets/118525453/90fb82a2-60a1-49ea-a8d2-2702bc451ab8)

<br>

RenzoFS is a distributed file system that allows users to query their file, using json written query. <br>
<br>
The distributed service is written in Go (1.20) and uses REST protocol as communication model, ensuring a remote access model, instead of upload-download access model. <br>
The main functionalities of the distributed service are: <br>
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


## Distributed Storage Service
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
```
localhost:8080/read/{dirname}/{filename}/?id=...
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
```
localhost:8080/delete/{dirname}/{filename}/?id=...&field=...
```
* Create a new remote directory (POST): <br>
```json
{
  "dir_to_create": "...."
}
```
* Delete a remote directory (DELETE): <br>
```
localhost:8080/deletedir/{dirname}
```
* Get a file informations (GET): <br>
```
localhost:8080/fileinfo/{dirname}/{filename}
```
<br>
Future changes : <br>
* Implementing a Microservices architecture;
* Implementing a Cluster-based file storage system;
* Implementing an Eventual Consistency system, based on a follower-leader architecture.
