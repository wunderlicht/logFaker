```
  ,,                                                                      
`7MM                      `7MM"""YMM          `7MM                        
  MM                        MM    `7            MM                        
  MM   ,pW"Wq.   .P"Ybmmm   MM   d    ,6"Yb.    MM  ,MP' .gP"Ya  `7Mb,od8 
  MM  6W'   `Wb :MI  I8     MM""MM   8)   MM    MM ;Y   ,M'   Yb   MM' "' 
  MM  8M     M8  WmmmP"     MM   Y    ,pm9MM    MM;Mm   8M""""""   MM     
  MM  YA.   ,A9 8M          MM       8M   MM    MM `Mb. YM.    ,   MM     
.JMML. `Ybmd9'   YMMMMMb  .JMML.     `Moo9^Yo..JMML. YA. `Mbmmd' .JMML.   
                6'     dP                                                 
                Ybmmmd'                                                   
```

# A trivial tool to fake log output on STDERR or STDOUT

It shuffels one of the 'systems' (emitters)  
httpd, builder, fluxcomp, core, or coffeemaker  
with one of the levels  
DEBUG, INFO, WARN, ERROR, or FATAL  
and one line from the [BOFH excuses](http://pages.cs.wisc.edu/~ballard/bofh/excuses)

## Build
```
go generate
go build
```
`generate` pulls the excuses from the web and builds `BOFHexcuses.go`.
An internet connection is required for building
logFaker. Using logFaker works offline, though.

## Usage
```
logFaker [options]
options:
    -c int
          short for -count (default 0)
    -count int
          number of generated fake lines (0 infinite) (default 0)
    -e string
          short for -emitter
    -emitter string
          use as emitter instead of changing it randomly
    -f int
          short for -freq (default 2)
    -freq int
          frequency of emitted fake lines (per second,  min. 1) (default 2)
    -stdout
          log to stdout instead of stderr
```
## Examples of generated log lines.
```
2020/10/28 09:49:27 WARN|coffeemaker|halon system went off and killed the operators.
2020/10/28 09:49:27 INFO|fluxcomp|fat electrons in the lines
2020/10/28 09:49:28 INFO|coffeemaker|Chewing gum on /dev/sd3c
2020/10/28 09:49:28 INFO|builder|Stale file handle (next time use Tupperware(tm)!)
2020/10/28 09:49:29 FATAL|core|the daemons! the daemons! the terrible daemons!
2020/10/28 09:49:29 FATAL|coffeemaker|It's not plugged in.
2020/10/28 09:49:30 FATAL|builder|There isn't any problem
2020/10/28 09:49:30 INFO|fluxcomp|Hot Java has gone cold
2020/10/28 09:49:31 INFO|httpd|(l)user error
[..]
```

