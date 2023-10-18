package status

const NOTAHIRERVEHICLE = 0
const PENDING = 1
const PAID = 3
const APPEAL = 4
const NOMINATED = 5
const ERROR = 6
const PROCESSING = 7
const REJECTED = 8
const WAITING_ON_INFO = 9
const TIME_ELAPSED = 10
const NO_ANSWER_UNABLE_TO_GET_INFO = 11

const API_PENDING_STATUS = 0
const API_NOMINATED_STATUS = 1
const API_PAID_STATUS = 2
const API_APPEAL_STATUS = 3
const API_DISMISSED_STATUS = 4
const API_EVIDENCE_STATUS = 5
const API_WAITING_STATUS = 6
const API_PROCESSING = 7

const API_TIME_ELAPSED = 100
const API_NO_DATA = 101

const API_ERROR_FATAL_STATUS = 999
const API_ERROR_RESEND_STATUS = 998

type ApiStatus struct {
	Code        int
	Description string
}

func ApiStatusFromLeaseResultStatus(status int) ApiStatus {

	switch status {

	case PENDING:
		return ApiStatus{API_PENDING_STATUS, "Pending"}

	case PROCESSING:
		return ApiStatus{API_PROCESSING, "Processing"}

	case PAID:
		return ApiStatus{API_PAID_STATUS, "Pay"}

	case APPEAL:
		return ApiStatus{API_APPEAL_STATUS, "Appeal"}

	case NOMINATED:
		return ApiStatus{API_NOMINATED_STATUS, "Nominate"}

	case REJECTED:
		return ApiStatus{API_DISMISSED_STATUS, "Dismissed"}

	case WAITING_ON_INFO:
		return ApiStatus{API_WAITING_STATUS, "Waiting"}

	case TIME_ELAPSED:
		return ApiStatus{API_TIME_ELAPSED, "Time Elapsed"}

	case NO_ANSWER_UNABLE_TO_GET_INFO:
		return ApiStatus{API_NO_DATA, "Unable to get Data"}

	case ERROR:
		return ApiStatus{API_ERROR_FATAL_STATUS, "Error Fatal"}

	default:
		return ApiStatus{status, "Unknown Status"}
	}

}
