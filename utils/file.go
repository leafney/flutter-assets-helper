/**
 * @Author:      leafney
 * @Date:        2021-01-24 17:17
 * @Project:     flutter-assets-helper
 * @Description:
 */

package utils

import "os"

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// IsDir reports whether d is a directory.
func IsDir(d string) (y bool) {
	if fi, err := os.Stat(d); err == nil {
		if fi.IsDir() {
			y = true
		}
	}
	return
}
