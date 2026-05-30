package status

/*
	1- SUC
	  - 10 : OK

	2- STATUS
	  - 20 : Accepted
	  - 21 : Processing

	3- CLIENT ERR
	  - 30 Client Error
	  - 31 Bad Request
	  - 32 Not Authorised
	  - 33 Expired Token
	  - 34 Forbidden
	  - 35 Duplicated
	  - 36 Conflict
	  - 37 Too many Request
	  - 38 Not Found

	4- SERVER ERR
	  - 40 Internal Server Error
*/
type Status int

const (
	StatusOK Status = 10

	StatusAccepted   Status = 20
	StatusProcessing Status = 21

	StatusBadRequest      Status = 31
	StatusUnauthorized    Status = 32
	StatusExpiredToken    Status = 33
	StatusForbidden       Status = 34
	StatusDuplicated      Status = 35
	StatusConflict        Status = 36
	StatusTooManyRequests Status = 37
	StatusNotFound        Status = 38

	StatusInternalError Status = 40
)
