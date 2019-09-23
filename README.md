# Overview

Custom Runtime go sample

[Custom Runtime manual](https://help.aliyun.com/document_detail/132044.html)

[基于custom runtime 打造 golang runtime](https://help.aliyun.com/document_detail/132053.html)

## Normal Invoke(No-HTTP-Trigger)

For No-HTTP-Trigger invoke, we provide a simple framework that favors your rapid development.

sample code:

```go
package main

import (
	"encoding/json"
	gr "github.com/awesome-fc/golang-runtime"
)

func initialize(ctx *gr.FCContext) error {
	fcLogger := gr.GetLogger().WithField("requestId", ctx.RequestID)
	fcLogger.Infoln("init golang!")
	return nil
}

func handler(ctx *gr.FCContext, event []byte) ([]byte, error) {
	fcLogger := gr.GetLogger().WithField("requestId", ctx.RequestID)
	b, err := json.Marshal(ctx)
	if err != nil {
		fcLogger.Error("error:", err)
	}
	fcLogger.Infof("hello golang! \ncontext = %s", string(b))
	return event, nil
}

func main() {
	gr.Start(handler, initialize)
}
```


## HTTP Trigger Invoke

Just implementing an HTTP server, Start the server with port = os.Getenv("FC_SERVER_PORT")

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func aHandler(w http.ResponseWriter, req *http.Request) {
	requestID := req.Header.Get("x-fc-request-id")
	fmt.Println(fmt.Sprintf("FC Invoke Start RequestId: %s", requestID))

	defer func() {
		fmt.Println(fmt.Sprintf(fcLogTailEndPrefix, requestID))
	}()

	// your logic
	w.Write([]byte(fmt.Sprintf("Hello, golang  http invoke!")))
}

func bHandler(w http.ResponseWriter, req *http.Request) {
	requestID := req.Header.Get("x-fc-request-id")
	fmt.Println(fmt.Sprintf("FC Invoke Start RequestId: %s", requestID))

	defer func() {
		fmt.Println(fmt.Sprintf(fcLogTailEndPrefix, requestID))
	}()

	// your logic
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	info := fmt.Sprintf("method =  %+v;\nheaders = %+v;\nbody = %+v", req.Method, req.Header, string(b))
	w.Write([]byte(fmt.Sprintf("Hello, golang  http invoke! detail:\n %s", info)))
}

func main() {
	fmt.Println("FunctionCompute go runtime inited.")
	http.HandleFunc("/a", aHandler) // 如果不使用自定义域名，则 path 为 /2016-08-15/proxy/$serviceName/$functionName/a
	http.HandleFunc("/b", bHandler)
	port := os.Getenv("FC_SERVER_PORT")
	if port == "" {
		port = "9000"
	}
	http.ListenAndServe(":"+port, nil)
}
```
