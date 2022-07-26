# Dicer

A dice rolling app

### Usage

```go
import "github.com/atallison/dicer"

func main() {
    roll, err := dicer.Roll("2d20")
    if err != nil {
    panic(err)
    }
    
    fmt.Println(roll.ToString())
}
```