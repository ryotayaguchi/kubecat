package main

import (
    "time"
)

func prepareForStop() {
    ENABLE_READINESS = false
    time.Sleep(time.Duration(TIME_FOR_PRESTOP) * time.Second)
}
