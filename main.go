package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "time"
)

func main() {
    searchPath := "."
	listFiles(searchPath)
}

func listFiles(searchPath string) {
    fis, err := ioutil.ReadDir(searchPath)
    if err != nil {
        panic(err)
    }
    file, err := os.OpenFile("./" + time.Now().Format("20060102150405") + ".txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
    for _, fi := range fis {
        searchFilePath := filepath.Join(searchPath, fi.Name())
        tab := ""
        if fi.IsDir() {
            displayFile(searchFilePath, tab, file)
			listFiles(searchFilePath)
        } else {
            if err != nil {
                panic(err)
            }
            displayFile(searchFilePath, tab, file)
        }
    }
    defer file.Close()
}

func displayFile(searchFilePath string, tab string, file *os.File) {
    countBackSlash := strings.Count(searchFilePath, "\\")
    if countBackSlash == 0 {
        fmt.Fprintln(file, searchFilePath)
    } else {
        for i:=0; i < countBackSlash; i++ {
            tab = tab + "\t"
        }
        filePath := strings.Split(searchFilePath, "\\")
        fileName := filePath[len(filePath)-1]
        fmt.Fprintln(file, tab + "-" + fileName)
    }
}