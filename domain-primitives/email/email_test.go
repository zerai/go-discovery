package email

import "testing"


func TestEmail(t *testing.T)  {
	
	t.Run("Returns internal value", func (t *testing.T)  {
		sut := Email{value: "pippo"}

		got := sut.Value()
		want := "pippo"

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("Returns a string value", func (t *testing.T)  {
		sut := Email{value: "pippo"}

		got := sut.String()
		want := "pippo"

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("Can be compared", func (t *testing.T)  {
		first := Email{value: "pippo"}
		second := Email{value: "pluto"}

		got := first.Equals(second)
		want := false

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}