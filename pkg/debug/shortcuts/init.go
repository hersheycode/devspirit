package writingwalk

import (
	"fmt"
	"io/fs"
)

//Debug
type D struct {
	L func(...any) (int, error)         //Line
	F func(string, ...any) (int, error) //Format
	S func(string, ...any) string       //String Format
	Quick
}

type Quick struct {
	P func(...any)
}

//New
func N() D {
	pf := func(s string, v ...any) (int, error) {
		if len(v) < 1 {
			return 0, nil
		}
		f := "%+v \n"
		if s != "" {
			f = s + "\n"
		}
		fmt.Printf(f, v...)
		return 0, nil
	}
	return D{
		L: fmt.Println,
		F: pf,
		S: fmt.Sprintf,
		// Quick: Quick{
		// 	P: pf,
		// },
	}
}

func (d D) PDirs(dirs []fs.DirEntry) {
	for _, dir := range dirs {
		d.F("Dir Name: %+v", dir.Name())
	}
}
