package dcs

import (
	"sync"

	"github.com/goalm/lib/utils"
)

var (
	Tables = make(map[string]*utils.TblCache)
	Enums  = make(map[string]*utils.Enum)
)

func init() {
	Tables = LoadTbls()
	Enums = loadEnum()
}

// fastcache
func LoadTbls() map[string]*utils.TblCache {
	m := make(map[string]*utils.TblCache)
	tblMap := utils.Conf.GetStringMapString("tableMaps")
	tblPaths := utils.Conf.GetStringSlice("tablePaths")
	for t, file := range tblMap {
		filePath := utils.GetFile(file, tblPaths)
		// m[t] = utils.LoadFacToFastCache(filePath, 1024*1024*2048)
		m[t] = utils.LoadFacToTblCache(filePath, 1024*1024*2048)
	}
	return m
}

// loadEnum loads all enums
func loadEnum() map[string]*utils.Enum {
	m := make(map[string]*utils.Enum)
	// load enums
	for _, enm := range utils.Enums {
		m[enm] = utils.LoadCsvToEnum(utils.GetEnumName(enm))
	}
	return m
}

func LocateFiles(files []string, paths []string) []string {
	s := make([]string, 0)
	for _, v := range files {
		path := utils.GetFile(v, paths)
		s = append(s, path)
	}
	return s
}

// load data files
func StreamGenericFiles[T any](filePaths []string, row T, dataChn chan *T) {
	wg := sync.WaitGroup{}
	for _, filePath := range filePaths {
		wg.Add(1)
		go func() {
			defer wg.Done()
			utils.StreamGenericCsvFile(filePath, row, dataChn)
		}()
	}
	wg.Wait()
	close(dataChn)
}
