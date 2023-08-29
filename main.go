package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"tailscale.com/hostinfo"
	"tailscale.com/tsnet"
)

var (
	hostname        = flag.String("hostname", envOr("TSNET_HOSTNAME", "what"), "hostname to use on your tailnet, TSNET_HOSTNAME in the environment")
	dataDir         = flag.String("data-location", dataLocation(), "where data is stored, defaults to DATA_DIR or ~/.config/tailscale/paste")
	tsnetLogVerbose = flag.Bool("tsnet-verbose", hasEnv("TSNET_VERBOSE"), "if set, have tsnet log verbosely to standard error")
	slogLevel       = flag.String("slog-level", envOr("SLOG_LEVEL", "INFO"), "log level")

	//go:embed static
	staticFiles embed.FS // your static assets

	//go:embed tmpl/*.html
	templateFiles embed.FS // your template files
)

func hasEnv(name string) bool {
	_, ok := os.LookupEnv(name)
	return ok
}

func dataLocation() string {
	if dir, ok := os.LookupEnv("DATA_DIR"); ok {
		return dir
	}
	dir, err := os.UserConfigDir()
	if err != nil {
		return os.Getenv("DATA_DIR")
	}
	return filepath.Join(dir, "tailscale", *hostname)
}

func envOr(key, defaultVal string) string {
	if result, ok := os.LookupEnv(key); ok {
		return result
	}
	return defaultVal
}

func main() {
	flag.Parse()

	hostinfo.SetApp("what")

	// Set log level, configure logger
	var programLevel slog.Level
	if err := (&programLevel).UnmarshalText([]byte(*slogLevel)); err != nil {
		fmt.Fprintf(os.Stderr, "invalid log level %s: %v, using INFO\n", *slogLevel, err)
		programLevel = slog.LevelInfo
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     programLevel,
	})))

	// Create the data directory if it doesn't exist.
	if err := os.MkdirAll(*dataDir, 0700); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(*dataDir, "tsnet"), 0700); err != nil {
		log.Fatal(err)
	}

	// Create the tsnet server.
	s := &tsnet.Server{
		Hostname: *hostname,
		Dir:      filepath.Join(*dataDir, "tsnet"),
		Logf:     func(string, ...any) {},
	}

	if *tsnetLogVerbose {
		s.Logf = log.Printf
	}

	// Acquire any resources such as a database client.

	// Start the tsnet server.
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	lc, err := s.LocalClient()
	if err != nil {
		log.Fatal(err)
	}

	tmpls := template.Must(template.ParseFS(templateFiles, "tmpl/*.html"))
	mux := http.NewServeMux()

	mux.Handle("/static/", http.FileServer(http.FS(staticFiles)))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userInfo, err := lc.WhoIs(r.Context(), r.RemoteAddr)
		if err != nil {
			slog.Error("can't get user info", "err", err, "remoteAddr", r.RemoteAddr)
			http.Error(w, "can't get user info", http.StatusInternalServerError)
			return
		}

		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			if err := tmpls.ExecuteTemplate(w, "404.html", map[string]any{
				"Title":    "what/",
				"UserInfo": userInfo,
				"Path":     r.URL.Path,
			}); err != nil {
				slog.Error("can't execute template", "err", err, "remoteAddr", r.RemoteAddr, "path", r.URL.Path)
				return
			}
			return
		}

		if err := tmpls.ExecuteTemplate(w, "index.html", map[string]any{
			"Title":    "what/",
			"UserInfo": userInfo,
		}); err != nil {
			slog.Error("can't execute template", "err", err, "remoteAddr", r.RemoteAddr, "path", r.URL.Path)
			return
		}
	})

	ln, err := s.ListenTLS("tcp", ":443")
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("listening", "addr", ln.Addr().String())
	log.Fatal(http.Serve(ln, mux))
}
