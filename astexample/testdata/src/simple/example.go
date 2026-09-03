package simple

type PersonalDetails struct {
	name, surname string
}

func NewPersonalDetails(name, surname string) PersonalDetails {
	return PersonalDetails{name: name, surname: surname}
}
