package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
)

// IPRateLimiter
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu *sync.RWMutex
	r rate.Limit
	b int
}

// NewIPRateLimiter
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	return i
}

// AddIP creates a new rate limiter and adds it to the ips map, using the IP address as the key
func (i *IPRateLimiter)AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)

	i.ips[ip] = limiter
	return limiter
}

// GetLimiter returns the rate limiter for the provided IP address if it exists. Otherwise calls AddIP to add IP address to the map
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]
	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}
	i.mu.Unlock()

	return limiter
}

var limiter = NewIPRateLimiter(1,5)


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	err := http.ListenAndServe(":8888", limitMiddleware(mux))
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := limiter.GetLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	// 某些消耗很高的数据库请求
	fmt.Println("alles")
	w.Write([]byte("alles gut"))
}


// Vegeta 压测工具
// echo "GET http://localhost:8888/" | vegeta attack -duration=10s -rate=100 |vegeta report
// https://www.hi-linux.com/posts/4650.html