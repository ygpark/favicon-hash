package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func murmurHash3(data []byte, seed uint32) int32 {
	const (
		c1 = 0xcc9e2d51
		c2 = 0x1b873593
		r1 = 15
		r2 = 13
		m  = 5
		n  = 0xe6546b64
	)

	hash := seed
	length := uint32(len(data))
	
	// Process 4-byte blocks
	nblocks := len(data) / 4
	for i := 0; i < nblocks; i++ {
		k := uint32(data[i*4]) |
			uint32(data[i*4+1])<<8 |
			uint32(data[i*4+2])<<16 |
			uint32(data[i*4+3])<<24

		k *= c1
		k = (k << r1) | (k >> (32 - r1)) // rotl32(k, r1)
		k *= c2

		hash ^= k
		hash = (hash << r2) | (hash >> (32 - r2)) // rotl32(hash, r2)
		hash = hash*m + n
	}

	// Process remaining bytes
	tail := data[nblocks*4:]
	var k1 uint32
	switch len(tail) {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= c1
		k1 = (k1 << r1) | (k1 >> (32 - r1)) // rotl32(k1, r1)
		k1 *= c2
		hash ^= k1
	}

	// Finalization
	hash ^= length
	hash ^= hash >> 16
	hash *= 0x85ebca6b
	hash ^= hash >> 13
	hash *= 0xc2b2ae35
	hash ^= hash >> 16

	return int32(hash)
}

func main() {
	// 플래그 정의
	hexOutput := flag.Bool("hex", false, "16진수로 출력")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "사용법: %s [옵션] <favicon_url>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\n옵션:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n예시:\n")
		fmt.Fprintf(os.Stderr, "  %s http://203.245.0.121/favicon.ico\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -hex http://203.245.0.121/favicon.ico\n", os.Args[0])
	}
	flag.Parse()

	// 명령행 인수 확인
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	faviconURL := flag.Arg(0)

	// HTTP GET 요청
	response, err := http.Get(faviconURL)
	if err != nil {
		log.Fatal("HTTP 요청 실패:", err)
	}
	defer response.Body.Close()

	// 응답 본문 읽기
	content, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("응답 읽기 실패:", err)
	}

	// Base64 인코딩 (Python codecs.encode와 동일하게 줄바꿈 포함)
	favicon := base64.StdEncoding.EncodeToString(content)
	
	// Python의 codecs.encode처럼 76자마다 줄바꿈 추가
	var result []byte
	for i := 0; i < len(favicon); i += 76 {
		end := i + 76
		if end > len(favicon) {
			end = len(favicon)
		}
		result = append(result, favicon[i:end]...)
		result = append(result, '\n')
	}
	faviconWithNewlines := string(result)

	// MurmurHash3 계산 (Python mmh3.hash()와 호환)
	hashValue := murmurHash3([]byte(faviconWithNewlines), 0)

	// 결과 출력
	if *hexOutput {
		fmt.Printf("%x\n", uint32(hashValue))
	} else {
		fmt.Println(hashValue)
	}
}