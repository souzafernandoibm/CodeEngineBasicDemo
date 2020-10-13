package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	db := "db"
	nsFile := "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
	if buf, _ := ioutil.ReadFile(nsFile); len(buf) > 0 {
		db = string(buf)
	}

	url := "http://cos.fun.cloud.ibm.com/" + db + "/customers?inc"

	for i := 0; i < 100; i++ {
		req, _ := http.NewRequest("PUT", url, nil)
		(&http.Client{}).Do(req)
	}
}
