package renamer

import (
	"os"
	"path/filepath"
	"testing"
)

func createTestFile(path string, t *testing.T) {
	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("Failed to create test file %s: %v", path, err)
	}
	defer file.Close()
	_, err = file.WriteString("test content")
	if err != nil {
		t.Fatalf("Failed to write to test file %s: %v", path, err)
	}
}

func createTestDir(path string, t *testing.T) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory %s: %v", path, err)
	}
}

func TestRename(t *testing.T) {
	tests := []struct {
		name       string
		params     RenamerParams
		setup      func() string
		teardown   func()
		shouldFail bool
	}{
		{
			name: "basic renaming",
			params: RenamerParams{
				Pattern:     "test",
				Replacement: "demo",
				Root:        "",
				DryRun:      false,
				Recursive:   false,
				Log:         false,
				LogFile:     "",
			},
			setup: func() string {
				dir := "testdata/basic"
				createTestDir(dir, t)
				createTestFile(filepath.Join(dir, "testfile.txt"), t)
				return dir
			},
			teardown: func() {
				os.Remove("gorename-log.txt")
				os.RemoveAll("testdata")
			},
			shouldFail: false,
		},
		{
			name: "recursive renaming",
			params: RenamerParams{
				Pattern:     "test",
				Replacement: "demo",
				Root:        "",
				DryRun:      false,
				Recursive:   true,
				Log:         false,
				LogFile:     "",
			},
			setup: func() string {
				dir := "testdata/recursive"
				createTestDir(filepath.Join(dir, "subdir"), t)
				createTestFile(filepath.Join(dir, "subdir", "testfile.txt"), t)
				return dir
			},
			teardown: func() {
				os.Remove("gorename-log.txt")
				os.RemoveAll("testdata")
			},
			shouldFail: false,
		},
		{
			name: "dry run",
			params: RenamerParams{
				Pattern:     "test",
				Replacement: "demo",
				Root:        "",
				DryRun:      true,
				Recursive:   false,
				Log:         false,
				LogFile:     "",
			},
			setup: func() string {
				dir := "testdata/dryrun"
				createTestDir(dir, t)
				createTestFile(filepath.Join(dir, "testfile.txt"), t)
				return dir
			},
			teardown: func() {
				os.Remove("gorename-log.txt")
				os.RemoveAll("testdata")
			},
			shouldFail: false,
		},
		{
			name: "log renames",
			params: RenamerParams{
				Pattern:     "test",
				Replacement: "demo",
				Root:        "",
				DryRun:      false,
				Recursive:   false,
				Log:         true,
				LogFile:     "testdata/testlog.txt",
			},
			setup: func() string {
				dir := "testdata/logging"
				createTestDir(dir, t)
				createTestFile(filepath.Join(dir, "testfile.txt"), t)
				return dir
			},
			teardown: func() {
				os.RemoveAll("testdata")
			},
			shouldFail: false,
		},
		{
			name: "different root directory",
			params: RenamerParams{
				Pattern:     "test",
				Replacement: "demo",
				Root:        "testdata/differentroot",
				DryRun:      false,
				Recursive:   false,
				Log:         false,
				LogFile:     "",
			},
			setup: func() string {
				dir := "testdata/differentroot"
				createTestDir(dir, t)
				createTestFile(filepath.Join(dir, "testfile.txt"), t)
				return dir
			},
			teardown: func() {
				os.Remove("gorename-log.txt")
				os.RemoveAll("testdata")
			},
			shouldFail: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := tt.setup()
			defer tt.teardown()

			tt.params.Root = root
			err := Rename(tt.params)

			if (err != nil) != tt.shouldFail {
				t.Fatalf("Rename() error = %v, shouldFail %v", err, tt.shouldFail)
			}

			if !tt.params.DryRun && !tt.shouldFail {
				// Verify the renaming for non-dry-run and non-failing cases
				files, err := os.ReadDir(root)
				if err != nil {
					t.Fatalf("Failed to read directory %s: %v", root, err)
				}
				for _, file := range files {
					if file.Name() == "testfile.txt" && !tt.params.DryRun {
						t.Fatalf("Expected file to be renamed, but found original name")
					}
				}
			}

			if tt.params.Log {
				if _, err := os.Stat(tt.params.LogFile); os.IsNotExist(err) {
					t.Fatalf("Expected log file %s to be created", tt.params.LogFile)
				}
			}
		})
	}
}
