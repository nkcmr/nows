# nows

a whitespace eliminating pipe program

## use cases

i mainly use this to make sure passwords i generate and encrypt do NOT contain any whitespace characters at all:

```bash
cat my_password.txt | nows | gpg
```

because i am never going to remember how to do this properly with `sed`

## installation

with `go` installed, do this:

```bash
go install -u -v github.com/nkcmr/nows
```

## usage

no arguments, but it MUST have data piped in, and then it pumps out the same input minus any whitespace characters:

```
password_generator | nows | <a program that can be sure not to recieve whitespace characters>
```

## license

```
MIT License

Copyright (c) 2020 Nicholas Comer

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
```

