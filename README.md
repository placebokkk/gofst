# Openfst Go wrapper (WIP)


### Install

- Install the latest [openfst](http://www.openfst.org/twiki/bin/view/FST/WebHome "Title"). Make sure the libfst.so or libfst.dylib(Mac)is in your /usr/local/lib

- go get github.com/placebokkk/gofst

### Test

```
git clone https://github.com/placebokkk/gofst.git
go test
```

### Usage

```
import (
    "github.com/placebokkk/gofst"
)
```

### Next (order by priority)
- support symbol table. wrap the symbol table in fst struct, similar to pyfst
- support more operations
- more real world use examples

I am new to golang and openfst. Welcome any comments and suggestions.