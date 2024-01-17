package enum

//go:generate stringer -type CodeStatus -linecomment -output code_enum.go
type CodeStatus int

const (
	SUCCESS                        CodeStatus = 200 // 成功
	ERROR                          CodeStatus = 500 // 错误
	ErrorInvalidParamsWithoutToken CodeStatus = 401 // 请求参数错误
	ErrorInvalidParams             CodeStatus = 400 // 不存在Token参数

	// 认证
	ErrorAuthParseTokenFail    CodeStatus = iota + 20001 // Token解析失败
	ErrorAuthCheckTokenTimeout                           // Token已超时
	ErrorAuthGenerateToken                               // Token生成失败
	ErrorAuthToken                                       // Token错误

	// 客户端错误
	ErrorUserGetInfo  CodeStatus = iota + 40001 // 获取到用户失败
	ErrorUserGetLogin                           // 获取到帐户失败
	ErrorUserRegName                            // 用户名输入格式错误
)
