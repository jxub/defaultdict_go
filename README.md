# defaultdict_go
##### Pythonic defaultdict in Golang :metal:

## Install

```zsh
$ go get github.com/jxub/defaultdict_go
```
:tada::shit::fire:

## Usage

```golang
import "github.com/jxub/defaultdict_go" dd

func main() {
    // create a defaultdict with default values of 0,
    // as int() constrictor yields in Python
    d := dd.NewDefaultDict(0)

    // Get and Set works as usual
    d.Set("bitcoin", 14364)
    d.Get("bitcoin") // -> 14364
    
    // getting a not initialized key yields the default value
    d.Get("dentacoin") // -> 0 (oh the irony ;))
}
```

PR's and other good stuff encouraged! Happy Hacking!