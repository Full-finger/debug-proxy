package handler

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type proxyReq struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body,omitempty"`
}

type proxyResp struct {
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
	TimeMs  int64             `json:"timeMs"`
	Error   string            `json:"error,omitempty"`
}

func ProxyRequest(c *gin.Context) {
	var req proxyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := &http.Client{Timeout: 30 * time.Second}
	start := time.Now()

	httpReq, _ := http.NewRequest(req.Method, req.URL, bytes.NewBufferString(req.Body))
	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}

	resp, err := client.Do(httpReq)
	ms := time.Since(start).Milliseconds()
	if err != nil {
		c.JSON(http.StatusOK, proxyResp{Error: err.Error(), TimeMs: ms})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	headers := make(map[string]string)
	for k, vs := range resp.Header {
		headers[k] = vs[0]
	}

	c.JSON(http.StatusOK, proxyResp{
		Status:  resp.StatusCode,
		Headers: headers,
		Body:    string(body),
		TimeMs:  ms,
	})
}
