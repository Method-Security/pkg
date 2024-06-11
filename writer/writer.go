// Copyright (c) 2024 Method Security. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package writer

import (
	"encoding/base64"
	"os"

	sig "github.com/Method-Security/pkg/signal"
	"github.com/palantir/pkg/datetime"
	"github.com/palantir/pkg/safejson"
	"github.com/palantir/pkg/safeyaml"
)

func Write(
	report any,
	config OutputConfig,
	startedAt datetime.DateTime,
	completedAt *datetime.DateTime,
	status int,
	errorMessage *string,
) error {
	var data []byte
	var err error
	switch config.Output.val {
	case JSON:
		data, err = safejson.Marshal(report)
	case YAML:
		data, err = safeyaml.Marshal(report)
	case SIGNAL:
		signal := sig.NewSignal(report, startedAt, completedAt, status, errorMessage)
		data, err = writeSignal(signal)
	}
	if err != nil {
		return err
	}
	return writeToFileOrStdout(data, config.FilePath)
}

func writeSignal(
	signal sig.Signal,
) ([]byte, error) {
	data, err := safejson.Marshal(signal.Report)
	if err != nil {
		return nil, err
	}
	encoded := base64.StdEncoding.EncodeToString(data)
	report := sig.NewSignalReport(encoded, signal)
	return safejson.Marshal(report)
}

func writeToFileOrStdout(data []byte, filePath *string) error {
	if filePath == nil {
		_, err := os.Stdout.Write(data)
		if err != nil {
			return err
		}
	} else {
		err := os.WriteFile(*filePath, data, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
