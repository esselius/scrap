package erratum

func Use(ro ResourceOpener, input string) (err error) {
	result, err := ro()
	if err != nil {
		switch err.(type) {
		case TransientError:
			return Use(ro, input)
		default:
			return err
		}
	}

	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case FrobError:
				err = r.(FrobError).inner
				result.Defrob(r.(FrobError).defrobTag)
			case error:
				err = r.(error)
			default:
				panic(r)
			}
		}
		result.Close()
	}()
	result.Frob(input)
	return err
}
