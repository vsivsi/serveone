# serveone

Super simple web server that serves a single file on port 8088 (or one you choose with -port).

The file then is available on the path of the base filename.

Build with golang 1.6:

```
go build serveone.go
```

Use:
```
% serveone -h

Usage:
 serveone [-port val] filename
  -port uint
    	Port to bind (default 8088)
```

Example:

```
% serveone -port 5678 LICENSE &
[1] 85431
Serving file LICENSE on port 5678

% curl localhost:5678/LICENSE
Copyright (C) 2016 by Vaughn Iverson

serveone is free software released under the MIT/X11 license:

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

% kill 85431
[1]+  Terminated: 15          serveone -port 5678 LICENSE
```
