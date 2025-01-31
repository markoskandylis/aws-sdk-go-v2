package ini

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidDataFiles(t *testing.T) {
	const expectedFileSuffix = "_expected"
	err := filepath.Walk(filepath.Join("testdata", "valid"),
		func(path string, info os.FileInfo, fnErr error) (err error) {
			if strings.HasSuffix(path, expectedFileSuffix) {
				return nil
			}

			if info.IsDir() {
				return nil
			}

			f, err := os.Open(path)
			if err != nil {
				t.Errorf("%s: unexpected error, %v", path, err)
			}

			defer func() {
				closeErr := f.Close()
				if err == nil {
					err = closeErr
				} else if closeErr != nil {
					err = fmt.Errorf("file close error: %v, original error: %w", closeErr, err)
				}
			}()

			tree, err := ParseAST(f)
			if err != nil {
				t.Errorf("%s: unexpected parse error, %v", path, err)
			}

			v := NewDefaultVisitor(path)
			err = Walk(tree, v)
			if err != nil {
				t.Errorf("%s: unexpected walk error, %v", path, err)
			}

			expectedPath := path + "_expected"
			e := map[string]interface{}{}

			b, err := ioutil.ReadFile(expectedPath)
			if err != nil {
				// ignore files that do not have an expected file
				return nil
			}

			err = json.Unmarshal(b, &e)
			if err != nil {
				t.Errorf("unexpected error during deserialization, %v", err)
			}

			for profile, tableIface := range e {
				p, ok := v.Sections.GetSection(profile)
				if !ok {
					t.Fatal("could not find profile " + profile)
				}

				table := tableIface.(map[string]interface{})
				for k, v := range table {
					switch e := v.(type) {
					case string:
						a := p.String(k)
						if e != a {
							t.Errorf("%s: expected %v, but received %v for profile %v", path, e, a, profile)
						}
					default:
						t.Errorf("unexpected type: %T", e)
					}
				}
			}

			return nil
		})
	if err != nil {
		t.Fatalf("Error while walking the file tree rooted at root, %d", err)
	}
}

func TestInvalidDataFiles(t *testing.T) {
	cases := []struct {
		path               string
		expectedParseError bool
		expectedWalkError  bool
	}{
		{
			path:               "./testdata/invalid/bad_syntax_1",
			expectedParseError: true,
		},
		{
			path:               "./testdata/invalid/bad_syntax_2",
			expectedParseError: true,
		},
		{
			path:               "./testdata/invalid/incomplete_section_profile",
			expectedParseError: true,
		},
		{
			path:               "./testdata/invalid/syntax_error_comment",
			expectedParseError: true,
		},
		{
			path:               "./testdata/invalid/invalid_keys",
			expectedParseError: true,
		},
		{
			path:               "./testdata/invalid/bad_section_name",
			expectedParseError: true,
		},
	}

	for i, c := range cases {
		t.Run(c.path, func(t *testing.T) {
			f, err := os.Open(c.path)
			if err != nil {
				t.Errorf("unexpected error, %v", err)
			}
			defer func() {
				closeErr := f.Close()
				if closeErr != nil {
					t.Errorf("unexpected file close error: %v", closeErr)
				}
			}()

			tree, err := ParseAST(f)
			if err != nil && !c.expectedParseError {
				t.Errorf("%d: unexpected error, %v", i+1, err)
			} else if err == nil && c.expectedParseError {
				t.Errorf("%d: expected error, but received none", i+1)
			}

			if c.expectedParseError {
				return
			}

			v := NewDefaultVisitor(c.path)
			err = Walk(tree, v)
			if err == nil && c.expectedWalkError {
				t.Errorf("%d: expected error, but received none", i+1)
			}
		})
	}
}
