package main

import (
	"fmt"

	"github.com/Yadier01/gol/filemanger"
	"github.com/Yadier01/gol/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanger.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		pricesJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		pricesJob.Process()
	}
}
