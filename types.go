package main

import (
    "github.com/google/uuid"
)
// UserAccount struct holds specific data pertaining to the user 
type UserAccount struct{
    ID string `json:"id"`
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    Progress uint8 `json:"progress"`
}

// newUserAccount constructor
func newUserAccount(firstName, lastName string) *UserAccount {
    return &UserAccount {
        // TODO Change ID creation later
        ID: uuid.New().String(),
        FirstName: firstName,
        LastName: lastName,
        Progress: 0,
    }
}
