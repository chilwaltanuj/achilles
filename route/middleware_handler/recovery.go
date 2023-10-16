package middlewareHandler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"runtime"

	"achilles/model"
	routeHelper "achilles/route/helper"

	"github.com/gin-gonic/gin"
)

const (
	recoveryStackSkip = 5
	defaultHTTPStatus = 500
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

// Recovery is a middleware that recovers from any panics.
// It writes the response status code to 500 and includes
// a JSON response with `success` as `false` and a `message` from the
// error which caused the panic. Details of the error are logged.
func Recovery() gin.HandlerFunc {
	return RecoveryMiddleware()
}

// RecoveryMiddleware is the core recovery middleware function.
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				stack := getStack(recoveryStackSkip)
				errorMessage := fmt.Sprintf("[Recovery] panic recovered: %s \n %s", r, stack)

				responseData := model.HttpResponseData{
					Success: false,
					Status:  defaultHTTPStatus,
					Message: errorMessage,
				}
				routeHelper.BuildAndSetHttpResponseInContext(c, responseData)
				routeHelper.UpdateRequestMetaDataInContext(c)
				routeHelper.RenderJsonResponse(c)
			}
		}()
		c.Next()
	}
}

// getStack returns a formatted stack frame.
func getStack(skip int) []byte {
	buf := new(bytes.Buffer)
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				// Handle the error or log it
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		fmt.Fprintf(buf, "\t%s: %s\n", getFunctionName(pc), getSource(lines, line))
	}
	return buf.Bytes()
}

// getSource returns a space-trimmed slice of the n'th line.
func getSource(lines [][]byte, n int) []byte {
	n--
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

// getFunctionName returns the name of the function containing the PC.
func getFunctionName(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	lastSlash := bytes.LastIndex(name, slash)
	if lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	firstDot := bytes.Index(name, dot)
	if firstDot >= 0 {
		name = name[firstDot+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}
