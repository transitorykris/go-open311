# go-open311

A golang client for talking to [Open311's API](http://wiki.open311.org/GeoReport_v2/).

## Warning

This package is a work in a progress and will be changing in backward incompatible ways without warning.

## Example

```golang
package main

import "github.com/transitorykris/go-open311"

func main() {
    client := open311.New(
        "yourapikey",
        "sfgov.org",
        "http://mobile311-dev.sfgov.org/open311/v2",
    )
}
```

## Documentation

See [https://godoc.org/github.com/transitorykris/go-keywords](https://godoc.org/github.com/transitorykris/go-open311)

## License

Copyright (c) 2017 Ahead by a Century, LLC

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