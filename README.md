
# goc

Wrapper to simulate symlinks with executables on Windows platform by calling a batch script.

## Example

Create a batch file `dir.cmd`.

```batch
@echo off
"c:\windows\system32\dir.exe" %*
```

Compile **goc** and rename `goc.exe` to `dir.exe`. Save `dir.exe` alongside `dir.cmd`.

Then you can remove `c:\windows\system32` from the `PATH` and use the new directory where is your `dir.exe` file instead.

When you call `dir.exe`, it spawns `dir.cmd` by passing arguments and forwarding stdout and stderr pipes.

## Why not just use `dir.cmd`?

Because you can't spawn `dir` (without suffix) if you have only a `dir.cmd` or `dir.bat` file in the `PATH`. It's necessary to spawn `dir.cmd` explicitly and it can be a problem when you don't have the source code. Then with **goc**, it's possible to spawn with just `dir`.

Note that **goc** is just a very trivial and limited wrapper. For example, stdin is not forwarded.

## License

The MIT License (MIT)

Copyright (c) 2016 Xcraft

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
