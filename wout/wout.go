package wout

import (
	"fmt"
	"io"

	"github.com/popmedic/go-logger/log"
)

type Wout struct{ Error error }

func (err Wout) Print(params ...interface{}) {
	if len(params) > 0 {
		if w, ok := params[0].(io.Writer); ok {
			fmt.Fprintf(w, err.Error.Error())
			for _, v := range params[1:] {
				log.Error(v)
			}
			log.Error(err)
		}
	}
}
