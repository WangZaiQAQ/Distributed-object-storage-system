package main

import (
	"../../src/lib/es"
	"../../src/lib/utils"
	"../apiServer/objects"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
遍历存储目录下的所有文件，并重新检验文件的哈希值，如果不一致则则打印日志
*/

func main() {
	files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/*")

	for i := range files {
		hash := strings.Split(filepath.Base(files[i]), ".")[0]
		verify(hash)
	}
}

func verify(hash string) {
	log.Println("verify", hash)
	size, e := es.SearchHashSize(hash)
	if e != nil {
		log.Println(e)
		return
	}
	stream, e := objects.GetStream(hash, size)
	if e != nil {
		log.Println(e)
		return
	}
	d := utils.CalculateHash(stream)
	if d != hash {
		log.Printf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	stream.Close()
}
