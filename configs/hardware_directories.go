/*
 * This file is part of arduino-cli.
 *
 * arduino-cli is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 *
 * Copyright 2017-2018 ARDUINO AG (http://www.arduino.cc/)
 */

package configs

import (
	"os"
	"path/filepath"

	"github.com/bcmi-labs/arduino-cli/paths"
)

// HardwareDirectories returns all paths that may contains hardware packages.
func HardwareDirectories() (paths.PathList, error) {
	res := paths.PathList{}

	if dir, err := PackagesFolder.Get(); err == nil {
		res = append(res, paths.New(dir))
	} else {
		return nil, err
	}

	if IsBundledInDesktopIDE() {
		bundledHardwareFolder := filepath.Join(*arduinoIDEDirectory, "hardware")
		if info, err := os.Stat(bundledHardwareFolder); err == nil && info.IsDir() {
			res = append(res, paths.New(bundledHardwareFolder))
		}
	}

	if dir, err := SketchbookFolder.Get(); err == nil {
		dir = filepath.Join(dir, "hardware")
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			res = append(res, paths.New(dir))
		}
	}

	return res, nil
}

// BundleToolsDirectories returns all paths that may contains bundled-tools.
func BundleToolsDirectories() (paths.PathList, error) {
	res := paths.PathList{}

	if IsBundledInDesktopIDE() {
		bundledToolsFolder := filepath.Join(*arduinoIDEDirectory, "hardware", "tools")
		if info, err := os.Stat(bundledToolsFolder); err == nil && info.IsDir() {
			res = append(res, paths.New(bundledToolsFolder))
		}
	}

	return res, nil
}
