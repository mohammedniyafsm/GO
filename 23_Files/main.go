package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ============================================
// 1. WHAT IS FILE HANDLING?
// ============================================
//
// File handling allows a program to store, read,
// update, copy, move, and delete data on disk.
//
// Real-world examples:
//
// - Save application logs.
// - Read configuration files.
// - Store generated reports.
// - Upload or process files.
// - Read user data from a text file.
//
// Important packages:
//
// os       = Files, directories, and operating-system operations.
// io       = General input/output utilities.
// bufio    = Buffered reading and writing.
// filepath = Portable file path operations.
//
// ============================================
// 2. CREATE AND WRITE A FILE
// ============================================

func createAndWriteFile() {
	fmt.Println("\n--- 2. Create and Write ---")

	data := []byte("Hello from Go!\nLearning file handling.")

	// Create the file if it does not exist.
	// If it already exists, replace its contents.
	//
	// 0644 is a permission mode:
	// Owner can read/write; others can read,
	// subject to operating-system permissions.
	err := os.WriteFile("example.txt", data, 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File created and written successfully")
}

// Expected output:
// --- 2. Create and Write ---
// File created and written successfully

// ============================================
// 3. READ THE ENTIRE FILE
// ============================================

func readEntireFile() {
	fmt.Println("\n--- 3. Read Entire File ---")

	// ReadFile returns the entire file as []byte.
	data, err := os.ReadFile("example.txt")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert bytes into a string for display.
	fmt.Println(string(data))
}

// Expected output:
// --- 3. Read Entire File ---
// Hello from Go!
// Learning file handling.

// ============================================
// 4. APPEND CONTENT TO A FILE
// ============================================

func appendToFile() {
	fmt.Println("\n--- 4. Append to File ---")

	// OpenFile lets us specify flags.
	//
	// O_APPEND = Write at the end of the file.
	// O_CREATE = Create the file if missing.
	// O_WRONLY = Open for writing only.
	file, err := os.OpenFile(
		"example.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Close the file when this function returns.
	defer file.Close()

	_, err = file.WriteString("\nNew appended line")

	if err != nil {
		fmt.Println("Error appending:", err)
		return
	}

	fmt.Println("Content appended successfully")
}

// Expected output:
// --- 4. Append to File ---
// Content appended successfully

// ============================================
// 5. OPEN AND READ A FILE
// ============================================

func openAndReadFile() {
	fmt.Println("\n--- 5. Open and Read ---")

	// Open opens an existing file for reading.
	file, err := os.Open("example.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the opened file into memory.
	data, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println(string(data))
}

// Expected output includes the appended line:
// --- 5. Open and Read ---
// Hello from Go!
// Learning file handling.
// New appended line

// ============================================
// 6. READ A FILE LINE BY LINE
// ============================================

func readLineByLine() {
	fmt.Println("\n--- 6. Read Line by Line ---")

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Scanner reads the file one line at a time.
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Scanner may stop because of a read error.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning:", err)
		return
	}
}

// Expected output:
// --- 6. Read Line by Line ---
// Hello from Go!
// Learning file handling.
// New appended line

// ============================================
// 7. CREATE A DIRECTORY
// ============================================

func createDirectory() {
	fmt.Println("\n--- 7. Create Directory ---")

	// Mkdir creates one directory.
	// MkdirAll creates parent directories if needed.
	err := os.MkdirAll("myfolder/reports", 0755)

	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	fmt.Println("Directory created successfully")
}

// Expected output:
// --- 7. Create Directory ---
// Directory created successfully

// ============================================
// 8. COPY A FILE
// ============================================

func copyFile(source, destination string) error {
	// Open the source file.
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	// Create or replace the destination file.
	dst, err := os.Create(destination)
	if err != nil {
		return err
	}

	// Close the destination and report a close error
	// if no earlier copy error occurred.
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return dst.Close()
}

func demonstrateCopy() {
	fmt.Println("\n--- 8. Copy File ---")

	err := copyFile("example.txt", "myfolder/reports/copy.txt")

	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

	fmt.Println("File copied successfully")
}

// Expected output:
// --- 8. Copy File ---
// File copied successfully

// ============================================
// 9. RENAME OR MOVE A FILE
// ============================================

func renameOrMoveFile() {
	fmt.Println("\n--- 9. Rename or Move ---")

	// Rename the copied file in the same directory.
	err := os.Rename(
		"myfolder/reports/copy.txt",
		"myfolder/reports/renamed.txt",
	)

	if err != nil {
		fmt.Println("Error renaming file:", err)
		return
	}

	fmt.Println("File renamed successfully")
}

// Expected output:
// --- 9. Rename or Move ---
// File renamed successfully
//
// Important:
// os.Rename can move a file to another directory
// on the same filesystem.
// Moving across filesystems may fail on some systems.
// Cross-filesystem moves may require copying and deleting.

// ============================================
// 10. GET FILE INFORMATION
// ============================================

func fileInformation() {
	fmt.Println("\n--- 10. File Information ---")

	info, err := os.Stat("example.txt")

	if err != nil {
		fmt.Println("Error getting information:", err)
		return
	}

	fmt.Println("Name:", info.Name())
	fmt.Println("Size in bytes:", info.Size())
	fmt.Println("Is directory:", info.IsDir())
	fmt.Println("Permissions:", info.Mode().Perm())
}

// Expected output starts with:
// --- 10. File Information ---
// Name: example.txt
// Size in bytes: 65
// Is directory: false
// Permissions: -rw-r--r--
//
// The size depends on the file's contents.
// The example above contains appended text;
// check the actual size when you run the program.

// ============================================
// 11. CHECK WHETHER A FILE EXISTS
// ============================================

func checkFileExists() {
	fmt.Println("\n--- 11. Check File Exists ---")

	_, err := os.Stat("example.txt")

	if err == nil {
		fmt.Println("File exists")
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist")
	} else {
		fmt.Println("Could not check file:", err)
	}
}

// Expected output:
// --- 11. Check File Exists ---
// File exists

// ============================================
// 12. DELETE A FILE
// ============================================

func deleteFile(path string) error {
	return os.Remove(path)
}

func demonstrateDelete() {
	fmt.Println("\n--- 12. Delete File ---")

	err := deleteFile("myfolder/reports/renamed.txt")

	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	fmt.Println("File deleted successfully")
}

// Expected output:
// --- 12. Delete File ---
// File deleted successfully

// ============================================
// 13. MAIN FUNCTION
// ============================================

func main() {
	// Run these examples in order because later
	// examples use files created by earlier examples.

	createAndWriteFile()
	readEntireFile()
	appendToFile()
	openAndReadFile()
	readLineByLine()
	createDirectory()
	demonstrateCopy()
	renameOrMoveFile()
	fileInformation()
	checkFileExists()
	demonstrateDelete()

	// Print the absolute path of the working directory.
	dir, err := os.Getwd()
	if err == nil {
		fmt.Println("\nWorking directory:", filepath.Clean(dir))
	}
}
