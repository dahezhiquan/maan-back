package errs

const (
	ParameterVerifyErrCode = 100002
)

// 系统/中间件错误 5xxxix
var (
	DBError     = NewError(500001, "DB错误")
	CacheError  = NewError(500002, "Cache错误")
	CopierError = NewError(500203, "模型转换错误")
)

// 校验错误 1xxxix

var (
	ParameterError = NewError(100001, "参数格式错误")
)

// 业务错误 2xxxix
var (

	// 全局错误

	ClockMovedBackwards      = NewError(200001, "系统时钟异常")
	WorkerIdExcessOfQuantity = NewError(200010, "ID超量")

	// NoticeNotExistError 公告错误

	NoticeNotExistError = NewError(210001, "公告不存在")
)

// 第三方应用错误 3xxxix

var ()
