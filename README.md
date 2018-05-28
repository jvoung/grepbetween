Overview
--------

This is a commandline tool that helps print text between a START marker and an END marker. It can also invert the printing, to print lines that are not between START/END. The START and END markers are regexes.

Example
-------

```
let START = "=+ ABCD =+"
let END = "=+ END =+"
```

and input
```
Some stuff
More stuff
==== ABCD ====
12345
==== END ====
Yet more stuff
```

Output is
```
12345
```

Or you can invert the match and get as output

```
Some stuff
More stuff
Yet more stuff
```

One can also specify multiple start and end regexes. In that case, any pair of start/end can match. In the future, this may become more specific and require specific start/end pairs to match.

Corner Cases
------------
It's currently undefined what happens if the start and end pairs are nested. Most likely, it will just stop matching after finding the first "end" that matches the first "start".