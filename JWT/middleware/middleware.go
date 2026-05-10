package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	util "github.com/luckxx24/RoomBooking/JWT/auth"
	"github.com/luckxx24/RoomBooking/cmd/jsonresponse"
	"github.com/luckxx24/RoomBooking/store"
)

type contextKey string

const UserIDKey contextKey = "UserIDKey"
const RoleUser contextKey = "RoleUser"

func AuthHandler(t *util.TokenUtil, st store.Storage) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			AuthHeader := r.Header.Get("Authorization")

			if AuthHeader == " " {
				jsonresponse.RespondWithNotfound(w, "gagal menemukan Header")
				return
			}

			Auth := strings.Split(AuthHeader, " ")

			if len(Auth) != 2 && Auth[1] != "Bearer" {
				jsonresponse.RespondWithBadRequest(w, "format salah")
				return
			}

			UserIDstr, err := t.ParseJWT(Auth[0])

			UserID, errs := uuid.Parse(UserIDstr)

			if errs != nil {
				jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal menparse ID %v", errs))
				return
			}

			if err != nil {
				jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal mend parse jwt %v", err))
				return
			}

			user, errsr := st.Users.GetUsers(r.Context(), UserID)

			if errsr != nil {
				jsonresponse.RespondWithNotfound(w, fmt.Sprintf("gagal menemukan user %v", errs))
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, UserID)
			ctx = context.WithValue(ctx, RoleUser, user.Role)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetIDFromContext(ctx context.Context) (string, bool) {
	ID, ok := ctx.Value(UserIDKey).(string)

	return ID, ok
}

func GetRoleFromContext(ctx context.Context) (string, bool) {
	Role, ok := ctx.Value(RoleUser).(string)

	return Role, ok

}
