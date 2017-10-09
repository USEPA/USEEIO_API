# USEEIO API
This projects contains an implementation of the USEEIO API in
[Go](https://golang.org/). It is the same RESTful API as provided by the 
[iomb.webapi](https://github.com/USEPA/IO-Model-Builder/tree/master/iomb/webapi)
package. However, the implementation in Go can be compiled to a single binary
file which is a full HTTP server that provides the data from one or more
input-output models and optionally static files. All data that is is stored in
files in a specific folder structure (see also the
[documentation of the file format](./doc/data_format.md)).

## Running the Server
The server application is just a single executable file. The folder with the
data files, the server port, and the location of the static files (HTML,
JavaScript, CSS) of the application can be specified via command line arguments:

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

TODO: On Linux (and macOS) ... Building a Docker image  ...

## License
This project is in the worldwide public domain, released under the 
[CC0 1.0 Universal Public Domain Dedication](https://creativecommons.org/publicdomain/zero/1.0/).

![Public Domain Dedication](https://licensebuttons.net/p/zero/1.0/88x31.png)
