package main

import (
    "fmt"
    "os"
    "flag"
    "log"
    "net/http"
)


type SrvCtx struct {
    Static string
    Port int
}

func buildIf(path string) (err error) {
    if _, err := os.Stat(path); os.IsNotExist(err) {
        if err := os.MkdirAll(path, 0755); err != nil {
            return err
        }
    } else if err != nil {
        return err
    }

    return nil
}

func (slf *SrvCtx) Populate(static string, port int) (err error) {

    if port > 65535 || port < 0 {
        return fmt.Errorf("[x] Invalid port number provided! %d", port)
    }

    slf.Port = port

    if err := buildIf(static); err != nil {
        return err
    }

    slf.Static = static

    return nil
}




func (slf *SrvCtx) Run() (err error) {
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", slf.Port),  http.FileServer(http.Dir(slf.Static))))
    return nil
}

func main() {
    var portnum int
    var staticFolder string
    var ctx SrvCtx

    flag.IntVar(&portnum, "p", 3030, "Port number to bind server to")
    flag.StringVar(&staticFolder, "s", "static", "Folder containing content to serve")
    flag.Parse()

    if err := ctx.Populate(staticFolder, portnum); err != nil {
        fmt.Println(err)
        return;
    }


    ctx.Run()

}
