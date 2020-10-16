# gostringlist

A class which implements convenience functions for working with lists of strings released under the MIT license. 

## Description

Go is a great language, but sometimes it's more than a little inconvenient. This module makes working with lists of strings a little easier--probably not at all with Go's level of Zen, but whatever. It's heavily influenced by Python's string management methods, so you'll probably find something helpful here.

## Usage

Create a StringList object instance with `myList := gostringlist.New()` and off you go. Helpful methods listed below:

### Access

- `Items` - the `[]string` slice used for internal storage. You'll be accessing this from time to time.
- `Copy()` - create a duplicate of the object
- `ToString` - converts the items into a pretty-printed string which is helpful for debugging

### Management

- `Insert` - inserts a string at the specified index. Good for one-off insertions, but performance isn't great if used a lot
- `Remove` - deletes the item at the specified index.
- `RemoveUnordered` - A faster item deletion method if you don't care about items' order
- `Join` - convenient access and syntactic sugar, no more, no less 
- `Filter` - takes a filtering function and creates a new StringList object based on the filter. If you miss Python's list comprehensions, you'll find this useful.
- `MatchFilter` - returns a new StringList containing all the items which match the supplied regular expression. Keep in mind that Go's regexes are their own special flavor.
- `ReplaceAllFilter` - performs a regular-expression search and replace and returns a new StringList object containing the results.

### Comparison

- `IsEqual`, `IsEmpty` - helpful comparison methods
- `IndexOf` - get the index of a string value or -1 if it doesn't exist
- `Contains` - returns true if the string passed to it is in the list

