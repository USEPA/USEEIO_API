# USEEIO API

The US Environmentally-Extended Input-Output Model Application Programming Interface (USEEIO API) is a web service that
provides pre-calculated [USEEIO](https://www.epa.gov/land-research/us-environmentally-extended-input-output-useeio-models)
model results and model metadata for USEEIO family models. The USEEIO API is part of the [USEEIO modeling framework](https://github.com/USEPA/useeio).

This repository contains the software to build the web service and documentation and to test the API. It is also
associated with a [Wiki](https://github.com/USEPA/USEEIO_API/wiki/) with instructions to build, test, deploy and use
 the USEEIO API. Live implementations of the USEEIO API are hosted elsewhere.

The USEEIO API web service is written in [Go](https://golang.org/) and can be built for serving on Windows and Linux machines. It can be further containerized
 in Docker container or deployed to a cloud foundry server. These web services can be built and served on a local machine or
 deployed to a remote server. USEEIO models make up the data for the web service. They are not hosted here, but must be built
  and exported in the [specified formats](format_specs/data_format.md) in order to be served by the
web services. USEEIO models are built in [useeior](https://github.com/USEPA/useeior/).

The web service can serve multiple USEEIO models. For each model, it serves available matrices and accompanying metadata, and also supports calculate of impact results for a given model with a user
defined demand vector and result perspective. More information about the provided components and metadata
 can be found in the API documentation.

The API documentation uses a standard API documentation specification, Swagger 2.0. It is built into HTML pages
using the `npm` and `gulp` Javscript package manager and build system.

The test suite is a set Python 3 pass/fail unit tests that extensive test all models on any local or remote web service.

The USEEIO API attempts to implement the [18F US federal government API best practices](https://github.com/18F/api-standards).

## Disclaimer

The United States Environmental Protection Agency (EPA) GitHub project code is provided on an "as is" basis
 and the user assumes responsibility for its use.  EPA has relinquished control of the information and no longer
  has responsibility to protect the integrity , confidentiality, or availability of the information.  Any
   reference to specific commercial products, processes, or services by service mark, trademark, manufacturer,
    or otherwise, does not constitute or imply their endorsement, recommendation or favoring by EPA.  The EPA seal
     and logo shall not be used in any manner to imply endorsement of any commercial product or activity by EPA or
      the United States Government.
