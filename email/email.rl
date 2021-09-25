package main

%% machine matcher;
%% write data;

func FindEmails(data []byte) []string {
	cs, p, pe, eof := 0, 0, len(data), len(data)
	emails := make([]string, 0)
	mark := 0
	_ = eof

	%%{
		action mark { mark = fpc }
		action captureEmail { emails = append(emails, string(data[mark:p])) }

		head = alnum | '_' | '.' | '+' | '-';
		org = alnum | '-';
		com = alnum | '-' | '.';

		emailFormat = head+ >mark '@' org+ '.' com+ %captureEmail;
		main := emailFormat;


		write init;
		write exec;
	}%%

	return emails
}
