package fetcher

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func TestFetch(t *testing.T) {
	const body = "<html><body>hello 城市</body></html>"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, body)
	}))
	defer ts.Close()

	got, err := Fetch(ts.URL)
	if err != nil {
		t.Fatalf("Fetch(%s) 返回错误: %v", ts.URL, err)
	}
	if string(got) != body {
		t.Errorf("内容不符\n期望: %q\n实际: %q", body, string(got))
	}
}

// TestFetchDecodesGBK 验证 determineEncoding 能按 meta charset 把 GBK 转成 UTF-8。
// 正文必须超过 1024 字节，否则 bufio.Peek(1024) 会失败并回退到 UTF-8。
func TestFetchDecodesGBK(t *testing.T) {
	const text = "城市列表"
	html := `<!DOCTYPE html><html><head><meta charset="gbk"></head><body>` +
		text + strings.Repeat("城", 600) + `</body></html>`

	encoded, err := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(html))
	if err != nil {
		t.Fatalf("编码成 GBK 失败: %v", err)
	}
	if len(encoded) <= 1024 {
		t.Fatalf("样本只有 %d 字节，Peek(1024) 会失败，测试无效", len(encoded))
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=gbk")
		w.Write(encoded)
	}))
	defer ts.Close()

	got, err := Fetch(ts.URL)
	if err != nil {
		t.Fatalf("Fetch(%s) 返回错误: %v", ts.URL, err)
	}
	if !strings.Contains(string(got), text) {
		t.Errorf("GBK 未被解码成 UTF-8，期望包含 %q\n实际: %q", text, string(got))
	}
}

func TestFetchRejectsNonOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	_, err := Fetch(ts.URL)
	if err == nil {
		t.Fatal("状态码 500 时应当返回错误，实际返回 nil")
	}
	if !strings.Contains(err.Error(), "500") {
		t.Errorf("错误信息应包含状态码 500，实际: %v", err)
	}
}

func TestFetchSetsUserAgent(t *testing.T) {
	var got string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Get("User-Agent")
		io.WriteString(w, "ok")
	}))
	defer ts.Close()

	if _, err := Fetch(ts.URL); err != nil {
		t.Fatalf("Fetch(%s) 返回错误: %v", ts.URL, err)
	}
	if !strings.Contains(got, "Mozilla") {
		t.Errorf("请求应带浏览器 User-Agent，实际: %q", got)
	}
}
