package dummy

import (
	"errors"
)

type RequestDTO struct {
	DummyId string `param:"dummy_id"`
}

type Request struct {
	DTO *RequestDTO
	Err error
}

func NewRequest(dto *RequestDTO) Request {
	return Request{
		DTO: dto,
	}
}

func (r *Request) Validate() error {
	if err := r.dummyIdRule(); err != nil {
		return err
	}
	return nil
}

func (r *Request) dummyIdRule() error {
	if r.DTO.DummyId == `""` {
		return errors.New("invalid_argument")
	}
	return nil
}
