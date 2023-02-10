package exceptions

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfErrorNotFound(err error) {
	if err != nil {
		panic(NewDataNotFoundError(err.Error()))
	}
}

func PanicIfErrorQuery(err error) {
	if err != nil {
		panic(NewDataNotFoundError(err.Error()))
	}
}
