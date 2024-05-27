
# Statistical Calculator

## Overview 

This project is designed to calculate key statistical measures from a given dataset. Specifically, it computes the following statistics:
>
    Average (Mean)
    Median
    Variance
    Standard Deviation
>

## Instructions
### Input File

The go program reads data from a file. The file path is passed as an argument to the program. The file should contain numerical data, with each value on a separate line, like this:

>
    189
    113
    121
    114
    145
    110
>

### Running the Program

The program should be executed from the command line like this:

>
    $ go run your-program.go data.txt
>

### Output Format

After reading the file and performing the calculations, the program prints the results in the following format, with values rounded to the nearest integer:

>
    Average: 132
    Median: 118
    Variance: 785
    Standard Deviation: 28
>

## Calculations

### Average (Mean)

The average is the sum of all values divided by the number of values.

### Median

The median is the middle value when the data is sorted in ascending order. If the number of values is even, the median is the average of the two middle values.

### Variance

Variance measures the spread of the data points from the mean. It is the average of the squared differences from the mean.

### Standard Deviation

The standard deviation is the square root of the variance, indicating how much the values deviate from the mean.

## Contributing

If you wish to contribute to this project, please fork the repository and create a pull request with your changes. Make sure to include tests for any new functionality.

## License

This project is licensed under the [MIT License](LICENSE).