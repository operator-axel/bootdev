package main

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

func newUser(name string, memType membershipType) User {
	newUser := User{
		Name: name,
		Membership: Membership{
			Type:             memType,
			MessageCharLimit: 100,
		},
	}
	if memType == TypePremium {
		newUser.Membership.MessageCharLimit = 1000
	}
	return newUser
}
