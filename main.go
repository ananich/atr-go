package main

import (
    "encoding/csv"
    "log"
    "os"
    "io"
    "github.com/shopspring/decimal"
)

const n = 14

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func main() {
  filePath := "AAPL.csv"
  if len(os.Args) > 1 {
    filePath = os.Args[1]
  }

  f, err := os.Open(filePath)
  if err != nil {
      log.Fatal("Unable to read input file " + filePath, err)
  }
  defer f.Close()

  csvReader := csv.NewReader(f)
  headers, err := csvReader.Read() // skip header line

  csvWriter := csv.NewWriter(os.Stdout)
  csvWriter.Write(append(headers, "ATR"))
  defer csvWriter.Flush()

  cp := decimal.Zero // previous close price
  i := 0
  atrp := decimal.Zero // previous ATR

  for {
    record, err := csvReader.Read()
    if err == io.EOF {
      break
    } else if err != nil {
      log.Fatal("Unable to parse CSV file " + filePath, err)
    }
    // date := record[0]
    // o := decimal.NewFromString(record[1])
    h, _ := decimal.NewFromString(record[2])
    l, _ := decimal.NewFromString(record[3])
    c, _ := decimal.NewFromString(record[4])

    // true range calculation
    tr := h.Sub(l)
    if cp != decimal.Zero {
      tr1 := h.Sub(cp).Abs()
      tr2 := l.Sub(cp).Abs()
      tr = decimal.Max(tr, tr1, tr2)
      cp = c
    }

    // average true range calculation
    atr := decimal.Zero
    if i == 0 {
      atr = tr
    } else {
      n1 := min(i, n)
      n_ := decimal.NewFromInt32(int32(n1))
      m_ := decimal.NewFromInt32(int32(n1-1))
      atr = atrp.Mul(m_).Add(tr).Div(n_)
    }
    atrp = atr
    i++

    record = append(record, atr.String())
    csvWriter.Write(record)
  }
}
