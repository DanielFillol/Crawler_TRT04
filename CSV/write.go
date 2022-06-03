package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Crawler_TRT04/Crawler"
	"os"
	"path/filepath"
)

const (
	fileName   = "Jurisprudência"
	folderName = "Results"
)

func ExportCsv(decisions []Crawler.Decision) error {
	var rows [][]string

	rows = append(rows, generateHeaders())

	for i := 0; i < len(decisions); i++ {
		rows = append(rows, generateRow(decisions[i]))
	}

	cf, err := createFile(folderName + "/" + fileName + ".csv")
	if err != nil {
		return err
	}

	defer cf.Close()

	w := csv.NewWriter(cf)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

// create csv file from operating system
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

// generate the necessary headers for csv file
func generateHeaders() []string {
	return []string{
		"Título",
		"Data",
		"Orgão Julgador",
		"Relator",
		"Ementa",
		"Link Inteiro Teor",
	}
}

// returns []string that compose the row in the csv file
func generateRow(result Crawler.Decision) []string {
	return []string{
		result.Title,
		result.Date,
		result.Court,
		result.Judge,
		result.DecisionText,
		result.LinkDecision,
	}
}
