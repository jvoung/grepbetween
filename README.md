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

Future Work
-----------
* extend this to handle multiple start and end pairs
* have options to include the start/end in output

Corner Cases
------------
It's currently undefined what happens if the start and end pairs are nested. Most likely, it will just stop matching after finding the first "end" that matches the first "start".