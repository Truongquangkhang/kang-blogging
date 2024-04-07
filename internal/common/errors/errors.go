package errors

const (
	// 400
	ERRCODE_BAD_REQUEST = 400_000
	ERRMSG_BAD_REQUEST  = "Request không hợp lệ"

	ERRCODE_IAM_BAD_REQUEST = 400_001
	ERRMSG_IAM_BAD_REQUEST  = "Không tìm thấy user IAM"

	ERRCODE_USE_VOUCHER_NOT_AVAILABLE = 400_002
	ERRMSG_USE_VOUCHER_NOT_AVAILABLE  = "Voucher hiện không còn khả dụng, vui lòng thử lại sau"

	ERRCODE_USE_VOUCHER_NOT_AVAILABLE_BLACKLIST = 400_003
	ERRMSG_USE_VOUCHER_NOT_AVAILABLE_BLACKLIST  = "Số điện thoại này đã bị đưa vào danh sách đen"

	ERRCODE_USE_VOUCHER_USED = 400_004
	ERRMSG_USE_VOUCHER_USED  = "Đã có người nhanh tay sử dụng voucher cuối cùng, vui lòng thử lại sau"

	ERRCODE_USE_VOUCHER_OUTDATED = 400_005
	ERRMSG_USE_VOUCHER_OUTDATED  = "Voucher đã hết hạn sử dụng"

	ERRCODE_VOUCHERS_FOR_SHOP_BAD_REQUEST = 400_006
	ERRMSG_VOUCHERS_FOR_SHOP_BAD_REQUEST  = "Yêu cầu không hợp lệ"

	ERRCODE_EXTRACT_QR_BAD_REQUEST = 400_007
	ERRMSG_EXTRACT_QR_BAD_REQUEST  = "Mã QR không hợp lệ"

	ERRCODE_GET_APPLICABLE_VOUCHER_BY_VOUCHER_CODE_BAD_REQUEST = 400_008
	ERRMSG_GET_APPLICABLE_VOUCHER_BY_VOUCHER_CODE_BAD_REQUEST  = "Voucher không khả dụng"

	// 401
	ERRCODE_UNAUTHORIZED = 401_000
	ERRMSG_UNAUTHORIZED  = "Yêu cầu không được phép"

	ERRCODE_IAM_UNAUTHORIZED = 401_001
	ERRMSG_IAM_UNAUTHORIZED  = "Yêu cầu IAM không được phép"

	ERRCODE_USE_VOUCHER_UNAUTHORIZED = 401_002
	ERRMSG_USE_VOUCHER_UNAUTHORIZED  = "Yêu cầu không được phép"

	// 404
	ERRCODE_NOT_FOUND = 404_000
	ERRMSG_NOT_FOUND  = "Không tìm thấy"

	ERRCODE_IAM_NOT_FOUND = 404_001
	ERRMSG_IAM_NOT_FOUND  = "Không tìm thấy user IAM"

	ERRCODE_GET_APPLICABLE_VOUCHER_BY_VOUCHER_CODE_NOT_FOUND = 404_002
	ERRMSG_GET_APPLICABLE_VOUCHER_BY_VOUCHER_CODE_NOT_FOUND  = "Voucher không tìm thấy"

	// 413
	ERRCODE_ENTITY_TOO_LARGE = 413_000
	ERRMSG_ENTITY_TOO_LARGE  = "Nội dung gửi lên quá lớn"

	// 415
	ERRCODE_UNSUPPORTED_MEDIA = 415_000
	ERRMSG_UNSUPPORTED_MEDIA  = "Phương tiện media không được hỗ trợ"

	// 500
	ERRCODE_INTERNAL_ERROR = 500_000
	ERRMSG_INTERNAL_ERROR  = "Có lỗi xảy ra vui lòng liên hệ với chúng tôi để được hỗ trợ"

	ERRCODE_3PARTY_ERROR = 500_001
	ERRMSG_3PARTY_ERROR  = "Có lỗi xảy ra ở service liên quan, xin vui lòng liên hệ để dược hỗ trợ"

	ERRCODE_USE_VOUCHER_INTERNAL_ERROR = 500_002
	ERRMSG_USE_VOUCHER_INTERNAL_ERROR  = "Có lỗi xảy ra vui lòng liên hệ với chúng tôi để được hỗ trợ"
)

