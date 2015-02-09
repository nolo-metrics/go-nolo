package nolo

type MeterPath string

func (p MeterPath) Name() string {
	return filepath.Base(string(p))
}

func (p MeterPath) Open() (*os.File, error) {
	return os.Open(string(p))
}

func (p MeterPath) IsDir() bool {
	fn, err := p.Open()
	if err != nil {
		return false
	}
	fi, err := fn.Stat()
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func (p MeterPath) Readdirnames() []MeterPath {
	fn, err := p.Open()
	if err != nil {
		return []MeterPath{}
	}
	fns, err := fn.Readdirnames(0)
	if err != nil {
		return []MeterPath{}
	}

	meter_paths := []MeterPath{}
	for _, fn := range fns {
		meter_paths = append(meter_paths, MeterPath(fn))
	}
	return meter_paths
}

func (p MeterPath) Expand() ([]MeterPath, error) {
	if p.IsDir() {
		meters := []MeterPath{}
		for _, f := range p.Readdirnames() {
			childfns, err := MeterPath(p + "/" + f).Expand()
			if err != nil {
				return nil, err
			}
			for _, childfn := range childfns {
				meters = append(meters, childfn)
			}
		}
		return meters, nil
	} else {
		// todo: check for executable
		return []MeterPath{p}, nil
	}
}

func (p MeterPath) Execute() (*nolo.Meter, error) {
	out, err := exec.Command(string(p)).Output()
	if err != nil {
		return nil, err
	}
	return nolo.Parse(p.Name(), string(out)), nil
}
