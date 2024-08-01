package http

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"connectrpc.com/connect"
	"github.com/manhrev/labeler/internal/const/header"
)

func (s Service) WithAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, v := range s.authExludedPathMap {
			if r.URL.Path == v {
				h.ServeHTTP(w, r)
				return
			}
		}
		errWritter := connect.NewErrorWriter()
		tokenStr := r.Header.Get(header.Authorization)
		claims, err := s.tokenService.Validate(tokenStr)
		if err != nil {
			s.logger.Errorf("Token validation error: %s", err.Error())
			connectErr := connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("token validation error: %s", err.Error()))
			if err = errWritter.Write(w, r, connectErr); err != nil {
				s.logger.Errorf("Cannot write error when token validation fail")
			}
			return
		}
		if claims == nil {
			s.logger.Infof("User not authenticated")
			connectErr := connect.NewError(connect.CodeUnauthenticated, nil)
			if err = errWritter.Write(w, r, connectErr); err != nil {
				s.logger.Errorf("Cannot write error when user not authenticated")
			}
			return
		}

		// check user exists, cache if needed
		userIDStr := claims.Subject
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			s.logger.Errorf("Cannot parse user id from header: %v", err)
			connectErr := connect.NewError(connect.CodeInternal, errors.New("user id malformed"))
			if err = errWritter.Write(w, r, connectErr); err != nil {
				s.logger.Errorf("Cannot write error when user id malformed")
			}
			return
		}

		_, err = s.userRepo.GetUserByID(r.Context(), int64(userID))
		if err != nil {
			s.logger.Errorf("User not found : %v", err)
			connectErr := connect.NewError(connect.CodeInternal, errors.New("cannot query user or user not active"))
			if err = errWritter.Write(w, r, connectErr); err != nil {
				s.logger.Errorf("Cannot write error when user not found")
			}
			return
		}

		r.Header.Add(header.UserID, userIDStr)
		h.ServeHTTP(w, r)
	})
}
