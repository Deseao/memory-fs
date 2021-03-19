package fs_test

import (
	"testing"

	"github.com/Deseao/memory-fs/internal/fs"
)

func TestNew(t *testing.T) {
	dir := fs.New("root")
	if dir == nil {
		t.Errorf("New directory was nil, expected it not to be nil")
	}
	if dir.Name != "root" {
		t.Errorf("Directory name was not %s, got %s", "root", dir.Name)
	}
}

func TestCreateDirectory(t *testing.T) {
	dir := fs.New("root")
	dir.CreateDirectory([]string{"1", "2", "3"})
	if len(dir.Subdirectories) != 1 {
		t.Errorf("Expected %d subdirectory, got %d", 1, len(dir.Subdirectories))
	}
	if dir.Subdirectories[0].Name != "1" {
		t.Errorf("Expected root to have a directory named %s but was %s", "1", dir.Subdirectories[0].Name)
	}

	dir1 := dir.Subdirectories[0]

	if len(dir1.Subdirectories) != 1 {
		t.Errorf("Expected \"1\" subdirectory to contain %d subdirectories itself but got %d", 1, len(dir1.Subdirectories))
	}
	if dir1.Subdirectories[0].Name != "2" {
		t.Errorf("Expected subdirectory to have name %s but was %s", "2", dir1.Subdirectories[0].Name)
	}
	dir2 := dir1.Subdirectories[0]
	if len(dir2.Subdirectories) != 1 {
		t.Errorf("Expected \"2\" subdirectory to contain %d subdirectories itself but got %d", 1, len(dir2.Subdirectories))
	}

	if dir2.Subdirectories[0].Name != "3" {
		t.Errorf("Expected subdirecto to have name %s but was %s", "3", dir2.Subdirectories[0].Name)
	}

}
