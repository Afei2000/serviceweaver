package main

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver/metrics"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

var (
	addCount = metrics.NewCounter(
		"add_count",
		"The number of times Adder.Add has been called",
	)
	addConcurrent = metrics.NewGauge(
		"add_concurrent",
		"The number of concurrent Adder.Add calls",
	)
	addSum = metrics.NewHistogram(
		"add_sum",
		"The sums returned by Adder.Add",
		[]float64{1, 10, 100, 1000, 10000},
	)
)

type app struct {
	weaver.Implements[weaver.Main]
	reverser weaver.Ref[Reverser]
	hello    weaver.Listener `weaver:"hello"`
}

func serve(ctx context.Context, app *app) error {

	fmt.Printf("hello listener available on %v\n", app.hello)

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		addCount.Add(1.0)
		addConcurrent.Add(1.0)
		defer addConcurrent.Sub(1.0)
		name := request.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		app.reverser.Get().Init(ctx)

		reversed, err := app.reverser.Get().Reverse(ctx, name)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(writer, "Hello, %s!\n", reversed)
	})
	return http.Serve(app.hello, nil)

	// Call the Reverse method.
	//var r Reverser = app.reverser.Get()
	//reversed, err := r.Reverse(ctx, "!dlroW ,olleH")
	//if err != nil {
	//	return err
	//}
	//fmt.Println(reversed)
	//return nil
}
