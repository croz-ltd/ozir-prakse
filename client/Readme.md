# Simple GoLang redirector

The "client" side of the summer internship selection task

Sends out a request to the server, receives a JSON number as response and redirects the user accordingly.
If running locally adjust the `http://server:8000` URL accordingly.

## Running the client

To build and run the client change directory to where your code is and run

```
go env -w GO111MODULE=auto
go build -o klijent
./klijent
```

If the server is up and running you should be able to go to `http://127.0.0.1:8090` and enjoy a relaxing web-comic.
