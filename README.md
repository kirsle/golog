# Go Log

This is Yet Another Logger for Go programs.

![Screenshot](https://raw.githubusercontent.com/kirsle/golog/master/screenshot.png)

This is a logging package designed for local interactive shells running text
based Go programs. To that end, this prints colorful log lines with customizable
themes.

The color options for the log lines are `NoColor` (default), `ANSIColor`
which limits the color codes to the standard 16 ANSI colors, and
`ExtendedColor` which supports the 256-color palette of `xterm` and other
modern terminal emulators. The theming engine supports defining colors using
hex codes, supported by [tomnomnom/xtermcolor](https://github.com/tomnomnom/xtermcolor).

This module is still a work in progress and will be extended and improved as I
use it for other personal Go projects.

# Usage

```go
package main

import "github.com/kirsle/golog"

var log golog.Logger

func init() {
    // Get a named logger and configure it. Note: you can call GetLogger any
    // number of times from any place in your codebase. It implements the
    // singleton pattern.
    log = golog.GetLogger("main")
    log.Configure(&golog.Config{
        Colors: golog.ExtendedColor,
        Theme: golog.DarkTheme,
    })
}

func main() {
    // The log functions work like `fmt.Printf`
    log.Debug("Running on %s", runtime.GOOS)
    log.Info("Hello, world!")
}
```

# License

```
The MIT License (MIT)

Copyright (c) 2017 Noah Petherbridge

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
