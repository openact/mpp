package sink

import (
	"bufio"
	"context"
	"fmt"
	"mpp/layout"
	"os"

	"github.com/openact/lib/utils"
)

var ctx = context.Background()

// MongoDB sink
func WriteToMongoDb[T any](out []T, db, tbl string) error {
	// gorm config
	cli := utils.GetMongoClient(ctx)
	defer func() {
		if err := cli.Close(ctx); err != nil {
			panic(err)
		}
	}()

	coll := cli.Database(db).Collection(tbl)
	coll.DropCollection(ctx)

	_, err := coll.InsertMany(ctx, out)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return nil
}

func WriteToFile(out []*layout.Output) error {
	// write to mp file
	mpFile := "./outputs/MP_FILE" // utils.GetFileName("mpFile")
	file, err := os.Create(mpFile + ".rpt")

	if err != nil {
		fmt.Println("error: ", err)
	}
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// write header - data fields of dcs.output
	fields := utils.FieldsToCsvString(out[0], "!")
	_, err = writer.WriteString(fields + "\n")
	for _, v := range out {
		_, err := writer.WriteString(utils.RecordToCsvString(v, "*") + "\n")
		if err != nil {
			fmt.Println("error: ", err)
		}
	}

	return nil
}

func WriteFromChn(outChn chan *layout.MpData) {
	mpFile := "./outputs/MP_FILE"
	file, err := os.Create(mpFile + ".rpt")

	if err != nil {
		fmt.Println("error: ", err)
	}
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for out := range outChn {
		_, err := writer.WriteString(utils.RecordToCsvString(out.Output, "*") + "\n")
		if err != nil {
			fmt.Println("error: ", err)
		}
	}
}
