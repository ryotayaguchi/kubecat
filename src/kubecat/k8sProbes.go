package main

func getLivenessState() bool {
    return (ENABLE_LIVENESS == true && getUptime() > float64(TIME_UNTIL_LIVENESS_TRUE))
}

func getReadinessState() bool {
    return (ENABLE_READINESS == true && getUptime() > float64(TIME_UNTIL_READINESS_TRUE))
}
