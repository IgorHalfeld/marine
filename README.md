# Golang Marine Traffic and Alert API

## Running

To run this project just enter the command bellow:

Running in the terminal if you have golang installed:
```
make run
```

Building and Creating a Container:
```
make docker-build
```

Running the Container locally:
```
make docker-run
```

## Requests 

GET http://localhost:80/traffic HTTP/1.1

###

GET http://localhost:80/general HTTP/1.1

###

POST http://localhost:80/alert HTTP/1.1
content-type: application/json

{
	"image_url": "https://firebasestorage.googleapis.com/v0/b/fala-ai-portohacksantos.appspot.com/o/1575788153051.jpg?alt=media&token=f6a0e2f0-6a40-4ffd-9c9b-f8ed58736bf8",
	"type": "2",
	"description": "3",
	"finger_print": "4",
	"lat": "5",
	"lon": "6"
}


## Running on Docker

The endpoint and the Port are external varibles setup in the Dockerfile, if change is needed it is ok, but they are required.

Thanks ;D