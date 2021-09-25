
//line email.rl:1
package main


//line email.rl:4

//line email.go:9
const matcher_start int = 1
const matcher_first_final int = 10
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
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 10:
		goto st_case_10
	case 9:
		goto st_case_9
	}
	goto st_out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 32 {
			goto st1
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st1
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr2
				}
			case data[p] >= 65:
				goto tr2
			}
		default:
			goto tr2
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr2:
//line email.rl:13
 mark = p 
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line email.go:100
		switch data[p] {
		case 43:
			goto st3
		case 95:
			goto st3
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st3
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st9
				}
			case data[p] >= 65:
				goto st9
			}
		default:
			goto st9
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
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
		if data[p] == 64 {
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
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
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
		if data[p] == 46 {
			goto st7
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st6
			}
		default:
			goto st6
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st8
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st8
			}
		default:
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 32 {
			goto tr10
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr10
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st8
				}
			case data[p] >= 65:
				goto st8
			}
		default:
			goto st8
		}
		goto st0
tr10:
//line email.rl:14
 emails = append(emails, string(data[mark:p])) 
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line email.go:256
		if data[p] == 32 {
			goto st10
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st10
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr2
				}
			case data[p] >= 65:
				goto tr2
			}
		default:
			goto tr2
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 43:
			goto st3
		case 64:
			goto st5
		case 95:
			goto st3
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st3
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st9
				}
			case data[p] >= 65:
				goto st9
			}
		default:
			goto st9
		}
		goto st0
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 8:
//line email.rl:14
 emails = append(emails, string(data[mark:p])) 
//line email.go:327
		}
	}

	_out: {}
	}

//line email.rl:22


	return emails
}
