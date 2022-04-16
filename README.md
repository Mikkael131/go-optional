# go-optional

Yet another port of Java Optionals to Go.

## Benchmarks
```bash
go test -run=. -bench=. -benchtime=20s -benchmem
```

```
goos: windows
goarch: amd64
pkg: github.com/mikkael131/go-optional
cpu: 12th Gen Intel(R) Core(TM) i7-12700K
Benchmark_Get_Present-20                1000000000               0.2057 ns/op          0 B/op          0 allocs/op
Benchmark_Get_Empty-20                  1000000000               0.2116 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Present-20          1000000000               0.1110 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Empty-20            1000000000               0.1093 ns/op          0 B/op          0 allocs/op
Benchmark_IfPresent_Present-20          1000000000               1.026 ns/op           0 B/op          0 allocs/op
Benchmark_IfPresent_Empty-20            1000000000               0.2207 ns/op          0 B/op          0 allocs/op
Benchmark_Else_Present-20               1000000000               0.2077 ns/op          0 B/op          0 allocs/op
Benchmark_Else_Empty-20                 1000000000               0.2039 ns/op          0 B/op          0 allocs/op
Benchmark_ElseGet_Present-20            1000000000               0.2138 ns/op          0 B/op          0 allocs/op
Benchmark_ElseGet_Empty-20              1000000000               0.2049 ns/op          0 B/op          0 allocs/op
Benchmark_ElseErr_Present-20            1000000000               0.2376 ns/op          0 B/op          0 allocs/op
Benchmark_ElseErr_Empty-20              1000000000               0.2212 ns/op          0 B/op          0 allocs/op
Benchmark_ElseZero_Present-20           1000000000               0.2024 ns/op          0 B/op          0 allocs/op
Benchmark_ElseZero_Empty-20             1000000000               0.2016 ns/op          0 B/op          0 allocs/op
Benchmark_Filter_Present-20             1000000000               3.810 ns/op           0 B/op          0 allocs/op
Benchmark_Filter_Empty-20               1000000000               1.830 ns/op           0 B/op          0 allocs/op
Benchmark_Map_Present-20                1000000000               3.682 ns/op           0 B/op          0 allocs/op
Benchmark_Map_Empty-20                  1000000000               1.839 ns/op           0 B/op          0 allocs/op
Benchmark_FlatMap_Present-20            1000000000               3.605 ns/op           0 B/op          0 allocs/op
Benchmark_FlatMap_Empty-20              1000000000               1.640 ns/op           0 B/op          0 allocs/op
```

