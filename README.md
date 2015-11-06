Example of how to compile a static website into a go package that can be bundled with the compiled binary.  Based on:

* https://github.com/jteeuwen/go-bindata
* https://github.com/elazarl/go-bindata-assetfs

Once you `go get elazarl/go-bindata-assetfs`, run `go-bindata-assetfs static/...` to build a package called `bindata_assefs.go`.  This will convert all the files and assets in the `static` directory into a go package and give a nice interface to working with them.  The assets and files are compressed and stored as a data structure in the package.

*For development, use `go-bindata-assetfs -debug static/...` the to serve the files directly, so that you don't have to compile them each time.*

Then, you can build or run it and mount the package as a file handler using `net/http`, like this:

```
package main

import (
	"net/http"
  "fmt"
)

func main() {
   http.Handle("/", http.FileServer(assetFS()))
   fmt.Println("Serving on port 8000")
   http.ListenAndServe(":8000", nil)
}
```

Here's an [example makefile](https://github.com/dreamersdw/webshare/blob/master/Makefile) of how you can do all this automatically:

```
run: build
	./main

build:
	go-bindata-assetfs static/...
	go build -o main *.go
```
