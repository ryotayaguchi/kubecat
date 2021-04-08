package main

import (
    "time"
)

var startTime time.Time

func initUptime() {
    startTime = time.Now()
}

func getUptime() float64 {
    return time.Since(startTime).Seconds()
}

