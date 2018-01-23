package rest

import (
    "bytes"
    "fmt"
    "net/http"
    "os"
    "io/ioutil"
)

func Get(url string) string {
    fmt.Println("HTTP GET: " + url)
    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
        return ""
    }
    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    b, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
        return ""
    }
    /* Get rid of trailing new line */
    if b[len(b)-1] == '\n' {
      return string(b[:len(b)-1])
    }
    return string(b)
}

func Post(url string, dataType string, data string) string {
    fmt.Println("HTTP POST: " + url)
    resp, err := http.Post(url, dataType, bytes.NewBufferString(data))
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
        return "Failed"
    }
    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    b, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
        return ""
    }
    fmt.Println(string(b))
    return "Posted"
}

func Delete(url string) {
    fmt.Println("HTTP DELETE: " + url)
}
