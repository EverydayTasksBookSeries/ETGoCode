package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	// 画一个大小为200x200的正方形
	size := 200
	step := size / 10
	imgRect := image.Rect(0, 0, size, size)
	img := image.NewNRGBA(imgRect)

	// 画出边框
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	// 填充黑色
	for x := 0; x < size; x += step {
		// 奇数块涂上阿根廷国旗的蓝色
		colorIceberg := color.NRGBA{117, 170, 219, 255}
		fill := &image.Uniform{colorIceberg}
		// 偶数块涂上白色
		if (x/step)%2 == 0 {
			fill = &image.Uniform{color.White}
		}
		draw.Draw(img, image.Rect(x, 0, x+step, size), fill, image.ZP, draw.Src)
	}
	// 打开输出文件
	outputFile, err := os.Create("D:/argentina.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// 调用png.Encode()来输出图片。参数是文件名和上面生成的Image对象
	png.Encode(outputFile, img)

}
