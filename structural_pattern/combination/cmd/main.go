package main

import "godesignpattern/structural_pattern/combination"

func main() {
	file1 := &combination.File{Name: "File1"}
	file2 := &combination.File{Name: "File2"}
	file3 := &combination.File{Name: "File3"}

	folder1 := &combination.Folder{
		Name: "Folder1",
	}
	folder1.Add(file1)

	folder2 := &combination.Folder{
		Name: "Folder2",
	}

	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
