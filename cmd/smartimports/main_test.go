package main

import "testing"

func Test_processData(t *testing.T) {
	tests := [...]struct {
		name, src, want string
	}{
		{
			name: "SeveralImportBlocksToOne",
			src: `package main

import "context"
import "os"
`,
			want: `package main

import (
	"context"
	"os"
)
`,
		},
		{
			name: "MergeImportSections",
			src: `package main

import (
	"context"

	"os"
)
`,
			want: `package main

import (
	"context"
	"os"
)
`,
		},
		{
			name: "SeparateStdlib",
			src: `package main

import (
	"context"
	"github.com/pkg/errors"
	"os"
)`,
			want: `package main

import (
	"context"
	"os"

	"github.com/pkg/errors"
)
`,
		},
		{
			name: "NonPreformattedImports",
			src: `package main

	import (
	"context"
  "os"

		"github.com/pkg/errors"

	"github.com/bradfitz/gomemcache"
	"fmt"
	)
`,
			want: `package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bradfitz/gomemcache"
	"github.com/pkg/errors"
)
`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := processData([]byte(test.src), getDefaultOpts())

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if string(got) != test.want {
				t.Errorf("unexpected result: got: %q, want: %q", string(got), test.want)
			}
		})
	}
}
