package main

import (
    "os"
    "log"
    "strconv"
)

var ENABLE_LIVENESS bool             // set false to force liveness probe return false
var ENABLE_READINESS bool            // set false to force readiness probe return false
var TIME_UNTIL_LIVENESS_TRUE int64  // time in seconds until liveness probe starts returning true
var TIME_UNTIL_READINESS_TRUE int64 // time in seconds until readiness probe starts returning true
var TIME_FOR_PRESTOP int64          // time in seconds required for completing prestop hook
var TIME_UNTIL_AUTO_STOP int64      // time in seconds until process automatically stops
var PORT_ADMIN int64                // port number for administrative API endpoints

func getEnvVarBool(s string) bool  {
    t, err := strconv.ParseBool(os.Getenv(s))
    if (err != nil) {
        log.Fatalf("FATAL: unable to parse %s", s)
    }
    log.Printf("INFO: %s = %t", s, t)
    return t
}

func getEnvVarInt64(s string) int64 {
    d, err := strconv.ParseInt(os.Getenv(s),10,64)
    if (err != nil) {
        log.Fatalf("FATAL: unable to parse %s", s)
    }
    log.Printf("INFO: %s = %d", s, d)
    return d
}

func initConfigParameters() {
    ENABLE_LIVENESS  = getEnvVarBool("ENABLE_LIVENESS")
    ENABLE_READINESS = getEnvVarBool("ENABLE_READINESS")
    TIME_UNTIL_LIVENESS_TRUE  = getEnvVarInt64("TIME_UNTIL_LIVENESS_TRUE")
    TIME_UNTIL_READINESS_TRUE = getEnvVarInt64("TIME_UNTIL_READINESS_TRUE")
    TIME_FOR_PRESTOP = getEnvVarInt64("TIME_FOR_PRESTOP")
    TIME_UNTIL_AUTO_STOP = getEnvVarInt64("TIME_UNTIL_AUTO_STOP")
    PORT_ADMIN = getEnvVarInt64("PORT_ADMIN")
}

