// Package graphics contains functionality to render and display parts of the ubongo game,
// like single blocks or solutions to problems
package graphics

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"path"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/tidwall/pinhole"

	"ubongo/base/array3d"
	"ubongo/base/vectorf"
	"ubongo/block"
	"ubongo/blockset"
	"ubongo/gamesolution"
)

// convertCoordinates is internally used to convert coordinates for the pinhole library,
// which has a coordinate system confined to +/-1 and also the z-axis is flipped,
// -1 is directed towards the viewer, y is up, x is right
// x, y, z are expected to be integer-array coordinates ranging from 0..maxDim-1
// The coordinates will be scaled such that maxDim equals 1 and everything will be
// moved by offset (before scaling)
func convertCoords(x1, y1, z1, x2, y2, z2 int, offset vectorf.V, maxDim float64) (xf1, yf1, zf1, xf2, yf2, zf2 float64) {
	xf1 = (float64(x1) + offset[0]) / maxDim
	yf1 = (float64(y1) + offset[1]) / maxDim
	zf1 = (-float64(z1) - offset[2]) / maxDim
	xf2 = (float64(x2) + offset[0]) / maxDim
	yf2 = (float64(y2) + offset[1]) / maxDim
	zf2 = (-float64(z2) - offset[2]) / maxDim

	return xf1, yf1, zf1, xf2, yf2, zf2
}

// drawBlock draws the given block at the given position to the pinhole object
func drawBlock(pn *pinhole.Pinhole, blockShape *array3d.A, blockColor block.BlockColor, pos vectorf.V, maxDim float64) {

	// implements the logical function that decides if an edge should be
	// shown based on the presence of a block at the two adjacient and the
	// diagonal volumes
	truthTable := func(adjacient1, adjacient2, diagonal bool) bool {
		if !adjacient1 && !adjacient2 {
			return true
		} else if adjacient1 && adjacient2 {
			return !diagonal
		} else {
			return diagonal
		}
	}

	// returns true if there is solid block at position px,py,z. Returns false
	// if empty of if the coordinates are outside of the volume
	get := func(px, py, pz int) bool {
		if px < 0 || py < 0 || pz < 0 ||
			px >= blockShape.DimX || py >= blockShape.DimY || pz >= blockShape.DimZ {
			return false
		}
		return blockShape.Get(px, py, pz) == 1
	}

	// evaluates if a specific edge of block (x,y,z) is visible or not
	showEdge := func(x, y, z, direction int) bool {
		switch direction {
		case 0: // top front
			return truthTable(get(x, y-1, z), get(x, y, z-1), get(x, y-1, z-1))
		case 1: // top right
			return truthTable(get(x, y-1, z), get(x+1, y, z), get(x+1, y-1, z))
		case 2: // top back
			return truthTable(get(x, y-1, z), get(x, y, z+1), get(x, y-1, z+1))
		case 3: // top left
			return truthTable(get(x, y-1, z), get(x-1, y, z), get(x-1, y-1, z))
		case 4: // front left
			return truthTable(get(x-1, y, z), get(x, y, z-1), get(x-1, y, z-1))
		case 5: // front right
			return truthTable(get(x, y, z-1), get(x+1, y, z), get(x+1, y, z-1))
		case 6: // back right
			return truthTable(get(x+1, y, z), get(x, y, z+1), get(x+1, y, z+1))
		case 7: // back left
			return truthTable(get(x-1, y, z), get(x, y, z+1), get(x-1, y, z+1))
		case 8: // down front
			return truthTable(get(x, y+1, z), get(x, y, z-1), get(x, y+1, z-1))
		case 9: // down right
			return truthTable(get(x, y+1, z), get(x+1, y, z), get(x+1, y+1, z))
		case 10: // down back
			return truthTable(get(x, y+1, z), get(x, y, z+1), get(x, y+1, z+1))
		case 11: // down left
			return truthTable(get(x, y+1, z), get(x-1, y, z), get(x-1, y+1, z))
		}
		return false
	}

	pn.Begin()

	// pinhole has a coordinate system confined to +/-1
	// also the z-axis is flipped, -1 is directed towards the viewer, y is up, x is right
	for x := 0; x < blockShape.DimX; x++ {
		for y := 0; y < blockShape.DimY; y++ {
			for z := 0; z < blockShape.DimZ; z++ {

				// draw the edges of the cube if the position is not empty
				if blockShape.Get(x, y, z) == 1 {
					// top-face, starting front, ccw: 0, 1, 2, 3
					// Y-direction, starting front left: 4, 5, 6, 7
					// bottom-face, starting fron, ccw: 8, 9, 10, 11
					if showEdge(x, y, z, 0) {
						pn.DrawLine(convertCoords(x, y, z, x+1, y, z, pos, maxDim))
					}
					if showEdge(x, y, z, 1) {
						pn.DrawLine(convertCoords(x+1, y, z, x+1, y, z+1, pos, maxDim))
					}
					if showEdge(x, y, z, 2) {
						pn.DrawLine(convertCoords(x+1, y, z+1, x, y, z+1, pos, maxDim))
					}
					if showEdge(x, y, z, 3) {
						pn.DrawLine(convertCoords(x, y, z+1, x, y, z, pos, maxDim))
					}
					if showEdge(x, y, z, 4) {
						pn.DrawLine(convertCoords(x, y, z, x, y+1, z, pos, maxDim))
					}
					if showEdge(x, y, z, 5) {
						pn.DrawLine(convertCoords(x+1, y, z, x+1, y+1, z, pos, maxDim))
					}
					if showEdge(x, y, z, 6) {
						pn.DrawLine(convertCoords(x+1, y, z+1, x+1, y+1, z+1, pos, maxDim))
					}
					if showEdge(x, y, z, 7) {
						pn.DrawLine(convertCoords(x, y, z+1, x, y+1, z+1, pos, maxDim))
					}
					if showEdge(x, y, z, 8) {
						pn.DrawLine(convertCoords(x, y+1, z, x+1, y+1, z, pos, maxDim))
					}
					if showEdge(x, y, z, 9) {
						pn.DrawLine(convertCoords(x+1, y+1, z, x+1, y+1, z+1, pos, maxDim))
					}
					if showEdge(x, y, z, 10) {
						pn.DrawLine(convertCoords(x+1, y+1, z+1, x, y+1, z+1, pos, maxDim))
					}
					if showEdge(x, y, z, 11) {
						pn.DrawLine(convertCoords(x, y+1, z+1, x, y+1, z, pos, maxDim))
					}
				}
			}
		}
	}

	pn.Colorize(blockColor.ToRGBA())

	pn.End()
}

