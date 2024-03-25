package main

type storage interface {
    CreateAccount(*UserAccount) error
    DeleteAccount(int) error
    UpdateAccount(*UserAccount) error
    GetAccountById(int) (*UserAccount error)
}

type PostgreeStore struct {
    db *sql.DB
}

func newPostgressStore()(*PostgressStore, error) {
    
}
