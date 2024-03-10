# dstparser
 
`dstparser` is a package written in go that provides ability to parse dst information into go type, and reflect go type back into lua script.

## install

```bash
go get github.com/dstgo/parser@latest
```

## usage
supported usage as follows

* modinfo.lua
* modoverride.lua
* leveldataoverride.lua
* cluster.ini
* player.txt

### modinfo
parse `modinfo.lua` directly
```go
import (
	"fmt"
	"github.com/dstgo/dstparser"
	"os"
)

func main() {
	bytes, err := os.ReadFile("workshop-123456789/modinfo.lua")
	if err != nil {
		panic(err)
	}
	info, err := dstparser.ParseModInfo(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(info.Name)
}
```
or use mod env
```go
package main
import (
    "fmt"
    "github.com/dstgo/dstparser"
    "os"
)

func main() {
    bytes, err := os.ReadFile("workshop-123456789/modinfo.lua")
    if err != nil {
        panic(err)
    }
    info, err := dstparser.ParseModInfoWithEnv(bytes, "workshop-123456789", "zh")
        if err != nil {
        panic(err)
    }
    fmt.Println(info.Name)
}

```

### modoverrides
supported parse `modoverrides.lua` to go type and reflecting back to lua script
```go
package main

import (
	"fmt"
	"github.com/dstgo/dstparser"
	"os"
)

func main() {
	// parse
	bytes, err := os.ReadFile("cluster/master/modoverrides.lua")
	if err != nil {
		panic(err)
	}
	modoverrides, err := dstparser.ParseModOverrides(bytes)
	if err != nil {
		panic(err)
	}
	for _, modoverride := range modoverrides {
		fmt.Println(modoverride.Id, modoverride.Enabled, len(modoverride.Items))
	}

	// serialize
	modOverrideLua, err := dstparser.ToModOverrideLua(modoverrides)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(modOverrideLua))
}
```

### leveldataoverrides
supported parsing `leveldataoverrides.lua` to go type and reflecting back to lua script
```go
package main

import (
	"fmt"
	"github.com/dstgo/dstparser"
	"os"
)

func main() {
	// parse
	bytes, err := os.ReadFile("cluster/master/leveldataoverrides.lua")
	if err != nil {
		panic(err)
	}
	overrides, err := dstparser.ParseLevelDataOverrides(bytes)
	if err != nil {
		panic(err)
	}
	masterLevelDataOverridesLua, err := dstparser.ToMasterLevelDataOverridesLua(overrides)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(masterLevelDataOverridesLua))
}
```

### cluster ini
supported parsing `cluster.ini `
```go
package main

import (
	"fmt"
	"github.com/dstgo/dstparser"
	"os"
)

func main() {
	// parse
	bytes, err := os.ReadFile("cluster/cluster/cluster.ini")
	if err != nil {
		panic(err)
	}
	clusterConfig, err := dstparser.ParseClusterInI(bytes)
	if err != nil{
		panic(err)
	}
	clusterInI, err := dstparser.ToClusterInI(clusterConfig)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(clusterInI))
}
```
or `server.ini`
```go
package main

import (
	"fmt"
	"github.com/dstgo/dstparser"
	"os"
)

func main() {
	// parse
	bytes, err := os.ReadFile("cluster/cluster/server.ini")
	if err != nil {
		panic(err)
	}
	serverConfig, err := dstparser.ParseServerInI(bytes)
	if err != nil{
		panic(err)
	}
	serverIni, err := dstparser.ToServerInI(serverConfig)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(serverIni))
}
```

### player
supported parsing `player.txt` and `server_chat_log.txt`
```go
package main

import (
	"fmt"
	"github.com/dstgo/dstparser"
	"os"
)

func main() {
	bytes, err := os.ReadFile("server_chat_log.txt")
	if err != nil {
		panic(err)
	}
	logs, err := dstparser.ParseServerChatLogs(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(logs)
}
```