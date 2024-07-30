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
	param       commandLineParameter
)

type commandLineParameter struct {
	freq    int64
	count   int64
	stdout  bool
	emitter string
}

func main() {
	parseCmdLineArguments()

	if param.stdout {
		log.SetOutput(os.Stdout)
	}

	emitterString := emitterName()

	if param.count > 0 {
		for ; param.count > 0; param.count-- {
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
	time.Sleep(time.Second / time.Duration(param.freq))
}

func parseCmdLineArguments() {
	var help bool
	flag.Int64Var(&param.freq, "freq", 2, "frequency of emitted fake lines (per second,  min. 1)")
	flag.Int64Var(&param.freq, "f", 2, "short for -freq")
	flag.Int64Var(&param.count, "count", 0, "number of generated fake lines (0 infinite) (default 0)")
	flag.Int64Var(&param.count, "c", 0, "short for -count (default 0)")
	flag.BoolVar(&param.stdout, "stdout", false, "log to stdout instead of stderr")
	flag.StringVar(&param.emitter, "emitter", "", "use as emitter instead of changing it randomly")
	flag.StringVar(&param.emitter, "e", "", "short for -emitter")
	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&help, "help", false, "show help")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func emitterName() func() string {
	// this is nice and saves the if in every log output.
	// Is it clear and easy to understand?
	// Fun and clever does not serve the purpose here.
	// Will be gone in the next commit.
	if param.emitter != "" {
		return func() string {
			return param.emitter
		}
	}

	return func() string {
		return app[rand.Intn(appLen)]
	}
}
