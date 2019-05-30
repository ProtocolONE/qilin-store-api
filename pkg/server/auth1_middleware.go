package server

import (
	jwtverifier "github.com/ProtocolONE/authone-jwt-verifier-golang"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"regexp"
)

const ErrorAuthHeaderInvalid = "Invalid authorization header"
const ErrorAuthFailed = "Unable to authenticate user"

func AuthOneJwtWithConfig(cfg *jwtverifier.JwtVerifier) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			userInfo, err := introspectToken(c, cfg)
			if err != nil {
				c.Set("user", nil)
				return err
			}

			c.Set("user", userInfo)
			return next(c)
		}
	}
}

func introspectToken(c echo.Context, cfg *jwtverifier.JwtVerifier) (*jwtverifier.UserInfo, error) {
	req := c.Request()
	auth := req.Header.Get("Authorization")
	if auth == "" {
		return nil, nil
	}

	r := regexp.MustCompile("Bearer ([A-z0-9_.-]{10,})")
	match := r.FindStringSubmatch(auth)
	if len(match) < 1 {
		return nil, &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: ErrorAuthHeaderInvalid,
		}
	}

	token, err := cfg.Introspect(c.Request().Context(), match[1])
	if err != nil {
		return nil, &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: errors.Wrap(err, ErrorAuthFailed).Error(),
		}
	}

	return &jwtverifier.UserInfo{UserID: token.Sub, Name: token.Username}, nil
}
