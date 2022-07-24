package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	// ErrorExistUser 成员错误
	ErrorExistUser      = 10002
	ErrorNotExistUser   = 10003
	ErrorFailEncryption = 10006
	ErrorNotCompare     = 10007

	ErrorAuthCheckTokenFail    = 30001 //token 错误
	ErrorAuthCheckTokenTimeout = 30002 //token 过期
	ErrorAuthToken             = 30003
	ErrorAuth                  = 30004
	ErrorDatabase              = 40001

	// ErrorExistType 分类错误
	ErrorExistType = 50001

	// ErrorExistTag 标签错误
	ErrorExistTag = 60001

	// ErrorExistBlog 博客错误
	ErrorExistBlog = 70001
)
