package sql_handler

type ISqlHandler interface {
	Get(name string)
}

type sqlHandler struct {
}

func NewISqlHandler() ISqlHandler {
	return &sqlHandler{}
}

func (sh *sqlHandler) Get(name string) {

}
