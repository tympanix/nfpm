package acceptance

import "testing"

func TestSimpleDeb(t *testing.T) {
	t.Run("amd64", func(t *testing.T) {
		accept(t, "simple_deb", "simple.yaml", "deb", "deb.dockerfile")
	})
	t.Run("i386", func(t *testing.T) {
		accept(t, "simple_deb_386", "simple.386.yaml", "deb", "deb.386.dockerfile")
	})
}

func TestComplexDeb(t *testing.T) {
	t.Run("amd64", func(t *testing.T) {
		accept(t, "complex_deb", "complex.yaml", "deb", "deb.complex.dockerfile")
	})
	t.Run("i386", func(t *testing.T) {
		accept(t, "complex_deb_386", "complex.386.yaml", "deb", "deb.386.complex.dockerfile")
	})
}

func TestMinDeb(t *testing.T) {
	t.Run("amd64", func(t *testing.T) {
		accept(t, "min_deb", "min.yaml", "deb", "deb.min.dockerfile")
	})
}
