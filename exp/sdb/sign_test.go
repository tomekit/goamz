package sdb_test

import (
	. "github.com/motain/gocheck"
	"github.com/tomekit/goamz/aws"
	"github.com/tomekit/goamz/exp/sdb"
)

// SimpleDB ReST authentication docs: http://goo.gl/CaY81

var testAuth = aws.Auth{"access-key-id-s8eBOWuU", "secret-access-key-UkQjTLd9", ""}

func (s *S) TestSignExampleDomainCreate(c *C) {
	method := "GET"
	params := map[string][]string{
		"Action":     {"CreateDomain"},
		"DomainName": {"MyDomain"},
		"Timestamp":  {"2011-08-20T07:23:57+12:00"},
		"Version":    {"2009-04-15"},
	}
	headers := map[string][]string{
		"Host": {"sdb.amazonaws.com"},
	}
	sdb.Sign(testAuth, method, "", params, headers)
	expected := "ot2JaeeqMRJqgAqW67hkzUlffgxdOz4RykbrECB+tDU="
	c.Assert(params["Signature"], DeepEquals, []string{expected})
}

// Do a few test methods which takes combinations of params
