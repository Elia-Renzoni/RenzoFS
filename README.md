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
this micro service handles access to remote files and directories, allowing you to store new changes, such as removing files or directories, or even editing and reading files. <br>
Query Format and API endpoints: <br>

* Write (POST) : <br>
```json
{
  "query_type": "insert",
  "user_name":"elia",
  "destination": "myFile",
  "query_content": ["value1", "44", "value2", "55.5"]
}
```
Endpoint 
```
localhost:8080/insert
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
Endpoint
```
localhost:8080/update
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
Endpoint
```
localhost:8080/createdir
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

