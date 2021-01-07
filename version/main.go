package version

import (
	"fmt"
	"runtime"
)

// GitCommit the git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

// Version the main version number that is being run at the moment.
var Version = "X.Y.Z"

// BuildDate datetime that binary was created
var BuildDate = ""

// GoVersion go runtime version
var GoVersion = runtime.Version()

// OsArch OS architecture
var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
