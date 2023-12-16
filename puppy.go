package puppy

import (
	dog "github.com/listonaj/GoToDog"
)

func Bark() string {
	return "woof"
}

func Barks() string {
	return "woof! woof! woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
