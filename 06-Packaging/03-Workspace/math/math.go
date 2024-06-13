package math

type MathPublic struct {
	A int // public
	B int // public
}

// public
func (m *MathPublic) Add() int {
	return m.A + m.B
}

type mathPrivate struct {
	a int // public
	b int // private
}

// We can export a private struct using a public function
func NewMathPrivate(a, b int) *mathPrivate {
	return &mathPrivate{a: a, b: b}
}

// public
func (m *mathPrivate) Add() int {
	return m.a + m.b
}

type MathMixed struct {
	A int // public
	b int // private
}

// public
func (m *MathMixed) Add() int {
	return m.A + m.getB()
}

// public
func (m *MathMixed) SetB(b int) {
	m.b = b
}

// private
func (m *MathMixed) getB() int {
	return m.b
}
