package internal

func (m Model) RenderDialog(w, h int, content string) string {
	v := applyWidth(content, w)
	v = centerX(v, m.width)
	v = centerY(v, m.height)
	return v
}
