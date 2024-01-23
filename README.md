# GO GO GADGET IMAGE API
VERY simple RESTful API to store and serve images.  
I am currently learning Go, please do not use this project for anything safety critical.

## Usage:
To build and run, just download this repository and run `docker-compose up`

## API
Server exposes 4 methods:
### GET /
Returns `{status: ok}` if server is running
### GET /img/\<id\>
Serves an image with a specified id
### POST /img/\<id\>
Stores image sent in request's body with a specified id
### DELETE /img/\<id\>
Deletes image with specified from the server