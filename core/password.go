package core

import (
	"crypto/cipher"
	"time"
)

const shortIDLength = 7

// PasswordBasic is basic of Password
type PasswordBasic struct {
	// Category of password
	Category string `cli:"c,category" usage:"category of password"`

	// Plain account and password
	PlainAccount  string `json:"-" cli:"u,account" usage:"account of password"`
	PlainPassword string `json:"-" cli:"-"`

	// Website address for web password
	Site string `cli:"site" usage:"website of password"`

	// Password tags
	Tags []string `cli:"tag" usage:"tags of password"`

	// Extension information: JSON base64 string
	Ext string `cli:"-"`
}

// Password represents entity of password
type Password struct {
	PasswordBasic

	// Unique id of password
	ID string `cli:"id" usage:"password id for updating"`

	// IVs
	AccountIV  []byte `cli:"-"`
	PasswordIV []byte `cli:"-"`

	// Ciphers
	CipherAccount  []byte `cli:"-"`
	CipherPassword []byte `cli:"-"`

	// Created time stamp
	CreatedAt int64 `cli:"-"`

	// Last updated time stamp
	LastUpdatedAt int64 `cli:"-"`
}

var passwordHeader = []string{"ID", "CATEGORY", "ACCOUNT", "PASSWORD", "UPDATED_AT"}

func (pw Password) get(i int) string {
	switch i {
	case 0:
		return pw.ShortID()
	case 1:
		return pw.Category
	case 2:
		return pw.PlainAccount
	case 3:
		return pw.PlainPassword
	case 4:
		return time.Unix(pw.LastUpdatedAt, 0).Format(time.RFC3339)
	}
	panic("unreachable")
}

func (pw Password) colCount() int {
	return 5
}

// NewEmptyPassword creates a empty Password entity
func NewEmptyPassword() *Password {
	return NewPassword("", "", "", "")
}

// NewPassword creates a Password entity
func NewPassword(category, account, passwd, site string) *Password {
	now := time.Now().Unix()
	pw := &Password{
		PasswordBasic: PasswordBasic{
			Category:      category,
			PlainAccount:  account,
			PlainPassword: passwd,
			Site:          site,
			Tags:          []string{},
		},
		AccountIV:      []byte{},
		PasswordIV:     []byte{},
		CipherAccount:  []byte{},
		CipherPassword: []byte{},
		CreatedAt:      now,
		LastUpdatedAt:  now,
	}
	return pw
}

// ShortID returns short length id string
func (pw *Password) ShortID() string {
	if len(pw.ID) > shortIDLength {
		return pw.ID[:shortIDLength]
	}
	return pw.ID
}

func (pw *Password) migrate(from *Password) {
	pw.PasswordBasic = from.PasswordBasic
	pw.PasswordBasic.Tags = make([]string, len(from.PasswordBasic.Tags))
	copy(pw.PasswordBasic.Tags, from.PasswordBasic.Tags)
}

// CheckPassword validate password string
func CheckPassword(passwd string) error {
	if len(passwd) < 6 {
		return errPasswordTooShort
	}
	return nil
}

func cfbEncrypt(block cipher.Block, iv, src []byte) []byte {
	cfb := cipher.NewCFBEncrypter(block, iv)
	dst := make([]byte, len(src))
	cfb.XORKeyStream(dst, src)
	return dst
}

func cfbDecrypt(block cipher.Block, iv, src []byte) []byte {
	cfb := cipher.NewCFBDecrypter(block, iv)
	dst := make([]byte, len(src))
	cfb.XORKeyStream(dst, src)
	return dst
}
