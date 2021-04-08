package main

import (
    "net/http"
    "io"
    "fmt"
    "log"
)

func httpHandlerLiveness(w http.ResponseWriter, _ *http.Request) {
    probe := getLivenessState()
    log.Printf("liveness: %t\n",probe)
    status := http.StatusOK
    body := "I'm alive. Please don't put me in a microwave.\n"
    if probe == false {
        status = http.StatusServiceUnavailable
        body = "I'm dead. Please don't talk to me.\n"
    }
    w.WriteHeader(status)
    io.WriteString(w, body)
}

func httpHandlerReadiness(w http.ResponseWriter, _ *http.Request) {
    probe := getReadinessState()
    log.Printf("readiness: %t\n",probe)
    status := http.StatusOK
    body := "I'm ready. You are allowed to bring my lunch here.\n"
    if probe == false {
        status = http.StatusServiceUnavailable
        body = "I'm unready. You lucky mouse. ignored at the moment.\n"
    }
    w.WriteHeader(status)
    io.WriteString(w, body)
}

func httpHandlerPrestop(w http.ResponseWriter, _ *http.Request) {
    log.Printf("starting preStop Hook\n")
    prepareForStop()
    log.Printf("finished preStop Hook\n")
    status := http.StatusOK
    body := "ok. I'm ready for running away.\n"
    w.WriteHeader(status)
    io.WriteString(w, body)
}

func httpHandlerUptime(w http.ResponseWriter, _ *http.Request) {
    uptime := getUptime()
    log.Printf("uptime: %f\n",uptime)
    status := http.StatusOK
    body := fmt.Sprintf("I already worked seriously for %0.0f sec\n",uptime)
    w.WriteHeader(status)
    io.WriteString(w, body)
}

func httpHandlerCpu(w http.ResponseWriter, _ *http.Request) {
    overloadCpu()
    status := http.StatusOK
    body := "ok. I ignited your expensive processors.\n"
    w.WriteHeader(status)
    io.WriteString(w, body)
}

func httpHandlerMemory(w http.ResponseWriter, _ *http.Request) {
    overloadMemory()
    status := http.StatusOK
    body := "ok. I scattered garbage in your memory space.\n"
    w.WriteHeader(status)
    io.WriteString(w, body)
}

func httpHandlerRoot(w http.ResponseWriter, _ *http.Request) {
    status := http.StatusOK
    body := "k ... meow\n"
    w.WriteHeader(status)
    io.WriteString(w, body)
}
