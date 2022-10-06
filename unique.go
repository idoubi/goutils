package goutils

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

// GenUuid: gen uuid with google/uuid package
func GenUuid() string {
	return uuid.New().String()
}

// GenSnowid: gen unique id with snowflake
func GenSnowid() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return ""
	}

	id := node.Generate()

	return id.String()
}
