package config

type constant struct {

	// DB Error
	DbQueryErrMsg  string
	DbUpdateErrMsg string
	DbInsertErrMsg string
	DbDeleteErrMsg string
}

var Constant = constant{

	// DB Error
	DbQueryErrMsg:  "Db Query Error!",
	DbUpdateErrMsg: "Db Update Error!",
	DbInsertErrMsg: "Db Insert Error!",
	DbDeleteErrMsg: "Db Delete Error!",
}
