<p align="center">
  <a href="https://github.com/mssola/user_agent/actions/workflows/ci.yml" title="Travis CI status for the default branch"><img src="https://github.com/mssola/user_agent/actions/workflows/ci.yml/badge.svg" alt="Build Status for the default branch" /></a>
  <a href="https://pkg.go.dev/github.com/mssola/user_agent" title="go.dev page"><img src="https://pkg.go.dev/badge/github.com/mssola/user_agent" alt="go.dev page" /></a>
  <a href="https://en.wikipedia.org/wiki/MIT_License" rel="nofollow"><img alt="MIT" src="https://img.shields.io/badge/license-MIT-blue.svg" style="max-width:100%;"></a>
</p>

---

UserAgent is a Go library that parses HTTP User Agents. As an example:

```go
package main

import (
    "fmt"

    "github.com/mssola/user_agent"
)

func main() {
    // The "New" function will create a new UserAgent object and it will parse
    // the given string. If you need to parse more strings, you can re-use
    // this object and call: ua.Parse("another string")
    ua := user_agent.New("Mozilla/5.0 (Linux; U; Android 2.3.7; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1")

    fmt.Printf("%v\n", ua.Mobile())   // => true
    fmt.Printf("%v\n", ua.Bot())      // => false
    fmt.Printf("%v\n", ua.Mozilla())  // => "5.0"
    fmt.Printf("%v\n", ua.Model())    // => "Nexus One"

    fmt.Printf("%v\n", ua.Platform()) // => "Linux"
    fmt.Printf("%v\n", ua.OS())       // => "Android 2.3.7"

    name, version := ua.Engine()
    fmt.Printf("%v\n", name)          // => "AppleWebKit"
    fmt.Printf("%v\n", version)       // => "533.1"

    name, version = ua.Browser()
    fmt.Printf("%v\n", name)          // => "Android"
    fmt.Printf("%v\n", version)       // => "4.0"

    // Let's see an example with a bot.

    ua.Parse("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

    fmt.Printf("%v\n", ua.Bot())      // => true

    name, version = ua.Browser()
    fmt.Printf("%v\n", name)          // => Googlebot
    fmt.Printf("%v\n", version)       // => 2.1
}
```

If you want to read the full API documentation simply check
[godoc](https://pkg.go.dev/github.com/mssola/user_agent).

## Installation

```
go get -u github.com/mssola/user_agent
```

## Contributing

Do you want to contribute with code, or to report an issue you are facing? Read
the [CONTRIBUTING.md](./CONTRIBUTING.md) file.

## [Changelog](https://pbs.twimg.com/media/DJDYCcLXcAA_eIo?format=jpg&name=small)

Read the [CHANGELOG.md](./CHANGELOG.md) file.

## License

```
Copyright (c) 2012-2023 Miquel Sabaté Solà

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
