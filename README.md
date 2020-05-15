# latest-version
Determine the latest available version of a module from [proxy.golang.org](proxy.golang.org).
Useful for tools that plan to check available upgrades.

## Installation
```
$ go get github.com/joshuabezaleel/latest-version
```

## Usage
``` go
import (
    latestver "github.com/joshuabezaleel/latest-version
)

func main() {
    latestVersion, _ := latestver.LatestVersion("github.com/hashicorp/raft")
    log.Println(latestVersion)
    // 1.1.2
}
```

## Prior Art
Highly inspired by [sindesorhus'](https://github.com/sindresorhus) [latest-version](https://github.com/sindresorhus/latest-version)

## License
MIT License