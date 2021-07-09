package markdown

import (
	"fmt"
	"reflect"
	"strings"

	gotkio "github.com/digisan/gotk/io"
	"github.com/digisan/gotk/slice/ti"
)

// "L", "M", "R"
func align(lmrGrp ...interface{}) string {

	mLMR := map[interface{}]string{
		"l": ":-----",
		"L": ":-----",
		"m": ":----:",
		"M": ":----:",
		"r": "-----:",
		"R": "-----:",
	}

	sb := &strings.Builder{}
	for _, lmr := range lmrGrp {
		sb.WriteString("| ")
		_, ok := mLMR[lmr]
		if !ok {
			lmr = "L"
		}
		sb.WriteString(mLMR[lmr])
		sb.WriteString(" ")
	}
	sb.WriteString("|")
	return sb.String()
}

func row(items ...interface{}) string {
	sb := &strings.Builder{}
	for _, item := range items {
		sb.WriteString("| ")
		sb.WriteString(fmt.Sprint(item))
		sb.WriteString(" ")
	}
	sb.WriteString("|")
	return sb.String()
}

func padding(slcGrp ...interface{}) int {
	nGrp := make([]int, len(slcGrp))
	for i, slc := range slcGrp {
		switch v := slc.(type) {
		case *[]interface{}:
			nGrp[i] = len(*v)
		default:
			panic("need more type to set align")
		}
	}
	m := ti.Max(nGrp...)
	for _, slc := range slcGrp {
		switch v := slc.(type) {
		case *[]interface{}:
			if d := m - len(*v); d > 0 {
				*v = append(*v, make([]interface{}, d)...)
			}
		}
	}
	return m
}

// MDTable : data is addr of struct, each struct field type is '*[]interface{}'
func MDTable(data interface{}, outfile string) {

	var (
		m      = 0
		v      = reflect.ValueOf(data)
		sb     = &strings.Builder{}
		fnames = []string{}
		cols   = [][]interface{}{}
		rows   = [][]interface{}{}
	)

	if v.Elem().Kind() == reflect.Struct {

		fields := []interface{}{}
		for i := 0; i < v.Elem().NumField(); i++ {
			fnames = append(fnames, v.Elem().Type().Field(i).Name)
			fields = append(fields, v.Elem().Field(i).Interface().(*[]interface{}))
		}
		m = padding(fields...)
		// fmt.Println(m)
		// for _, field := range fields {
		//  fmt.Println(field)
		// }
		// fmt.Println(fnames)

		for _, field := range fields {
			col := []interface{}{}
			for i := 0; i < m; i++ {
				col = append(col, (*field.(*[]interface{}))[i])
			}
			cols = append(cols, col)
		}
	}

	for i := 0; i < m; i++ {
		row := []interface{}{}
		for _, col := range cols {
			row = append(row, col[i])
		}
		// fmt.Println(row)
		rows = append(rows, row)
	}

	for i, item := range rows[0] {
		if item == nil || item == "" {
			rows[0][i] = fnames[i]
		}
	}

	for i := 0; i < len(rows); i++ {
		switch i {
		case 1:
			sb.WriteString(align(rows[i]...))
			sb.WriteString("\n")
		default:
			sb.WriteString(row(rows[i]...))
			sb.WriteString("\n")
		}
	}

	content := strings.ReplaceAll(sb.String(), "<nil>", "")
	gotkio.MustWriteFile(outfile, []byte(content))
}
