package dcs

import (
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

/* big cache
func LoadTbls() map[string]*bigcache.BigCache {
	m := make(map[string]*bigcache.BigCache)
	tblMap := utils.Conf.GetStringMapString("tableMaps")
	tblPaths := utils.Conf.GetStringSlice("tablePaths")
	for t, file := range tblMap {
		filePath := utils.GetFile(file, tblPaths)
		m[t] = utils.LoadFacToBigCache(filePath)
	}
	return m
}
*/

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
