# RenzoFS


![Diagramma senza titolo drawio](https://github.com/Elia-Renzoni/RenzoFS/assets/118525453/3cd2ba5c-f996-4379-98b5-5c254a446978)
<br>

RenzoFS is distributed file system with a web based interface that allows user to see their remote file system and query their file, using json written query. <br>
<br>
The backend is written in Go (1.20) and uses RESTful API to interact with the web server, specially with the remote file system, wich is accessed by client only on remote access model, instead of upload-download access model. <br>
The main functionalities of the API are: <br>
|Functionality|
|-------------|
|Sign in user|
|Sign out user|
|Add user's friend|
|Create a remote directory|
|Create empty csv files|
|Quering files|
|Provide statistic of the file system usage|
|Provide info about who, and when, changed files or query them|
|File sharing|

<br>
Query Format: <br>

 
