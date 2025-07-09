package main

import (
	"context"
	"net/http"
	"net/url"
	"sync"
	"time"

	auth "github.com/ynoacamino/infrastructure/services/auth/gen/auth"
	authsvr "github.com/ynoacamino/infrastructure/services/auth/gen/http/auth/server"
	"goa.design/clue/debug"
	"goa.design/clue/log"
	goahttp "goa.design/goa/v3/http"
)

func handleHTTPServer(ctx context.Context, u *url.URL, authEndpoints *auth.Endpoints, wg *sync.WaitGroup, errc chan error, dbg bool) {
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	var mux goahttp.Muxer
	mux = goahttp.NewMuxer()
	if dbg {
		debug.MountPprofHandlers(debug.Adapt(mux))
		debug.MountDebugLogEnabler(debug.Adapt(mux))
	}

	var authServer *authsvr.Server
	eh := errorHandler(ctx)
	authServer = authsvr.New(authEndpoints, mux, dec, enc, eh, nil)

	authsvr.Mount(mux, authServer)

	var handler http.Handler = mux
	if dbg {
		handler = debug.HTTP()(handler)
	}
	handler = log.HTTP(ctx)(handler)
	// handler = addReverseProxyHeaders(handler)

	srv := &http.Server{
		Addr:              u.Host,
		Handler:           handler,
		ReadHeaderTimeout: time.Second * 60,
		ReadTimeout:       time.Second * 60,
		WriteTimeout:      time.Second * 60,
		IdleTimeout:       time.Second * 120,
	}

	for _, m := range authServer.Mounts {
		log.Printf(ctx, "HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
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

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			log.Printf(ctx, "failed to shutdown: %v", err)
		}
	}()
}

func errorHandler(logCtx context.Context) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		log.Printf(logCtx, "ERROR: %s", err.Error())
	}
}

// func addReverseProxyHeaders(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if realIP := r.Header.Get("X-Forwarded-For"); realIP != "" {
// 			r.Header.Set("X-Real-IP", realIP)
// 		}

// 		w.Header().Set("X-Content-Type-Options", "nosniff")
// 		w.Header().Set("X-Frame-Options", "DENY")
// 		w.Header().Set("X-XSS-Protection", "1; mode=block")
// 		w.Header().Set("Vary", "Origin, Accept-Encoding")

// 		next.ServeHTTP(w, r)
// 	})
// }
