package main

import (
	"fmt"
	"hash"
	"math"
	"github.com/spaolacci/murmur3"
)

//aliases
var log = fmt.Println
// var lg = fmt.Print
var cin = fmt.Scanln

type Bloom struct {
	m      uint32 //size of bitset
	k      uint32 //no. of hash fns
	bitset []bool
	hashes []hash.Hash32
}


func New(n uint32, p float64) *Bloom {
	m := M(n, p)
	k := K(m, n)
	bitset := make([]bool, m)
	hashes := make([]hash.Hash32, k)

	var i uint32;
	for ; i<k ; i++ {
		hashes[i] = murmur3.New32WithSeed(i)
	}

	bloom := Bloom{
		m: m,
		k: k,
		bitset: bitset,
		hashes: hashes,
	}

	return &bloom
}

func (b *Bloom) Add(key []byte) {
	for _, h := range b.hashes {
		h.Reset()
		_ ,err := h.Write(key)

		if err != nil{
			log("Error writing to hasher")
			return
		}

		idx := h.Sum32() % b.m
		b.bitset[idx] = true
	}
}

func (b *Bloom) Contains(key []byte) bool{
	for _, h := range b.hashes {
		h.Reset()
		_ ,err := h.Write(key)

		if err != nil{
			log("Error writing to hasher")
		}

		idx := h.Sum32() % b.m
		if !b.bitset[idx] {
			return false
		}
	}
	return true
}

func M(n uint32, p float64) uint32 {
	return uint32(math.Ceil(-((float64(n) * math.Log(p)) / (math.Pow(math.Log(2), 2)))))
}

func K(m, n uint32) uint32 {
	return uint32(math.Ceil(math.Log(2) * float64(m) / float64(n)))
}

func main() {

	log("Bloom filter")
	var n uint32
	var p float64
	log("Define estimated no. of userIDs and tolerable fp(0-1): ")
	cin(&n, &p)
	b := New(n, p)

	// add
	b.Add([]byte("username1"))
	b.Add([]byte("username2"))

	//check
	log("=========Check=========")
	log(b.Contains([]byte("username1")))
	log(b.Contains([]byte("username2")))

	log(b.Contains([]byte("username3")))




}
