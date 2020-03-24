package junction

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func testCreate() error {
	target, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd: %w", err)
	}
	mountPt := filepath.Join(os.TempDir(), "junctionTest")

	err = Create(target, mountPt)
	if err != nil {
		return fmt.Errorf("Create: %w", err)
	}
	defer os.Remove(mountPt)

	link, err := os.Readlink(mountPt)
	if err != nil {
		return fmt.Errorf("os.Readlink: %w", err)
	}
	if target != link {
		return fmt.Errorf("os.Readlink: linked path differs : '%s' != '%s'", target, link)
	}

	main_go := filepath.Join(mountPt, "main.go")
	fd, err := os.Open(main_go)
	if err != nil {
		return fmt.Errorf("mount may be fake because '%s' can not be open", main_go)
	}
	fd.Close()
	return nil
}

func TestCreate(t *testing.T) {
	if err := testCreate(); err != nil {
		t.Fatal(err.Error())
	}
}
