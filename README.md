Simple math with dates
======================

[![Build Status](https://travis-ci.org/wuzuf/datemath.svg?branch=master)](https://travis-ci.org/wuzuf/datemath) 
[![Go Report Card](https://goreportcard.com/badge/github.com/wuzuf/datemath)](https://goreportcard.com/report/github.com/wuzuf/datemath) 
[![Code Coverage](https://scrutinizer-ci.com/g/wuzuf/datemath/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/wuzuf/datemath/?branch=master) 
[![GoDoc](https://godoc.org/github.com/wuzuf/datemath?status.svg)](https://godoc.org/github.com/wuzuf/datemath)

Inspired from https://www.elastic.co/guide/en/elasticsearch/reference/current/common-options.html#date-math

## Getting started
Install:
```
go get github.com/wuzuf/datemath
```

Import:
```go
import "github.com/wuzuf/datemath"
```

## Example
```go
package main

import (
	"fmt"
	"time"

	"github.com/wuzuf/datemath"
)

func main() {
	now := datemath.Eval("now")
	today := datemath.Eval("now/d")
	thismonday := datemath.Eval("now/w")
	lastmonday := datemath.Eval("now/w-w")
	fmt.Printf("Now: %s\n", now.Format(time.RubyDate))
	fmt.Printf("Today: %s\n", today.Format(time.RubyDate))
	fmt.Printf("This Monday: %s\n", thismonday.Format(time.RubyDate))
	fmt.Printf("Last Monday: %s\n", lastmonday.Format(time.RubyDate))
}
```

Checkout the tests for more examples.

## Contributing
Please feel free to make suggestions, create issues, fork the repository and send pull requests!

## License

Copyright (c) 2019, Gabriel de Labachelerie
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:
    * Redistributions of source code must retain the above copyright
      notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
      notice, this list of conditions and the following disclaimer in the
      documentation and/or other materials provided with the distribution.
    * Neither the name of datemath nor the
      names of its contributors may be used to endorse or promote products
      derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.