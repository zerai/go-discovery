package email

type Email struct{
	value string
}

func NewEmail(email string) *Email {
    return &Email{
		value: email,
	}
}

func (e Email) Value() string {
	return e.value
}

func (e Email) String() string {
	return e.value
}

func (e Email) Equals(other Email) bool {
	if e.value == other.value {
		return true
	}
	
	return false
}

