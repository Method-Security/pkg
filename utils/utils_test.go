package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// Test the getEntriesFromFiles function
func TestGetEntriesFromFiles(t *testing.T) {
	// Create temporary directory
	tmpDir, err := ioutil.TempDir("", "test_files")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create temporary test files
	testFile1 := filepath.Join(tmpDir, "test1.txt")
	testFile2 := filepath.Join(tmpDir, "test2.txt")

	content1 := "line1\nline2\nline3"
	content2 := "line4\nline5\nline6"

	err = ioutil.WriteFile(testFile1, []byte(content1), 0644)
	if err != nil {
		t.Fatalf("Failed to write to test file1: %v", err)
	}

	err = ioutil.WriteFile(testFile2, []byte(content2), 0644)
	if err != nil {
		t.Fatalf("Failed to write to test file2: %v", err)
	}

	// Call getEntriesFromFiles
	paths := []string{testFile1, testFile2}
	entries, err := getEntriesFromFiles(paths)
	if err != nil {
		t.Fatalf("Error getting entries from files: %v", err)
	}

	// Expected entries
	expectedEntries := []string{"line1", "line2", "line3", "line4", "line5", "line6"}

	// Verify output
	if len(entries) != len(expectedEntries) {
		t.Errorf("Expected %d entries, got %d", len(expectedEntries), len(entries))
	}
	for i, entry := range entries {
		if entry != expectedEntries[i] {
			t.Errorf("Expected entry '%s', got '%s'", expectedEntries[i], entry)
		}
	}
}

// Test error case: Invalid file path
func TestGetEntriesFromFiles_InvalidPath(t *testing.T) {
	_, err := getEntriesFromFiles([]string{"nonexistent.txt"})
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
}
