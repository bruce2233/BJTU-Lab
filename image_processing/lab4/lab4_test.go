package lab4

import (
	"image/color"
	"testing"
	qrcode "github.com/skip2/go-qrcode"
)

func TestLab4(t *testing.T) {
	qrcode.WriteColorFile("https://example.org", qrcode.Medium, 256, color.Black, color.White, "qr.png")

}
