# rslicer
Fast and simple rune slicer. No memory allocations.
Just tell borders for runes and get byte positions in a string.
Support negative indexes.

get module - `go get github.com/HoskeOwl/rslicer`

Available two functions:
* GetRuneRange - return byte position for
* GetSlice - return string slice

## Examples
Can be used with ascii

```
package main

import "github.com/HoskeOwl/rslicer"

func main() {
	s := "Some sentence."
	begin, end, err := rslicer.GetRuneRange(s, 0, 4)
    if err != nil{
        fmt.Printf("Got some error: %v", err)
        return
    }
    fmt.Printf("Result: '%v'", s[begin:end]) //Result: 'Some'
}
```

Can be used with unicode

```
package main

import "github.com/HoskeOwl/rslicer"

func main() {
	s := "Some 日本語 symbols."
	begin, end, err := rslicer.GetRuneRange(s, 5, 8)
    if err != nil{
        fmt.Printf("Got some error: %v", err)
        return
    }
    fmt.Printf("Result: '%v'", s[begin:end]) //Result: '日本語'
}
```

And with negative indexes too.

```
package main

import "github.com/HoskeOwl/rslicer"

func main() {
	s := "Some 日本語 symbols."
	begin, end, err := rslicer.GetRuneRange(s, 5, -9)
    if err != nil{
        fmt.Printf("Got some error: %v", err)
        return
    }
    fmt.Printf("Result: '%v'", s[begin:end]) //Result: '日本語'
}
```

Or directry get slice.

```
package main

import "github.com/HoskeOwl/rslicer"

func main() {
	s := "Some 日本語 symbols."
	out, err := rslicer.GetRuneSlice(s, 5, -9)
    if err != nil{
        fmt.Printf("Got some error: %v", err)
        return
    }
    fmt.Printf("Result: '%v'", out) //Result: '日本語'
}
```


## Library Benchmark
```
goos: linux
goarch: amd64
pkg: github.com/HoskeOwl/rslicer
cpu: AMD Ryzen 7 4700U with Radeon Graphics  

# slice through runes transformation
BenchmarkSmallThroughRunes-8                	41294881	        30.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkMediumThroughRunes-8               	 2165407	       542.9 ns/op	     176 B/op	       1 allocs/op
BenchmarkLargeThroughRunes-8                	  137636	      7809 ns/op	    2000 B/op	       2 allocs/op

# slice through library
BenchmarkPositiveSmallWithFunction-8        	48914659	        21.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkNegativeSmallWithFunction-8        	40990228	        24.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkPositiveMediumWithFunction-8       	16005105	        76.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkNegativeMediumWithFunction-8       	14637120	        80.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkPositiveLargeWithFunction-8        	  727724	      1551 ns/op	       0 B/op	       0 allocs/op
BenchmarkNegativeLargeWithFunction-8        	  936205	      1275 ns/op	       0 B/op	       0 allocs/op
BenchmarkNegativeLargeSliceWithFunction-8   	  845863	      1269 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/HoskeOwl/rslicer	16.561s

```

