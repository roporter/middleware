package recovery

import (
	"io"
	"os"
	"time"

	"github.com/kataras/iris"
)

type recovery struct {
	//out optional output to log any panics
	out io.Writer
}

func (r recovery) Serve(ctx *iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			r.out.Write([]byte("[IRIS:" + time.Now().String() + "] Recovery from panic \n"))
			//ctx.Panic just sends  http status 500 by default, but you can change it by: iris.OnPanic(func( c *iris.Context){})
			ctx.Panic()
		}
	}()
	ctx.Next()
}

// New restores the server on internal server errors (panics)
// receives an optional writer, the default is the os.Stderr if no out writer given
// returns the middleware
func New(out ...io.Writer) iris.HandlerFunc {
	r := recovery{os.Stderr}
	if out != nil && len(out) == 1 {
		r.out = out[0]
	}
	return r.Serve
}
