package ops

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/durbanlegend/transform/parser"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func TestCacheFile(*testing.T) {
	home := os.Getenv("HOME")

	jsontype := parser.Unmarshal(home + "/person.json")
	CacheFile(jsontype)

	fmt.Printf("Unmarshalled %s: %v\n", reflect.TypeOf(jsontype), jsontype)
	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}
	db, err := leveldb.OpenFile(jsontype.Outfile, o)
	CheckError(err)

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		//// Seek-then-iterate
		//startKey := []byte("Don key 4")
		//for ok := iter.Seek(startKey); ok; ok = iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		//key := iter.Key()
		//value := iter.Value()
		//fmt.Printf("Key=%s; Value=%s\n", key, value)
	}
	iter.Release()
	err = iter.Error()
	CheckError(err)

	// Iterate over subset of database content:
	iter = db.NewIterator(&util.Range{Start: []byte("Thompson"), Limit: []byte("Thompsonz")}, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("Key=%s; Value=%s\n", key, value)
	}
	iter.Release()
	err = iter.Error()
	CheckError(err)

	defer db.Close()
}
