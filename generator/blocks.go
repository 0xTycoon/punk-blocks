package main

import (
	"encoding/binary"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
)

type block struct {
	i image.Image
}

const pngHeader = "\x89PNG\r\n\x1a\n"

var alphaChanDetected = errors.New("alpha channel present")

/**
* getPunkBlock returns a single punk building block pic 24x24
 */
func (b *block) getPunkBlock(blockID int) image.Image {
	if blockID > 132 {
		blockID = 0
	}
	x := blockID % 10 * 24
	y := (140 * blockID) / 1400 * 24
	ret := b.i.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(x, y, x+24, y+24))
	return ret
}

func (b *block) load(path string) (img *image.RGBA, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	var ret image.Image
	ret, err = png.Decode(f)
	img = imageToRGBA(ret)
	if err != nil {
		return
	}
	b.i = img
	return
}

/**
optimalPngCompress writes out a png that uses max zlib compression
*/
func optimalPngCompress(w io.Writer, m image.Image) error {
	var e png.Encoder
	e.CompressionLevel = png.BestCompression
	return e.Encode(w, m)

}

func buildPalette(src image.Image) (color.Palette, error) {
	pal := color.Palette([]color.Color{})
	rgba := imageToRGBA(src)
	colors := make(map[color.Color]bool)
	var detected error
	for x := 0; x < 24; x++ {
		for y := 0; y < 24; y++ {
			c := rgba.At(x+rgba.Rect.Min.X, y+rgba.Rect.Min.Y)
			if _, _, _, a := c.RGBA(); a > 0 {
				if a < 65535 {
					detected = alphaChanDetected
				}
			}
			if _, ok := colors[c]; !ok {
				colors[c] = true
				pal = append(pal, c)
			}
		}
	}
	return pal, detected
}

/**
* optimizeImage returns an image that uses an indexed palette when possible
 */
func optimizeImage(src image.Image) (image.Image, error) {
	var pal color.Palette
	var err error
	if pal, err = buildPalette(src); err != nil {
		return src, err
	}
	rgba := imageToRGBA(src)
	img := image.NewPaletted(image.Rect(0, 0, 24, 24), pal)
	draw.Draw(img, img.Bounds(), rgba, src.Bounds().Min, draw.Src)
	return img, nil
}

func imageToRGBA(src image.Image) *image.RGBA {

	// No conversion needed if image is an *image.RGBA.
	if dst, ok := src.(*image.RGBA); ok {
		return dst
	}
	// Use the image/draw package to convert to *image.RGBA.
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
	return dst
}

type ihdr struct {
	width  int32
	height int32
	bd     int // bit depth
	ct     int // color type
	z      int // compression method
	f      int // filter method
	i      int // interlace method
}

type pngFile struct {
	h      ihdr
	p      color.Palette
	chunks [][]byte
}

/**
* parsePng is just a simple PNG validation routine PoC
 */
func parsePng(data []byte) (*pngFile, error) {
	pngf := new(pngFile)
	var pos int32
	if string(data[pos:8]) != pngHeader {
		return pngf, errors.New("not a png")
	}
	pos = 8
	for {
		chunkLen := int32(binary.BigEndian.Uint32(data[pos : pos+4]))
		pos += 4
		chunkType := string(data[pos : pos+4])
		pos += 4
		payload := data[pos : pos+chunkLen]
		pos += chunkLen + 4 // ignore 4 byte CRC32
		switch chunkType {
		case "IHDR":
			pngf.h.width = int32(binary.BigEndian.Uint32(payload[0:4]))
			pngf.h.height = int32(binary.BigEndian.Uint32(payload[4:8]))
			pngf.h.bd = int(payload[8])
			pngf.h.ct = int(payload[9])
			pngf.h.z = int(payload[10])
			pngf.h.f = int(payload[11])
			pngf.h.i = int(payload[12])

		case "PLTE":
			pngf.p = make(color.Palette, 256)
			var i int
			for offset := 0; offset < len(payload); offset += 3 {
				pngf.p[i] = color.RGBA{payload[offset], payload[offset+1], payload[offset+2], 0xff}
				i++
			}
		case "tRNS":
			// todo grayscale?
			if pngf.h.ct == 3 { // palette-ed
				n := len(payload)
				for i := 0; i < n; i++ {
					rgba := pngf.p[i].(color.RGBA)
					pngf.p[i] = color.NRGBA{rgba.R, rgba.G, rgba.B, payload[i]}
				}
			}
		case "IDAT":
			pngf.chunks = append(pngf.chunks, payload) // do not decompress
		case "IEND":
			return pngf, nil
		}
		if int(pos) > len(data) {
			break
		}
	}
	return pngf, nil
}
