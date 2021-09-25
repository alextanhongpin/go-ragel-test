package parser

%%{
	machine phone;
	# main := any;


	# Sets the start point for the current value.
	action mark { m.pb = m.p }

	# Writes international code.
	action set_intcode {
		phone.IntCode = m.string()
	}

	# Writes area code.
	action set_area {
		phone.AreaCode = m.string()
	}

	# Return area-related error.
	action err_area {
		m.err = fmt.Errorf("invalid area code, expected 200...999")
	}

	# Writes number and replaces all hypens and spaces.
	action set_number {
		n := strings.ReplaceAll(m.string(), "-", "")
		phone.Number = strings.ReplaceAll(n, " ", "")
	}

	# Returns an error when phone with incorrect format.
	action err_phone {
		m.err = fmt.Errorf("invalid phone format: %s", m.data)
	}

	# Sets default international code.
	action set_defaults {
		if phone.IntCode == "" {
			phone.IntCode = "1"
		}
	}

	sp = ' ';
	hy = (sp* | '-');
	hbl = ( hy | '(' );
	hbr = ( hy | ')' );

	# Defines an international code: +1, 1
	int = '+' ? '1' > mark %set_intcode;

	# Defines a range 200..999
	nxx = ('2'..'9' . digit{2});

	# Defines an area code with separators: (NXX),-NXX-
	area = hbl ? nxx >mark %set_area @err(err_area) hbr?;

	# Defines phone number in range 2000000..9999999, with an optional separators.
	number = (nxx hy? digit{2} hy? digit{2}) >mark %set_number @err(err_phone);

	# Defines main machine for phone.
	main := int? sp* area sp* number @set_defaults;

	write data;

	access m.;
	variable p m.p;
	variable pe m.pe;
	variable eof m.eof;
	variable data m.data;
}%%

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

	%% write init;
	%% write exec;

	return phone, m.err
}
