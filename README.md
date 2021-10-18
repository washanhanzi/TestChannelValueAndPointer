# TestChannelValueAndPointer

- cpu: Intel(R) Core(TM) i7-8700B CPU @ 3.20GHz

| type                           | iter   | ns/op       | B/op       | allocs/op   |
| ------------------------------ | ------ | ----------- | ---------- | ----------- |
| Benchmark/500byte_value-12     | 295989 | 4211 ns/op  | 672 B/op   | 4 allocs/op |
| Benchmark/500byte_pointer-12   | 569635 | 3389 ns/op  | 1685 B/op  | 6 allocs/op |
| Benchmark/1000byte_value-12    | 371932 | 4773 ns/op  | 679 B/op   | 4 allocs/op |
| Benchmark/1000byte_pointer-12  | 688429 | 6162 ns/op  | 2718 B/op  | 6 allocs/op |
| Benchmark/10000byte_value-12   | 100243 | 14913 ns/op | 624 B/op   | 4 allocs/op |
| Benchmark/10000byte_pointer-12 | 174306 | 21985 ns/op | 21231 B/op | 6 allocs/op |
