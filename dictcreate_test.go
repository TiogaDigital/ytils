package ytils

import (
	"log"
	"testing"
)

func TestDictCreate(t *testing.T) {

	filenameofCSV := "./testdata.csv"
	results := DictCreate(filenameofCSV)

	log.Printf("***\n*** Results of test: \n%v \n***\n", results)

}
