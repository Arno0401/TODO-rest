package utils

import "errors"

var (
	UserExistsError = errors.New("пользователь с таким логином уже существует")

	InternalError = errors.New("внутренняя ошибка")

	UserNotFoundError = errors.New("пользователь не найден")
)
