package errs

import "errors"

var NotFoundError = errors.New("not found")

var InvalidSignatureError = errors.New("invalid signature")

var ExpiredSignatureError = errors.New("expired signature")
