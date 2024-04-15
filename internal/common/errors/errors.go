package errors

import "errors"

const (
	// 400
	ERRCODE_BAD_REQUEST = 400_000
	ERRMSG_BAD_REQUEST  = "Bad Request"

	ERRCODE_INVALID_PARAMS_REQUEST = 400_001
	ERRMSG_INVALID_PARAMS_REQUEST  = "Invalid Params"

	ERRCODE_IAM_BAD_REQUEST = 400_002
	ERRMSG_IAM_BAD_REQUEST  = "User IAM Bad Request"

	ERRCODE_MISSING_MERCHANT_ID_REQUEST = 400_003
	ERRMSG_MISSING_MERCHANT_ID_REQUEST  = "Missing Merchant ID"

	ERRCODE_PROJECT_NOT_DELETABLE = 400_004
	ERRMSG_PROJECT_NOT_DELETABLE  = "Project Not Deletable"

	ERRCODE_POLYGON_NOT_DELETABLE = 400_005
	ERRMSG_POLYGON_NOT_DELETABLE  = "Polygon Not Deletable"

	ERRCODE_NOT_KML_FILE = 400_006
	ERRMSG_NOT_KML_FILE  = "Not a KML file"

	ERRCODE_BAD_KML_FILE = 400_008
	ERRMSG_BAD_KML_FILE  = "Bad KML file"

	ERRCODE_KML_FILE_NO_POLY = 400_009
	ERRMSG_KML_FILE_NO_POLY  = "File does not contain polygons"

	ERRCODE_POLYGON_CONFLICT = 400_010
	ERRMSG_POLYGON_CONFLICT  = "Polygons must not overlapped with others"

	ERRCODE_SUBMISSION_NOT_EDITABLE = 400_011
	ERRMSG_SUBMISSION_NOT_EDITABLE  = "Submission Not Editable"

	ERRCODE_MISMATCH_IMAGE = 400_012
	ERRMSG_MISMATCH_IMAGE  = "This image was not expected"

	ERRCODE_DUPLICATED_IMAGE = 400_013
	ERRMSG_DUPLICATED_IMAGE  = "This image was already uploaded"

	ERRCODE_LOYALTY_BAD_REQUEST = 400_014
	ERRMSG_LOYALTY_BAD_REQUEST  = "Loyalty bad request"

	// 401
	ERRCODE_UNAUTHORIZED = 401_000
	ERRMSG_UNAUTHORIZED  = "Unauthorized Error"

	ERRCODE_IAM_UNAUTHORIZED = 401_001
	ERRMSG_IAM_UNAUTHORIZED  = "Unauthorized IAM Error"

	// 403
	ERRCODE_FORBIDDEN = 403_000
	ERRMSG_FORBIDDEN  = "You are not allowed to perform this action."

	// 404
	ERRCODE_NOT_FOUND = 404_000
	ERRMSG_NOT_FOUND  = "Not Found"

	ERRCODE_IAM_NOT_FOUND = 404_001
	ERRMSG_IAM_NOT_FOUND  = "User IAM Not Found"

	ERRCODE_PROJECT_NOT_FOUND = 404_002
	ERRMSG_PROJECT_NOT_FOUND  = "Project Not Found"

	ERRCODE_POLYGON_NOT_FOUND = 404_003
	ERRMSG_POLYGON_NOT_FOUND  = "Polygon Not Found"

	ERRCODE_SUBMISSION_NOT_FOUND = 404_004
	ERRMSG_SUBMISSION_NOT_FOUND  = "Submission Not Found"

	ERRCODE_TREE_SPECIES_NOT_FOUND = 404_005
	ERRMSG_TREE_SPECIES_NOT_FOUND  = "Tree Species Not Found"

	ERRCODE_PLOT_NOT_FOUND = 404_006
	ERRMSG_PLOT_NOT_FOUND  = "Plot Not Found"

	ERRCODE_IMAGE_NOT_FOUND = 404_007
	ERRMSG_IMAGE_NOT_FOUND  = "Image Not Found"

	ERRCODE_MEMBER_LOYALTY_NOT_FOUND = 404_008
	ERRMSG_MEMBER_LOYALTY_NOT_FOUND  = "Member Loyalty Not Found"

	ERRCODE_IAM_USER_NOT_FOUND = 404_009
	ERRMSG_IAM_USER_NOT_FOUND  = "Iam User not found"

	ERRCODE_FILE_NOT_FOUND = 404_010
	ERRMSG_FILE_NOT_FOUND  = "File Not Found"

	ERRCODE_LOCATION_NOT_FOUND = 404_011
	ERRMSG_LOCATION_NOT_FOUND  = "Location Not Found"

	ERRCODE_NOT_ALLOWED = 405_000
	ERRMSG_NOT_ALLOWED  = "Not Allowed"

	ERRCODE_UPDATE_FORM_SUBMISSION_NOT_ALLOWED = 405_001
	ERRMSG_UPDATE_FORM_SUBMISSION_NOT_ALLOWED  = "Update form-submission not allowed while processing"

	// 409
	ERRCODE_ALREADY_EXIST = 409_000
	ERRMSG_ALREADY_EXIST  = "Already Exist"

	// 413
	ERRCODE_ENTITY_TOO_LARGE = 413_000
	ERRMSG_ENTITY_TOO_LARGE  = "Entity Too Large Error"

	ERRCODE_KML_FILE_TOO_LARGE = 413_001
	ERRMSG_KML_FILE_TOO_LARGE  = "Uploaded Kml File Too Large"

	// 415
	ERRCODE_UNSUPPORTED_MEDIA = 415_000
	ERRMSG_UNSUPPORTED_MEDIA  = "Unsupported Media Type"

	//422
	ERRCODE_MISSING_MERCHANT_ID = 422_000
	ERRMSG_MISSING_MERCHANT_ID  = "Merchant ID not found"

	//423
	ERRCODE_OVERLAP_SCHEDULE = 423_000
	ERRMSG_OVERLAP_SCHEDULE  = "Overlapping Schedule"

	// 500
	ERRCODE_INTERNAL_ERROR = 500_000
	ERRMSG_INTERNAL_ERROR  = "Internal Server Error"

	ERRCODE_3PARTY_ERROR = 500_001
	ERRMSG_3PARTY_ERROR  = "Internal Server Error"

	ERRCODE_SUCCESS = 0
	ERRMSG_SUCCESS  = "Success"
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

func (s BaseError) BaseErrorCode() int32 {
	return s.errorCode / 1000
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

func NewLoyaltyBadRequestError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_LOYALTY_BAD_REQUEST,
		errorMessage: message,
	}
}

