package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/suchy1105/sensoremu/api"
	"github.com/suchy1105/sensoremu/sensor"
)

func main() {
	run()
}
func run() {
	m := api.NewMeasurement()

	var frontendRouter chi.Router = chi.NewRouter()
	var wg sync.WaitGroup

	frontendRouter.Route("/frontend", api.FrontendAPI(m))
	apiServer := http.Server{
		Addr:           ":8888",
		Handler:        frontendRouter,
		ReadTimeout:    360 * time.Second,
		WriteTimeout:   360 * time.Second,
		MaxHeaderBytes: 1 << 20,
		// discard error logs
		//ErrorLog: golog.New(ioutil.Discard, "", 0),
	}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		log.Log().Msgf("starting backend api server on %s", apiServer.Addr)
		err := apiServer.ListenAndServe()
		if err != nil {
			if err != http.ErrServerClosed {
				log.Warn().Err(err).Caller().Msg("error while closing api server")
			}
		}
	}(&wg)

	//go timer()

	fmt.Println("lisetner")

	fmt.Println("go")
	s := sensor.NewSensor()
	for {
		s.Getmeasurement()
		m.Measurement=s.Value
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}
