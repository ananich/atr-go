# ATR indicator calculation example

Here is an example of calculation of average true range indicator using GO lang.

It reads input data from the file specified in the first command line parameter.
The file has to be in CSV format. Historical data from Yahoo! Finance can be
used for testing purposes. By default output is printed to STDOUT which may
be redirected to a file:

```
go run . AAPL.csv > AAPL+ATR.csv
```

Output contains data in CSV format similar to input plus additional column with
calculated ATR indicator.

 * Computational complexity: linear
 * Memory complexity: const

![Test1](https://github.com/ananich/atr-go/blob/bf7afdbe19a2a8eb0fc24efa595a83a72c46e320/Screen%20Shot%202023-03-11%20at%2011.03.51%20PM.png)


