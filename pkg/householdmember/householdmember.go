// Package householdmember provides functionality to manage household members.
//
// It includes the ability to create a new household member with a name and phone number.
package householdmember

type householdMember struct {
	name        string
	phoneNumber string
}

func newHouseholdMember(name, phoneNumber string) householdMember {
	newMember := householdMember{
		name:        name,
		phoneNumber: phoneNumber,
	}
	return newMember
}
