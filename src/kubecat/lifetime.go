package main

import (
    "os"
    "log"
    "time"
)

func startLifetimeChecker() {
    ticker := time.NewTicker(time.Second)
    go func() {
        for {
            select {
                case <-ticker.C:
                    if (TIME_UNTIL_AUTO_STOP != 0 && getUptime() > float64(TIME_UNTIL_AUTO_STOP)) {
                        log.Printf("INFO: I do NOT work for longer than %d sec for saving my healthy. Ciao churu!", TIME_UNTIL_AUTO_STOP)
                        prepareForStop()
                        os.Exit(0)
                    }
            }
        }
    }()
}
