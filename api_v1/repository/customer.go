package repository

type Customer struct {
	// CONN *sql.CONN,
}

func (repository Customer) Get(params interface{}) (interface{}, error) {
	return params, nil
}

func (repository Customer) All(params interface{}) ([]interface{}, error) {
	return nil, nil
}

func (repository Customer) Put(params interface{}, where interface{}) error {
	return nil
}

func (repository Customer) Post(params interface{}) error {
	return nil
}

func (repository Customer) Delete(params interface{}) error {
	return nil
}
