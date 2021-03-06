//line machine.rl:1
package parser

import (
	"fmt"
	"strings"
)

//line machine.go:7
const phone_start int = 1
const phone_first_final int = 20
const phone_error int = 0

const phone_en_main int = 1


//line machine.rl:71


// Machine contains variables used by Ragel auto generated code.
// See Ragel docs for details.
type Machine struct {
	// Data to process.
	data []byte

	// Data pointer.
	p int

	// Data end pointer.
	pe int

	// Current state.
	cs int

	// End of file pointer.
	eof int

	// Start of current date block.
	pb int

	// Current error.
	err error
}


// Phone represents parser result and will be filled by Ragel actions.
type Phone struct {
	IntCode  string
	AreaCode string
	Number   string
}


// New initialise new Machine structure.
func New() *Machine {
	return &Machine{}
}

// string returns current parsed variable. Variable m.pb
// should be updated by "mark" action, while m.p is a current
// parser position inside m.data.
func (m *Machine) string() string {
	return string(m.data[m.pb:m.p])
}

// Parse takes a slice of bytes as an input and fills Phone struct
// in case input is a phone number in one of valid formats.
func (m *Machine) Parse(input []byte) (*Phone, error) {
	m.data = input
	m.pe = len(input)
	m.p = 0
	m.pb = 0
	m.err = nil
	m.eof = len(input)

	phone := &Phone{}


//line machine.go:77
	{
	 m.cs = phone_start
	}

//line machine.rl:132

//line machine.go:84
	{
	if ( m.p) == ( m.pe) {
		goto _test_eof
	}
	switch  m.cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 20:
		goto st_case_20
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	}
	goto st_out
	st_case_1:
		switch ( m.data)[( m.p)] {
		case 32:
			goto st2
		case 40:
			goto st3
		case 43:
			goto st18
		case 45:
			goto st3
		case 49:
			goto tr4
		}
		if 50 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr5
		}
		goto tr0
tr0:
//line machine.rl:22

		m.err = fmt.Errorf("invalid area code, expected 200...999")

	goto st0
tr8:
//line machine.rl:33

		m.err = fmt.Errorf("invalid phone format: %s", m.data)

	goto st0
//line machine.go:163
st_case_0:
	st0:
		 m.cs = 0
		goto _out
tr24:
//line machine.rl:12

		phone.IntCode = m.string()

	goto st2
	st2:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof2
		}
	st_case_2:
//line machine.go:179
		switch ( m.data)[( m.p)] {
		case 32:
			goto st2
		case 40:
			goto st3
		case 45:
			goto st3
		}
		if 50 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr5
		}
		goto tr0
tr25:
//line machine.rl:12

		phone.IntCode = m.string()

	goto st3
	st3:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof3
		}
	st_case_3:
//line machine.go:203
		if 50 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr5
		}
		goto tr0
tr5:
//line machine.rl:9
 m.pb = m.p
	goto st4
tr26:
//line machine.rl:12

		phone.IntCode = m.string()

//line machine.rl:9
 m.pb = m.p
	goto st4
	st4:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof4
		}
	st_case_4:
//line machine.go:225
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st5
		}
		goto tr0
	st5:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof5
		}
	st_case_5:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st6
		}
		goto tr0
	st6:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr9
		case 41:
			goto tr9
		case 45:
			goto tr9
		}
		if 50 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr10
		}
		goto tr8
tr9:
//line machine.rl:17

		phone.AreaCode = m.string()

	goto st7
	st7:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof7
		}
	st_case_7:
//line machine.go:267
		if ( m.data)[( m.p)] == 32 {
			goto st7
		}
		if 50 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr12
		}
		goto tr8
tr12:
//line machine.rl:9
 m.pb = m.p
	goto st8
tr10:
//line machine.rl:17

		phone.AreaCode = m.string()

//line machine.rl:9
 m.pb = m.p
	goto st8
	st8:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof8
		}
	st_case_8:
//line machine.go:292
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st9
		}
		goto tr8
	st9:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof9
		}
	st_case_9:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st10
		}
		goto tr8
	st10:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch ( m.data)[( m.p)] {
		case 32:
			goto st11
		case 45:
			goto st17
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st12
		}
		goto tr8
	st11:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof11
		}
	st_case_11:
		if ( m.data)[( m.p)] == 32 {
			goto st11
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st12
		}
		goto tr8
	st12:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof12
		}
	st_case_12:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st13
		}
		goto tr8
	st13:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch ( m.data)[( m.p)] {
		case 32:
			goto st14
		case 45:
			goto st16
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st15
		}
		goto tr8
	st14:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof14
		}
	st_case_14:
		if ( m.data)[( m.p)] == 32 {
			goto st14
		}
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st15
		}
		goto tr8
	st15:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof15
		}
	st_case_15:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr22
		}
		goto tr8
tr22:
//line machine.rl:38

		if phone.IntCode == "" {
			phone.IntCode = "1"
		}

	goto st20
	st20:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof20
		}
	st_case_20:
//line machine.go:391
		goto st0
	st16:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof16
		}
	st_case_16:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st15
		}
		goto tr8
	st17:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof17
		}
	st_case_17:
		if 48 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto st12
		}
		goto tr8
	st18:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof18
		}
	st_case_18:
		if ( m.data)[( m.p)] == 49 {
			goto tr4
		}
		goto st0
tr4:
//line machine.rl:9
 m.pb = m.p
	goto st19
	st19:
		if ( m.p)++; ( m.p) == ( m.pe) {
			goto _test_eof19
		}
	st_case_19:
//line machine.go:429
		switch ( m.data)[( m.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 45:
			goto tr25
		}
		if 50 <= ( m.data)[( m.p)] && ( m.data)[( m.p)] <= 57 {
			goto tr26
		}
		goto tr0
	st_out:
	_test_eof2:  m.cs = 2; goto _test_eof
	_test_eof3:  m.cs = 3; goto _test_eof
	_test_eof4:  m.cs = 4; goto _test_eof
	_test_eof5:  m.cs = 5; goto _test_eof
	_test_eof6:  m.cs = 6; goto _test_eof
	_test_eof7:  m.cs = 7; goto _test_eof
	_test_eof8:  m.cs = 8; goto _test_eof
	_test_eof9:  m.cs = 9; goto _test_eof
	_test_eof10:  m.cs = 10; goto _test_eof
	_test_eof11:  m.cs = 11; goto _test_eof
	_test_eof12:  m.cs = 12; goto _test_eof
	_test_eof13:  m.cs = 13; goto _test_eof
	_test_eof14:  m.cs = 14; goto _test_eof
	_test_eof15:  m.cs = 15; goto _test_eof
	_test_eof20:  m.cs = 20; goto _test_eof
	_test_eof16:  m.cs = 16; goto _test_eof
	_test_eof17:  m.cs = 17; goto _test_eof
	_test_eof18:  m.cs = 18; goto _test_eof
	_test_eof19:  m.cs = 19; goto _test_eof

	_test_eof: {}
	if ( m.p) == ( m.eof) {
		switch  m.cs {
		case 1, 2, 3, 4, 5, 19:
//line machine.rl:22

		m.err = fmt.Errorf("invalid area code, expected 200...999")

		case 20:
//line machine.rl:27

		n := strings.ReplaceAll(m.string(), "-", "")
		phone.Number = strings.ReplaceAll(n, " ", "")

		case 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
//line machine.rl:33

		m.err = fmt.Errorf("invalid phone format: %s", m.data)

//line machine.go:482
		}
	}

	_out: {}
	}

//line machine.rl:133

	return phone, m.err
}
