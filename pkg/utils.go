package utils

import (
	"crypto/sha512"
	"math/rand"
	"unicode/utf8"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func HashPassword(password string) []byte {
	hashPassword := sha512.Sum512([]byte(password))
	passwordByteSlice := hashPassword[:]
	return passwordByteSlice
}

func RandStringRunes(seed int) string {
	symbols := make([]rune, seed)
	for i := range symbols {
		symbols[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(symbols)
}

func ValidatorComment(text string) bool {
	return utf8.RuneCountInString(text) < MaxTextComment
}

const (
	DefaultOffset  = 0
	DefaultLimit   = 10
	MaxRetries     = 3
	MaxTextComment = 2000
)

const (
	InternalError    = "Internal Server Error"
	BadRequest       = "Bad Request"
	ConvertedIdError = "Converted id error"
	PaginatorError   = "limit or offset exceeded"
	SessionNull      = "Session is null"
	ParentIdRequired = "parentID is required"
	CommentIsTooBig  = "Comment is too big"
)

const (
	CommentsByPostId      = "GetCommentsByPostId"
	CommentsByCommentId   = "GetCommentsByCommentId"
	GetPost               = "GetPost"
	CreatePost            = "CreatePost"
	CreateComment         = "CreateComment"
	GetPosts              = "GetPosts"
	PostOrCommentNotFound = "Post or comment not found"
)
