package main

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
	"reflect"
)

const (
	JPEG = 1
	GIF  = 2
	PNG  = 3
)

type Image struct {
	Format int
	Object image.Image
	Width  int
	Height int
}

func NewImage(data []byte) (self *Image, err error) {

	var i image.Image
	var w, h int
	var f int

	// detect format
	if reflect.DeepEqual(data[0:2], []byte{0xff, 0xd8}) {
		f = JPEG
	} else if reflect.DeepEqual(data[0:2], []byte{0x47, 0x49}) {
		f = GIF
	} else if reflect.DeepEqual(data[0:2], []byte{0x89, 0x50}){
		f = PNG
	} else {
		err = errors.New("unsupported format")
	}

	r := bytes.NewReader(data)

	switch f {
		case JPEG:
			i, err = jpeg.Decode(r)
			if err == nil {
				r.Seek(0, 0)
				c, err := jpeg.DecodeConfig(r)
				if err == nil {
					w = c.Width
					h = c.Height
				}
			}
		case GIF:
			i, err = gif.Decode(r)
			if err == nil {
				r.Seek(0, 0)

				c, err := gif.DecodeConfig(r)
				if err == nil {
					w = c.Width
					h = c.Width
				}
			}
		case PNG:
			i, err = png.Decode(r)
			if err == nil {
				r.Seek(0, 0)
				c, err := png.DecodeConfig(r)
				if err == nil {
					w = c.Width
					h = c.Width
				}
			}
	}

	return &Image{
		Format: f,
		Object: i,
		Width:  w,
		Height: h,
	}, err
}
