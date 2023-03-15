package main

import (
	"ananich/atr-go/indicators"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func p(x float64) string {
	return fmt.Sprintf("%.2f", math.Round(x*100)/100)
}

func main() {
	filePath := "test/testdata/AAPL.csv"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	headers, _ := csvReader.Read() // skip header line

	csvWriter := csv.NewWriter(os.Stdout)
	csvWriter.Write(append(headers[:5], "Volume", "ATR"))
	defer csvWriter.Flush()

	atr := indicators.NewATR(14)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Unable to parse CSV file "+filePath, err)
		}

		// date := record[0]
		t := record[0]
		o, _ := strconv.ParseFloat(record[1], 32)
		h, _ := strconv.ParseFloat(record[2], 32)
		l, _ := strconv.ParseFloat(record[3], 32)
		c, _ := strconv.ParseFloat(record[4], 32)
		v, _ := strconv.ParseInt(record[6], 10, 32)

		atr.Update(o, h, l, c, v)

		line := []string{t, p(o), p(h), p(l), p(c), fmt.Sprintf("%d", v), p(atr.Value())}
		csvWriter.Write(line)
	}
}
