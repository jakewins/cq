package types_test

import (
	"errors"

	_ "github.com/wfreeman/cq"
	"github.com/wfreeman/cq/types"
	. "launchpad.net/gocheck"
)

func (s *TypesSuite) TestQueryMapStringCypherValueParam(c *C) {
	stmt := prepareTest("with {0} as test return test")
	rows, err := stmt.Query(
		types.MapStringCypherValue{
			map[string]types.CypherValue{
				"key1": types.CypherValue{Val: "1", Type: types.CypherString},
				"key2": types.CypherValue{Val: 2, Type: types.CypherInt},
			}})
	c.Assert(err, IsNil)

	rows.Next()
	var test types.MapStringCypherValue
	err = rows.Scan(&test)
	c.Assert(err, IsNil)
	c.Assert(test, DeepEquals,
		types.MapStringCypherValue{
			map[string]types.CypherValue{
				"key1": types.CypherValue{Val: "1", Type: types.CypherString},
				"key2": types.CypherValue{Val: 2, Type: types.CypherInt},
			}})
}

func (s *TypesSuite) TestQueryStringCypherValueMapParam(c *C) {
	stmt := prepareTest("with {0} as test return test")
	rows, err := stmt.Query(map[string]types.CypherValue{
		"key1": types.CypherValue{Val: "1", Type: types.CypherString},
		"key2": types.CypherValue{Val: 2, Type: types.CypherInt},
	})
	c.Assert(err, IsNil)

	rows.Next()
	var test types.MapStringCypherValue
	err = rows.Scan(&test)
	c.Assert(err, IsNil)
	c.Assert(test, DeepEquals,
		types.MapStringCypherValue{
			map[string]types.CypherValue{
				"key1": types.CypherValue{Val: "1", Type: types.CypherString},
				"key2": types.CypherValue{Val: 2, Type: types.CypherInt},
			}})
}

func (s *TypesSuite) TestQueryMapStringCypherValue(c *C) {
	rows := prepareAndQuery(`return {key1:"1",key2:2}`)
	rows.Next()
	var test types.MapStringCypherValue
	err := rows.Scan(&test)
	c.Assert(err, IsNil)
	c.Assert(test, DeepEquals,
		types.MapStringCypherValue{
			map[string]types.CypherValue{
				"key1": types.CypherValue{Val: "1", Type: types.CypherString},
				"key2": types.CypherValue{Val: 2, Type: types.CypherInt},
			}})
}

func (s *TypesSuite) TestQueryBadMapStringCypherValue(c *C) {
	rows := prepareAndQuery(`return 1`)
	rows.Next()
	var test types.MapStringCypherValue
	err := rows.Scan(&test)
	c.Assert(err, DeepEquals, errors.New("sql: Scan error on column index 0: cq: invalid Scan value for *types.MapStringCypherValue: int"))
}

func (s *TypesSuite) TestQueryNullMapStringCypherValue(c *C) {
	rows := prepareAndQuery("return null")
	rows.Next()
	var test types.MapStringCypherValue
	err := rows.Scan(&test)
	c.Assert(err, DeepEquals, errors.New("sql: Scan error on column index 0: cq: scan value is null"))
}
