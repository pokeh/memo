### testing out local packages

use `replace`

```go.mod
module github.com/pokeh/someotherrepo

require (
    github.com/pokeh/memo
)

replace github.com/pokeh/memo => ../memo
```
