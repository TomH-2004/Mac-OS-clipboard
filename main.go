package main

import (
    "fmt"
    "time"

    "github.com/atotto/clipboard"
    "github.com/gen2brain/beeep"
)

func main() {
    var lastClipboardContent string

    for {
        clipboardContent, err := clipboard.ReadAll()
        if err != nil {
            fmt.Println("Error reading clipboard:", err)
        }

        if clipboardContent != lastClipboardContent {
            fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
            fmt.Println()
            fmt.Println(clipboardContent)
            fmt.Println()

 
            err := beeep.Notify("Clipboard Updated", clipboardContent, "")
            if err != nil {
                fmt.Println("Error notifying:", err)
            }

            lastClipboardContent = clipboardContent
        }

        time.Sleep(time.Second * 1)
    }
}
