package bookrepo_test

import (
	"context"
	"grpc-server/src/comm/test"
	"grpc-server/src/repo/mongodb/user/bookrepo"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s grpc-server/src/repo/mongodb/user/bookrepo

func Test(t *testing.T) {

	test.Setup()

	g := goblin.Goblin(t)

	g.Describe("bookrepo", func() {

		g.It("Insert", func() {

			ctx := context.Background()
			// insert
			bookrepo.Insert(ctx, "book name", 100, "user name")
		})

	})

}
