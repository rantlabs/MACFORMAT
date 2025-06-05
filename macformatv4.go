package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "regexp"
    "strings"
)

// Function to convert MAC address to different formats
func formatMAC(mac string) (string, string, string) {
    normalizedMAC := strings.ToUpper(strings.ReplaceAll(mac, ":", ""))
    normalizedMAC = strings.ReplaceAll(normalizedMAC, "-", "")
    normalizedMAC = strings.ReplaceAll(normalizedMAC, ".", "")

    if len(normalizedMAC) != 12 {
        return "", "", ""
    }

    dashFormat := fmt.Sprintf("%s-%s-%s-%s-%s-%s", normalizedMAC[0:2], normalizedMAC[2:4], normalizedMAC[4:6], normalizedMAC[6:8], normalizedMAC[8:10], normalizedMAC[10:12])
    dotFormat := fmt.Sprintf("%s.%s.%s", normalizedMAC[0:4], normalizedMAC[4:8], normalizedMAC[8:12])
    colonFormat := fmt.Sprintf("%s:%s:%s:%s:%s:%s", normalizedMAC[0:2], normalizedMAC[2:4], normalizedMAC[4:6], normalizedMAC[6:8], normalizedMAC[8:10], normalizedMAC[10:12])

    return dashFormat, dotFormat, colonFormat
}

// Function to process lines containing MAC addresses
func processLine(line string) string {
    re := regexp.MustCompile(`([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})|[0-9A-Fa-f]{4}\.[0-9A-Fa-f]{4}\.[0-9A-Fa-f]{4}`)
    matches := re.FindAllString(line, -1)

    for _, mac := range matches {
        dashFormat, dotFormat, colonFormat := formatMAC(mac)
        if dashFormat != "" && dotFormat != "" && colonFormat != "" {
            line += fmt.Sprintf(" DASH: %s DOT: %s COLON: %s", dashFormat, dotFormat, colonFormat)
        }
    }

    return line
}

// Function to display usage information
func usage() {
    fmt.Println("Usage:")
    fmt.Println("  -file <filename>\tSpecify a file containing MAC addresses")
    fmt.Println("  -mac <address>\tSpecify a single MAC address")
    fmt.Println("  -output <filename>\tSpecify an output file (default is stdout)")
    fmt.Println("  STDIN\t\t\tRead MAC addresses from standard input")
}

func main() {
    fileName := flag.String("file", "", "File containing MAC addresses")
    macAddress := flag.String("mac", "", "Single MAC address to format")
    outputFile := flag.String("output", "", "Output file")
    flag.Parse()

    var output *os.File
    var err error

    if *outputFile != "" {
        output, err = os.Create(*outputFile)
        if err != nil {
            fmt.Println("Error creating output file:", err)
            return
        }
        defer output.Close()
    } else {
        output = os.Stdout
    }

    if *macAddress == "" && *fileName == "" && !isInputFromPipe() {
        fmt.Println("Error: No input provided. Please provide input via -file, -mac, or STDIN.")
        usage()
        return
    }

    writer := bufio.NewWriter(output)
    defer writer.Flush()

    if *macAddress != "" {
        dashFormat, dotFormat, colonFormat := formatMAC(*macAddress)
        if dashFormat != "" && dotFormat != "" && colonFormat != "" {
            fmt.Fprintf(writer, "DASH: %s DOT: %s COLON: %s\n", dashFormat, dotFormat, colonFormat)
        } else {
            fmt.Fprintln(writer, "Invalid MAC address format")
        }
        return
    }

    if *fileName != "" {
        file, err := os.Open(*fileName)
        if err != nil {
            fmt.Println("Error opening file:", err)
            return
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            fmt.Fprintln(writer, processLine(scanner.Text()))
        }

        if err := scanner.Err(); err != nil {
            fmt.Println("Error reading file:", err)
        }
        return
    }

    if isInputFromPipe() {
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            fmt.Fprintln(writer, processLine(scanner.Text()))
        }

        if err := scanner.Err(); err != nil {
            fmt.Println("Error reading STDIN:", err)
        }
    }
}

// Function to check if input is from a pipe
func isInputFromPipe() bool {
    fileInfo, _ := os.Stdin.Stat()
    return (fileInfo.Mode() & os.ModeCharDevice) == 0
}
