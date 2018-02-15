package wout

import (
	"fmt"
	"io"
	"log"
)

type wout struct{ error }

func (err wout) Print(params ...interface{}) {
	if len(params) > 0 {
		if w, ok := params[0].(io.Writer); ok {
			fmt.Fprintf(w, err.Error())
			for _, v := range params[1:] {
				log.Print(v)
			}
			log.Println(err)
		}
	}
}
