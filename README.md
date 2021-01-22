
<p align="center">
  <a href="https://travis-ci.org/mssola/user_agent" title="Travis CI status for the master branch"><img src="https://travis-ci.org/mssola/user_agent.svg?branch=master" alt="Build Status for master branch" /></a>
  <a href="http://godoc.org/github.com/mssola/user_agent" title="godoc.org page"><img src="https://godoc.org/github.com/mssola/user_agent?status.png" alt="godoc.org page" /></a>
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
    ua := user_agent.New("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11")

    fmt.Printf("%v\n", ua.Mobile())   // => false
    fmt.Printf("%v\n", ua.Bot())      // => false
    fmt.Printf("%v\n", ua.Mozilla())  // => "5.0"

    fmt.Printf("%v\n", ua.Platform()) // => "X11"
    fmt.Printf("%v\n", ua.OS())       // => "Linux x86_64"

    name, version := ua.Engine()
    fmt.Printf("%v\n", name)          // => "AppleWebKit"
    fmt.Printf("%v\n", version)       // => "537.11"

    name, version = ua.Browser()
    fmt.Printf("%v\n", name)          // => "Chrome"
    fmt.Printf("%v\n", version)       // => "23.0.1271.97"

    // Let's see an example with a bot.

    ua.Parse("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

    fmt.Printf("%v\n", ua.Bot())      // => true

    name, version = ua.Browser()
    fmt.Printf("%v\n", name)          // => Googlebot
    fmt.Printf("%v\n", version)       // => 2.1
}
```

If you want to read the full API documentation simply check
[godoc](http://godoc.org/github.com/mssola/user_agent).

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
Copyright (c) 2012-2021 Miquel Sabaté Solà

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
