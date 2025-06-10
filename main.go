package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "os"

    "github.com/ledongthuc/pdf"
)

func main() {
    input := flag.String("input", "", "PDF file path")
    output := flag.String("output", "", "Output JSON file path")
    flag.Parse()

    if *input == "" || *output == "" {
        fmt.Println("Usage: -input <file.pdf> -output <file.json>")
        os.Exit(1)
    }

    file, err := os.Open(*input)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
        os.Exit(1)
    }
    defer file.Close()

    fileStat, err := file.Stat()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error getting file info:", err)
        os.Exit(1)
    }

    reader, err := pdf.NewReader(file, fileStat.Size())
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error creating PDF reader:", err)
        os.Exit(1)
    }

    var content string
    for i := 1; i <= reader.NumPage(); i++ {
        page := reader.Page(i)
        text, err := page.GetPlainText(nil)
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error reading page:", err)
            continue
        }
        content += text
    }

    data := map[string]string{"content": content}
    jsonData, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
        os.Exit(1)
    }

    err = os.WriteFile(*output, jsonData, 0644)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error writing output:", err)
        os.Exit(1)
    }

    fmt.Println("PDF content successfully written to", *output)
}
