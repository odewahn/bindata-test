Example of how to compile a static website into a go package that can be bundled with the compiled binary.  I'll use this to build a local GUI that can be shipped and served directly from the executable itself. Based on:

* https://github.com/jteeuwen/go-bindata
* https://github.com/elazarl/go-bindata-assetfs

To use it, first run `go-bindata-assetfs gui/...` to build a go package called `bindata_assefs.go`.  This will convert all the assets in the `gui` directory into blobs, and give a nice interface to working with them.  *I should also explore the `--debug` mode where it serves the files themselves, versus compiling them.*

Then, you can build or run it and serve the package as a site using `net/http`, like this:

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
	./webshare

build:
	go-bindata-assetfs  static/...
	go build
```
