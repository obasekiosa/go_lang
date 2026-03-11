package main

import (
	"container/list"
	"crypto/aes"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	// "github.com/Workiva/go-datastructures/queue"
	"github.com/zyedidia/generic/heap"
	"golang.org/x/crypto/argon2"
)

type Item struct {
	val int
}

func (i Item) Compare(other Item) bool {
	return i.val < other.val
}

type Pq []*Item

func (pq Pq) Len() int {
	return len(pq)
}

// func (pq Pq) Less(i, j int) bool {
// 	return pq[i].val < pq[j].val
// }

func (pq Pq) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// func (pq Pq) Pop

func randomNumber() {
	var b [32]byte
	rand.Read(b[:])

	fmt.Printf("%d\n", b)
	// fmt.Println(int32(b)) // convert []byte to int/float etc

	msg := []byte("hello")
	key := make([]byte, 32)
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)
	tag := mac.Sum(nil)

	mac2 := hmac.New(sha256.New, key)
	mac2.Write(msg)

	ok := hmac.Equal(tag, mac2.Sum(nil))

	fmt.Println(ok)

	p_hash := argon2.IDKey([]byte("password"), []byte("salt"), 2, 32*1064, 8, 32)

	fmt.Printf("%X\n", p_hash)

	cip, _ := aes.NewCipher(key)
	base := [32]byte{'m', 'e', 's', 's', 'a', 'g', 'e'}
	src := base[:]

	fmt.Println("original", src, string(src))
	dst := make([]byte, 256)
	cip.Encrypt(dst, src)

	fmt.Println("Encrypt ->", dst)

	dcr := make([]byte, 32)
	cip.Decrypt(dcr, dst)
	fmt.Println("decrypt ->", dcr, string(dcr))
	fmt.Println(string(src) == string(dcr))

	// chacha20poly1305

	pub, priv, _ := ed25519.GenerateKey(rand.Reader)

	sig := ed25519.Sign(priv, msg)
	ok = ed25519.Verify(pub, msg, sig)

	// fmt.Println(ok, sig, "\n", pub, "\n", priv)

	// conn, _ := tls.Dial("tcp", "www.google.com:443", nil)

	// req := "GET / HTTP/1.1\r\nHost: www.google.com\r\n\r\n"
	// wrote, _ := conn.Write([]byte(req))
	// fmt.Println("Wrote", wrote)

	// buff := make([]byte, 4096)
	// n, _ := conn.Read(buff)
	// fmt.Println(string(buff[:n]))

	// resp, _ := http.Get("https://www.google.com")

	// fmt.Println(resp.Header)

	l := list.New()

	l.PushBack(3)
	l.PushBack(true)
	item := l.Front()

	fmt.Println(item.Value)

	// var i heap.Interface
	// heap.Init()

	// slices.Sort(dcr)
	// fmt.Println(dcr)

	// pq := queue.NewPriorityQueue(10, false) // true = min-heap
	// var it queue.Item
	// it = Item{val: 2}
	// pq.Put(it)

	pq := heap.New(Item.Compare)

	for _, v := range dcr {
		pq.Push(Item{int(v)})
	}

	fmt.Println(pq)

	for range pq.Size() {
		v, e := pq.Pop()
		fmt.Println(v, e, pq.Size())
	}
}
