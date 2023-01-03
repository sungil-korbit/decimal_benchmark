```
cd ./benches 
go test -bench=. -benchtime=1x    
```

```text
BenchmarkAdd10k-10                             1            777667 ns/op
BenchmarkAddSelf-10                            1              9667 ns/op
BenchmarkAddSimple-10                          1              9208 ns/op
BenchmarkAddOne-10                             1             19958 ns/op
BenchmarkAddTwo-10                             1              8791 ns/op
BenchmarkAddOneHundred-10                      1              7584 ns/op
BenchmarkAddPointZeroOne-10                    1              5834 ns/op
BenchmarkAddNegativePointFive-10               1              8417 ns/op
BenchmarkAddPi-10                              1              7500 ns/op
BenchmarkAddNegativePi-10                      1              5709 ns/op
BenchmarkSub10k-10                             1            725875 ns/op
BenchmarkSubSelf-10                            1              3750 ns/op
BenchmarkSubSimple-10                          1              1041 ns/op
BenchmarkSubOne-10                             1              5042 ns/op
BenchmarkSubTwo-10                             1              3208 ns/op
BenchmarkSubOneHundred-10                      1              4833 ns/op
BenchmarkSubPointZeroOne-10                    1              4000 ns/op
BenchmarkSubNegativePointFive-10               1              3292 ns/op
BenchmarkSubPi-10                              1              5416 ns/op
BenchmarkSubNegativePi-10                      1              5125 ns/op
BenchmarkDiv10k-10                             1           3284333 ns/op
BenchmarkDivSelf-10                            1              7042 ns/op
BenchmarkDivSimple-10                          1              6959 ns/op
BenchmarkDivOne-10                             1              7125 ns/op
BenchmarkDivTwo-10                             1              9792 ns/op
BenchmarkDivOneHundred-10                      1              6459 ns/op
BenchmarkDivPointZeroOne-10                    1              6125 ns/op
BenchmarkDivNegativePointFive-10               1              6208 ns/op
BenchmarkDivPi-10                              1             12292 ns/op
BenchmarkDivNegativePi-10                      1             10333 ns/op
BenchmarkDivUnderflow-10                       1              8917 ns/op
BenchmarkMul10k-10                             1              6000 ns/op
BenchmarkMulSelf-10                            1              8792 ns/op
BenchmarkMulSimple-10                          1              5459 ns/op
BenchmarkMulOne-10                             1              3500 ns/op
BenchmarkMulTwo-10                             1              3250 ns/op
BenchmarkMulOneHundred-10                      1              3250 ns/op
BenchmarkMulPointZeroOne-10                    1              3125 ns/op
BenchmarkMulNegativePointFive-10               1              3125 ns/op
BenchmarkMulPi-10                              1              3416 ns/op
BenchmarkMulNegativePi-10                      1              3083 ns/op

```