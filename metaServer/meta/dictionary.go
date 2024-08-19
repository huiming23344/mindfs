package meta

import (
	"fmt"
	"strings"
)

type Directory struct {
	Name      string
	SubDirs   map[string]*Directory
	Files     map[string]*File
	ParentDir *Directory
	Status    Status
}

// NewDirectory 创建一个新的目录实例
func NewDirectory(name string) *Directory {
	return &Directory{
		Name:    name,
		SubDirs: make(map[string]*Directory),
		Files:   make(map[string]*File),
	}
}

// AddDir 添加子目录
func (d *Directory) AddDir(name string) error {
	if _, exists := d.SubDirs[name]; exists {
		return fmt.Errorf("directory '%s' already exists", name)
	}
	d.SubDirs[name] = NewDirectory(name)
	d.SubDirs[name].ParentDir = d
	return nil
}

// AddFile 添加文件
func (d *Directory) AddFile(name string) error {
	if _, exists := d.Files[name]; exists {
		return fmt.Errorf("file '%s' already exists", name)
	}
	d.Files[name] = &File{Name: name}
	return nil
}

// FindDir 查找目录
func (d *Directory) FindDir(path string) (*Directory, error) {
	parts := strings.Split(path, "/")
	currentDir := d
	for _, part := range parts {
		if part == "" { // 忽略空部分，例如路径开始的斜杠
			continue
		}
		nextDir, exists := currentDir.SubDirs[part]
		if !exists {
			return nil, fmt.Errorf("directory '%s' not found", part)
		}
		currentDir = nextDir
	}
	return currentDir, nil
}

// Print 打印目录结构
func (d *Directory) Print(indent string) {
	fmt.Println(indent + d.Name + "/")
	for _, dir := range d.SubDirs {
		dir.Print(indent + "  ")
	}
	for _, file := range d.Files {
		fmt.Println(indent + "  " + file.Name)
	}
}

// DeleteFile 删除文件
func (d *Directory) DeleteFile(fileName string) error {
	if _, exists := d.Files[fileName]; !exists {
		return fmt.Errorf("file '%s' not found", fileName)
	}
	delete(d.Files, fileName)
	return nil
}

// DeleteDir 删除目录
func (d *Directory) DeleteDir(dirName string) error {
	if _, exists := d.SubDirs[dirName]; !exists {
		return fmt.Errorf("directory '%s' not found", dirName)
	}
	delete(d.SubDirs, dirName)
	return nil
}

// Move 移动文件或目录
func (d *Directory) Move(sourcePath, destPath string) error {
	// 解析源路径和目标路径
	srcParts := strings.Split(sourcePath, "/")
	destParts := strings.Split(destPath, "/")

	// 获取源文件或目录
	srcDir, srcName, err := d.getDirAndNameFromPath(srcParts)
	if err != nil {
		return err
	}

	// 获取目标目录
	destDir, _, err := d.getDirAndNameFromPath(destParts)
	if err != nil {
		return err
	}

	// 检查源文件或目录是否存在
	if _, fileExists := srcDir.Files[srcName]; fileExists {
		// 移动文件
		if _, exists := destDir.Files[srcName]; exists {
			return fmt.Errorf("file '%s' already exists in destination", srcName)
		}
		destDir.Files[srcName] = srcDir.Files[srcName]
		delete(srcDir.Files, srcName)
	} else if _, dirExists := srcDir.SubDirs[srcName]; dirExists {
		// 移动目录
		if _, exists := destDir.SubDirs[srcName]; exists {
			return fmt.Errorf("directory '%s' already exists in destination", srcName)
		}
		destDir.SubDirs[srcName] = srcDir.SubDirs[srcName]
		srcDir.SubDirs[srcName].ParentDir = destDir
		delete(srcDir.SubDirs, srcName)
	} else {
		return fmt.Errorf("source '%s' not found", sourcePath)
	}

	return nil
}

// getDirAndNameFromPath 根据路径获取目录和文件/目录名
func (d *Directory) getDirAndNameFromPath(parts []string) (*Directory, string, error) {
	var currentDir *Directory = d
	var name string
	for i, part := range parts {
		if part == "" {
			continue
		}
		if i == len(parts)-1 {
			name = part
		} else {
			var exists bool
			currentDir, exists = currentDir.SubDirs[part]
			if !exists {
				return nil, "", fmt.Errorf("path '%s' not found", strings.Join(parts[:i+1], "/"))
			}
		}
	}
	return currentDir, name, nil
}
