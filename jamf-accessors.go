package jamf

func (d *Department) SetName(v string) {
	d.Name = &v
}

func (d *Department) SetId(v string) {
	d.Id = &v
}
