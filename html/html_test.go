package html

import "testing"

func TestGenerate(t *testing.T) {
	_, err := Generate(50)
	if err == nil {
		t.Error(err)
	}
	_, err = Generate(30)
	if err != nil {
		t.Error("fixme")
	}
}