compared with similar packages:
- [leighmcculloch](https://github.com/leighmcculloch/go-optional) generics
- [markphelps](https://github.com/markphelps/optional) primitives
- [moznion](https://github.com/moznion/go-optional) generics
```
# Get()
Benchmark_Get_Present_mikkael131-20                     1000000000               0.2057 ns/op          0 B/op          0 allocs/op
Benchmark_Get_Present_leighmcculloch-20                 1000000000               1.483 ns/op           0 B/op          0 allocs/op
Benchmark_Get_Present_markphelps-20                     1000000000               0.1075 ns/op          0 B/op          0 allocs/op
Benchmark_Get_Present_moznion-20                        1000000000               0.2047 ns/op          0 B/op          0 allocs/op
Benchmark_Get_Empty_mikkael131-20                       1000000000               0.2116 ns/op          0 B/op          0 allocs/op
Benchmark_Get_Empty_leighmcculloch-20                   1000000000               0.9375 ns/op          0 B/op          0 allocs/op
Benchmark_Get_Empty_markphelps-20                       1000000000               0.09633 ns/op         0 B/op          0 allocs/op
Benchmark_Get_Empty_moznion-20                          1000000000               0.2078 ns/op          0 B/op          0 allocs/op

# IsPresent()
Benchmark_IsPresent_Present_mikkael131-20               1000000000               0.1110 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Present_leighmcculloch-20           1000000000               0.1064 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Present_markphelps-20               1000000000               0.1016 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Present_moznion-20                  1000000000               0.1012 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Empty_mikkael131-20                 1000000000               0.1093 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Empty_leighmcculloch-20             1000000000               0.1053 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Empty_markphelps-20                 1000000000               0.1066 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Empty_moznion-20                    1000000000               0.1038 ns/op          0 B/op          0 allocs/op

# IfPresent()
Benchmark_IfPresent_Present_mikkael131-20               1000000000               1.026 ns/op           0 B/op          0 allocs/op
Benchmark_IfPresent_Present_leighmcculloch-20           1000000000               1.227 ns/op           0 B/op          0 allocs/op
Benchmark_IfPresent_Present_markphelps-20               1000000000               0.1113 ns/op          0 B/op          0 allocs/op
Benchmark_IfPresent_Empty_mikkael131-20                 1000000000               0.2207 ns/op          0 B/op          0 allocs/op
Benchmark_IfPresent_Empty_leighmcculloch-20             1000000000               0.2159 ns/op          0 B/op          0 allocs/op
Benchmark_IfPresent_Empty_markphelps-20                 1000000000               0.1123 ns/op          0 B/op          0 allocs/op

# Else()
Benchmark_Else_Present_mikkael131-20                    1000000000               0.2077 ns/op          0 B/op          0 allocs/op
Benchmark_Else_Present_leighmcculloch-20                1000000000               2.869 ns/op           0 B/op          0 allocs/op
Benchmark_Else_Present_markphelps-20                    1000000000               0.1184 ns/op          0 B/op          0 allocs/op
Benchmark_Else_Present_moznion-20                       1000000000               0.2129 ns/op          0 B/op          0 allocs/op
Benchmark_Else_Empty_mikkael131-20                      1000000000               0.2039 ns/op          0 B/op          0 allocs/op
Benchmark_Else_Empty_leighmcculloch-20                  1000000000               2.172 ns/op           0 B/op          0 allocs/op
Benchmark_Else_Empty_markphelps-20                      1000000000               0.1073 ns/op          0 B/op          0 allocs/op
Benchmark_Else_Empty_moznion-20                         1000000000               0.2109 ns/op          0 B/op          0 allocs/op

# ElseGet()
Benchmark_ElseGet_Present_mikkael131-20                 1000000000               0.2138 ns/op          0 B/op          0 allocs/op
Benchmark_ElseGet_Present_leighmcculloch-20             1000000000               1.236 ns/op           0 B/op          0 allocs/op
Benchmark_ElseGet_Present_moznion-20                    1000000000               0.2293 ns/op          0 B/op          0 allocs/op
Benchmark_ElseGet_Empty_mikkael131-20                   1000000000               0.2049 ns/op          0 B/op          0 allocs/op
Benchmark_ElseGet_Empty_leighmcculloch-20               1000000000               1.641 ns/op           0 B/op          0 allocs/op
Benchmark_ElseGet_Empty_moznion-20                      1000000000               0.2123 ns/op          0 B/op          0 allocs/op

# ElseZero()
Benchmark_ElseZero_Present_mikkael131-20                1000000000               0.2024 ns/op          0 B/op          0 allocs/op
Benchmark_ElseZero_Present_leighmcculloch-20            1000000000               2.896 ns/op           0 B/op          0 allocs/op
Benchmark_ElseZero_Empty_mikkael131-20                  1000000000               0.2016 ns/op          0 B/op          0 allocs/op
Benchmark_ElseZero_Empty_leighmcculloch-20              1000000000               2.109 ns/op           0 B/op          0 allocs/op

# Filter()
Benchmark_Filter_Present_mikkael131-20                  1000000000               3.810 ns/op           0 B/op          0 allocs/op
Benchmark_Filter_Present_moznion-20                     1000000000               4.020 ns/op           0 B/op          0 allocs/op
Benchmark_Filter_Empty_mikkael131-20                    1000000000               1.830 ns/op           0 B/op          0 allocs/op
Benchmark_Filter_Empty_moznion-20                       1000000000               1.803 ns/op           0 B/op          0 allocs/op
```