package types

import "fmt"

type Errors []error

func (e *Errors) Add(err error) {
	*e = append(*e, err)
}

func (e *Errors) ToError(header error) error {
	if len(*e) == 0 {
		return nil
	}

	var err error
	if header != nil {
		err = fmt.Errorf("%w\n", header)
	}

	for _, val := range *e {
		err = fmt.Errorf(
			"%w\n%w",
			err,
			val,
		)
	}
	return err
}
