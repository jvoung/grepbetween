Overview
--------

This is a commandline tool that helps *extract* or *filter out* lines of
text between a START marker and an END marker.
The START and END markers are specified as regexes.

Example
-------

`let START = "=+ ABCD =+"`

`let END = "=+ END =+"`

and input
```
Some stuff
More stuff ==== ABCD ====
==== ABCD ====
==== ABCD ====
12345
==== END ====
Yet more stuff
```

Output is
```
Some stuff
More stuff ==== ABCD ====
Yet more stuff
```

Or you can invert the match and get as output

```
==== ABCD ====
12345
```

Future Work
-----------
* extend this to handle multiple start and end pairs
* have options to include the start/end in output

Corner Cases
------------
It's currently undefined what happens if the start and end pairs are nested. Most likely, it will just stop matching after finding the first "end" that matches the first "start".