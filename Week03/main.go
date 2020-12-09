package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
    "os"
    "os/signal"
    "syscall"
	"time"
)

func main() {
    g := new(errgroup.Group)

    // initialize http server metadata
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		fmt.Println("r.Method = ", r.Method)
		fmt.Println("r.URL = ", r.URL)
		fmt.Println("r.Header = ", r.Header)
		fmt.Println("r.Body = ", r.Body)
	})

	//http server definition
	serv := &http.Server{
			Addr:    "0.0.0.0:8880",
			Handler: mux,
	}

	//context definition
	ctx := context.Background()

	//signal definition
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    //signal goroutine , mark application done
    go func() {
        sig := <-sigs
        fmt.Println("receive system signal :",sig)

        done <- true
    }()

    g.Go(func() error {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()

			err := serv.ListenAndServe()

			return err
    })


	// Wait for all http server start
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully start all http server.")
	}

	select {
	case <-done:
		fmt.Println("exiting")
		serv.Shutdown(ctx);
		fmt.Println("exit")
	}
}
