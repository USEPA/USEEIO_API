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
To build the USEEIO API from source, you must have generated data for the USEEIO models to be provided
in the specified [format](./doc/data_format.md). This may be performed using the [IO Model Builder
matio module](https://github.com/USEPA/IO-Model-Builder/blob/master/iomb/matio.py).

You need to have the Go compiler installed in order to build the USEEIO API.
On Windows, you then simply run the `make.bat` file from the command line.

TODO: On Linux (and macOS) ... Building a Docker image  ...

## License
This project is in the worldwide public domain, released under the 
[CC0 1.0 Universal Public Domain Dedication](https://creativecommons.org/publicdomain/zero/1.0/).

![Public Domain Dedication](https://licensebuttons.net/p/zero/1.0/88x31.png)
