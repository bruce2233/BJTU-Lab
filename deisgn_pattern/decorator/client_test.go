package decorator

import "testing"

func TestDecorator(t *testing.T) {
	loginExtendedImpl := LoginExtendedImpl{UserImpl{}}
	teacher := Teacher{loginExtendedImpl}
	teacher.login()
	teacher.logout()
	student := Student{loginExtendedImpl}
	student.login()
	student.logout()
	graduate := Graduate{loginExtendedImpl}
	graduate.login()
	graduate.logout()
}