// RenderBlock creates an image of a single block
func RenderBlock(block *block.B, width, height int, rx, ry, rz float64) *image.RGBA {
	pn := pinhole.New()

	shape := block.Shapes[0]
	bb := shape.GetBoundingBox()
	maxDim := float64(bb.Max())
	offset := shape.GetCenterOfGravity().Flip()

	drawBlock(pn, shape, block.Color, offset, maxDim)

	pn.Translate(0, 0, 0)
	pn.Rotate(rx, ry, rz)

	opt := pinhole.ImageOptions{
		BGColor:   color.Black,
		LineWidth: 1.0,
		Scale:     0.9}

	return pn.Image(width, height, &opt)
}

// RenderBlockset renders all blocks of the set to an image each, storing it in the given path
// returns a list of filenames
func RenderBlockset(blocks *blockset.S, dir string, width, height int) []string {
	files := make([]string, 0)
	for _, b := range blocks.AsSlice() {
		img := RenderBlock(b, width, height, math.Pi/4, -math.Pi/8, 0)
		filename := fmt.Sprintf("%s_%s.png", b.Color, b.Name)
		fullpath := path.Join(dir, filename)
		err := SaveAsPng(img, fullpath)
		if err == nil {
			files = append(files, fullpath)
		}
	}
	return files
}

// RenderSolution creates an image the given game solution
// The rotation can be given with the rx, ry, rz parameters
func RenderSolution(gs *gamesolution.S, width, height int, rx, ry, rz, explode float64) *image.RGBA {

	pn := pinhole.New()

	gameCog := gs.GetCenterOfGravity()
	bb := gs.GetBoundingBox()
	maxDim := float64(bb.Max())

	for i, block := range gs.Blocks {
		shapeIdx := gs.ShapeIndex[i]
		shape := block.Shapes[shapeIdx]

		pos := gs.Shifts[i].AsVectorf().Sub(gameCog)
		explodeOffset := pos.Sub(gameCog).Mult(explode)

		drawBlock(pn, shape, block.Color, pos.Add(explodeOffset), maxDim)
	}

	pn.Translate(0, 0, 0)
	pn.Rotate(rx, ry, rz)
	//pn.Rotate(math.Pi/2-0.3+rx, ry, 0.0+rz)

	opt := pinhole.ImageOptions{
		BGColor:   color.Black,
		LineWidth: 1.0,
		Scale:     0.9}

	return pn.Image(width, height, &opt)
}

// SaveAsPng save the given image as png-file to disk
func SaveAsPng(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}

// Visualize shows a gamesolution on the screen using the Fyne library
func Visualize(gs *gamesolution.S) {

	updateImage := func(win fyne.Window, sol *gamesolution.S, w, h int, rx, ry, rz float64) {
		img := RenderSolution(sol, w, h, rx, ry, rz, 0.1)
		win.SetContent(canvas.NewImageFromImage(img))
	}

	a := app.New()
	w := a.NewWindow("Ubongo")
	imgWidth := 800
	imgHeight := 600

	w.Resize(fyne.NewSize(float32(imgWidth), float32(imgHeight)))
	updateImage(w, gs, imgWidth, imgHeight, 0, 0, 0)

	var RX, RY, RZ float64 = math.Pi / 2, 0.0, 0.0
	go func() {
		const minFrameTime = 1.0 / 60 // min time to show one frame in seconds
		const speedRx = 0.0           // radians per second
		const speedRy = 0.2           // radians per second
		const speedRz = 0.0           // radians per second
		lastFrame := time.Now()
		for range time.Tick(time.Millisecond) {
			timePassed := float64(time.Since(lastFrame).Seconds())
			if timePassed >= minFrameTime {
				updateImage(w, gs, imgWidth, imgHeight, RX, RY, RZ)
				RX += speedRx * timePassed
				RY += speedRy * timePassed
				RZ += speedRz * timePassed
				lastFrame = time.Now()
			}
		}
	}()

	w.ShowAndRun()
}
