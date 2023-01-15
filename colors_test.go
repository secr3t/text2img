package text2img

import "testing"

func TestPickColor(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(PickColor())
	}
}
