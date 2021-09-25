package main

%% machine matcher;
%% write data;

func FindEmails(data []byte) []string {
	cs, p, pe, eof := 0, 0, len(data), len(data)
	emails := make([]string, 0)
	mark := 0
	_ = eof

	%%{
		action mark { mark = p }
		action add { emails = append(emails, string(data[mark:p])) }

		email = space* (alnum+ >mark ('.' | '+' | '_' | '-')? alnum+ '@' alnum+ . '.' . alnum+) %add . space+ >/add ;
		main := email+;


		write init;
		write exec;
	}%%

	return emails
}
