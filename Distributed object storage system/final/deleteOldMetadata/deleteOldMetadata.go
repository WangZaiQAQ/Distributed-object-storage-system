package main

import (
	"../../src/lib/es"
	"log"
)

// 函数的作用是删除所有版本计数大于 MIN_VERSION_COUNT 的元数据
const MIN_VERSION_COUNT = 5

func main() {
	buckets, e := es.SearchVersionStatus(MIN_VERSION_COUNT + 1)
	if e != nil {
		log.Println(e)
		return
	}
	for i := range buckets {
		bucket := buckets[i]
		for v := 0; v < bucket.Doc_count-MIN_VERSION_COUNT; v++ {
			es.DelMetadata(bucket.Key, v+int(bucket.Min_version.Value))
		}
	}
}
