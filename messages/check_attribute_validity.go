package messages

import "fmt"

func (m *Messages) CheckAttributeValidity(exists bool) {
	if !exists {
		fmt.Println(AttributeNotFound)
		return
	}
}
