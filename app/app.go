package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	db := "db"
	nsFile := "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
	if buf, _ := ioutil.ReadFile(nsFile); len(buf) > 0 {
		db = string(buf)
	}

	url := "http://cos.fun.cloud.ibm.com/" + db + "/customers"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, _ := http.Get(url)

		count := "None yet"
		if res != nil {
			buf, _ := ioutil.ReadAll(res.Body)
			if len(buf) > 0 {
				count = string(buf)
			}
		}

		fmt.Fprintf(w, "Customers: %s\n", count)
	})

	http.ListenAndServe(":8080", nil)
}
