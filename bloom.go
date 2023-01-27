package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"hash"
	"hash/fnv"
	"math"
)

type BloomFilter struct {
	bitArray []bool
	k        int         // number of hash functions
	n        int         // number of inserted elements
	m        int         // size of bit array
	h        hash.Hash64 // the hash function to use
}

func NewBloomFilter(capacity int, falsePositiveRate float64) (*BloomFilter, error) {
	if capacity <= 0 {
		return nil, errors.New("filter capacity must be more than 0")
	}

	if falsePositiveRate <= 0 || falsePositiveRate >= 1 {
		return nil, errors.New("falsePositiveRate must be between 0 and 1")
	}

	m := int(math.Ceil(-float64(capacity) * math.Log(falsePositiveRate) / math.Pow(math.Log(2), 2)))
	k := int(math.Ceil(float64(m) / float64(capacity) * math.Log(2)))

	return &BloomFilter{
		bitArray: make([]bool, m),
		k:        k,
		n:        0,
		m:        m,
		h:        fnv.New64(),
	}, nil
}

func (bf *BloomFilter) Add(item any) error {
	b, err := toBytes(item)
	if err != nil {
		return err
	}

	for i := 0; i < bf.k; i++ {
		bf.h.Reset()
		bf.h.Write(b)
		bf.h.Write([]byte{byte(i)})
		h := bf.h.Sum64() % uint64(bf.m)
		bf.bitArray[h] = true
	}
	bf.n++
	return nil
}

func (bf *BloomFilter) Exist(item any) (bool, error) {
	b, err := toBytes(item)
	if err != nil {
		return false, err
	}

	for i := 0; i < bf.k; i++ {
		bf.h.Reset()
		bf.h.Write(b)
		bf.h.Write([]byte{byte(i)})
		h := bf.h.Sum64() % uint64(bf.m)
		if !bf.bitArray[h] {
			return false, err
		}
	}
	return true, nil
}

func toBytes(item any) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(item)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
