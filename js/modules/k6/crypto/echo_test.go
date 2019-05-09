package crypto

import (
	"context"
	"testing"

	"github.com/dop251/goja"
	"github.com/loadimpact/k6/js/common"
	"github.com/stretchr/testify/assert"
)

func makeRuntime() *goja.Runtime {
	rt := goja.New()
	rt.SetFieldNameMapper(common.FieldNameMapper{})
	ctx := context.Background()
	ctx = common.WithRuntime(ctx, rt)
	rt.Set("crypto", common.Bind(rt, New(), &ctx))
	return rt
}

func TestEcho(t *testing.T) {
	rt := makeRuntime()

	_, err := common.RunString(rt, `
	const options = { type: "platinum" };
	const result = crypto.echo(options);
	if (result !== "platinum") {
		throw new Error("Distorted echo: " + result);
	}`)
	assert.NoError(t, err)
}
