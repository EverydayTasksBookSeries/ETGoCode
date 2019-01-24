package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func changeColor(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	res := image.NewRGBA(bounds)
	white := color.RGBA{255, 255, 255, 255}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			c := img.At(i, j)
			if c == white {
				// 如果是白色，设置成黑色
				res.SetRGBA(i, j, color.RGBA{0, 0, 0, 255})
			} else {
				// 如果不是白色，设置成红色
				res.SetRGBA(i, j, color.RGBA{255, 0, 0, 255})
			}
		}
	}
	return res
}

func main() {
	// 打开PNG文件
	f, err := os.Open("D:/argentina.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	newImg := changeColor(img)

	// 打开目标文件
	fout, err := os.Create("D:/acmilan.png")
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	png.Encode(fout, newImg)
}
