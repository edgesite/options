package options

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitCamelCase(t *testing.T) {
	assert.EqualValues(t, SplitCamelCase("HaProxy"), []string{"Ha", "Proxy"})
	assert.EqualValues(t, SplitCamelCase("HAProxy"), []string{"HA", "Proxy"})
	// assert.EqualValues(t, SplitCamelCase("HAProxy!"), []string{"HA", "Proxy!"})
	assert.EqualValues(t, SplitCamelCase("hHaProxy"), []string{"h", "Ha", "Proxy"})
}

func TestSet(t *testing.T) {
	v := struct {
		Uint       uint
		Uint8      uint8
		Uint16     uint16
		Uint32     uint32
		Uint64     uint64
		Int        int
		Int8       int8
		Int16      int16
		Int32      int32
		Int64      int64
		String     string
		Bool_True  bool
		Bool_Yes   bool
		Bool_1     bool
		Bool_False bool
	}{}
	getField := func(i int) reflect.Value { return reflect.ValueOf(&v).Elem().FieldByIndex([]int{i}) }
	set(getField(0), "1")
	set(getField(1), "1")
	set(getField(2), "1")
	set(getField(3), "1")
	set(getField(4), "1")
	set(getField(5), "1")
	set(getField(6), "1")
	set(getField(7), "1")
	set(getField(8), "1")
	set(getField(9), "1")
	set(getField(10), "1")
	set(getField(11), "true")
	set(getField(12), "yes")
	set(getField(13), "1")
	set(getField(14), "false")
	assert.Equal(t, v.Uint, uint(1))
	assert.Equal(t, v.Uint8, uint8(1))
	assert.Equal(t, v.Uint16, uint16(1))
	assert.Equal(t, v.Uint32, uint32(1))
	assert.Equal(t, v.Uint64, uint64(1))
	assert.Equal(t, v.Int, int(1))
	assert.Equal(t, v.Int8, int8(1))
	assert.Equal(t, v.Int16, int16(1))
	assert.Equal(t, v.Int32, int32(1))
	assert.Equal(t, v.Int64, int64(1))
	assert.Equal(t, v.String, "1")
	assert.Equal(t, v.Bool_True, true)
	assert.Equal(t, v.Bool_Yes, true)
	assert.Equal(t, v.Bool_1, true)
	assert.Equal(t, v.Bool_False, false)
}

func TestParseEnv(t *testing.T) {
	v := struct {
		Foo string `options:"auto"`
		Bar string `env:"bar"`
	}{}
	os.Setenv("FOO", "foo")
	os.Setenv("bar", "bar")
	Parse(&v, true, false)
	assert.Equal(t, v.Foo, "foo")
	assert.Equal(t, v.Bar, "bar")
}

func TestRequired(t *testing.T) {
	v := struct {
		Foo string `options:"auto required"`
	}{}
	os.Setenv("FOO", "foo")
	Parse(&v, true, false)
	assert.Panics(t, func() {
		v := struct {
			Bar string `options:"auto required"`
		}{}
		Parse(&v, true, false)
	})
}
