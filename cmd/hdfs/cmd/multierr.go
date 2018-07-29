package cmd

import "strings"

type multierr struct {
	errs []error
}

func (m multierr) Error() string {
	out := make([]string, len(m.errs))
	for i, x := range m.errs {
		out[i] = x.Error()
	}
	return strings.Join(out, "\n")
}

func (m *multierr) AddErr(err error) {
	m.errs = append(m.errs, err)
}

func (m multierr) IsError() error {
	if len(m.errs) == 0 {
		return nil
	}
	return m
}
