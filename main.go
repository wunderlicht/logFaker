package main

//go:generate go run generators/BOFHgenerator.go
//go:generate go fmt

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
)

var logLevel = [...]string{
	"DEBUG", "INFO", "WARN", "ERROR", "FATAL",
}

var app = [...]string{
	"httpd", "builder", "fluxcomp", "core", "coffeemaker",
}

var (
	bofhLen     = len(BOFH)
	logLevelLen = len(logLevel)
	appLen      = len(app)
	freq        int64
	count       int64
	stdout      bool
	emitter     string
)

func main() {
	parseCmdLineArguments()

	if stdout {
		log.SetOutput(os.Stdout)
	}

	emitterString := emitterName()

	rand.Seed(time.Now().UnixNano())

	if count > 0 {
		for ; count > 0; count-- {
			logAndWait(emitterString())
		}
		return //job done
	}

	//we are here? loop forever!
	for {
		logAndWait(emitterString())
	}
}

func logAndWait(emitter string) {
	log.Printf("%s|%s|%s\n", logLevel[rand.Intn(logLevelLen)], emitter, BOFH[rand.Intn(bofhLen)])
	time.Sleep(time.Second / time.Duration(freq))
}

func parseCmdLineArguments() {
	flag.Int64Var(&freq, "freq", 2, "frequency of emitted fake lines (per second,  min. 1)")
	flag.Int64Var(&freq, "f", 2, "short for -freq")
	flag.Int64Var(&count, "count", 0, "number of generated fake lines (0 infinite) (default 0)")
	flag.Int64Var(&count, "c", 0, "short for -count (default 0)")
	flag.BoolVar(&stdout, "stdout", false, "log to stdout instead of stderr")
	flag.StringVar(&emitter, "emitter", "", "use as emitter instead of changing it randomly")
	flag.StringVar(&emitter, "e", "", "short for -emitter")
	flag.Parse()
}

func emitterName() func() string {
	// this is nice and saves the if in every log output.
	// Is it clear and easy to understand?
	// Fun and clever does not serve the purpose here.
	// Will be gone in the next commit.
	if emitter != "" {
		return func() string {
			return emitter
		}
	}

	return func() string {
		return app[rand.Intn(appLen)]
	}
}
