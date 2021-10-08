package docxe

type Command struct {
	extract      bool
	update       bool
	overwriteDir bool
	removeDir    bool
}

func (c *Command) normalize() {
	if c.update {
		c.extract = false
		c.overwriteDir = false
	}

	if c.extract {
		c.update = false
		c.removeDir = false
	}
}

func NewCommand(extract, update, overwriteDir, removeDir bool) Command {
	var config = Command{
		extract:      extract,
		update:       update,
		overwriteDir: overwriteDir,
		removeDir:    removeDir,
	}

	config.normalize()
	return config
}

func (c *Command) Extract() bool {
	return c.extract
}

func (c *Command) Update() bool {
	return c.update
}

func (c *Command) OverwriteDir() bool {
	return c.overwriteDir
}

func (c *Command) RemoveDir() bool {
	return c.removeDir
}
