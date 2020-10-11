# Logger

Log to file by date for go lang.

### Example
```go

log := logger.New(path, filename prefix, log prefix, print out log)
```

```go
package main

import (
	"github.com/weihung/logger"
)

func main() {
	log := logger.New("./log", "api", "", true)
	log.Println("This is an example")
}
``` 
