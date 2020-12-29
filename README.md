![Go](https://github.com/kylhuk/goTWED/workflows/Go/badge.svg) [![codecov](https://codecov.io/gh/kylhuk/goTWED/branch/main/graph/badge.svg?token=V56U12ZLT1)](https://codecov.io/gh/kylhuk/goTWED)
# goTWED: Go bindings for the [Time Warp Edit Distance][1] algorithm

Go bindings for the C library of P. Marteau's TWED algorithm.

If you have any suggestions, either open an issue or a pull request with the changes. Keep in mind updating the test cases if necessary.

Have fun!

# Requirements
You need the gcc binary in your PATH variable. On Windows, the easiest way is to install [TDM-GCC][2].

# Getting started
```go
go get -u github.com/kylhuk/goTWED
```

I have added an example file to get you started as fast as possible. It is stored under `example/example.go`
```go
package main

import (
	"fmt"
	"github.com/kylhuk/goTWED"
)

func main() {

	ta := []float64{0, 0, 1, 1, 2, 3, 5, 2, 0, 1, -0.1}
	tsa := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tb := []float64{0, 1, 1, 1, 2, 3, 5, 2, 0, 1, -0.1}
	tsb := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nu := float64(1.1)
	lambda := float64(2.2)
	degree := int32(1)

	dist := goTWED.TWED(ta, tsa, tb, tsb, nu, lambda, degree)

	if dist != -1 {
		fmt.Println("The calculated distance is: ", dist)
	}
}
```

# Citidation
Marteau, P.; F. (2009). "Time Warp Edit Distance with Stiffness Adjustment for Time Series Matching". IEEE Transactions on Pattern Analysis and Machine Intelligence. 31 (2): 306–318.

[1]: https://github.com/pfmarteau/TWED
[2]: https://github.com/jmeubank/tdm-gcc

# License
Beware, this project has two different licenses. The license for the original C code of P. Marteau (MIT License) is included in the file `twed.c_LICENSE`. The Go part is licensed under the Apache License 2.0, see `goTWED.go_LICENSE` 
