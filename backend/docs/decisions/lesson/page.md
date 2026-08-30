# Element
#### Problem
A structure of a page might be complex: containing nesting, different types of elements, various layouts, etc. Providing support for elements with dynamic styles and content is hard, since we must predefine them all.
#### Solution
In order to provide such a flexibility and avoid implementation overhead, the server side abstracts information about page altogether using 2 data types:
1. **Element** - an abstract page element, bearing *type* and *config* determining how element will be rendered; also has a reference to the parent element if one exists;
2. **Tree** - a tree composed of elements according to their parent elements references.
In other words, backend knows nothing except the type of an element and what element is the parent of the former one. This removes the responsibility to handle and validate different elements, their settings, content, etc.

