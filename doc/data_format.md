## Data formats

The data required for serving USEEIO models via the USEEIO API are matrix files and metadata files. These data are exported from an [useeior model](https://github.com/USEPA/useeior/blob/master/format_specs/model.md).

### Matrix Files
 The series of matrices include model component, result and price adjustment matrices. Optionally, models may also have data quality matrices. All matrices are stored as binary data with a `.bin` extension.

### CSV Files
Indicators, sectors, flows, and other metadata are stored in plain CSV files.
In general, these CSV files should have the following format:

* The first row contains the column headers. It is ignored when reading the
  files.
* Commas (`,`) are used as column separators.
* Character strings have to be enclosed in double quotes (`"`) if they contain 
  a comma or multiple lines. In other cases the quoting is optional.
* Numbers or Boolean values are never enclosed in quotes. The decimal separator
  is a point (`.`).
* The file encoding is UTF-8 without byte-order mark

The columns of these files are specified in the respective sections below.

#### Data quality matrices
The API supports data quality matrices associated with any model result matrix. These matrices are optionally included and not present in all models. Each data quality matrix consists of data quality entries for values in the same position of its associated matrix. These entries are in the form of 5 data quality scores (values for each ranging from 1 to 5) for the five [EPA LCA flow data quality indicators](https://cfpub.epa.gov/si/si_public_record_report.cfm?dirEntryId=321834) in this order: (Data reliability, Temporal correlation, Technological correlation, Technological correlation). For example '(3,1,1,4,5)'.  Each data quality matrix is named with the matrix name followed by `_dqi` like `B_dqi` which provides data quality scores for the `B` matrix. Data quality matrices are stored as [CSV files](#CSV-files).

## Common metadata for All Models

### models.csv
The file `models.csv` which is located in the root folder of the data directory
contains the metadata of the available IO models and should have the following
columns:

```
Column   Field     Description
0        ID            The ID of the model. This is also the name of the folder
                       that contains the model data.
1        Name          A user friendly name of the model.
2        Location      The location code of the model.
3        Description   A description of the model
4        Sector_Schema The reference schema used for the sector codes, e.g. 'BEA_2012_Summary', which corresponds to a schema in the sectorcrosswalk.csv. Uses model name if the model has a unique set of codes.
```

### sectorcrosswalk.csv
A correspondence file between various reference schemas for sector codes, applicable to all models. Each column has a title of the schema name plus code, e.g., 'BEA_2012_Summary', and values are all available codes in that schema. Currently various [BEA IO](https://www.bea.gov/data/industries/input-output-accounts-data) and [NAICS](https://www.census.gov/cgi-bin/sssd/naics/naicsrch) schemas are represented. 

## metadata Specific to Models

### indicators.csv
The file `indicators.csv` contains the metadata of the indicators in the model.
It should have the following columns:

```
Column  Field       Description
0       Index       The zero-based matrix index of the indicator.
1       ID          The ID of the indicator (typically the same as in the 
                    input-output model builder).
2       Name        The name of the indicator.
3       Code        The code of the indicator.
4       Unit        The indicator unit in which results are calculated.
5       Group       The indicator group which should be exactly one of these values:
                    Impact Potential, Resource Use, Waste Generated,
                    Economic & Social, Chemical Releases
6       SimpleUnit  A simplified version of the Unit. 
7       SimpleName  A simplified version of Name.
```

### demands.csv
The file `demands.csv` contains the information of all available demand vectors
of the model. It should have the following columns:

```
Column   Field    Description
0        ID       The ID of the demand vector. The ID + the file extension should
                  be the name of the file under the demands folder that contains
                  the data of the demand vector.
1        Year     The year of the demand vector.
2        Type     'Production' or 'Consumption'.
3        System   'Complete' or the name of the sub-system.
4        Location The location code.
```

### flows.csv
The file `flows.csv` contains the metadata of the flows in the model.
It should have the following columns:

```
Column  Field       Description
0       Index       The zero-based matrix index of the flow.
1       ID          The ID of the flow.
2       Name        The name of the flow, e.g., 'Sulfuric acid'.
3       Category    The category of the flow, generally the primary environmental compartment, e.g. 'air'.
4       Sub-Category A sub-category of the category, e.g., 'low-population density'.
5       Unit        The flow unit, e.g., 'kg'.
6       UUID        A 36-digit hexadecimal ID for the flow.
```

### sectors.csv
The file `sectors.csv` contains the metadata of the sectors in the model.
It should have the following columns:

```
Column  Field         Description
0       Index         The zero-based matrix index of the sector.
1       ID            The ID of the sector.
2       Name          The name of the sector, e.g., 'Wood products'.
3       Code          The code of the sector in the Sector_Schema.
4       Location      The location code.
5       Description   A text description of the sector, optional.
```

### years.csv
The file `years.csv` contains the years for model data that is in model [sector-by-year](https://github.com/USEPA/useeior/blob/master/format_specs/model.md#sector-by-year) format. It should have the following columns:

```
Column  Field         Description
0       Index         The zero-based matrix index of the sector.
1       ID            The ID of the flow.
```
