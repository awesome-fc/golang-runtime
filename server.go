package golangruntime

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime/debug"
)

var initialize func(ctx *FCContext) error

var handler func(ctx *FCContext, event []byte) ([]byte, error)

func invokeHandler(w http.ResponseWriter, req *http.Request) {
	requestID := req.Header.Get(fcRequestID)
	fmt.Println(fmt.Sprintf(fcLogTailStartPrefix, requestID))
	defer func() {
		if r := recover(); r != nil {
			//if the function is abnormal, try ger error info as response
			w.Header().Set(fcStatus, "404")
			w.Write([]byte(fmt.Sprintf("Error: %+v;\nStack: %s", r, string(debug.Stack()))))
		}
		fmt.Println(fmt.Sprintf(fcLogTailEndPrefix, requestID))
	}()

	event, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fcCtx := NewFromContext(req)

	resp, err := handler(fcCtx, event)
	if err != nil {
		w.Header().Set(fcStatus, "404")
		w.Write([]byte(fmt.Sprintf("Error: %+v;", err)))
		return
	}
	w.Write(resp)
}

func initializeHandler(w http.ResponseWriter, req *http.Request) {
	requestID := req.Header.Get(fcRequestID)
	fmt.Println(fmt.Sprintf(fcInitializeLogTailStartPrefix, requestID))
	defer func() {
		fmt.Println(fmt.Sprintf(fcLogInitializeTailEndPrefix, requestID))
	}()
	fcCtx := NewFromContext(req)
	if initialize == nil {
		panic("this function doesn't have initialize")
	}
	err := initialize(fcCtx)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(""))
}

func handle(w http.ResponseWriter, req *http.Request) {
	controlPath := req.Header.Get(fcControlPath)
	if controlPath == "/initialize" {
		initializeHandler(w, req)
	} else {
		invokeHandler(w, req)
	}
}

// Start ...
func Start(h func(ctx *FCContext, event []byte) ([]byte, error), init func(ctx *FCContext) error) {
	handler = h
	initialize = init
	fmt.Println("FunctionCompute go runtime inited.")
	initLogger()
	http.HandleFunc("/", handle)
	port := os.Getenv("FC_SERVER_PORT")
	if port == "" {
		port = "9000"
	}
	http.ListenAndServe(":"+port, nil)
}
