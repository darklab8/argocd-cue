package logus

import (
	"os"

	"github.com/darklab8/go-typelog/typelog"
)

// we aren't allowed outputing to stdout stuff of warning/info/debug level because it will break info received by argocd
var LogStdout *typelog.Logger = typelog.NewLogger("argocue", typelog.WithLogLevel(typelog.LEVEL_ERROR))

var LogFile *typelog.Logger

func init() {

	f, err := os.OpenFile("/tmp/argocue_log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	LogStdout.CheckError(err, "failed to opened")

	LogFile = typelog.NewLogger("argocue", typelog.WithLogLevel(typelog.LEVEL_DEBUG), typelog.WithIoWriter(f))
}
