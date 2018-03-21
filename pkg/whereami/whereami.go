// Copyright 2016 - by Jim Lawless
// License: MIT / X11
// See: http://www.mailsend-online.com/license2016.php
//
// This code may not conform to popular Go coding idioms

// custom by jungmu

package whereami

import (
	"fmt"
	"runtime"
	"strings"
)

// WhereAmI : return a string containing the file name, function name
// and the line number of a specified entry on the call stack
func WhereAmI(depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	fullFunctionName := strings.Split(runtime.FuncForPC(function).Name(), "/")
	functionName := fullFunctionName[len(fullFunctionName)-1]
	return fmt.Sprintf("%s  Func: %s Line: %d \t", chopPath(file), functionName, line)
}

// WhereAmIforError : return a string containing the file name, function full name
// and the line number of a specified entry on the call stack
func WhereAmIforError(depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	fullFunctionName := strings.Split(runtime.FuncForPC(function).Name(), "/")
	return fmt.Sprintf("%s  Func: %s Line: %d \t", chopPath(file), fullFunctionName, line)
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}
