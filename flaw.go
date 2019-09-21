package flaw

import (
	"database/sql"

	"github.com/go-pg/pg"
	"github.com/jinzhu/gorm"
)

//CustomError : custom error
func CustomError(debugMsg ...string) *Error {
	return newError(nil, CustomCode, debugMsg...)
}

//UnmarshalError : error occured while unmarshal
func UnmarshalError(err error, debugMsg ...string) *Error {
	return newError(err, UnmarshalCode, debugMsg...)
}

//MarshalError : error occured while unmarshal
func MarshalError(err error, debugMsg ...string) *Error {
	return newError(err, MarshalCode, debugMsg...)
}

//MiscError : error occured while processing
func MiscError(err error, debugMsg ...string) *Error {
	return newError(err, MiscCode, debugMsg...)
}

//ConnError : error occured while creating connection
func ConnError(err error, debugMsg ...string) *Error {
	return newError(err, ConnCode, debugMsg...)
}

//SelectError : error occured while select query
func SelectError(err error, debugMsg ...string) *Error {
	if err == pg.ErrNoRows {
		return NotFoundError(debugMsg...)
	}
	return newError(err, SelectCode, debugMsg...)
}

//SelectIgnoreNoRow : error occured while select query by ignoring no row error
func SelectIgnoreNoRow(err error, debugMsg ...string) error {
	if err != pg.ErrNoRows && err != sql.ErrNoRows && err != gorm.ErrRecordNotFound {
		return newError(err, SelectCode, debugMsg...)
	}
	return nil
}

//CreateError : error occured while creating table
func CreateError(err error, debugMsg ...string) *Error {
	return newError(err, CreateCode, debugMsg...)
}

//InsertError : error occured while insert query
func InsertError(err error, debugMsg ...string) *Error {
	return newError(err, InsertCode, debugMsg...)
}

//UpdateError : error occured while update query
func UpdateError(err error, debugMsg ...string) *Error {
	return newError(err, UpdateCode, debugMsg...)
}

//DeleteError : error occured while delete query
func DeleteError(err error, debugMsg ...string) *Error {
	return newError(err, DeleteCode, debugMsg...)
}

//TxError : error occured while starting transaction
func TxError(err error, debugMsg ...string) *Error {
	return newError(err, TxCode, debugMsg...)
}

//NotFoundError : error occured when id not found
func NotFoundError(debugMsg ...string) *Error {
	if len(debugMsg) == 0 {
		debugMsg = append(debugMsg, "Record not exists")
	}
	return newError(nil, NotFoundCode, debugMsg...)
}

//BadReqError : error occured while validating request
//like while typecasting request, fk in request dosn't exists
func BadReqError(err error, debugMsg ...string) *Error {
	return newError(err, BadReqCode, debugMsg...)
}

//ForbiddenErr : unauthorized access
func ForbiddenErr(debugMsg ...string) *Error {
	return newError(nil, ForbiddenCode, debugMsg...)
}
