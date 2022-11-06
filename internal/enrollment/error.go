package enrollment

import (
	"errors"
	"fmt"
)

var ErrUserIDRequired = errors.New("user id is required")
var ErrCourseIDRequired = errors.New("course id is required")
var ErrStatusRequired = errors.New("status is required")

type ErrNotFound struct {
	EnrollmentsID string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("enrollment '%s' doesn't exist", e.EnrollmentsID)
}

type ErrUserNotExist struct {
	UserID string
}

func (e ErrUserNotExist) Error() string {
	return fmt.Sprintf("user '%s' doesn't exist", e.UserID)
}

type ErrCourseNotExist struct {
	CourseID string
}

func (e ErrCourseNotExist) Error() string {
	return fmt.Sprintf("course '%s' doesn't exist", e.CourseID)
}