func NewLoyaltyBadRequestDefaultError() BaseError {
	return NewLoyaltyBadRequestError(ERRMSG_LOYALTY_BAD_REQUEST)
}

func NewInvalidParamsRequestError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_INVALID_PARAMS_REQUEST,
		errorMessage: message,
	}
}

func NewInvalidParamsRequestDefaultError() BaseError {
	return NewInvalidParamsRequestError(ERRMSG_INVALID_PARAMS_REQUEST)
}

func NewMissingMerchantIdRequestError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_MISSING_MERCHANT_ID_REQUEST,
		errorMessage: message,
	}
}

func NewMissingMerchantIdRequestDefaultError() BaseError {
	return NewMissingMerchantIdRequestError(ERRMSG_MISSING_MERCHANT_ID_REQUEST)
}

func NewProjectNotDeletableError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_PROJECT_NOT_DELETABLE,
		errorMessage: message,
	}
}

func NewProjectNotDeletableDefaultError() BaseError {
	return NewProjectNotDeletableError(ERRMSG_PROJECT_NOT_DELETABLE)
}

func NewPolygonNotDeletableError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_POLYGON_NOT_DELETABLE,
		errorMessage: message,
	}
}

func NewPolygonNotDeletableDefaultError() BaseError {
	return NewPolygonNotDeletableError(ERRMSG_POLYGON_NOT_DELETABLE)
}

func NewNotKmlFileError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_NOT_KML_FILE,
		errorMessage: message,
	}
}

