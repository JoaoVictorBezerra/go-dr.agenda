package customErrors

import "errors"

var StmtCloseError = errors.New("error while closing statements")
var RowsCloseError = errors.New("error while closing rows")
