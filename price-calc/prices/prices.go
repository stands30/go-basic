package prices

import (
	"fmt"

	"example.com/priceCal/conversion"
	"example.com/priceCal/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		// fmt.Println(err)
		return err
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	result := make(map[string]string)
	err := job.LoadData()
	if err != nil {
		// fmt.Println(err)
		return err
	}
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
	job.IOManager.WriteResult(job)
	return nil
}
func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
