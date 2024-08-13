package entity

type Category struct {
	ID       string
	Name     string
	CourseID []string
}

func (c *Category) AddCourse(id string) {
	c.CourseID = append(c.CourseID, id)
}

type Course struct {
	ID         string
	Name       string
	CategoryID string
}
