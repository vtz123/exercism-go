package erratum

func Use(o ResourceOpener, input string) (err error) {
	//
	//var (
	//	mr Resource
	//	err error
	//)
	mr, err := o()

	if err != nil {
		switch err.(type) {
		case TransientError:
			return Use(o, input)
		default:
			return err
		}
	}

	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case FrobError:
				err = r.(FrobError).inner
				mr.Defrob(r.(FrobError).defrobTag)
			case error:
				err = r.(error)
			default:
				panic(r)
			}
		}
		mr.Close()
	}()

	mr.Frob(input)

	return
}