package lib

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Itob method, transfer int32 to bytes.
func Itob(v int32) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// Logger method
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// Md5 method
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// GetID method
func GetID(path string) int32 {
	p := strings.Split(path, "/")
	id, _ := strconv.Atoi(p[len(p)-1])
	return int32(id)
}

//GetRandom method
func GetRandom() int32 {
	rand.Seed(int64(time.Now().UnixNano()))
	return int32(rand.Int31n(65536))
}

// CallMethod method
func CallMethod(item string) []byte {
	cmd := exec.Command("sh", "-c", "cd python-sdk && python3 console.py call Company 0x1f494c56c3ad1e6738f3500d19499cd3541160ea "+item)
	outfile, err := os.Create("out")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	cmd.Wait()

	b, err := ioutil.ReadFile("out")
	if err != nil {
		fmt.Print(err)
	}
	return b
}

// SendtxMethod method
func SendtxMethod(item string) string {
	cmd := exec.Command("sh", "-c", "cd python-sdk && python3 console.py sendtx Company 0x1f494c56c3ad1e6738f3500d19499cd3541160ea "+item)
	outfile, err := os.Create("out")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	cmd.Wait()
	b, err := ioutil.ReadFile("out")
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
