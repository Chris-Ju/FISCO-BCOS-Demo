package lib

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"log"
	"net/http"
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

//GetDID method
func GetDID(path string) (int32, int32) {
	p := strings.Split(path, "/")
	articleid, _ := strconv.Atoi(p[len(p)-2])
	userid, _ := strconv.Atoi(p[len(p)-1])
	return int32(articleid), int32(userid)
}
