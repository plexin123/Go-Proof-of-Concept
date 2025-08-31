package cmd

type Bear struct {
	age          int32
	color        string
	foodProvider FoodProvider
}

func NewBear(age int32, color string, foodProvider FoodProvider) *Bear {
	return &Bear{
		age:          age,
		color:        color,
		foodProvider: foodProvider,
	}
}
func NewDefaultBear() *Bear {
	return NewBear(
		12, "green", &HoneyProvider{},
	)
}

// FoodProvider could be considered as the label;
type FoodProvider interface {
	GetFood() string
}

// This is just a shell, that needs to declare the function GetFood() if it wants to use that interface !
type HoneyProvider struct{}

type AppleProvider struct{}

func (h *AppleProvider) GetFood() int32 {
	return 123
}

func (h *HoneyProvider) GetFood() string {
	return "Honey"
}

func (b *Bear) eatSomething(times int32) *Bear {
	sample := Bear{age: 12, color: "green"}
	return &sample
}
