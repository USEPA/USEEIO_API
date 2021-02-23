# USEEIO API

The US Environmentally-Extended Input-Output Model Application Programming Interface (USEEIO API) is a web service that
provides pre-calculated [USEEIO](https://cfpub.epa.gov/si/si_public_record_report.cfm?dirEntryId=336332)
model results and model metadata for USEEIO family models.

This repository contains the software to build the web service and documentation and to test the API. It is also
associated with a [Wiki](https://github.com/USEPA/USEEIO_API/wiki/) with instructions to build, test, deploy and use
 the USEEIO API. Live implementations of the USEEIO API are hosted elsewhere.

There are two versions of the USEEIO API web service, a [Go](https://golang.org/) and Python version. Both provide the same
 endpoints (e.g. `api/models`) and serve the results in JSON format.  The Go version can be further containerized
 in Docker container or deployed to a cloud foundry server. These web services can be built and served on a local machine or
 deployed to a remote server. USEEIO models make up the data for the web service. They are not hosted here, but must be built
  and exported in the [specified formats](./doc/data_format.md) in order to be served by the
web services. See the [instructions](https://github.com/USEPA/USEEIO_API/wiki/Build#export-the-model-files-and-create-the-modelscsv-metadata-file-first)
 on preparing USEEIO models for the web service.

The web services provide the following USEEIO model components and metadata

| Item | Type | Description |
| --- | --- | --- |
| models | metadata | A list of USEEIO models provided |
| A | component | The direct requirements matrix for a given model |
| A_d | component | The direct requirements matrix with only US (domestic)inputs |
| B | component | The satellite/flow matrix for a given model |
| C | component | The characterization factor matrix for a given model |
| D | component | The direct impact matrix for a given model |
| L | component | The Leontief inverse matrix for a given model |
| L | component | The Leontief inverse matrix for a given model |
| M | component | The direct + indirect flow per dollar matrix for a given model |
| N | component | The direct + indirect impact per dollar matrix for a given model |
| B_dqi | component | Data quality scores for the B matrix for a given model |
| D_dqi | component | Data quality scores for the D matrix for a given model |
| N_dqi | component | Data quality scores for the U matrix for a given model |
| sectors | metadata | A list of sectors in a given model|
| flows | metadata | A list of flows in a given model|
| indicators | metadata | A list of indicators in a given model|
| demands | metadata | A list of demand vectors available for use with a given model|

The web service also supports calculate of life cycle impact results for a given model with a user
defined demand vector and result perspective. More information about the provided components and metadata
 can be found in the API documentation.

The API documentation uses a standard API documentation specification, Swagger 2.0. It is built into HTML pages
using the `npm` and `gulp` Javscript package manager and build system.

The test suite is a set Python 3 pass/fail unit tests that extensive test all models on any local or remote web service.

The USEEIO API attempts to implement the [18F US federal government API best practices](https://github.com/18F/api-standards).

## Contact information

For more information, contact [Wesley Ingwersen](https://github.com/WesIngwersen).

## Disclaimer

The United States Environmental Protection Agency (EPA) GitHub project code is provided on an "as is" basis
 and the user assumes responsibility for its use.  EPA has relinquished control of the information and no longer
  has responsibility to protect the integrity , confidentiality, or availability of the information.  Any
   reference to specific commercial products, processes, or services by service mark, trademark, manufacturer,
    or otherwise, does not constitute or imply their endorsement, recommendation or favoring by EPA.  The EPA seal
     and logo shall not be used in any manner to imply endorsement of any commercial product or activity by EPA or
      the United States Government.
