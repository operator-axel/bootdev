package main

type authenticationInfo struct {
	username string
	password string
}

func (ai authenticationInfo) getBasicAuth() string {
	return "Authorization: Basic " + ai.username + ":" + ai.password
}