func NewNotKmlFileDefaultError() BaseError {
	return NewNotKmlFileError(ERRMSG_NOT_KML_FILE)
}

func NewBadKmlFileError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_BAD_KML_FILE,
		errorMessage: message,
	}
}

func NewBadKmlFileDefaultError() BaseError {
	return NewBadKmlFileError(ERRMSG_BAD_KML_FILE)
}

func NewKmlFileNoPolyError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_KML_FILE_NO_POLY,
		errorMessage: message,
	}
}

func NewKmlFileNoPolyDefaultError() BaseError {
	return NewKmlFileNoPolyError(ERRMSG_KML_FILE_NO_POLY)
}

func NewConflictPolyError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_POLYGON_CONFLICT,
		errorMessage: message,
	}
}

func NewConflictPolyDefaultError() BaseError {
	return NewConflictPolyError(ERRMSG_POLYGON_CONFLICT)
}

func NewSubmissionNotEditableError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_SUBMISSION_NOT_EDITABLE,
		errorMessage: message,
	}
}

func NewSubmissionNotEditableDefaultError() BaseError {
	return NewSubmissionNotEditableError(ERRMSG_SUBMISSION_NOT_EDITABLE)
}

func NewMismatchImageError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_MISMATCH_IMAGE,
		errorMessage: message,
	}
}

func NewMismatchImageDefaultError() BaseError {
	return NewMismatchImageError(ERRMSG_MISMATCH_IMAGE)
}

func NewDuplicatedImageError(message string) BaseError {
	return BaseError{
		error:        "incorrect-input",
		errorCode:    ERRCODE_DUPLICATED_IMAGE,
		errorMessage: message,
	}
}

func NewDuplicatedImageDefaultError() BaseError {
	return NewDuplicatedImageError(ERRMSG_DUPLICATED_IMAGE)
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

func NewForbiddenError(message string) BaseError {
	return BaseError{
		error:        "forbidden",
		errorCode:    ERRCODE_FORBIDDEN,
		errorMessage: message,
	}
}

func NewForbiddenDefaultError() BaseError {
	return NewForbiddenError(ERRMSG_FORBIDDEN)
}

func NewNotFoundError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_NOT_FOUND,
		errorMessage: message,
	}
}

func NewNotFoundDefaultError() BaseError {
	return NewNotFoundError("")
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

func NewLoyaltyNotFoundError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_MEMBER_LOYALTY_NOT_FOUND,
		errorMessage: message,
	}
}

func NewLoyaltyNotFoundDefaultError() BaseError {
	return NewLoyaltyNotFoundError(ERRMSG_MEMBER_LOYALTY_NOT_FOUND)
}

func NewProjectNotFoundError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_PROJECT_NOT_FOUND,
		errorMessage: message,
	}
}

func NewProjectNotFoundDefaultError() BaseError {
	return NewProjectNotFoundError(ERRMSG_PROJECT_NOT_FOUND)
}

func NewPolygonNotFoundError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_POLYGON_NOT_FOUND,
		errorMessage: message,
	}
}

func NewPolygonNotFoundDefaultError() BaseError {
	return NewPolygonNotFoundError(ERRMSG_POLYGON_NOT_FOUND)
}

func NewSubmissionNotFoundError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_SUBMISSION_NOT_FOUND,
		errorMessage: message,
	}
}

func NewSubmissionNotFoundDefaultError() BaseError {
	return NewSubmissionNotFoundError(ERRMSG_SUBMISSION_NOT_FOUND)
}

func NewTreeSpeciesNotFoundError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_TREE_SPECIES_NOT_FOUND,
		errorMessage: message,
	}
}

func NewTreeSpeciesNotFoundDefaultError() BaseError {
	return NewTreeSpeciesNotFoundError(ERRMSG_TREE_SPECIES_NOT_FOUND)
}

func NewPlotNotFoundError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_PLOT_NOT_FOUND,
		errorMessage: message,
	}
}

func NewPlotNotFoundDefaultError() BaseError {
	return NewPlotNotFoundError(ERRMSG_PLOT_NOT_FOUND)
}

