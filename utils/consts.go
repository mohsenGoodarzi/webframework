package utils

type DBOperationResult int

const (
	FAILED    DBOperationResult = 1000
	SUCCSSED  DBOperationResult = 1001
	NO_EFFECT DBOperationResult = 1002
)
