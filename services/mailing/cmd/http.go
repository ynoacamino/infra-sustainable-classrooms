package main

import (
	"context"
	"net/http"
	"net/url"
	"sync"
	"time"

	mailingsvr "github.com/ynoacamino/infra-sustainable-classrooms/services/mailing/gen/http/mailing/server"
	mailing "github.com/ynoacamino/infra-sustainable-classrooms/services/mailing/gen/mailing"
	"goa.design/clue/debug"
	"goa.design/clue/log"
	goahttp "goa.design/goa/v3/http"
)

func handleHTTPServer(ctx context.Context, u *url.URL, mailingEndpoints *mailing.Endpoints, wg *sync.WaitGroup, errc chan error, dbg bool) {
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	var mux goahttp.Muxer = goahttp.NewMuxer()
	if dbg {
		debug.MountPprofHandlers(debug.Adapt(mux))
		debug.MountDebugLogEnabler(debug.Adapt(mux))
	}

	var mailingServer *mailingsvr.Server
	eh := errorHandler(ctx)
	mailingServer = mailingsvr.New(mailingEndpoints, mux, dec, enc, eh, nil)

	mailingsvr.Mount(mux, mailingServer)

	var handler http.Handler = mux
	if dbg {
		handler = debug.HTTP()(handler)
	}
	handler = log.HTTP(ctx)(handler)

	srv := &http.Server{
		Addr:              u.Host,
		Handler:           handler,
		ReadHeaderTimeout: time.Second * 60,
		ReadTimeout:       time.Second * 60,
		WriteTimeout:      time.Second * 60,
		IdleTimeout:       time.Second * 120,
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		go func() {
			log.Printf(ctx, "HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		log.Printf(ctx, "shutting down HTTP server at %q", u.Host)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}()
}

func errorHandler(ctx context.Context) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		log.Printf(ctx, "ERROR: %s", err.Error())
		if _, ok := err.(goahttp.Statuser); !ok {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		}
	}
}
