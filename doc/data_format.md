## Data formats

The data required for USEEIO models are matrix files and metadata files.

### Matrix Files
The web-service of the USEEIO API uses matrices and
meta-information that are exported from an `iomb` model. The series of matrices
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

The satellite matrix `B` is an `elementary flow x sector` matrix and contains in
each column `i` the amounts of emissions and resources - given in the reference
units of the respective elementary flows - that are directly related to 1 USD
output from sector `i`:

```
       sectors
      +-------+
flows |       |
      |     B |
      +-------+
```

In the matrix `C`, each column `k` contains the characterization factors of
the different LCIA categories related to one reference unit of flow `k`:

```
                  flows
                +-------+
LCIA categories |       |
                |     C |
                +-------+
```

From the matrices `B` and `C` the direct impact matrix `D` can be calculated
via:

```
D = C * B
``` 

The matrix `D` contains in each column `i` the impact assessment result that is
related to the direct emissions and resources for 1 USD output from sector `i`:

```
                 sectors
                +-------+
LCIA categories |       |
                |     D |
                +-------+
```

The Leontief inverse `L` is calculated via:

```
L = (I - A)^-1
```

`L` is also a `sector x sector` matrix and contains in each column `i` the
total requirements of the respective sectors inputs to produce 1 USD of output
from sector `i`:

```
         sectors
        +-------+
sectors |       |
        |     L |
        +-------+
```

With the direct impacts `D` and the total requirements `L` the matrix `U` which
contains the upstream totals can be calculated via:

```
U = D * L
```

The matrix `U` is a `LCIA category x sector` matrix and contains in each column
`i` the total impact assessment result related to the direct and indirect 
emissions and resources that are required to produce 1 USD output of sector `i`:

```
                 sectors
                +-------+
LCIA categories |       |
                |     U |
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

`U_dqi` provides data quality scores for the `U` matrix.


### CSV Files
Indicators, sectors, flows, and other meta-data are stored in plain CSV files.
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

## Meta-Data

### models.csv
The file `models.csv` which is located in the root folder of the data directory
contains the meta-data of the available IO models and should have the following
columns:

```
Column   Field     Description
0        ID            The ID of the model. This is also the name of the folder
                       that contains the model data.
1        Name          A user friendly name of the model.
2        Location      The location code of the model.
3        Description   A description of the model
```

### indicators.csv
The file `indicators.csv` contains the meta-data of the indicators in the tools.
It should have the following columns:

```
Column  Field    Description
0       Index    The zero-based matrix index of the indicator.
1       ID       The ID of the indicator (typically the same as in the 
                 input-output model builder).
2       Name     The name of the indicator that is displayed in the tool.
3       Code     The code of the indicator used in charts and tables.
4       Unit     The indicator unit in which results are calculated.
5       Group    The indicator group which should be exactly one of these values:
                 Impact Potential, Resource Use, Waste Generated,
                 Economic & Social, Chemical Releases
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
