package golog

import . "fmt"
import . "regexp"
import . "strings"
import "bytes"

type Term interface {
    // Functor returns the term's name
    Functor() string

    // Arity returns the number of arguments a term has. An atom has 0 arity.
    Arity() int

    // Arguments returns a slice of this term's arguments, if any
    Arguments() []Term

    // String provides a string representation of a term
    String() string

    // Indicator() provides a "predicate indicator" representation of a term
    Indicator() string
}

type Structure struct {
    Func    string
    Args    []Term
}
func (self *Structure) Functor() string {
    return self.Func
}
func (self *Structure) Arity() int {
    return len(self.Args)
}
func (self *Structure) Arguments() []Term {
    return self.Args
}
func (self *Structure) String() string {
    // an atom
    quotedFunctor := QuoteFunctor(self.Functor())
    if self.Arity() == 0 {
        return quotedFunctor
    }

    var buf bytes.Buffer
    Fprintf(&buf, "%s(", quotedFunctor)
    arity := self.Arity()
    for i := 0; i<arity; i++ {
        if i>0 {
            Fprintf(&buf, ", ")
        }
        Fprintf(&buf, "%s", self.Arguments()[i])
    }
    Fprintf(&buf, ")")
    return buf.String()
}
func (self *Structure) Indicator() string {
    return Sprintf("%s/%d", self.Functor(), self.Arity())
}


// NewTerm creates a new term with the given functor and optional arguments
func NewTerm(functor string, arguments ...Term) Term {
    return &Structure{
        Func:   functor,
        Args:   arguments,
    }
}


// QuoteFunctor returns a canonical representation of a term's name
// by quoting characters that require quoting
func QuoteFunctor(name string) string {
    needsQuote, err := MatchString(`\W`, name)
    maybePanic(err)
    if needsQuote {
        escapedName := Replace(name, `'`, `\'`, -1)
        return Sprintf("'%s'", escapedName)
    }

    return name
}

func maybePanic(err error) {
    if err != nil {
        panic(err)
    }
}
