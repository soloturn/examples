some examples on how the compile and execution
speed of different programming languages behave
for primitive 1..1000 loop. as well it allows
to compare the size of the executable created
with default paramters.

```
time cc c-loop.c -o c-loop
time ./c-loop

time go build go-loop.go
time ./go-loop

time python py-loop.py

time rustc rs-loop.rs
time ./rs-loop

ls -la
```

it shows that c and rust have the same speed,
go takes 30% and more time, python is 10 times
slower.

while a c executable is 15KB, the rust executable
is 8MB, go 2MB.
