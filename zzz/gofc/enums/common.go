package enum

//go:generate stringer -type CommonStatus -linecomment -output common_enum.go
type CommonStatus int32

// invalid err code
const (
	ErrInvalidCustom   CommonStatus = errOffsetInvalid + iota // 自定义参数
	ErrInvalidParam                                           // 无效的参数
	ErrInvalidData                                            // 无效的数据
	ErrInvalidAccount                                         // 无效的账号
	ErrInvalidPhone                                           // 无效的手机号
	ErrInvalidEmail                                           // 无效的邮箱
	ErrInvalidUrl                                             // 无效的网址
	ErrInvalidImg                                             // 无效的图片
	ErrInvalidSign                                            // 无效的签名
	ErrInvalidDate                                            // 无效的日期
	ErrInvalidIdcardno                                        // 无效的身份证号
	ErrInvalidBirthday                                        // 无效的生日格式
	ErrInvalidPassword                                        // 密码只支持6-20位
)

// common err code
const (
	ErrCommonUnknownError    CommonStatus = errOffsetCommon + iota // 未知错误
	ErrCustomMsg                                                   // 自定义错误信息
	ErrCommonUnknownRequest                                        // 未知请求
	ErrCommonRequestFrequent                                       // 请求太频繁，请稍后再试
	ErrCommonMarshal                                               // 网络传输失败，请稍后重试
	ErrCommonUnmarshal                                             // 网络传输失败，请稍后重试
	ErrCommonRemotecall                                            // 服务通信失败，请稍后重试
	ErrCommonMysql                                                 // 获取数据失败，请稍后再试
	ErrCommonRedis                                                 // 获取数据失败，请稍后再试
	ErrCommonUploadFile                                            // 上传文件失败
)
