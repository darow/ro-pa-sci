package server

import "errors"

var (
	ErrNotFoundInContext = errors.New("пользователь не найден в контексте запроса")
	ErrForbidden         = errors.New("нет доступа")
)
