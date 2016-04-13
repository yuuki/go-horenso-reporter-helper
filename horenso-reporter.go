package horensoreporter

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Songmu/horenso"
)

// Get horenso report via Stdin
func GetReport() (*horenso.Report, error) {
	j, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}

	var r horenso.Report
	json.Unmarshal(j, &r)

	return &r, nil
}
