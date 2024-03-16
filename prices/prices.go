package prices

import (
	"fmt"

	"github.com/Yadier01/gol/conversion"
	"github.com/Yadier01/gol/filemanger"
)

type TaxIncludedPricesJob struct {
	IOManager         filemanger.FileManger `json:"-"`
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPricesJob) Process() {
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		TaxIncludedPrices := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludedPrices)
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
}

func (job *TaxIncludedPricesJob) LoadData() {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
	}
	prices, err := conversion.StringToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return
	}
	job.InputPrices = prices
}

func NewTaxIncludedPriceJob(fm filemanger.FileManger, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
