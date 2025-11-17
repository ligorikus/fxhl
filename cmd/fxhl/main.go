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

var (
	baseOrthoMin = -8.0
	baseOrthoMax = 8.0
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)

	monitor := glfw.GetPrimaryMonitor()
	//videoMode := monitor.GetVideoMode()

	window, err := glfw.CreateWindow(600, 600, "Testing", monitor, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()

	window.SetMouseButtonCallback(mouseButtonCallback)
	window.SetScrollCallback(scrollCallback)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	hexTextures := [gamemap.HexNameCount]uint32{}
	hexes := [gamemap.HexNameCount]Hexagone{}
	for key, value := range gamemap.NewHexagonCoordinatesList() {
		img := texture.NewImage("assets/Hexes/Map" + gamemap.HexNameList[gamemap.HexName(key)] + "Hex.TGA")
		rgba := texture.NewRgba(img)

		hexTextures[key] = texture.NewTexture(1, img, rgba)
		defer gl.DeleteTextures(1, &hexTextures[key])

		hexes[key] = getHexagoneCoords(float32(value.GetX()), float32(value.GetY()), 1.0)
	}

	gl.ClearColor(0.2, 0.2, 0.2, 1.0)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(baseOrthoMin, baseOrthoMax, baseOrthoMin, baseOrthoMax, -1.0, 1.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		for key, value := range hexes {
			value.drawHexagone(hexTextures[key])
		}
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

func getHexagoneCoords(x, y float32, radius float32) Hexagone {
	var cx, cy float32

	cx = x * (radius + 0.5)
	cy = 2*y*float32(math.Sin(math.Pi/3)) + x*float32(math.Sin(math.Pi/3))

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

func mouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	buttonName := ""
	switch button {
	case glfw.MouseButtonLeft:
		buttonName = "Left"
	case glfw.MouseButtonRight:
		buttonName = "Right"
	case glfw.MouseButtonMiddle:
		buttonName = "Middle"
	default:
		buttonName = fmt.Sprintf("Button %d", button)
	}

	actionName := ""
	switch action {
	case glfw.Press:
		actionName = "Pressed"
	case glfw.Release:
		actionName = "Released"
	}

	fmt.Printf("Mouse %s button %s (Mods: %v)\n", buttonName, actionName, mods)
}

func scrollCallback(w *glfw.Window, xoff, yoff float64) {
	fmt.Printf("Scroll event: x-offset=%.2f, y-offset=%.2f\n", xoff, yoff)
}
