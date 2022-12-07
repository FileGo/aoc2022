package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	typeFile = 1
	typeDir  = 2
)

type filesytemEntry struct {
	absolutePath       string
	parent             *filesytemEntry
	parentAbsolutePath string
	children           []*filesytemEntry
	name               string
	fileType           int
	size               int64
	totalSize          int64
}

func sortByName(slice []*filesytemEntry) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].name < slice[j].name
	})
}

func sortBySize(slice []*filesytemEntry) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].totalSize < slice[j].totalSize
	})
}

func getDirTotalSize(input *filesytemEntry, totalSize *int64) {
	if input.fileType == typeFile {
		return
	}

	for _, child := range input.children {
		*totalSize += child.size

		if len(child.children) > 0 {
			getDirTotalSize(child, totalSize)
		}
	}

}

func main() {
	buf, err := os.ReadFile("puzzle.dat")
	if err != nil {
		panic(err)
	}

	var allFsEntries []*filesytemEntry

	rootFs := filesytemEntry{
		fileType:     typeDir,
		absolutePath: "/",
		name:         "/",
	}

	allFsEntries = append(allFsEntries, &rootFs)

	currentDir := "/"
	previousCommand := ""
	for _, line := range strings.Split(string(buf), "\n") {
		if line[0] == '$' {
			// Command
			if strings.Contains(line, "cd") {
				dir := strings.Replace(line, "$ cd ", "", -1)
				if dir == ".." {
					if currentDir == "/" {
						continue
					}

					lastSlash := strings.LastIndex(currentDir, "/")
					currentDir = currentDir[:lastSlash]
					if currentDir == "" {
						currentDir = "/"
					}
				} else if dir == "." {
					continue
				} else {
					if dir[0] == '/' {
						if currentDir == "/" {
							currentDir = dir
						} else {
							currentDir += dir
						}
					} else {
						if currentDir == "/" {
							currentDir += dir
						} else {
							currentDir += "/" + dir
						}
					}
				}
			}

			previousCommand = line[2:]
		} else {
			// Output
			if strings.Contains(previousCommand, "ls") {
				var newEntry filesytemEntry

				fields := strings.Split(line, " ")
				f0 := fields[0]
				f1 := fields[1]

				newEntry.name = f1
				if currentDir == "/" {
					newEntry.absolutePath = currentDir + f1
				} else {
					newEntry.absolutePath = currentDir + "/" + f1
				}

				if f0 == "dir" {
					newEntry.fileType = typeDir
				} else {
					newEntry.fileType = typeFile
					newEntry.size, err = strconv.ParseInt(f0, 10, 64)
					if err != nil {
						panic(err)
					}
				}

				// Find parent
				for _, fsEntry := range allFsEntries {
					if fsEntry.absolutePath == currentDir {
						newEntry.parent = fsEntry
						fsEntry.children = append(fsEntry.children, &newEntry)
						sortByName(fsEntry.children)
					}
				}

				allFsEntries = append(allFsEntries, &newEntry)
			}
		}
	}

	sortByName(allFsEntries)

	// Calculate total sizes
	for _, fsEntry := range allFsEntries {
		if fsEntry.fileType == typeDir {
			getDirTotalSize(fsEntry, &fsEntry.totalSize)
		}
	}

	// Part 1
	var totalSum int64

	for _, fsEntry := range allFsEntries {
		if fsEntry.fileType == typeDir {
			if fsEntry.totalSize < 1e5 {
				totalSum += fsEntry.totalSize
			}
		}
	}

	fmt.Printf("Total sizes (part1): %d\n", totalSum)

	// Part 2
	fmt.Println("Part 2")
	sortBySize(allFsEntries)

	for _, fsEntry := range allFsEntries {
		if fsEntry.fileType == typeDir {

			if 7e7-rootFs.totalSize+fsEntry.totalSize >= 3e7 {
				fmt.Printf("Path: %s Size: %d\n", fsEntry.absolutePath, fsEntry.totalSize)
				break
			}
		}
	}
}
