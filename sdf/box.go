//-----------------------------------------------------------------------------
/*

 */
//-----------------------------------------------------------------------------

package sdf

//-----------------------------------------------------------------------------

type Box3 struct {
	Min, Max V3
}

type Box2 struct {
	Min, Max V2
}

//-----------------------------------------------------------------------------

// create a new Box with a given center and size
func NewBox3(center, size V3) Box3 {
	half := size.MulScalar(0.5)
	return Box3{center.Sub(half), center.Add(half)}
}

// create a new Box with a given center and size
func NewBox2(center, size V2) Box2 {
	half := size.MulScalar(0.5)
	return Box2{center.Sub(half), center.Add(half)}
}

//-----------------------------------------------------------------------------

// are the boxes equal?
func (a Box3) Equals(b Box3, tolerance float64) bool {
	return (a.Min.Equals(b.Min, tolerance) && a.Max.Equals(b.Max, tolerance))
}

// are the boxes equal?
func (a Box2) Equals(b Box2, tolerance float64) bool {
	return (a.Min.Equals(b.Min, tolerance) && a.Max.Equals(b.Max, tolerance))
}

//-----------------------------------------------------------------------------

// return a box that encloses two boxes
func (a Box3) Extend(b Box3) Box3 {
	return Box3{a.Min.Min(b.Min), a.Max.Max(b.Max)}
}

// return a box that encloses two boxes
func (a Box2) Extend(b Box2) Box2 {
	return Box2{a.Min.Min(b.Min), a.Max.Max(b.Max)}
}

//-----------------------------------------------------------------------------

// translate a box a distance v
func (a Box3) Translate(v V3) Box3 {
	return Box3{a.Min.Add(v), a.Max.Add(v)}
}

// translate a box a distance v
func (a Box2) Translate(v V2) Box2 {
	return Box2{a.Min.Add(v), a.Max.Add(v)}
}

//-----------------------------------------------------------------------------

// return the size of the box
func (a Box3) Size() V3 {
	return a.Max.Sub(a.Min)
}

// return the size of the box
func (a Box2) Size() V2 {
	return a.Max.Sub(a.Min)
}

// return the center of the box
func (a Box3) Center() V3 {
	return a.Min.Add(a.Size().MulScalar(0.5))
}

// return the center of the box
func (a Box2) Center() V2 {
	return a.Min.Add(a.Size().MulScalar(0.5))
}

//-----------------------------------------------------------------------------

// Return a slice of box vertices
func (a Box2) Vertices() V2Set {
	v := make([]V2, 4)
	v[0] = a.Min                // bl
	v[1] = V2{a.Max.X, a.Min.Y} // br
	v[2] = V2{a.Min.X, a.Max.Y} // tl
	v[3] = a.Max                // tr
	return v
}

//-----------------------------------------------------------------------------
