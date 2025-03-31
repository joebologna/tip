package main

import (
	"embed"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//go:embed icons.yaml assets/icons/*
var embeddedFiles embed.FS

type TileParams struct {
	iconRec   IconRec
	tileSize  float32
	labelSize fyne.Size
	icons     Icons
	tap       func(string)
}

type Tile struct {
	widget.BaseWidget
	TileParams
}

func NewTile(iconRec IconRec, tileSize float32, labelSize fyne.Size, tap func(string)) *Tile {
	tile := &Tile{
		TileParams: TileParams{
			iconRec:   iconRec,
			tileSize:  tileSize,
			labelSize: labelSize,
			tap:       tap,
		},
	}
	tile.SetText(iconRec.Name)
	tile.ExtendBaseWidget(tile)
	return tile
}

func (tile *Tile) Tapped(_ *fyne.PointEvent) {
	tile.tap(tile.iconRec.Name)
}

func (tile *Tile) SetText(text string) {
	tile.iconRec.Name = text
	tile.Refresh()
}

type tileRenderer struct {
	tile   *Tile
	bg     *canvas.Rectangle
	labels [2]*canvas.Text
	icon   *widget.Icon
}

func (r *tileRenderer) MinSize() fyne.Size {
	return fyne.NewSquareSize(r.tile.tileSize)
}

func (r *tileRenderer) Layout(_ fyne.Size) {
	r.bg.Resize(fyne.NewSquareSize(r.tile.tileSize))
	r.icon.Resize(fyne.NewSquareSize(r.tile.tileSize))
	r.labels[0].SetMinSize(r.tile.labelSize)
	mid := r.tile.tileSize / 2
	var o float32 = float32(len(r.labels[0].Text)) / 2.0
	r.labels[0].Move(fyne.NewPos(mid-o, r.tile.tileSize-r.labels[0].TextSize*1.5))
}

func (*tileRenderer) Destroy() {}

func (*tileRenderer) Visible() bool { return true }

func (r *tileRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bg, r.icon, r.labels[0]}
}

func (r *tileRenderer) Refresh() {
	r.labels[0].Text = r.tile.iconRec.Name
	r.labels[0].Refresh()
	fmt.Println("refresh called")
}

func (tile *Tile) CreateRenderer() fyne.WidgetRenderer {
	l := canvas.NewText(tile.iconRec.Name, color.Black)
	l.Alignment = fyne.TextAlignCenter
	l.TextStyle = fyne.TextStyle{Bold: true}
	r := &tileRenderer{
		tile:   tile,
		bg:     canvas.NewRectangle(color.White),
		labels: [2]*canvas.Text{l},
		icon:   tile.getIcon(),
	}
	return r
}

func (tile *Tile) getIcon() *widget.Icon {
	iconData, err := embeddedFiles.ReadFile(tile.iconRec.Path)
	if err != nil {
		return widget.NewIcon(theme.WindowCloseIcon())
	}
	resource := fyne.NewStaticResource(tile.iconRec.Name, iconData)
	icon := widget.NewIcon(resource)
	return icon
}

var _ fyne.WidgetRenderer = (*tileRenderer)(nil)
