package model

import "errors"

//根据业务逻辑，需要自定义一些错误
var (
	ERROR_USER_NOTESISTS = errors.New("用户不存在..")
	ERROR_USER_ESISTS    = errors.New("用户已经存在..")
	ERROR_USER_PWD       = errors.New("密码不正确..")
)
