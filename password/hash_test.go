package password

import (
	"testing"
)

func TestHashPassword_AngryMonkey(t *testing.T) {
	expected := "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="
	password := "angryMonkey"
	actual := HashPassword512(password)

	if actual != expected {
		t.Fatalf("%s != %s", actual, expected)
	}
}
