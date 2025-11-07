package paste_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/paste"
)

func ExamplePaste_basic() {
	// paste file1.txt file2.txt
	// Note: This requires actual files to paste together
	yup.MustRun(
		Paste("testdata/file1.txt", "testdata/file2.txt"),
	)
	// Output:
	//
}

