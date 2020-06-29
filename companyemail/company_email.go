package companyemail

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	//CompanyEmailMinLength : company email minimum lenght
	CompanyEmailMinLength = 5

	//CompanyEmailMaxLength :  company email minimum lenght
	CompanyEmailMaxLength = 255
)

var (
	companyDomains = [...]string{"acme.com", "acme.org", "acme-store.com"}

	ErrInvalidEmail              = errors.New("Invalid Email.")
	errCompanyEmailTooShort      = errors.New(fmt.Sprintf("Invalid email: min length allowed is %d", CompanyEmailMinLength))
	errCompanyEmailTooLong       = errors.New(fmt.Sprintf("Invalid email: max length allowed is %d", CompanyEmailMaxLength))
	errCompanyEmailInvalidDomain = errors.New(fmt.Sprintf("Invalid email: this email contains an invalid domain name."))

	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

//CompanyEmail : value object
type CompanyEmail interface {
	String() string
	Value() string
	Equals(o CompanyEmail) bool
}

type companyEmail struct {
	value string
}

func (p *companyEmail) String() string {
	return p.value
}

func (p *companyEmail) Value() string {
	return p.value
}

func (p *companyEmail) Equals(o CompanyEmail) bool {
	if p.Value() == o.Value() {
		return true
	}
	return false
}

func isValidDomainName(emailAddress string) bool {
	components := strings.Split(emailAddress, "@")
	emailDomain := components[1]
	for _, domain := range companyDomains {
		if emailDomain == domain {
			return true
		}
	}
	return false
}

// Here, NewCompanyEmail returns an interface, and not the companyEmail struct itself
/*   func NewCompanyEmail(name string) CompanyEmail {
	return companyEmail{
	  value: name,
	}
  }
*/

// Here, NewCompanyEmail returns a pointer to companyEmail struct
/* func NewCompanyEmail(emailAddress string) (CompanyEmail, error) {

	emailAddress = strings.TrimSpace(emailAddress)
	if len(emailAddress) == 0 {
		return nil, ErrInvalidEmail
	}

	if len(emailAddress) < CompanyEmailMinLength {
		return nil, errCompanyEmailTooShort
	}

	if len(emailAddress) > CompanyEmailMaxLength {
		return nil, errCompanyEmailTooLong
	}

	if !emailRegexp.MatchString(emailAddress) {
		return nil, ErrInvalidEmail
	}

	return &companyEmail{
		value: emailAddress,
	}, nil
} */

func NewCompanyEmail(emailAddress string) (CompanyEmail, error) {
	emailAddress = strings.TrimSpace(emailAddress)
	if len(emailAddress) == 0 {
		return nil, ErrInvalidEmail
	}

	if len(emailAddress) < CompanyEmailMinLength {
		return nil, errCompanyEmailTooShort
	}

	if len(emailAddress) > CompanyEmailMaxLength {
		return nil, errCompanyEmailTooLong
	}

	if !emailRegexp.MatchString(emailAddress) {
		return nil, ErrInvalidEmail
	}

	if !isValidDomainName(emailAddress) {
		return nil, errCompanyEmailInvalidDomain
	}

	return &companyEmail{
		value: emailAddress,
	}, nil
}
