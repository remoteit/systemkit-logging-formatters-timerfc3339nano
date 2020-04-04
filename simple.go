package formatters

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	logging "github.com/codemodify/systemkit-logging"
)

type timeRFC3339NanoFormatter struct {
	timeFormatter string
}

const maxFormattedTimeLength = 30

// NewTimeRFC3339NanoFormatter -
func NewTimeRFC3339NanoFormatter() logging.Logger {
	return &timeRFC3339NanoFormatter{
		timeFormatter: time.RFC3339Nano,
	}
}

func (thisRef timeRFC3339NanoFormatter) Log(logEntry logging.LogEntry) logging.LogEntry {
	var formattedTime = logEntry.Time.UTC().Format(thisRef.timeFormatter)

	// reformat the time and fill-in with zeros for the nano seconds
	if len(formattedTime) < maxFormattedTimeLength {
		var spacesCount = maxFormattedTimeLength - len(formattedTime)

		var newV = fmt.Sprintf("%"+strconv.Itoa(spacesCount+1)+"v", "Z")
		newV = strings.Replace(newV, " ", "0", spacesCount)

		formattedTime = strings.Replace(
			formattedTime,
			"Z",
			newV,
			1,
		)
	}

	// format the log line
	logEntry.Message = fmt.Sprintf(
		"%s | %s | %s",
		formattedTime,
		logEntry.Type,
		logEntry.Message,
	)

	return logEntry
}
