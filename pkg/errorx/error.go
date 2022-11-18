package errorx

import "errors"

// TODO error 标准化，最后需要整理一下
var (
	/*
		ErrorX cli errors
	*/
	ErrPrintAndExit    = errors.New("print and exit")
	ErrNoInputConfig   = errors.New("have no input config")
	ErrNoInputFilepath = errors.New("have no input filepath")

	/*
		ErrorX runtime errors
	*/
	ErrPackageLogsFailed = errors.New("package logs failed")

	/*
		ErrorX internal errors
	*/
	ErrConfigInvalid             = errors.New("config invalid")
	ErrEnvKeyInvalid             = errors.New("env key invalid")
	ErrEnvLookupFailed           = errors.New("env lookup failed")
	ErrRevisionNotCommit         = errors.New("revision not commit")
	ErrSnapshotNotFoundStatusSet = errors.New("snapshot not found status set")

	/*
		ErrorX file errors
	*/
	ErrFileNotExisted   = errors.New("file not existed")
	ErrFileHasExisted   = errors.New("file has existed")
	ErrFileTypeNotMatch = errors.New("file type not match")
)
