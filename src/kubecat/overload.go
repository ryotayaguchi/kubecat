package main

func overloadCpu() {
    for i := 0; i < 128; i++ {
        go func() {
            for {}
        }()
    }
}

func overloadMemory() {
    var m [][]byte
    for i := 0; i < 1024; i++ {
        m = append(m, make([]byte, 0, 1024*1024*1024))
    }
}

