
godoctest
=========

`godoctest` spares you some drudgery by generating boilerplate around your
table-driven tests.

Inside a function, immediately following the declaration, write a comment of the
form `@test = <your test table>`. Each row of the table should consist of N + M
elements, where N is the number of parameters and M is the number of return
values. `godoctest` derives the types for the test table from the function
declaration and generates all the usual boilerplate for a test function. See the
example below.

Here's your source code:

``` go
func fibonacci(n int) int {
	/*
		@test = {
			{1, 1},
			{2, 1},
			{3, 2},
			{7, 13},
			{11, 89},
		}
	*/
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
```

Here's the test generated by `godoctest` (slightly redacted for readability):

``` go
func Test_fibonacci_gdt1(t *testing.T) {
    tests := []struct{
        f0 int
        e0 int
    }{
        {1, 1},
        {2, 1},
        {3, 2},
        {7, 13},
        {11, 89},
    }
    for _, test := range tests {
        r0 := fibonacci(test.f0)
        assert.Equal(t, test.e0, r0)
    }
}
```