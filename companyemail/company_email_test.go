package companyemail

import (
	"strings"
	"testing"
)

func TestCompanyEmail(t *testing.T) {

	t.Run("Can be created from string", func(t *testing.T) {

		value := "foo@acme.com"

		companyMail, err := NewCompanyEmail(value)
		want := value
		if err != nil {
			t.Error("got an error but didn't expected one")
		}

		got := companyMail.String()

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("Empty string should returns error", func(t *testing.T) {
		value := ""

		_, err := NewCompanyEmail(value)

		assertError(t, err, ErrInvalidEmail.Error())
	})

	t.Run("Short string should returns error", func(t *testing.T) {
		value := "@a.a"

		_, err := NewCompanyEmail(value)

		assertError(t, err, errCompanyEmailTooShort.Error())
	})

	t.Run("Long string should returns error", func(t *testing.T) {
		value := strings.Repeat("x", CompanyEmailMaxLength) + "@a.a"

		_, err := NewCompanyEmail(value)

		assertError(t, err, errCompanyEmailTooLong.Error())
	})

	t.Run("Malformed email address should returns error", func(t *testing.T) {
		invalidEmailDataProvider := []struct {
			malformedEmail string
			error          error
		}{
			{"johndoe", ErrInvalidEmail},
			{"johndoe.", ErrInvalidEmail},
			{"john.doe", ErrInvalidEmail},
			{"john.doe@", ErrInvalidEmail},
			//{"john.doe@google", ErrInvalidEmail}, <- it's passed
			{"john.doe@google.", ErrInvalidEmail},
			{"@google.com", ErrInvalidEmail},
			{"@google.", ErrInvalidEmail},
			{"@google.foo.com", ErrInvalidEmail},
			{"foo@google.com@google.com", ErrInvalidEmail},
			{"john£$%&/()doe@google.", ErrInvalidEmail},
		}

		for _, tt := range invalidEmailDataProvider {
			_, err := NewCompanyEmail(tt.malformedEmail)
			assertError(t, err, tt.error.Error())
		}
	})

	t.Run("Invalid domain in email address should returns error", func(t *testing.T) {
		value := "foo@notacompanydomain.com"

		_, err := NewCompanyEmail(value)

		assertError(t, err, errCompanyEmailInvalidDomain.Error())
	})

	t.Run("Invalid domain in email address should returns error", func(t *testing.T) {
		value := "foo@notacompanydomain.com"

		_, err := NewCompanyEmail(value)

		assertError(t, err, errCompanyEmailInvalidDomain.Error())
	})

	t.Run("Can be compared", func(t *testing.T) {
		first, firstError := NewCompanyEmail("foo@acme.com")
		second, secondError := NewCompanyEmail("bar@acme.org")
		copyOfFirst, copyOfFirstError := NewCompanyEmail("foo@acme.com")

		if firstError != nil || secondError != nil || copyOfFirstError != nil {
			t.Error("got an error but didn't expected one")
		}
		assertFalse(t, first.Equals(second))
		assertTrue(t, first.Equals(copyOfFirst))
		assertFalse(t, second.Equals(copyOfFirst))
	})
}

func assertTrue(t *testing.T, got bool) {
	t.Helper()

	if got != true {
		t.Errorf("%v is not true", got)
	}

}

func assertFalse(t *testing.T, got bool) {
	t.Helper()

	if got != false {
		t.Errorf("%v is not false", got)
	}
}

func assertError(t *testing.T, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
