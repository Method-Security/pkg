// Copyright (c) 2024 Method Security. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package writer

import (
	"encoding/json"
	"os"
)

func Write(report any, outputFilePath *string) error {
	fileData, err := json.Marshal(report)
	if err != nil {
		return err
	}

	if outputFilePath == nil {
		_, err = os.Stdout.Write(fileData)
		if err != nil {
			return err
		}
	} else {
		err = os.WriteFile(*outputFilePath, fileData, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
