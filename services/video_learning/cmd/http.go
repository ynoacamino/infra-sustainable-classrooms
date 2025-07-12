package main

import (
	"context"
	"net/http"
	"net/url"
	"sync"
	"time"

	videolearningsvr "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/http/video_learning/server"
	videolearning "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/video_learning"
	"goa.design/clue/debug"
	"goa.design/clue/log"
	goahttp "goa.design/goa/v3/http"
)

func handleHTTPServer(ctx context.Context, u *url.URL, videoLearningEndpoints *videolearning.Endpoints, wg *sync.WaitGroup, errc chan error, dbg bool) {
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	var mux goahttp.Muxer = goahttp.NewMuxer()
	if dbg {
		debug.MountPprofHandlers(debug.Adapt(mux))
		debug.MountDebugLogEnabler(debug.Adapt(mux))
	}

	var videoLearningServer *videolearningsvr.Server
	eh := errorHandler(ctx)
	videoLearningServer = videolearningsvr.New(videoLearningEndpoints, mux, dec, enc, eh, nil, nil, nil)

	videolearningsvr.Mount(mux, videoLearningServer)

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

	for _, m := range videoLearningServer.Mounts {
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