func NewImageNotFoundError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_IMAGE_NOT_FOUND,
		errorMessage: message,
	}
}

func NewImageNotFoundDefaultError() BaseError {
	return NewImageNotFoundError(ERRMSG_IMAGE_NOT_FOUND)
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

func NewKmlFileTooLargeError(message string) BaseError {
	return BaseError{
		error:        "entity-too-large",
		errorCode:    ERRCODE_KML_FILE_TOO_LARGE,
		errorMessage: message,
	}
}

func NewKmlFileTooLargeDefaultError() BaseError {
	return NewKmlFileTooLargeError(ERRMSG_KML_FILE_TOO_LARGE)
}

func NewUnsupportedMediaError(message string) BaseError {
	return BaseError{
		error:        "unsupported-media",
		errorCode:    ERRCODE_UNSUPPORTED_MEDIA,
		errorMessage: message,
	}
}

func NewUnsupportedMediaDefaultError() BaseError {
	return NewUnsupportedMediaError(ERRMSG_UNSUPPORTED_MEDIA)
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

func NewAlreadyExistError(message string) BaseError {
	return BaseError{
		error:        "already-exists",
		errorCode:    ERRCODE_ALREADY_EXIST,
		errorMessage: message,
	}
}

func NewAlreadyExistErrorDefaultError() BaseError {
	return NewAlreadyExistError(ERRMSG_ALREADY_EXIST)
}

func NewOverlapScheduleError(message string) BaseError {
	return BaseError{
		error:        "overlap-schedule",
		errorCode:    ERRCODE_OVERLAP_SCHEDULE,
		errorMessage: message,
	}
}

func NewOverlapScheduleErrorDefaultError() BaseError {
	return NewOverlapScheduleError(ERRMSG_OVERLAP_SCHEDULE)
}

func NewIamUserNotFoundError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_IAM_USER_NOT_FOUND,
		errorMessage: message,
	}
}

func NewIamUserNotFoundErrorDefaultError() BaseError {
	return NewIamUserNotFoundError(ERRMSG_IAM_USER_NOT_FOUND)
}

func IsLoyaltyNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	var e BaseError
	if errors.As(err, &e) {
		return e.errorCode == ERRCODE_MEMBER_LOYALTY_NOT_FOUND
	}
	return false
}

func NewFileNotFoundError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_FILE_NOT_FOUND,
		errorMessage: message,
	}
}

func NewNotAllowedError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_NOT_ALLOWED,
		errorMessage: message,
	}
}

func NewFileNotFoundDefaultError() BaseError {
	return NewIamUserNotFoundError(ERRMSG_FILE_NOT_FOUND)
}

func NewNotAllowedDefaultError() BaseError {
	return NewNotAllowedError(ERRMSG_NOT_ALLOWED)
}

func NewUpdateFormSubmissionNotAllowedDefaultError() BaseError {
	return BaseError{
		error:        "not-allowed",
		errorCode:    ERRCODE_UPDATE_FORM_SUBMISSION_NOT_ALLOWED,
		errorMessage: ERRMSG_UPDATE_FORM_SUBMISSION_NOT_ALLOWED,
	}
}

func NewLocationNotFoundError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_LOCATION_NOT_FOUND,
		errorMessage: message,
	}
}

func NewLocationNotFoundDefaultError() BaseError {
	return NewLocationNotFoundError(ERRMSG_LOCATION_NOT_FOUND)
}

func NewMissingMerchantIdError(message string) BaseError {
	return BaseError{
		error:        "not-found",
		errorCode:    ERRCODE_MISSING_MERCHANT_ID,
		errorMessage: message,
	}
}

func NewMissingMerchantIdDefaultError() BaseError {
	return NewLocationNotFoundError(ERRMSG_MISSING_MERCHANT_ID)
}

func NewInvalidDateFormatError(message string) BaseError {
	return BaseError{
		error:        "invalid",
		errorCode:    ERRCODE_INTERNAL_ERROR,
		errorMessage: message,
	}
}

func NewInvalidDateFormatDefaultError() BaseError {
	return NewInvalidDateFormatError(ERRMSG_INTERNAL_ERROR)
}
