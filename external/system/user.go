package system

import (
	"os/user"
)

//UserRetriever retrieve the current user from somewhere
type UserRetriever interface {
	Current() (*user.User, error)
}

//OSUserRetriever holds the OS current logged user capabilities
type OSUserRetriever struct {
}

//Current retrieves the current OS logged user
func (userRetriever OSUserRetriever) Current() (*user.User, error) {
	return user.Current()
}
