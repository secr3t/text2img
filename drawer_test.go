package text2img

import (
	"image/png"
	"os"
	"testing"
)

func TestDraw(t *testing.T) {
	d, err := NewDrawer(Params{})
	if err != nil {
		t.Fatal(err.Error())
	}

	img, err := d.Draw("랜덤 이미지 생성 테스트")
	if err != nil {
		t.Fatal(err.Error())
	}
	file, err := os.Create("test.png")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer file.Close()
	if err = png.Encode(file, img); err != nil {
		t.Fatal(err.Error())
	}
}
