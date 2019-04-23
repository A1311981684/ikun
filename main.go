package main

import (
	"code.google.com/p/graphics-go/graphics"
	"fmt"
	"image"
	"image/gif"
	"os"
	"time"
)

func main() {
	//Get gif name from command line argument or provide nothing to use default kun kun gif
	//从命令行参数获取指定的GIF，或者不提供命令行参数使用默认的坤坤gif
	var selectedGIF string
	if len(os.Args) == 2 {
		selectedGIF = os.Args[1]
	} else if len(os.Args) == 1 {
		selectedGIF = "./BBB.gif"
	}
	//Open specific gif
	//打开特定的gif
	f, err := os.Open(selectedGIF)
	if err != nil {
		panic(err)
	}
	//Decode gif frame data
	//解码gif帧数据
	frames, err := gif.DecodeAll(f)
	if err != nil {
		panic(err)
	}
	//Don't forget to close opened file at the end of the programme execution
	//程序最后记得关闭打开的文件
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()
	//Perform the traversal of all frames
	//执行每帧遍历
	for i := 0; i < 5; i++ {
		for _, v := range frames.Image {
			//Resize image in case it is too big to display
			//重新调整图像大小以免过大显示不全
			ni := resizeImage(v, 100)
			//Convert a image into a char image
			//转化图像为字幅图像并显示
			makeCharImage(ni)
			//Delay : delay between two frames. Modify this if needed
			//延时  : 两帧之间的延时，如果有何必要就修改之
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func resizeImage(img image.Image, scale int) *image.RGBA {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newImg := image.NewRGBA(image.Rect(0, 0, scale, scale*dy/dx))
	graphics.Scale(newImg, img)
	return newImg
}
func makeCharImage(img image.Image) {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	//Mapping pixel to a single char string
	//映射像素信息到单字符string
	arrs := []string{"B", "W", "G", "Q", "$", "O", "C", "?", "7", "0", "–", "–", "–", "–", "."}
	//A slice stores all chars of one char image
	//Note that if you don't use this frame by frame buffer method and instead, you want to print line by line, you may
	//be not comfortable about what is going to happen: not continuous output or even worse, making dog eyes blind
	//一个存有整帧字符图像字符的slice
	//如果不采用这种缓冲方式而选择一行行输出,可能会闪瞎狗眼
	var singleFrame []string
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			colorRgb := img.At(j, i)
			_, g, _, _ := colorRgb.RGBA()
			avg := uint8(g >> 8)
			num := avg / 18
			singleFrame = append(singleFrame, arrs[num])
			if j == dx-1 {
				//At the end of a line, append a return
				//末尾加入回车换行
				singleFrame = append(singleFrame, "\n")
			}
		}
	}
	//Output a frame
	//输出一帧
	fmt.Printf("%s", singleFrame)
}
