package main

//go:generate go run generators/BOFHgenerator.go
//go:generate go fmt

import (
	"flag"
	"log"
	"math/rand"
	"time"
)

var logLevel = [...]string{
	"DEBUG", "INFO", "WARN", "ERROR", "FATAL",
}

var app = [...]string{
	"httpd", "builder", "fluxcomp", "core", "coffeemaker",
}

func main() {
	var (
		bofhLen     = len(BOFH)
		logLevelLen = len(logLevel)
		appLen      = len(app)
		freq        int64
		count       int64
	)
	flag.Int64Var(&freq, "freq", 2, "frequency of emitted fake lines (per second,  min. 1)")
	flag.Int64Var(&count, "c", 0, "number of generated fake lines (0 infinite)")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	if count > 0 {
		for ; count > 0; count-- {
			log.Printf("%s|%s|%s\n", logLevel[rand.Intn(logLevelLen)], app[rand.Intn(appLen)], BOFH[rand.Intn(bofhLen)])
			time.Sleep(time.Second / time.Duration(freq))
		}
		return //job done
	}

	//we are here? loop forever!
	for {
		log.Printf("%s|%s|%s\n", logLevel[rand.Intn(logLevelLen)], app[rand.Intn(appLen)], BOFH[rand.Intn(bofhLen)])
		time.Sleep(time.Second / time.Duration(freq))
	}
}
