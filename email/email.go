
//line email.rl:1
package main


//line email.rl:4

//line email.go:9
const matcher_start int = 1
const matcher_first_final int = 6
const matcher_error int = 0

const matcher_en_main int = 1


//line email.rl:5

func FindEmails(data []byte) []string {
	cs, p, pe, eof := 0, 0, len(data), len(data)
	emails := make([]string, 0)
	mark := 0
	_ = eof

	
//line email.go:26
	{
	cs = matcher_start
	}

//line email.go:31
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
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
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 43:
			goto tr0
		case 95:
			goto tr0
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr0
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr0
				}
			case data[p] >= 65:
				goto tr0
			}
		default:
			goto tr0
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr0:
//line email.rl:13
 mark = p 
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line email.go:91
		switch data[p] {
		case 43:
			goto st2
		case 64:
			goto st3
		case 95:
			goto st2
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st2
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st2
				}
			case data[p] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 45 {
			goto st4
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st4
			}
		default:
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 45:
			goto st4
		case 46:
			goto st5
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st4
			}
		default:
			goto st4
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st6
				}
			case data[p] >= 65:
				goto st6
			}
		default:
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st6
				}
			case data[p] >= 65:
				goto st6
			}
		default:
			goto st6
		}
		goto st0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 6:
//line email.rl:14
 emails = append(emails, string(data[mark:p])) 
//line email.go:222
		}
	}

	_out: {}
	}

//line email.rl:26


	return emails
}
