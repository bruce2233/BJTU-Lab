package image_processing

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

func image() {
	// This example uses png.Decode which can only decode PNG images.
	catFile, err := os.Open("E:\\学习资料\\数字图像处理\\2.png")
	outputFile, err := os.OpenFile("E:\\学习资料\\数字图像处理\\output.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer catFile.Close()

	// Consider using the general image.Decode as it can sniff and decode any registered image format.
	img, err := png.Decode(catFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(img)

	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			if level == 5 {
				level--
			}
			fmt.Print(levels[level])
			fmt.Fprint(outputFile, levels[level])
		}
		fmt.Print("\n")
		fmt.Fprint(outputFile, "\n")

	}
}
