// Copyright (c) 2024 Method Security. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package writer

import (
	"strings"
)

type Format struct {
	val Format_Value
}

type Format_Value string

const (
	JSON    Format_Value = "json"
	YAML    Format_Value = "yaml"
	SIGNAL  Format_Value = "signal"
	UNKNOWN Format_Value = "unknown"
)

func Format_Values() []Format_Value {
	return []Format_Value{
		JSON,
		YAML,
		SIGNAL,
	}
}

func New_Format(value Format_Value) Format {
	return Format{val: value}
}

func (f Format) IsUnknown() bool {
	switch f.val {
	case JSON, YAML, SIGNAL:
		return false
	}
	return true
}

func (f Format) String() string {
	return string(f.val)
}

func (f *Format) UnmarshalText(text []byte) error {
	switch v := strings.ToUpper(string(text)); v {
	default:
		*f = New_Format(Format_Value(v))
	case "JSON":
		*f = New_Format(JSON)
	case "YAML":
		*f = New_Format(YAML)
	case "SIGNAL":
		*f = New_Format(SIGNAL)
	case "UNKNOWN":
		*f = New_Format(UNKNOWN)
	}
	return nil
}

func (f Format) MarshalText() ([]byte, error) {
	return []byte(f.String()), nil
}
