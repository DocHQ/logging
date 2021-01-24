package console

import (
	// TODO Copy default format and just remove import
	"log"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var terminalSupportsColor bool = false

var (
	// So we write to the correct outputs,
	// Some applications require this to identify when an error has occured
	errLog = log.New(os.Stderr, "", 0)
	outLog = log.New(os.Stdout, "", 0)
)

func levelToName(l uint32) string {
	switch l {
	case 0:
		if terminalSupportsColor {
			return " \u001b[32mDEBUG\u001b[0m "
		}
		return " DEBUG "
	case 1:
		if terminalSupportsColor {
			return " \u001b[34mINFO\u001b[0m  "
		}
		return "INFO  "
	case 2:
		if terminalSupportsColor {
			return " \u001b[33mWARN\u001b[0m  "
		}
		return " WARN  "
	case 3:
		if terminalSupportsColor {
			return " \u001b[31mERROR\u001b[0m "
		}
		return " ERROR "
	case 4:
		if terminalSupportsColor {
			return " \u001b[31mFATAL\u001b[0m "
		}
		return " FATAL "
	default:
		return " N/A "
	}
}

func init() {
	var outString string = "0"

	// Check weather the current terminal supports colours
	out, _ := exec.Command("/usr/bin/tput", "colors").Output()

	outString = strings.TrimSpace(string(out))

	if outString == "" {
		return
	}

	colours, _ := strconv.Atoi(outString)

	if colours > 8 {
		terminalSupportsColor = true
	}
}

type Logger struct{}

func (t Logger) Log(i interface{}, fields map[string]interface{}, level uint32, verbose bool) {
	// Get the current time
	timestring := time.Now().Format("Mon Jan _2 15:04:05")

	// If the type is a structure format it nicly for output
	if t := reflect.TypeOf(i); t.Kind() == reflect.Struct {
		i = "\n" + spew.Sdump(i)
	}

	// Check the current debug status
	if verbose {
		// Get the runtime caller
		pc, _, line, _ := runtime.Caller(2)
		details := runtime.FuncForPC(pc)

		if level <= 2 {
			outLog.Printf("%s[%s] [%s#%d] %s \n", timestring, levelToName(level), details.Name(), line, i)
		} else {
			errLog.Printf("%s[%s] [%s#%d] %s \n", timestring, levelToName(level), details.Name(), line, i)
		}
	} else {
		if level <= 2 {
			outLog.Printf("%s[%s] %s \n", timestring, levelToName(level), i)
		} else {
			errLog.Printf("%s[%s] %s \n", timestring, levelToName(level), i)
		}
	}

	// If is a fatal call
	if level == 4 {
		os.Exit(1)
	}
}
