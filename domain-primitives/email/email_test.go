package email

import "testing"


func TestEmail(t *testing.T)  {
	
	t.Run("Returns internal value", func (t *testing.T)  {
		value := "foo@example.com"
		
		email := Email{value: value}

		got := email.Value()
		want := value

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("Returns a string value", func (t *testing.T)  {
		value := "foo@example.com"
		
		email := Email{value: value}

		got := email.String()
		want := value

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("Can be compared", func (t *testing.T)  {
		first := Email{value: "foo@example.com"}
		second := Email{value: "bar@example.com"}
		copyOfFirst := Email{value: "foo@example.com"}

		assertTrue(t, first.Equals(copyOfFirst))
		assertFalse(t, first.Equals(second))
		assertFalse(t, second.Equals(copyOfFirst))
	})
}

func assertTrue(t *testing.T, got bool)  {
	t.Helper()

	if got != true {
		t.Errorf("%v is not true", got)
	}

}

func assertFalse(t *testing.T, got bool)  {
	t.Helper()

	if got != false {
		t.Errorf("%v is not false", got)
	}
}