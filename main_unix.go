// +build !windows

package junction

import (
	"os"
)

func create(target, mountPt string) error {
	return os.Symlink(target, mountPt)
}
