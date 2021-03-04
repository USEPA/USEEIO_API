## Data formats

The data required for serving USEEIO models via the USEEIO API are matrix files and metadata files.

### Matrix Files
The web service of the USEEIO API uses matrices and
meta-information that are exported from an `iomb` model.


 The series of matrices
include Data matrices and Data quality matrices.

#### Data matrices
The matrix `A` is a `sector x sector` matrix and contains in each column `i` the
direct sector inputs that are required to produce 1 USD dollar of output from
sector `i`:

```
         sectors
        +-------+
sectors |       |
        |     A |
        +-------+
```

The related `A_d` matrix provide direct sector inputs per dollar output that are only from the US.

The satellite matrix `B` is a `flow x sector` matrix and contains in
each column `i` the amount of a flow given in the reference
units of the respective flow per 1 USD output from sector `i`:

```
       sectors
      +-------+
flows |       |
      |     B |
      +-------+
```

In the matrix `C`, each column `k` contains the characterization factors of
the different indicators related to one reference unit of flow `k`:

```
                  flows
                +-------+
indicators      |       |
                |     C |
                +-------+
```

From the matrices `B` and `C` the direct impact matrix `D` can be calculated
via:

```
D = C * B
``` 

The matrix `D` contains in each column `i` the direct impact result per USD output from sector `i`:

```
                 sectors
                +-------+
indicators      |       |
                |     D |
                +-------+
```

The Leontief inverse `L` is calculated via:

```
L = (I - A)^-1
```

`L` is also a `sector x sector` matrix and contains in each column `i` the
total requirements of the respective sectors inputs per 1 USD of output
from sector `i`:


```
         sectors
        +-------+
sectors |       |
        |     L |
        +-------+
```

The related `L_d` matrix provides direct + indirect sector inputs per dollar output that are only from the US.

The domestic Leontief inverse `L_d` is calculated via:
```
L_d = (I - A_d)^-1`
```

With the `B` and the total requirements `L`, the matrix `M` which
contains the direct + indirect flow totals can be calculated via:

```
M = B * L
```

The matrix `M` is a `flow x sector` matrix and contains in each column
`i` the direct and indirect emissions and resources per 1 USD output of sector `i`:

```
                 sectors
                +-------+
flows           |       |
                |     M |
                +-------+
```


With the direct impacts `D` and the total requirements `L`, the matrix `N` which
contains the direct + indirect impact per dollar totals can be calculated via:

```
N = D * L
```

The matrix `N` is a `indicator x sector` matrix and contains in each column
`i` the direct and indirect impact result per 1 USD output of sector `i`:

```
                 sectors
                +-------+
indicators      |       |
                |     N |
                +-------+
```

#### Data quality matrices
There are also 3 data quality matrices associated with three of the matrices above.
Each matrix consists of data quality etnries for values in the same position of its associated matrix.
 These entries are in the form of 5 data quality scores (values for each ranging from 1 to 5) for the five
 EPA LCA flow data quality indicators in this order: (Data reliability, Temporal correlation, Technological correlation,
         Technological correlation). For example '(3,1,1,4,5)'.

`B_dqi` provides data quality scores for the `B` matrix.

`D_dqi` provides data quality scores for the `D` matrix.

`N_dqi` provides data quality scores for the `N` matrix.


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
1       ID            The ID of the flow.
2       Name          The name of the sector, e.g., 'Wood products'.
3       Code          The code of the sector in the Sector_Schema.
4       Location      The location code.
5       Description   A text description of the sector, optional.
