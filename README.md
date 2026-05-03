# IP2C Client Module For Go

[![Go Version](https://img.shields.io/badge/Go-1.25%2B-blue)](https://go.dev/)
[![Test Status](https://github.com/alex-cos/ip2c/actions/workflows/test.yml/badge.svg)](https://github.com/alex-cos/ip2c/actions/workflows/test.yml)
[![Lint Status](https://github.com/alex-cos/ip2c/actions/workflows/lint.yml/badge.svg)](https://github.com/alex-cos/ip2c/actions/workflows/lint.yml)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/alex-cos/ip2c)](https://goreportcard.com/report/github.com/alex-cos/ip2c)


A lightweight Go client for the **ip2c.org** IP geolocation API. This module allows you to quickly resolve an IPv4 or IPv6 address to its corresponding country code and country name. It supports customizable HTTP clients, timeouts, and context-based cancellation for full control over network requests.

Official Documentation is here:
[about.ip2c.org](https://about.ip2c.org/#about)

## Install

With Go installed, you can install with command line interface:

```bash
  go get github.com/alex-cos/ip2c
```

## Usage

### Basic example

```go
package main

import (
  "fmt"
  "log"

  "github.com/alex-cos/ip2c"
)

func main() {
  client := ip2c.New()

  resp, err := client.Check("94.238.20.184")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Country: %s (%s)\n", resp.CountryName, resp.CountryCode)
}
```

### With custom timeout

```go
package main

import (
  "fmt"
  "log"
  "time"

  "github.com/alex-cos/ip2c"
)

func main() {
  client := ip2c.NewWithTimeout(10 * time.Second)

  resp, err := client.Check("94.238.20.184")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Country: %s (%s)\n", resp.CountryName, resp.CountryCode)
}
```

### With context (cancellable)

```go
package main

import (
  "context"
  "fmt"
  "log"
  "time"

  "github.com/alex-cos/ip2c"
)

func main() {
  client := ip2c.New()

  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  resp, err := client.CheckWithContext(ctx, "94.238.20.184")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Country: %s (%s)\n", resp.CountryName, resp.CountryCode)
}
```

### With custom HTTP client

```go
package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/alex-cos/ip2c"
)

func main() {
  httpClient := &http.Client{
    Transport: &http.Transport{
      MaxIdleConns: 10,
    },
  }

  client := ip2c.NewWithClient(httpClient)

  resp, err := client.Check("94.238.20.184")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Country: %s (%s)\n", resp.CountryName, resp.CountryCode)
}
```

## Error handling

```go
resp, err := client.Check("127.0.0.1")
if err != nil {
  switch {
  case errors.Is(err, ip2c.ErrLocalhost):
    log.Println("localhost check not supported")
  case errors.Is(err, ip2c.ErrInvalidIP):
    log.Println("invalid IP address")
  case errors.Is(err, ip2c.ErrNotFound):
    log.Println("IP not found in database")
  default:
    log.Fatalf("unexpected error: %v", err)
  }
}
```
