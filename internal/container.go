package bot

type Container struct {
	value interface{}
	frozen bool
}

func NewContainer(value interface{}) Container {
	return Container{value, false}
}

func (c *Container) Set(value interface{}) bool {
	ok := !c.frozen
	if (ok) {
		c.value = value
	}
	return ok
}

func (c *Container) Get() interface{} {
	return c.value
}

func (c *Container) Frozen() bool {
	return c.frozen
}

func (c *Container) Freeze() {
	c.frozen = true
}
