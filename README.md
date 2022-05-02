# go-optional 

[![Software License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
![Go version](https://img.shields.io/github/go-mod/go-version/mikkael131/go-optional)
[![Release](https://img.shields.io/github/release/mikkael131/go-optional.svg)](https://github.com/mikkael131/go-optional/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/mikkael131/go-optional)](https://goreportcard.com/report/github.com/mikkael131/go-optional)
[![codecov](https://codecov.io/gh/mikkael131/go-optional/branch/main/graph/badge.svg?token=UUSWVIW16A)](https://codecov.io/gh/mikkael131/go-optional)
[![Github workflow](https://github.com/mikkael131/go-optional/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/mikkael131/go-optional/actions/workflows/go.yml)

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
Benchmark_Init_Present-20                       1000000000               0.09934 ns/op         0 B/op          0 allocs/op
Benchmark_Init_Empty-20                         1000000000               0.1068 ns/op          0 B/op          0 allocs/op
Benchmark_Init_OfPtr_Value-20                   1000000000               0.1037 ns/op          0 B/op          0 allocs/op
Benchmark_Init_OfPtr_Nil-20                     1000000000               0.1021 ns/op          0 B/op          0 allocs/op
Benchmark_Get_Present-20                        1000000000               0.2047 ns/op          0 B/op          0 allocs/op
Benchmark_Get_Empty-20                          1000000000               0.2097 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Present-20                  1000000000               0.1110 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Empty-20                    1000000000               0.1057 ns/op          0 B/op          0 allocs/op
Benchmark_IfPresent_Present-20                  1000000000               1.045 ns/op           0 B/op          0 allocs/op
Benchmark_IfPresent_Empty-20                    1000000000               0.2312 ns/op          0 B/op          0 allocs/op
Benchmark_IfPresentOrElse_Present-20            1000000000               1.842 ns/op           0 B/op          0 allocs/op
Benchmark_IfPresentOrElse_Empty-20              1000000000               1.813 ns/op           0 B/op          0 allocs/op
Benchmark_Else_Present-20                       1000000000               0.2040 ns/op          0 B/op          0 allocs/op
Benchmark_Else_Empty-20                         1000000000               0.2290 ns/op          0 B/op          0 allocs/op
Benchmark_ElseGet_Present-20                    1000000000               0.2144 ns/op          0 B/op          0 allocs/op
Benchmark_ElseGet_Empty-20                      1000000000               0.2227 ns/op          0 B/op          0 allocs/op
Benchmark_ElseErr_Present-20                    1000000000               0.2041 ns/op          0 B/op          0 allocs/op
Benchmark_ElseErr_Empty-20                      1000000000               0.2138 ns/op          0 B/op          0 allocs/op
Benchmark_ElseZero_Present-20                   1000000000               0.2282 ns/op          0 B/op          0 allocs/op
Benchmark_ElseZero_Empty-20                     1000000000               0.2112 ns/op          0 B/op          0 allocs/op
Benchmark_Filter_Present-20                     1000000000               3.902 ns/op           0 B/op          0 allocs/op
Benchmark_Filter_Empty-20                       1000000000               1.856 ns/op           0 B/op          0 allocs/op
Benchmark_Map_Present-20                        1000000000               3.955 ns/op           0 B/op          0 allocs/op
Benchmark_Map_Empty-20                          1000000000               1.825 ns/op           0 B/op          0 allocs/op
Benchmark_FlatMap_Present-20                    1000000000               3.900 ns/op           0 B/op          0 allocs/op
Benchmark_FlatMap_Empty-20                      1000000000               1.657 ns/op           0 B/op          0 allocs/op
Benchmark_Ptr_Present-20                        1000000000               0.2555 ns/op          0 B/op          0 allocs/op
Benchmark_Ptr_Empty-20                          1000000000               0.2431 ns/op          0 B/op          0 allocs/op
Benchmark_Val_Present-20                        1000000000               0.1129 ns/op          0 B/op          0 allocs/op
Benchmark_Val_Empty-20                          1000000000               0.1041 ns/op          0 B/op          0 allocs/op
Benchmark_MarshalJSON_Present-20                303242358               79.23 ns/op           32 B/op          2 allocs/op
Benchmark_MarshalJSON_Empty-20                  1000000000               1.034 ns/op           0 B/op          0 allocs/op
Benchmark_UnmarshalJSON_String-20               140326406              170.0 ns/op           264 B/op          7 allocs/op
Benchmark_UnmarshalJSON_EmptyString-20          365910396               65.89 ns/op          184 B/op          3 allocs/op
Benchmark_UnmarshalJSON_NullString-20           1000000000               2.459 ns/op           0 B/op          0 allocs/op
```

compared with similar packages:
- [leighmcculloch](https://github.com/leighmcculloch/go-optional) generics
- [markphelps](https://github.com/markphelps/optional) primitives
- [moznion](https://github.com/moznion/go-optional) generics
```
# Init
Benchmark_Init_Present_mikkael131-20                    1000000000               0.09934 ns/op         0 B/op          0 allocs/op
Benchmark_Init_Present_leighmcculloch-20                1000000000               0.09648 ns/op         0 B/op          0 allocs/op
Benchmark_Init_Present_markphelps-20                    1000000000               0.1072 ns/op          0 B/op          0 allocs/op
Benchmark_Init_Present_moznion-20                       1000000000               0.1013 ns/op          0 B/op          0 allocs/op
Benchmark_Init_Empty_mikkael131-20                      1000000000               0.1068 ns/op          0 B/op          0 allocs/op
Benchmark_Init_Empty_leighmcculloch-20                  1000000000               0.1045 ns/op          0 B/op          0 allocs/op
Benchmark_Init_Empty_markphelps-20                      1000000000               0.1098 ns/op          0 B/op          0 allocs/op
Benchmark_Init_Empty_moznion-20                         1000000000               0.1119 ns/op          0 B/op          0 allocs/op
Benchmark_Init_OfPtr_Value_mikkael131-20                1000000000               0.1037 ns/op          0 B/op          0 allocs/op
Benchmark_Init_OfPtr_Value_leighmcculloch-20            1000000000               0.2011 ns/op          0 B/op          0 allocs/op
Benchmark_Init_OfPtr_Nil_mikkael131-20                  1000000000               0.1021 ns/op          0 B/op          0 allocs/op
Benchmark_Init_OfPtr_Nil_leighmcculloch-20              1000000000               0.2181 ns/op          0 B/op          0 allocs/op

# Get()
Benchmark_Get_Present_mikkael131-20                     1000000000               0.2047 ns/op          0 B/op          0 allocs/op
Benchmark_Get_Present_leighmcculloch-20                 1000000000               1.236 ns/op           0 B/op          0 allocs/op
Benchmark_Get_Present_markphelps-20                     1000000000               0.1008 ns/op          0 B/op          0 allocs/op
Benchmark_Get_Present_moznion-20                        1000000000               0.2122 ns/op          0 B/op          0 allocs/op
Benchmark_Get_Empty_mikkael131-20                       1000000000               0.2097 ns/op          0 B/op          0 allocs/op
Benchmark_Get_Empty_leighmcculloch-20                   1000000000               0.9403 ns/op          0 B/op          0 allocs/op
Benchmark_Get_Empty_markphelps-20                       1000000000               0.1093 ns/op          0 B/op          0 allocs/op
Benchmark_Get_Empty_moznion-20                          1000000000               0.2090 ns/op          0 B/op          0 allocs/op

# IsPresent()
Benchmark_IsPresent_Present_mikkael131-20               1000000000               0.1110 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Present_leighmcculloch-20           1000000000               0.1055 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Present_markphelps-20               1000000000               0.1123 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Present_moznion-20                  1000000000               0.1043 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Empty_mikkael131-20                 1000000000               0.1057 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Empty_leighmcculloch-20             1000000000               0.1132 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Empty_markphelps-20                 1000000000               0.1023 ns/op          0 B/op          0 allocs/op
Benchmark_IsPresent_Empty_moznion-20                    1000000000               0.1067 ns/op          0 B/op          0 allocs/op

# IfPresent()
Benchmark_IfPresent_Present_mikkael131-20               1000000000               1.045 ns/op           0 B/op          0 allocs/op
Benchmark_IfPresent_Present_leighmcculloch-20           1000000000               1.039 ns/op           0 B/op          0 allocs/op
Benchmark_IfPresent_Present_markphelps-20               1000000000               0.1033 ns/op          0 B/op          0 allocs/op
Benchmark_IfPresent_Empty_mikkael131-20                 1000000000               0.2312 ns/op          0 B/op          0 allocs/op
Benchmark_IfPresent_Empty_leighmcculloch-20             1000000000               0.2481 ns/op          0 B/op          0 allocs/op
Benchmark_IfPresent_Empty_markphelps-20                 1000000000               0.1088 ns/op          0 B/op          0 allocs/op

# Else()
Benchmark_Else_Present_mikkael131-20                    1000000000               0.2040 ns/op          0 B/op          0 allocs/op
Benchmark_Else_Present_leighmcculloch-20                1000000000               2.671 ns/op           0 B/op          0 allocs/op
Benchmark_Else_Present_markphelps-20                    1000000000               0.1048 ns/op          0 B/op          0 allocs/op
Benchmark_Else_Present_moznion-20                       1000000000               0.2066 ns/op          0 B/op          0 allocs/op
Benchmark_Else_Empty_mikkael131-20                      1000000000               0.2290 ns/op          0 B/op          0 allocs/op
Benchmark_Else_Empty_leighmcculloch-20                  1000000000               2.159 ns/op           0 B/op          0 allocs/op
Benchmark_Else_Empty_markphelps-20                      1000000000               0.1069 ns/op          0 B/op          0 allocs/op
Benchmark_Else_Empty_moznion-20                         1000000000               0.2038 ns/op          0 B/op          0 allocs/op

# ElseGet()
Benchmark_ElseGet_Present_mikkael131-20                 1000000000               0.2144 ns/op          0 B/op          0 allocs/op
Benchmark_ElseGet_Present_leighmcculloch-20             1000000000               1.130 ns/op           0 B/op          0 allocs/op
Benchmark_ElseGet_Present_moznion-20                    1000000000               0.2158 ns/op          0 B/op          0 allocs/op
Benchmark_ElseGet_Empty_mikkael131-20                   1000000000               0.2227 ns/op          0 B/op          0 allocs/op
Benchmark_ElseGet_Empty_leighmcculloch-20               1000000000               1.636 ns/op           0 B/op          0 allocs/op
Benchmark_ElseGet_Empty_moznion-20                      1000000000               0.2069 ns/op          0 B/op          0 allocs/op

# ElseZero()
Benchmark_ElseZero_Present_mikkael131-20                1000000000               0.2282 ns/op          0 B/op          0 allocs/op
Benchmark_ElseZero_Present_leighmcculloch-20            1000000000               2.684 ns/op           0 B/op          0 allocs/op
Benchmark_ElseZero_Empty_mikkael131-20                  1000000000               0.2112 ns/op          0 B/op          0 allocs/op
Benchmark_ElseZero_Empty_leighmcculloch-20              1000000000               2.118 ns/op           0 B/op          0 allocs/op

# Filter()
Benchmark_Filter_Present_mikkael131-20                  1000000000               3.902 ns/op           0 B/op          0 allocs/op
Benchmark_Filter_Present_moznion-20                     1000000000               3.659 ns/op           0 B/op          0 allocs/op
Benchmark_Filter_Empty_mikkael131-20                    1000000000               1.856 ns/op           0 B/op          0 allocs/op
Benchmark_Filter_Empty_moznion-20                       1000000000               1.773 ns/op           0 B/op          0 allocs/op
```