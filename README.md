##localtime

设置时区会默认去找`ZONEINFO`环境变量，找不到就去找`GOROOT/lib/time/zoneinfo.zip`

### 用法
```go
import "github.com/Icepo/localtime"
```
### go mod
```go
require github.com/Icepo/localtime v1.0.0
```