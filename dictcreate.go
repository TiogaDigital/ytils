package ytils

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func DictCreate(filenameofCSV string) string { // returns the two columns as a dictionary in go format
	file, err := os.Open(filenameofCSV)
	if err != nil {
		// err is printable
		// elements passed are separated by space automatically
		log.Println("Error:", err)
		return "nothing"
	}
	// automatically call Close() at the end of current method
	defer file.Close()
	//
	writer := bytes.NewBufferString("")

	reader := csv.NewReader(file)
	// // options are available at:
	// // http://golang.org/src/pkg/encoding/csv/reader.go?s=3213:3671#L94
	// reader.Comma = ','
	lineCount := 0
	for {
		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("Error:", err)
			return writer.String()
		}
		// record is an array of string so is directly printable
		//log.Println("Record", lineCount, "/", len(record), "fields", "is", record)
		lineItem := "\"" + record[0] + "\": \"" + record[1] + "\",\n"
		writer.WriteString(lineItem)
		// // and we can iterate on top of that
		// for i := 0; i < len(record); i++ {
		// 	log.Println(" ", record[i])
		// }
		// log.Println()
		lineCount += 1
	}

	return writer.String()
}
