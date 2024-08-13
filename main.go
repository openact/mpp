package main

import (
	"fmt"
	"mpp/dcs"
	"mpp/layout"
	"mpp/sink"
	"time"

	//_ "net/http/pprof"

	"github.com/goalm/lib/etl"
	"github.com/goalm/lib/utils"
)

func main() {
	/*
		go func() {
			_ = http.ListenAndServe("localhost:6060", nil)
		}()
	*/

	//for testing
	start := time.Now()

	srcData := utils.Conf.GetStringSlice("data")
	srcPaths := dcs.LocateFiles(srcData, utils.Conf.GetStringSlice("dataPaths"))
	outPath := utils.Conf.GetString("outPaths.modelPoints")
	fmt.Println("outPath: ", outPath)
	in := make(chan *layout.Input)
	//go producer.LoadSrcTbls(srcPaths, in)
	//go producer.LoadGenericFiles(srcPaths, layout.Input{}, in)
	go utils.StreamGenericFiles(srcPaths, layout.Input{}, in)

	fmt.Println("loading data ... used", time.Since(start))
	utils.InitializePath(outPath)
	// producer functions
	producerFn := dcs.ProducerFac(in)

	stageSinkFn := func(in *layout.MpData) (out *layout.MpData, err error) {
		err = sink.WriteMp(outPath, in.Output)
		//err = sink.WriteMpBufio(outPath, in.Output)
		if err != nil {
			fmt.Println("error: ", err)
		}
		return in, nil
	}

	producer := etl.NewProducer(
		producerFn,
		etl.Concurrency(10),
		etl.Name("producer"))

	stage1 := etl.NewStage(
		dcs.StageFn,
		etl.Concurrency(10),
		etl.Name("stage1"))

	stageSink := etl.NewStage(
		stageSinkFn,
		etl.Concurrency(10),
		etl.Name("stageSink"))

	err := etl.Do(producer, stage1, stageSink)

	if err != nil {
		fmt.Println("error: ", err)
	}

	// sink functions
	//sink.WriteToMongoDb(outp, "ModelPoints", "test")
	if err != nil {
		fmt.Println("error: ", err)
	}

	elapsed := time.Since(start)
	fmt.Println("elapsed: ", elapsed)
}
