package utils

import (
	"os"
	"strconv"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

func CreateAccountObject(id uint) *v1.ObjectReference {
	return &v1.ObjectReference{
		ObjectType: os.Getenv("SPICE_DB_PREFIX") + "/user",
		ObjectId:   strconv.FormatUint(uint64(id), 10),
	}
}

func CreateDocumentObject(id uint) *v1.ObjectReference {
	return &v1.ObjectReference{
		ObjectType: os.Getenv("SPICE_DB_PREFIX") + "/document",
		ObjectId:   strconv.FormatUint(uint64(id), 10),
	}
}

func CreateFolderObject(id uint) *v1.ObjectReference {
	return &v1.ObjectReference{
		ObjectType: os.Getenv("SPICE_DB_PREFIX") + "/document",
		ObjectId:   strconv.FormatUint(uint64(id), 10),
	}
}
