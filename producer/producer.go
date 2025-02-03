package producer

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"mpp/layout"
	"os"
	"strconv"
	"sync"

	"github.com/openact/lib/utils"

	"github.com/jszwec/csvutil"
)

func LoadSrcTbls(filePaths []string, dataChn chan *layout.Input) {
	wg := sync.WaitGroup{}
	for _, filePath := range filePaths {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//LoadSrcTbl(filePath, dataChn)
			LoadMpFile(filePath, dataChn)
		}()
	}
	wg.Wait()
	close(dataChn)
}

func LoadGenericFiles[T any](filePaths []string, row T, dataChn chan *T) {
	wg := sync.WaitGroup{}
	for _, filePath := range filePaths {
		wg.Add(1)
		go func() {
			defer wg.Done()
			LoadGenericFile(filePath, row, dataChn)
		}()
	}
	wg.Wait()
	close(dataChn)
}

func LoadMpFile(filePath string, dataChn chan *layout.Input) {
	file, err := os.Open(filePath)
	fmt.Println(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 1024*1024)
	// Skip the header lines
	l := 0
	for {
		peek, err := reader.Peek(1)
		if err != nil {
			if err.Error() == "EOF" {
				log.Printf("Reading completed for %s, %s lines in total, no MP data found", filePath, strconv.Itoa(l))
				return
			}
		}

		l++
		if string(peek) != "!" && string(peek) != "*" {
			// Skip the line
			_, err := reader.ReadString('\n')
			if err != nil {
				if err.Error() == "EOF" {
					log.Printf("Reading completed for %s, %s lines in total (last line w/o line break), no MP data found", filePath, strconv.Itoa(l))
					return
				}
			}
		} else {
			break
		}
	}
	csvReader := csv.NewReader(reader)
	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Printf("Error occured reading %s, %v", filePath, err.Error())
		return
	}

	for {
		record := layout.Input{}
		err = dec.Decode(&record)
		if err != nil {
			log.Printf("Error occured reading %s, %v", filePath, err.Error())
			break
		}
		if record.PROD_NAME == "" {
			record.PROD_NAME, err = utils.FilePathToName(filePath)
			if err != nil {
				log.Fatal(err)
			}
		}
		dataChn <- &record
	}
}

func LoadGenericFile[T any](filePath string, row T, dataChn chan *T) {
	file, err := os.Open(filePath)
	fmt.Println(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 1024*1024)
	// Skip the header lines
	l := 0
	for {
		peek, err := reader.Peek(1)
		if err != nil {
			if err.Error() == "EOF" {
				log.Printf("Reading completed for %s, %s lines in total, no MP data found", filePath, strconv.Itoa(l))
				return
			}
		}

		l++
		if string(peek) != "!" && string(peek) != "*" {
			// Skip the line
			_, err := reader.ReadString('\n')
			if err != nil {
				if err.Error() == "EOF" {
					log.Printf("Reading completed for %s, %s lines in total (last line w/o line break), no MP data found", filePath, strconv.Itoa(l))
					return
				}
			}
		} else {
			break
		}
	}
	csvReader := csv.NewReader(reader)
	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Printf("Error occured reading %s, %v", filePath, err.Error())
		return
	}

	for {
		record := row
		err = dec.Decode(&record)
		if err != nil {
			log.Printf("Error occured reading %s, %v", filePath, err.Error())
			break
		}
		//todo: move to downstream
		//if record.PROD_NAME == "" {
		//	record.PROD_NAME, err = utils.FilePathToName(filePath)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//}
		dataChn <- &record
	}
}
