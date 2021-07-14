package errs

import "github.com/chaosnote/go-kernel/net/http/response"

// 伺服器錯誤清單
const (
	Err0000 response.Status = "0000" // 伺服器內部錯誤
	Err0001 response.Status = "0001" // 客端參數錯誤
	Err0002 response.Status = "0002" // 帳號密碼錯誤
	Err0003 response.Status = "0003" // 非法請求
)
