package validate

type Validatable interface {
	Validate() error
}

func Do[V Validatable](val V) (V, error) {
	err := val.Validate()
	if err != nil {
		return *new(V), err
	}
	return val, nil
}

func FromPtr[V any](val *V, err error) (V, error) {
	return *val, err
}
