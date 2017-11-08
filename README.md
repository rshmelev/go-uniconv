# go-uniconv

Cast "anything to anything".

```golang

import( . github.com/rshmelev/go-uniconv/cast )

a := Cast(1).Str() // fmt.Sprint(value) 
b := Cast(2).Float32()
c := Cast("stringsome").Bytes()
d := Cast([]int{1,2,3}).Strings()
e := Cast([]string{"1","2","3"}).Ints()
f := Cast(2343424343242).Time()

// var TimeFormats = []string{"1 2 2006", "1 2 2006 15 4 5", "2006 1 2 15 4 5", "2006 1 2", "15 4 5 Jan 2, 2006 MST"}
g := Cast("1-2-2006").Time()

```

there is also "First" as an additional sublibrary

```
import( . github.com/rshmelev/go-uniconv/first )

a := FirstStr("hello") // "hello"
b := FirstStr() // ""
c := FirstStr(a...) // depends on emptiness of a

```

Note: it is definitely slower than usual conv methods, so it is kind of universal thing that may be useful on non-frequently-repeating parts of code

author: rshmelev@gmail.com

