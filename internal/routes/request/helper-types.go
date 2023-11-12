package request

import "time"

type DateTime time.Time

func (ct *DateTime) UnmarshalParam(param string) error {
	if param == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02T15:04", param)
	if err != nil {
		return err
	}
	*ct = DateTime(t)
	return nil
}
