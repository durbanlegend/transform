package ops

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"github.com/durbanlegend/transform/parser"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	//"reflect"
	//"strconv"
	"strings"
	"time"
)

func CheckError(err error) {
	if err != nil {
		log.Fatalf("Get: %v", err)
	}
}

// Write file to a leveldb database for later reading
func CacheFile(trans parser.TransType) {

	f, err := os.Open(trans.Infile)
	CheckError(err)

	delim := ";"
	scanner := bufio.NewScanner(f)
	defer f.Close()

	// Set up a Bloom filter for cache
	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}

	err = os.RemoveAll(trans.Outfile)
	CheckError(err)

	db, err := leveldb.OpenFile(trans.Outfile, o)
	CheckError(err)

	// open message log
	fe, err := os.Create("/dev/stderr")
	CheckError(err)

	ew := bufio.NewWriter(fe)
	defer ew.Flush()

	start := time.Now()

	i := 0

	batch := new(leveldb.Batch)
	var buffer bytes.Buffer
	keylen := len(trans.Keynames)
	fmt.Printf("keylen=%v; keyNames=%s\n", keylen, trans.Keynames)
	keyPositions := make([]int, keylen)
	for i, keyname := range trans.Keynames {
		for j, field := range trans.Inrec.Infields {
			//fmt.Printf("i=%v; j=%v; field.Name=%s; keyname=%s\n", i, j, field.Name, keyname)
			if field.Name == keyname {
				keyPositions[i] = j
			}
		}
	}
	fmt.Printf("keyPositions=%v\n", keyPositions)

	for scanner.Scan() {
		i++
		CheckError(scanner.Err())

		arr := strings.Split(scanner.Text(), delim)

		buffer.Reset()

		//keylen := len(trans.Keynames)
		for _, k := range keyPositions {
			buffer.WriteString(arr[k])
			buffer.WriteString(delim)
		}
		//fmt.Printf("key=%s\n", buffer.String())
		batch.Put(buffer.Bytes(), scanner.Bytes())
	}
	err = db.Write(batch, nil)
	CheckError(err)
	defer db.Close()

	end := time.Now()
	delta := end.Sub(start)
	//fmt.Println("Full transform took”, delta)
	fmt.Fprintln(ew, "Elapsed:", delta)
	fmt.Fprintln(ew, "Rows processed:", i)
}
