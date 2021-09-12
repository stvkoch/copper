package repository

type Repository interface {
	Get(where interface{}) (interface{}, error)
	All(where interface{}) ([]interface{}, error)
	Put(params interface{}, where interface{}) error
	Post(params interface{}) error
	Delete(where interface{}) error
}
