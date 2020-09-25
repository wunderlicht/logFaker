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

#A trivial tool to fake log output on STDOUT

```
Usage of logFaker:
  -c int
        number of generated fake lines (0 infinite)
  -freq int
        frequency of emitted fake lines (per second,  min. 1) (default 2)

```
Examples of generated log lines.
```
fluxcomp|DEBUG|Chewing gum on /dev/sd3c
coffeemaker|WARN|Fanout dropping voltage too much, try cutting some of those little traces
core|FATAL|Traceroute says that there is a routing problem in the backbone.  It's not our problem.
core|ERROR|NOTICE: alloc: /dev/null: filesystem full
coffeemaker|INFO|High line impedance.
httpd|FATAL|(l)user error
fluxcomp|FATAL|Just type 'mv * /dev/null'.
[..]
```

It shuffels one of the 'systems'  
httpd, builder, fluxcomp, core, or coffeemaker  
with one of the levels  
DEBUG, INFO, WARN, ERROR, or FATAL  
and one line from the [BOFH excuses](http://pages.cs.wisc.edu/~ballard/bofh/excuses)