package main

import (
	"fmt"
	"sync"
	"time"
)

type File struct {
	Name   string
	SizeMB int
}

type Downloader struct {
	Files []File
}

func (self *Downloader) AddFile(name string, size int) {
	file := File{Name: name, SizeMB: size}
	self.Files = append(self.Files, file)
}

func downloader(name string, size int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Downloading %v (%vMB)...\n", name, size)
	for i := 0; i < size; i++ {
		time.Sleep(1 * time.Millisecond)
	}
}

func (self *Downloader) DownloadAll() {
	var wg sync.WaitGroup

	for i := range self.Files {
		wg.Add(1)
		go downloader(self.Files[i].Name, self.Files[i].SizeMB, &wg)
	}

	wg.Wait()
	for i := range self.Files {
		fmt.Printf("Download finished: %v\n", self.Files[i].Name)
	}
}

func main() {
	var downloader Downloader

	downloader.AddFile("image.png", 12)
	downloader.AddFile("bigtext.txt", 10)
	downloader.AddFile("app.exe", 231)
	downloader.AddFile("test.zip", 172)

	downloader.DownloadAll()

}
