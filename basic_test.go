package paste_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/paste"
)

func ExamplePaste_basic() {
	// paste file1.txt file2.txt
	// Note: This requires actual files to paste together
	gloo.MustRun(
		Paste("testdata/file1.txt", "testdata/file2.txt"),
	)
	// Output:
	//
}
