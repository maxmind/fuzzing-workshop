package tomlfuzz

// In this exercise, we are going to fuzz github.com/BurntSushi/toml, looking
// for input that will crash it. You should first set up your environment:
//
// 1. Open your favorite terminal
// 2. Run `export GO111MODULES=off` to disable Go modules if you have them
//    enabled.
// 3. Install go-fuzz by running
//    `go get github.com/dvyukov/go-fuzz/go-fuzz && go get github.com/dvyukov/go-fuzz/go-fuzz-build`
// 4. Install the TOML library with `go get github.com/BurntSushi/toml`
//
// You are now read to fill in the missing pieces. Please complete all of the
// TODO items below using the "Complex Usage" example code from
// https://github.com/BurntSushi/toml#more-complex-usage

// TODO: Add the types for the structs that you will decode the data to.
//
// type tomlConfig struct {...

// Fuzz is the function that `go-fuzz` will call when fuzzing. It should run
// the code you are interested in using the data passed in as a []byte. The
// return value should be 1 if the fuzzer should increase the priority of the
// input and 0 otherwise. For our purposes, you should return 0 if the parser
// returned an error cleanly and 1 if there was no error, indicating that it
// was parsed successfully.
func Fuzz(data []byte) int {
	// TODO: Add a tomlConfig variable to decode the data into.

	// TODO: Decode `data` into the variable as shown in the example. You will
	// need to convert `data` to a string, e.g., `string(data)` when doing this.

	// TODO: Return 0 if there is an error returned and 1 otherwise.
}

// Now that you have filled in all of the TODO items, you are almost ready to
// run `go-fuzz`.
//
// Before doing this, you first need to build the binaries need by it by running
// `go-fuzz-build .` in the current directory. After doing this, you should have
// a `tomlfuzz-fuzz.zip` file.
//
// We now need to seed the corpus so that `go-fuzz` has a starting point when
// generating input data. To do this, copy the example TOML document from
// https://github.com/BurntSushi/toml#more-complex-usage to a file in the `corpus`
// subdirectory.
//
// You are now ready to run `go-fuzz`! Do this by running
// `go-fuzz -bin=tomlfuzz-fuzz.zip`.
//
// You should now see lines like the following in your terminal:
//
// 2019/04/03 11:58:49 workers: 8, corpus: 1 (2s ago), crashers: 0, restarts: 1/2861, execs: 37194 (3757/sec), cover: 2402, uptime: 9s
//
// Among other things, this tells you the current corpus size and the number of
// crashers (e.g., panics) found.
//
// As `go-fuzz` runs, it will be constantly mutating the corpus and saving
// those mutations that cause it to reach new parts of the code.

// When a crasher is found, it will be placed in the `crasher` subdirectory
// along with the output from the crash. If you wanted to do extensive testing,
// you might let `go-fuzz` run for hours or days, but for our example, let it
// run for a few minutes.

// Did you find any crashers? Based on the results, would you be comfortable
// using this library to parse user-inputted TOML?
