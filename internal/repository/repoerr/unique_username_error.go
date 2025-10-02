package repoerr

import "fmt"

type UniqueUsernameError struct {
	Username string
}

func (ue *UniqueUsernameError) Error() string {
	return fmt.Sprintf("Username '%s' is already taken.", ue.Username)
}
