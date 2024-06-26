package mturk_test

import (
	. "github.com/motain/gocheck"
	"github.com/tomekit/goamz/aws"
	"github.com/tomekit/goamz/exp/mturk"
)

// Mechanical Turk REST authentication docs: http://goo.gl/wrzfn

var testAuth = aws.Auth{"user", "secret", ""}

// == fIJy9wCApBNL2R4J2WjJGtIBFX4=
func (s *S) TestBasicSignature(c *C) {
	params := map[string]string{}
	mturk.Sign(testAuth, "AWSMechanicalTurkRequester", "CreateHIT", "2012-02-16T20:30:47Z", params)
	expected := "b/TnvzrdeD/L/EyzdFrznPXhido="
	c.Assert(params["Signature"], Equals, expected)
}
