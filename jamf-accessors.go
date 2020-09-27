package jamf

func (d *Department) SetName(v string) {
	d.Name = &v
}

func (d *Department) SetId(v string) {
	d.Id = &v
}

func (d *Department) GetName() string {
	if d == nil || d.Name == nil {
		return ""
	}
	return *d.Name
}

func (d *Department) GetId() string {
	if d == nil || d.Id == nil {
		return ""
	}
	return *d.Id
}
