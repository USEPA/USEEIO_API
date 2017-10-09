# Data Storage Format
The server-side data is stored as plain files in a specific folder structure.
The path to this folder is a start parameter of the server... 

## File Formats

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

### Matrix Files
...

## Meta-Data

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
