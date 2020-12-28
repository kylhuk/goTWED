# goTWED: Go bindings for the [Time Warp Edit Distance][1] algorithm

Go bindings for the C library of P. Marteau's TWED algorithm.

If you have any suggestions, either open an issue or a pull request with the changes. Keep in mind updating the test cases if necessary.

Have fun!

# Requirements
* You need the gcc binary in your PATH variable. On Windows, the easies way is to install TDM-GCC[2]

# Getting started
```go
go get -u github.com/kylhuk/goTWED
```

```go

```

# Citidation
Marteau, P.; F. (2009). "Time Warp Edit Distance with Stiffness Adjustment for Time Series Matching". IEEE Transactions on Pattern Analysis and Machine Intelligence. 31 (2): 306–318.

[1]: https://github.com/pfmarteau/TWED
[2]: https://github.com/jmeubank/tdm-gcc

# License
Beware, this project has two different licenses. The license for the original C code of P. Marteau (MIT License) is included in the file `twed.c_LICENSE`. The Go part is licensed under the Apache License 2.0, see `goTWED.go_LICENSE` 
