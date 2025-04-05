# 1 Billion Row Challenge

Essentially a slimmed down fork of [Gunnar Morling's 1 Billion Row Challenge](https://github.com/gunnarmorling/1brc/tree/main), a fun practice of data aggregation with a 1 billion row file.

## Setup

Large files would be awful to commit to and pull down from Github, so instead all `input/*.txt` files are being ignored from git changes. To setup for the project:

1. Ensure you have Java 17 or newer installed and available on your PATH.
2. CD into the `inputs` directory.
3. Run `java CreateMeasurements.java <number of rows>`
  - This took about 9 minutes to generate 1 billion rows for me.
  - `CreateMeasurementsFast.java` is noticeably faster, but doesn't run safety checks or batches and dies if it runs out of memory for larger amounts.
4. This may be repeated (likely with renaming the resulting file in between) to have different sample sizes to test with.

## Input/Output Example

The text file contains temperature values for a range of weather stations. Each row is one measurement in the format <string: station name>;<double: measurement>, with the measurement value having exactly one fractional digit. The following shows ten rows as an example:

```plaintext
Hamburg;12.0
Bulawayo;8.9
Palembang;38.8
St. John's;15.2
Cracow;12.6
Bridgetown;26.9
Istanbul;6.2
Roseau;34.4
Conakry;31.2
Istanbul;23.0
```

The task is to write a program which reads the file, calculates the min, mean, and max temperature value per weather station, and emits the results on stdout like this (i.e. sorted alphabetically by station name, and the result values per station in the format <min>/<mean>/<max>, rounded to one fractional digit):

```json
{Abha=-23.0/18.0/59.2, Abidjan=-16.2/26.0/67.3, Abéché=-10.0/29.4/69.0, Accra=-10.1/26.4/66.4, Addis Ababa=-23.7/16.0/67.0, Adelaide=-27.8/17.3/58.5, ...}
```
