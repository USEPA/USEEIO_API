# SMM-Tool Backend
This folder contains an implementation of the API for the SMM-Tool in
[Go](https://golang.org/). It is the same API as provided by the 
[iomb.webapi](https://github.com/USEPA/IO-Model-Builder/tree/master/iomb/webapi)
package. However, the implementation in Go can be compiled to a single static
binary file that is very easy to deploy. It just uses the `http` package from
standard library which directly contains an HTTP server that is ready for
[production](https://stackoverflow.com/questions/30832195/using-gos-http-server-for-production).

## Running the Backend Server
The backend server is just a single executable file. You should be able to start
it with a double click. The folder with the data files, the server port, and the
location of the static files (HTML, JavaScript, CSS) of the application can be
specified via command line arguments

```bash
app -data <folder with data files> -static <folder with static files> -port <server port>
```
All arguments are optional and the defaults are the same like starting the
application in the following way:

```bash
app -data data -static static -port 5000
```

## Building from Source
You need to have the Go compiler installed in order to build the backend server.
On Windows, you then simply run the `make.bat` file in the `backend` folder from
the command line:

```batch
cd backend
make.bat
```

On Linux (and macOS?) ... coming soon

## Building a Docker Image
... coming soon
