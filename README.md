# aspell check

Small repo to extend [go-aspell](https://github.com/trustmaster/go-aspell) 
for use with more than one word.

## Install

Check [go-aspell](https://github.com/trustmaster/go-aspell) on how to make a 
proper setup of the that library. `go-aspell-check` doesn't require further 
setup.

### GNU aspell

First make sure aspell library and headers are installed on your system.

#### Debian/Ubuntu
On Debian/Ubuntu you can install it this way:
```
sudo apt-get install aspell libaspell-dev
```

#### macOS
On mac you can use `brew` for installing aspell
```
brew install aspell
```


## Usage

```go
package main

import (
	"fmt"
	gac "github.com/pjnr1/go-aspell-check"
)

func main() {
	opts := map[string]string{"lang": "en_US"}
	speller := gac.NewSpeller(opts)

	test := "Hey there. Hope you are haaving a blast!"
	fmt.Println(test)
	fmt.Println(speller.CheckWithFeedback(test))
}
```

which should give you the output

```
Hey there. Hope you are haaving a blast!
                        ~~~~~~~         
```

