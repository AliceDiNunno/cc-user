package request

import "github.com/AliceDiNunno/cc-user/src/core/domain"

type AccessTokenRequest struct {
	Mail     string `binding:"required"`
	Password string `binding:"required"`
	OtpCode  string
}

func (r AccessTokenRequest) ToDomain() domain.AccessTokenRequest {
	return domain.AccessTokenRequest{
		Mail:     r.Mail,
		Password: r.Password,
		OtpCode:  r.OtpCode,
	}
}
