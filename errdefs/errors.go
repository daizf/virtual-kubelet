package errdefs

const (
	Unknown              = "Unknown"
	ParamsAdapterFailed  = "ParamsAdapterFailed"
	NotSupport           = "NotSupport"
	ResourceNotFound     = "ResourceNotFound"
	ResourceNotReady     = "ResourceNotReady"
	ProviderInvokeFailed = "ProviderInvokeFailed"
	FailedScheduling     = "FailedScheduling"
)

type RetryType int32

const (
	RtForEver      RetryType = 1
	RtEndWithTimes RetryType = 2
	RtNoRetry      RetryType = 3
	RtUnknow       RetryType = 4
)

func GetRetryType(err interface{}) RetryType {
	if pe, ok := err.(*ProviderError); ok {
		return pe.Retry
	}
	return RtUnknow
}

type ProviderError struct {
	Reason  string
	Message string
	OrigErr error
	Retry   RetryType
}

var _ error = &ProviderError{}

// Error implements the Error interface.
func (e *ProviderError) Error() string {
	return e.Message
}

func NewParamsAdapterFailedNotRetry(err error) *ProviderError {
	return &ProviderError{
		Reason:  ParamsAdapterFailed,
		Message: err.Error(),
		OrigErr: err,
		Retry:   RtNoRetry,
	}
}

func NewResourceNotFoundRetry(err error) *ProviderError {
	return &ProviderError{
		Reason:  ResourceNotFound,
		Message: err.Error(),
		OrigErr: err,
		Retry:   RtForEver,
	}
}

func NewResourceNotReadyRetry(err error) *ProviderError {
	return &ProviderError{
		Reason:  ResourceNotReady,
		Message: err.Error(),
		OrigErr: err,
		Retry:   RtForEver,
	}
}

func NewNotSupportNotRetry(err error) *ProviderError {
	return &ProviderError{
		Reason:  NotSupport,
		Message: err.Error(),
		OrigErr: err,
		Retry:   RtNoRetry,
	}
}

func NewProviderFailedRetry(err error) *ProviderError {
	return &ProviderError{
		Reason:  ProviderInvokeFailed,
		Message: err.Error(),
		OrigErr: err,
		Retry:   RtForEver,
	}
}

func NewFailedSchedulingRetry(err error) *ProviderError {
	return &ProviderError{
		Reason:  FailedScheduling,
		Message: err.Error(),
		OrigErr: err,
		Retry:   RtForEver,
	}
}

func NewProviderFailedRetryTimes(err error) *ProviderError {
	return &ProviderError{
		Reason:  ProviderInvokeFailed,
		Message: err.Error(),
		OrigErr: err,
		Retry:   RtEndWithTimes,
	}
}

func NewProviderFailedNotRetry(err error) *ProviderError {
	return &ProviderError{
		Reason:  ProviderInvokeFailed,
		Message: err.Error(),
		OrigErr: err,
		Retry:   RtNoRetry,
	}
}
