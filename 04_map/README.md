Map VS Struct

Map
* All keys must be the same type.
* All values must be the same type.
* Keys are indexed - we can iterate over them.
* [v] Use to represent a collection of related properties.
* [v] Don't need to know all the keys at compile time.
* Reference Type! (Object in some external memory space, no pointers required to change these)


Struct
* Values can be of different type.
* Keys don't support indexing.
* You need to all the different fields at compile time.
* Use to represent a "thing" with a lot of different properties.
* Value Type! (Use pointers to change these in function)