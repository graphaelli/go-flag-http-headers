package headerflag

import (
	"flag"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGoodHeaders(t *testing.T) {
	cases := []struct {
		argv []string
		want http.Header
	}{
		{
			argv: []string{"Authorization=Bearer mytoken"},
			want: http.Header{
				"Authorization": []string{"Bearer mytoken"},
			},
		},
		{
			argv: []string{
				"Authorization=ApiKey foo",
				"Content-Type=application/json",
			},
			want: http.Header{
				"Authorization": []string{"ApiKey foo"},
				"Content-Type":  []string{"application/json"},
			},
		},
		{
			argv: []string{
				"X-Test-1=ok1",
				"X-Test-1=ok2",
			},
			want: http.Header{
				"X-Test-1": []string{"ok1", "ok2"},
			},
		},
	}

	for _, c := range cases {
		set := flag.NewFlagSet(t.Name(), flag.ContinueOnError)
		got := New()
		set.Var(got, "header", "header, can be repeated")

		argc := len(c.argv)
		argv := make([]string, argc*2)
		for i := 0; i < argc; i++ {
			argv[i*2] = "-header"
		}
		for i := 0; i < argc; i++ {
			argv[i*2+1] = c.argv[i]
		}

		err := set.Parse(argv)
		require.NoError(t, err)
		require.Equal(t, c.want, got.Headers())
	}
}

func TestBadHeaders(t *testing.T) {
	set := flag.NewFlagSet(t.Name(), flag.ContinueOnError)
	got := New()
	set.Var(got, "header", "header, can be repeated")
	err := set.Parse([]string{"-header", "foo"})
	require.Error(t, err)
}
