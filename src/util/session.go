package util

import (
	"encoding/gob"
	"aws-lambda-go/src/types"

	"github.com/gorilla/sessions"
)

const SessionName = "github_session"

var Store = sessions.NewCookieStore([]byte("secret"))

func init() {
	gob.Register(types.ContextKey(""))
}
