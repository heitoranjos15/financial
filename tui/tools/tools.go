package tools

type Field struct {
  Input string
  Value string
}

type Tools struct {
  Cursor int
  Fields []Field
}

func (t *Tools) GetCursor() int {
  return t.Cursor
}

func (t *Tools) SetCursor(value int) {
  t.Cursor = value
}

func (t *Tools) GetFields() []Field {
  return t.Fields
}

func (t *Tools) SetFields(fields []Field) {
  t.Fields = fields
}

func (f Field) GetInput() string {
  return f.Input
}

func (f Field) GetValue() string {
  return f.Value
}

func (t *Tools) CursorActions(key string) {
  switch key {
  case "up", "k":
    if t.Cursor > 0 {
      t.Cursor -= 1
    }
  case "down", "j":
    if t.Cursor < len(t.Fields) -1 {
      t.Cursor += 1
    }
  }
}
