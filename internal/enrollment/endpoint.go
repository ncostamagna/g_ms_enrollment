package enrollment

import (
	"context"

	"github.com/ncostamagna/g_ms_client/meta"
	"github.com/ncostamagna/g_ms_enrollment_ex/pkg/response"
)

//Endpoints struct
type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		Create Controller
	}

	CreateReq struct {
		UserID   string `json:"user_id"`
		CourseID string `json:"course_id"`
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
		Meta   *meta.Meta  `json:"meta,omitempty"`
	}
)

//MakeEndpoints handler endpoints
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
	}
}

func makeCreateEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateReq)

		if req.UserID == "" {
			return nil, response.BadRequest(ErrUserIDRequired.Error())
		}

		if req.CourseID == "" {
			return nil, response.BadRequest(ErrCourseIDRequired.Error())
		}

		enroll, err := s.Create(ctx, req.UserID, req.CourseID)
		if err != nil {

			return nil, response.InternalServerError(err.Error())
		}

		return response.Created("success", enroll, nil), nil

	}
}
