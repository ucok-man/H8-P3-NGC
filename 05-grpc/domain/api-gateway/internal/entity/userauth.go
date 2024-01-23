package entity

import "golang.org/x/crypto/bcrypt"

type UserAuth struct {
	Username string
	Password string
}

func (p *UserAuth) SetPassword(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	p.Password = string(hash)
	return nil
}

func (p *UserAuth) MatchesPassword(plaintextPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(plaintextPassword))
	if err != nil {
		return err
	}

	return nil
}
