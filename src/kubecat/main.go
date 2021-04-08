package main

import (
    "net/http"
    "log"
    "strconv"
)

func main() {
    initUptime()
    initConfigParameters()
    startLifetimeChecker()

    http.HandleFunc("/", httpHandlerRoot)
    http.HandleFunc("/liveness/", httpHandlerLiveness)
    http.HandleFunc("/readiness/", httpHandlerReadiness)
    http.HandleFunc("/prestop/", httpHandlerPrestop)
    http.HandleFunc("/cpu/", httpHandlerCpu)
    http.HandleFunc("/memory/", httpHandlerMemory)
    http.HandleFunc("/uptime/", httpHandlerUptime)
    log.Fatal(http.ListenAndServe(":" + strconv.FormatInt(PORT_ADMIN, 10), nil))
}

