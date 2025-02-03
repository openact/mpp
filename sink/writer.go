package sink

import (
	"bufio"
	"fmt"
	"log"
	"mpp/layout"
	"os"
	"sync"

	"github.com/openact/lib/utils"
	"github.com/schollz/progressbar/v3"
)

// todo: dont allow concurrent create files
var creatMu sync.Mutex

type BufioWriter struct {
	file *bufio.Writer
	mu   sync.Mutex
}

type Writer struct {
	file *os.File
	mu   sync.Mutex
}

func (w *Writer) Write(str string) (n int, err error) {
	w.mu.Lock()
	n, err = w.file.WriteString(str + "\n")
	w.mu.Unlock()
	return
}

var bar *progressbar.ProgressBar = progressbar.Default(-1, "Processing records")

func RegisterWriter(path string, out *layout.Output) (writer *Writer) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	// check if writer already exists
	mpName := out.PROD_NAME
	if w, exists := SyncWriters.Load(mpName); exists {
		writer = w.(*Writer)
		return
	}
	creatMu.Lock()
	defer creatMu.Unlock()
	// Double-check if writer was created while waiting for the lock
	if w, exists := SyncWriters.Load(mpName); exists {
		writer = w.(*Writer)
		return
	}
	// create the new file and writer
	f, err := os.Create(path + "/" + mpName + ".rpt")
	if err != nil {
		log.Fatalln("Error creating file ", mpName, err)
	}
	writer = &Writer{file: f, mu: sync.Mutex{}}
	SyncWriters.Store(mpName, writer)
	_, err = writer.Write(utils.FieldsToCsvString(out, "!"))
	return
}

func WriteMp(path string, out *layout.Output) error {
	w := RegisterWriter(path, out)
	// w.mu.Lock()
	// defer w.mu.Unlock()
	// write the record
	_, err := w.Write(utils.RecordToCsvString(out, "*"))

	bar.Add(1)
	if err != nil {
		fmt.Printf("path is %s, value is %v err is %v\n", path, out, err)
	}
	return nil
}

var SyncWriters sync.Map

var BufioWriters sync.Map

func RegisterBufioWriter(path string, mp *layout.Output) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	mpName := mp.PROD_NAME
	_, exists := BufioWriters.Load(mpName)
	if exists {
		return nil
	}

	nw, err := os.Create(path + "/" + mpName + ".rpt")
	if err != nil {
		log.Fatalln("Couldn't open the filename file", err)
	}

	bw := &BufioWriter{file: bufio.NewWriterSize(nw, 1024)}
	_, err = bw.file.WriteString(utils.FieldsToCsvString(mp, "!") + "\n")
	if err != nil {
		log.Fatalln("Couldn't write the header", err)
	}
	BufioWriters.Store(mpName, bw)
	return nil
}

func WriteMpBufio(path string, out *layout.Output) error {
	w, exists := BufioWriters.Load(out.PROD_NAME)
	if !exists {
		log.Printf("No bufio writer found for %s", out.PROD_NAME)
	}
	bw := w.(*BufioWriter)
	bw.mu.Lock()
	content := utils.RecordToCsvString(out, "*") + "\n"
	_, err := bw.file.WriteString(content)
	if err != nil {
		fmt.Printf("path is %s, value is %v err is %v\n", path, out, err)
	}
	// bw.file.Flush() flush to be used at the end of the process
	bw.mu.Unlock()
	return nil
}