type BaseError struct {
	error        string
	errorCode    int32
	errorMessage string
}

func (s BaseError) Error() string {
	return s.error
}

func (s BaseError) ErrorCode() int32 {
	return s.errorCode
}

func (s BaseError) ErrorMessage() string {
	return s.errorMessage
}

func NewError(code int32, message string) BaseError {
	return BaseError{
		error:        message,
		errorCode:    code,
		errorMessage: message,
	}
}

func NewBadRequestError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_BAD_REQUEST,
		errorMessage: message,
	}
}

func NewBadRequestDefaultError() BaseError {
	return NewBadRequestError(ERRMSG_BAD_REQUEST)
}

func NewIAMBadRequestError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_IAM_BAD_REQUEST,
		errorMessage: message,
	}
}

func NewIAMBadRequestDefaultError() BaseError {
	return NewIAMBadRequestError(ERRMSG_IAM_BAD_REQUEST)
}

func NewUseVoucherNotAvailableError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_USE_VOUCHER_NOT_AVAILABLE,
		errorMessage: message,
	}
}

func NewUseVoucherNotAvailableDefaultError() BaseError {
	return NewUseVoucherNotAvailableError(ERRMSG_USE_VOUCHER_NOT_AVAILABLE)
}

func NewUseVoucherNotAvailableBlacklistError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_USE_VOUCHER_NOT_AVAILABLE_BLACKLIST,
		errorMessage: message,
	}
}

func NewUseVoucherNotAvailableBlacklistDefaultError() BaseError {
	return NewUseVoucherNotAvailableError(ERRMSG_USE_VOUCHER_NOT_AVAILABLE_BLACKLIST)
}

func NewUseVoucherUsedError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_USE_VOUCHER_USED,
		errorMessage: message,
	}
}

func NewUseVoucherUsedDefaultError() BaseError {
	return NewUseVoucherUsedError(ERRMSG_USE_VOUCHER_USED)
}

func NewUseVoucherOutdatedError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_USE_VOUCHER_OUTDATED,
		errorMessage: message,
	}
}

func NewUseVoucherOutdatedDefaultError() BaseError {
	return NewUseVoucherOutdatedError(ERRMSG_USE_VOUCHER_OUTDATED)
}

func NewVouchersForShopBadRequestError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_VOUCHERS_FOR_SHOP_BAD_REQUEST,
		errorMessage: message,
	}
}

func NewVouchersForShopBadRequestDefaultError(message string) BaseError {
	return NewVouchersForShopBadRequestError(ERRMSG_VOUCHERS_FOR_SHOP_BAD_REQUEST)
}

func NewVouchersForShopBadRequestQrExtractDefaultError() BaseError {
	return NewVouchersForShopBadRequestError(ERRMSG_EXTRACT_QR_BAD_REQUEST)
}

func NewExtractQrBadRequestError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_EXTRACT_QR_BAD_REQUEST,
		errorMessage: message,
	}
}

func NewExtractQrBadRequestDefaultError() BaseError {
	return NewExtractQrBadRequestError(ERRMSG_EXTRACT_QR_BAD_REQUEST)
}

func NewGetApplicableVoucherByVoucherCodeBadRequestError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_GET_APPLICABLE_VOUCHER_BY_VOUCHER_CODE_BAD_REQUEST,
		errorMessage: message,
	}
}

