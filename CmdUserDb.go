package main

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/urfave/cli"
)

func CmdUserDb(c *cli.Context) {
	fn := c.String("i")
	fmt.Println("open file ", fn)
	db, err := bolt.Open(fn, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	tx, err := db.Begin(true)
	if err != nil {
		panic(err)
	}
	fmt.Println("sheet id:")
	sheetIDBucket, err := tx.CreateBucketIfNotExists([]byte("sheetID"))
	if err != nil {
		panic(err)
	}
	sheetIDBucket.ForEach(func(k []byte, v []byte) error {
		fmt.Println(string(k), string(v))
		return nil
	})
	refreshTokenBucket, err := tx.CreateBucketIfNotExists([]byte("refreshToken"))
	if err != nil {
		panic(err)
	}
	fmt.Println("refresh token:")
	refreshTokenBucket.ForEach(func(k []byte, v []byte) error {
		fmt.Println(string(k), string(v))
		return nil
	})
}
