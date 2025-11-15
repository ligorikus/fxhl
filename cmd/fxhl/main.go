package main

import (
	"fmt"
	"log"
	"math"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/ligorikus/fxhl/internal/gamemap"
	"github.com/ligorikus/fxhl/internal/texture"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	fmt.Println(gamemap.NewHexagonCoordinatesList())
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	monitor := glfw.GetPrimaryMonitor()
	//videoMode := monitor.GetVideoMode()

	window, err := glfw.CreateWindow(600, 600, "Testing", monitor, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	img := texture.NewImage("assets/Hexes/MapDeadlandsHex.TGA")
	rgba := texture.NewRgba(img)

	img2 := texture.NewImage("assets/Hexes/MapMarbanHollowHex.TGA")
	rgba2 := texture.NewRgba(img2)

	deadlands := texture.NewTexture(1, img, rgba)
	marban := texture.NewTexture(2, img2, rgba2)
	defer gl.DeleteTextures(2, &marban)
	defer gl.DeleteTextures(1, &deadlands)

	hex := getHexagoneCoords(0.0, 0.0, 1.0)
	hex2 := getHexagoneCoords(1.0, 1.0, 1.0)
	gl.ClearColor(0.2, 0.2, 0.2, 1.0)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-5.0, 5.0, -5.0, 5.0, -1.0, 1.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		hex.drawHexagone(deadlands)
		hex2.drawHexagone(marban)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func (hex *Hexagone) drawHexagone(texture uint32) {
	gl.Enable(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.Begin(gl.TRIANGLE_FAN)

	gl.TexCoord2f(0.0, 0.5)
	gl.Vertex3f(hex.left.x, hex.left.y, 0.0)

	gl.TexCoord2f(0.25, 0.0)
	gl.Vertex3f(hex.topLeft.x, hex.topLeft.y, 0.0)

	gl.TexCoord2f(0.75, 0.0)
	gl.Vertex3f(hex.topRight.x, hex.topRight.y, 0.0)

	gl.TexCoord2f(1.0, 0.5)
	gl.Vertex3f(hex.right.x, hex.right.y, 0.0)

	gl.TexCoord2f(0.75, 1.0)
	gl.Vertex3f(hex.bottomRight.x, hex.bottomRight.y, 0.0)

	gl.TexCoord2f(0.25, 1.0)
	gl.Vertex3f(hex.bottomLeft.x, hex.bottomLeft.y, 0.0)
	gl.End()

	gl.Disable(gl.TEXTURE_2D)
}

type Vertex struct {
	x float32
	y float32
}

type Hexagone struct {
	left        Vertex
	right       Vertex
	topLeft     Vertex
	topRight    Vertex
	bottomLeft  Vertex
	bottomRight Vertex
}

func getHexagoneCoords(x, y float64, radius float32) Hexagone {

	cx := float32(x*math.Cos(math.Pi/3)+y*math.Sin(math.Pi/3)) * radius
	cy := float32(y*math.Cos(math.Pi/3)-x*math.Sin(math.Pi/3)) * radius

	right := Vertex{
		x: cx + radius,
		y: cy,
	}

	topRight := Vertex{
		x: cx + radius*float32(math.Cos(math.Pi/3)),
		y: cy + radius*float32(math.Sin(math.Pi/3)),
	}

	topLeft := Vertex{
		x: cx + radius*float32(math.Cos(2*math.Pi/3)),
		y: cy + radius*float32(math.Sin(2*math.Pi/3)),
	}

	left := Vertex{
		x: cx - radius,
		y: cy,
	}

	bottomLeft := Vertex{
		x: cx + radius*float32(math.Cos(4*math.Pi/3)),
		y: cy + radius*float32(math.Sin(4*math.Pi/3)),
	}

	bottomRight := Vertex{
		x: cx + radius*float32(math.Cos(5*math.Pi/3)),
		y: cy + radius*float32(math.Sin(5*math.Pi/3)),
	}

	return Hexagone{
		left:        left,
		right:       right,
		topLeft:     topLeft,
		topRight:    topRight,
		bottomLeft:  bottomLeft,
		bottomRight: bottomRight,
	}
}