func NewGetApplicableVoucherByVoucherCodeBadRequestDefaultError() BaseError {
	return NewGetApplicableVoucherByVoucherCodeBadRequestError(
		ERRMSG_GET_APPLICABLE_VOUCHER_BY_VOUCHER_CODE_BAD_REQUEST)
}

func NewAuthorizationError(message string) BaseError {
	return BaseError{
		error:        "unauthorized",
		errorCode:    ERRCODE_UNAUTHORIZED,
		errorMessage: ERRMSG_UNAUTHORIZED,
	}
}

func NewAuthorizationDefaultError() BaseError {
	return NewAuthorizationError(ERRMSG_UNAUTHORIZED)
}

func NewIAMAuthorizationError(message string) BaseError {
	return BaseError{
		error:        "unauthorized",
		errorCode:    ERRCODE_IAM_UNAUTHORIZED,
		errorMessage: message,
	}
}

func NewIAMAuthorizationDefaultError() BaseError {
	return NewIAMAuthorizationError(ERRMSG_IAM_UNAUTHORIZED)
}

func NewNotFoundError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_NOT_FOUND,
		errorMessage: message,
	}
}

func NewNotFoundDefaultError() BaseError {
	return NewNotFoundError(ERRMSG_NOT_FOUND)
}

func NewIAMNotFoundError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_IAM_NOT_FOUND,
		errorMessage: message,
	}
}

func NewIAMNotFoundDefaultError() BaseError {
	return NewIAMNotFoundError(ERRMSG_IAM_NOT_FOUND)
}

func NewGetApplicableVoucherByVoucherCodeNotFoundError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_GET_APPLICABLE_VOUCHER_BY_VOUCHER_CODE_NOT_FOUND,
		errorMessage: message,
	}
}

func NewGetApplicableVoucherByVoucherCodeNotFoundDefaultError() BaseError {
	return NewGetApplicableVoucherByVoucherCodeNotFoundError(
		ERRMSG_GET_APPLICABLE_VOUCHER_BY_VOUCHER_CODE_NOT_FOUND)
}

func NewEntityTooLargeError(message string) BaseError {
	return BaseError{
		error:        "entity-too-large",
		errorCode:    ERRCODE_ENTITY_TOO_LARGE,
		errorMessage: message,
	}
}

func NewEntityTooLargeDefaultError() BaseError {
	return NewEntityTooLargeError(ERRMSG_ENTITY_TOO_LARGE)
}

func NewUnsuppportedMediaError(message string) BaseError {
	return BaseError{
		error:        "unsupported-media",
		errorCode:    ERRCODE_UNSUPPORTED_MEDIA,
		errorMessage: message,
	}
}

func NewUnsuppportedMediaDefaultError() BaseError {
	return NewUnsuppportedMediaError(ERRMSG_UNSUPPORTED_MEDIA)
}

func NewInternalErrorError(message string) BaseError {
	return BaseError{
		error:        "internal-error",
		errorCode:    ERRCODE_INTERNAL_ERROR,
		errorMessage: message,
	}
}

func NewInternalErrorDefaultError() BaseError {
	return NewInternalErrorError(ERRMSG_INTERNAL_ERROR)
}

func NewUseVoucherInternalErrorError(message string) BaseError {
	return BaseError{
		error:        "internal-error",
		errorCode:    ERRCODE_USE_VOUCHER_INTERNAL_ERROR,
		errorMessage: message,
	}
}

func NewUseVoucherInternalErrorDefaultError() BaseError {
	return NewUseVoucherInternalErrorError(ERRMSG_USE_VOUCHER_INTERNAL_ERROR)
}

func NewThirdPartyError(message string) BaseError {
	return BaseError{
		error:        "3party-error",
		errorCode:    ERRCODE_3PARTY_ERROR,
		errorMessage: message,
	}
}

func NewThirdPartyErrorDefaultError() BaseError {
	return NewThirdPartyError(ERRMSG_3PARTY_ERROR)
}
