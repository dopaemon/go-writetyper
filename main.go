package writetyper

import (
    "fmt"
    "time"
)

func writetyper(text string) {
    for _, char := range text {
        fmt.Print(string(char))
        time.Sleep(50 * time.Millisecond)
    }
    fmt.Println()
}
