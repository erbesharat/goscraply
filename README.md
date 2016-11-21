# goscraply
A simple package to scrap websites.

## Supported websites
* [Goodread](https://goodread.com)
* [Listchallenges](http://listchallenges.com)

## Usage
```go
package main

import(
    "fmt"
    "github.com/erbesharat/ebscraply"
)

func main(){
    books := ebscraply.goodreads("horror")
    fmt.Println(books)
}
```
## Todo
* Add other websites
* Test

## License
[GPLv3](https://www.gnu.org/licenses/gpl-3.0.txt)
