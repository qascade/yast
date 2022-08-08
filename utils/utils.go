/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/

package utils

// This Package contains some helper functions for logging.
import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func LogTrace() {
	log.Print(TraceMsgOffset("", 1))
}

func LogTraceMsg(msg string) {
	log.Print(TraceMsgOffset(msg, 1))
}

func LogTraceMsgOffset(msg string, callStackIdxOffset int) {
	log.Print(TraceMsgOffset(msg, callStackIdxOffset+1))
}

func Trace() string {
	return TraceMsgOffset("", 1)
}

func TraceMsg(msg string) string {
	return TraceMsgOffset(msg, 1)
}

func TraceMsgOffset(msg string, callStackIdxOffset int) string {
	var traceMsg string
	callStackIdx := 2 + callStackIdxOffset
	pc := make([]uintptr, 15)
	n := runtime.Callers(callStackIdx, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	traceMsg += fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function)
	if msg != "" {
		traceMsg += fmt.Sprintf(" %s ", msg)
	}
	traceMsg += "\n"
	return traceMsg
}

func LogUnimplementedFunc() {
	callStackIdx := 2
	pc := make([]uintptr, 15)
	n := runtime.Callers(callStackIdx, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	functionNameParts := strings.Split(frame.Function, ".")
	simpleFuncName := functionNameParts[len(functionNameParts)-1]

	logMsg := fmt.Sprintf("TODO - Implement function %s at %s:%d", simpleFuncName, frame.File, frame.Line)
	log.Print(logMsg)
}
