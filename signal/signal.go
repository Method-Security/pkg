// Copyright (c) 2024 Method Security. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package signal

import (
	"github.com/palantir/pkg/datetime"
)

type SignalReport struct {
	Content      string             `json:"content" yaml:"content"`
	StartedAt    datetime.DateTime  `json:"started_at" yaml:"started_at"`
	CompletedAt  *datetime.DateTime `json:"completed_at,omitempty" yaml:"completed_at,omitempty"`
	Status       int                `json:"status" yaml:"status"`
	ErrorMessage *string            `json:"error_message,omitempty" yaml:"error_message,omitempty"`
}

type Signal struct {
	Report       any
	StartedAt    datetime.DateTime
	CompletedAt  *datetime.DateTime
	Status       int
	ErrorMessage *string
}

func NewSignal(report any, startedAt datetime.DateTime, completedAt *datetime.DateTime, status int, errorMessage *string) Signal {
	return Signal{
		Report:       report,
		StartedAt:    startedAt,
		CompletedAt:  completedAt,
		Status:       status,
		ErrorMessage: errorMessage,
	}
}

func NewSignalReport(content string, signal Signal) SignalReport {
	return SignalReport{
		Content:      content,
		StartedAt:    signal.StartedAt,
		CompletedAt:  signal.CompletedAt,
		Status:       signal.Status,
		ErrorMessage: signal.ErrorMessage,
	}
}
