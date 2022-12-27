package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	CMD              = "$"
	CD               = "cd"
	LS               = "ls"
	DIR              = "dir"
	PARENT_DIR_ALIAS = ".."
	AVAILABLE_SPACE  = 70000000
	SPACE_NEEDED     = 30000000
)

type (
	Directory struct {
		name            string
		parentDirectory *Directory
		subdirectories  map[string]*Directory
		files           []*File
	}

	File struct {
		name string
		size int
	}
)

func main() {
	lines := openInputFile()
	fs := constructFs(lines)
	root := fs.NavigateToRootDir()

	log.Printf("Part 1: The sum of the size of directories w/ size <= 100000 is %v", root.SumDirectorySizeUnderMaxSize(100000))

	spaceAvailable := AVAILABLE_SPACE - root.CalculateSize()
	freeSpaceNeededForUpdate := SPACE_NEEDED - spaceAvailable
	log.Printf("Part 2: The smallest directory over 30000000 is %v", root.DetermineSmallestEligibleDirSize(freeSpaceNeededForUpdate))
}

/*
Constructs the layout of the filesystem based on commands and ls output given in input file
*/
func constructFs(lines []string) *Directory {
	var currentDir *Directory
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, CMD):
			if strings.HasPrefix(line, fmt.Sprintf("%v %v", CMD, CD)) {
				currentDir = changeDirectories(currentDir, line)
			}
		case strings.HasPrefix(line, DIR):
			dirName := strings.Split(line, " ")[1]
			currentDir.AddSubdirectory(dirName)
		default:
			splitLine := strings.Split(line, " ")
			size, _ := strconv.Atoi(splitLine[0])
			name := splitLine[1]
			file := &File{name, size}
			currentDir.AddFile(file)
		}
	}
	return currentDir
}

/*
Given some directory currentDir and cd cmd, return the Directory object associated w/ the named dir in the cmd (EX 'abcd' for cmd $ cd abcd)
The name of the directory given in cmd must either be a subdirectory of currentDir or must be '..' indicating
that the directory should be changed to the parent directory of currentDir
*/
func changeDirectories(currentDir *Directory, cmdStr string) *Directory {
	var dir *Directory

	cmd := strings.Split(cmdStr, " ")
	cdDirName := cmd[2]

	if currentDir != nil {
		if cdDirName == PARENT_DIR_ALIAS {
			dir = currentDir.parentDirectory
		} else {
			dir = currentDir.subdirectories[cdDirName]
		}
	} else {
		// Initialize root directory
		dir = &Directory{cdDirName, nil, map[string]*Directory{}, []*File{}}
	}
	return dir
}

func (dir *Directory) AddFile(file *File) {
	dir.files = append(dir.files, file)
}

func (dir *Directory) AddSubdirectory(subDirName string) {
	dir.subdirectories[subDirName] = &Directory{subDirName, dir, map[string]*Directory{}, []*File{}}
}

/*
Calculates the size of given directory based on the size of its child files and subdirectories
*/
func (dir *Directory) CalculateSize() int {
	var dirSize int
	for _, file := range dir.files {
		dirSize += file.size
	}
	for _, subDir := range dir.subdirectories {
		dirSize += subDir.CalculateSize()
	}
	return dirSize
}

func (dir *Directory) GetFullPath() string {
	return dir.getFullPathHelper(dir.name)
}

func (dir *Directory) getFullPathHelper(accum string) string {
	if dir.parentDirectory == nil {
		return accum
	}
	return dir.parentDirectory.getFullPathHelper(fmt.Sprintf("%v/%v", dir.parentDirectory.name, accum))
}

/*
Traverses directory structure and returns root directory - IE directory w/ no parent
*/
func (dir *Directory) NavigateToRootDir() *Directory {
	if dir.parentDirectory != nil {
		return dir.parentDirectory.NavigateToRootDir()
	}
	return dir
}

func (dir *Directory) SumDirectorySizeUnderMaxSize(maxSize int) int {
	var sum int
	if dir.CalculateSize() <= maxSize {
		sum += dir.CalculateSize()
	}
	for _, subdir := range dir.subdirectories {
		sum += subdir.SumDirectorySizeUnderMaxSize(maxSize)
	}
	return sum
}

func (dir *Directory) DetermineSmallestEligibleDirSize(minSize int) int {
	return dir.DetermineSmallestEligibleDirSizeHelper(minSize, dir.CalculateSize())
}

func (dir *Directory) DetermineSmallestEligibleDirSizeHelper(minSize int, minObserved int) int {
	dirSize := dir.CalculateSize()
	if dirSize >= minSize {
		println(dirSize, dir.name)

	}
	if dirSize >= minSize && dirSize < minObserved {
		minObserved = dirSize
	}

	for _, subdir := range dir.subdirectories {
		minObserved = subdir.DetermineSmallestEligibleDirSizeHelper(minSize, minObserved)
	}
	return minObserved
}

/*
Parses input file for AOC 2022 day_7
*/
func openInputFile() []string {
	data, err := os.ReadFile("resources/input")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
