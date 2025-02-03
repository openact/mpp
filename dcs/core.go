package dcs

import (
	"log"
	"mpp/layout"
	"reflect"
	"strconv"
	"strings"

	"github.com/openact/lib/utils"
	"github.com/samber/lo"
)

func PopulateInputFields(in *layout.MpData) {
	// output = input
	val := reflect.ValueOf(in.Output)
	typ := reflect.TypeOf(in.Output)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		varType := field.Tag.Get("type")
		if varType == "input" {
			varDef := field.Tag.Get("def")
			value, ok := utils.FindFieldByName(in.Input, varDef)
			if !ok {
				panic("field not found")
			}
			val.Field(i).Set(value)
		}
	}
}

func ReadCache(tblName string, idx ...string) string {
	key := idx[0]
	for _, v := range idx[1 : len(idx)-1] {
		key = key + ":" + v
	}

	tbl, ok := Tables[tblName]
	if !ok {
		log.Printf("Table %s not loaded\n", tblName)
	}
	subKeys := tbl.SubKeys
	val := tbl.Caches.Get(nil, []byte(key))
	valStrS := strings.Split(string(val), ",")

	n := lo.IndexOf(subKeys, idx[len(idx)-1])

	if n == -1 {
		// log.Printf("Key %s not found in %s\n", idx[len(idx)-1], tblName)
		return ""
	}

	if len(valStrS) <= n {
		return ""
	}

	return valStrS[n]
}

func ReadCacheInt(tblName string, idx ...string) int {
	str := ReadCache(tblName, idx...)
	if str == "" {
		return 0
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}

func ReadCacheFloat(tblName string, idx ...string) float64 {
	str := ReadCache(tblName, idx...)
	if str == "" {
		return 0
	}
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return val
}

func ReadCacheString(tblName string, idx ...string) string {
	return ReadCache(tblName, idx...)
}

// todo: add "Name" to enum for error handlings
func PopulateStringEnum(in *layout.Input, tgt []string, enum *utils.Enum) {
	if enum.Size() == 0 {
		log.Println("enum size is 0")
		log.Fatalf("enum size is 0")
	}

	for i := 0; i < enum.Size(); i++ {
		str, ok := enum.IntToStr(i)
		if !ok {
			// log.Printf("enum %s index %v not found", enum.Name, i)
		} else {
			val := reflect.ValueOf(in).Elem().FieldByName(str)
			if !val.IsValid() {
				// log.Printf("field %s not found, face value %s was used", enum.Name, str)
				tgt[i] = str
			} else {
				tgt[i] = val.String()
			}
		}
	}
}

func PopulateIntEnum(in *layout.Input, tgt []int, enum *utils.Enum) {
	if enum.Size() == 0 {
		log.Println("enum size is 0")
		log.Fatalf("enum size is 0")
	}

	for i := 0; i < enum.Size(); i++ {
		str, ok := enum.IntToStr(i)
		if !ok {
			// log.Printf("enum %s index %v not found", enum.Name, i)
		} else {
			val := reflect.ValueOf(in).Elem().FieldByName(str)
			if !val.IsValid() {
				// log.Printf("field %s not found, face value %v was used", enum.Name, str)
				v, err := strconv.Atoi(str)
				if err != nil {
					panic(err)
				}
				tgt[i] = v
			} else {
				tgt[i] = int(val.Int())
			}
		}
	}
}

// todo: output can map to input (+-x/)
// todo: read fac
