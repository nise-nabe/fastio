package fastio

type InOut struct {
	in []byte
	*bufio.Writer
}

func NewInOut(r io.Reader, w io.Writer) *InOut {
	in, _ := ioutil.ReadAll(r)
	return &InOut{in, bufio.NewWriter(w)}
}

func (s *InOut) next() (r int) {
	buf := s.in
	p := 1
	for (buf[0] < '0' || '9' < buf[0]) && buf[0] != '-' {
		buf = buf[1:]
	}
	if buf[0] == '-' {
		p = -1
		buf = buf[1:]

	}
	for '0' <= buf[0] && buf[0] <= '9' {
		r = 10*r + int(buf[0]-'0')
		buf = buf[1:]
	}
	r *= p
	s.in = buf
	return
}

func (s *InOut) nextLine() (r string) {
	buf := s.in
	for buf[0] == '\n' {
		buf = buf[1:]
	}
	p := 0
	for buf[p] != '\n' {
		p++
	}
	r = string(buf[0:p])
	s.in = buf[p:]
	return
}

func (s *InOut) nextStr() (r string) {
	buf := s.in
	for buf[0] == '\n' || buf[0] == ' ' {
		buf = buf[1:]
	}
	p := 0
	for buf[p] != '\n' && buf[p] != ' ' {
		p++
	}
	r = string(buf[0:p])
	s.in = buf[p:]
	return
}

func (s *InOut) print(os ...interface{}) {
	for _, o := range os {
		switch o.(type) {
		case byte:
			s.WriteByte(o.(byte))
		case string:
			s.WriteString(o.(string))
		case int:
			s.WriteString(strconv.Itoa(o.(int)))
		case int64:
			s.WriteString(strconv.FormatInt(o.(int64), 10))
		default:
			s.WriteString(fmt.Sprint(o))
		}
	}
}

func (s *InOut) println(os ...interface{}) {
	for _, o := range os {
		s.print(o)
	}
	s.print("\n")
}

func (s *InOut) printlnNow(o interface{}) {
	fmt.Println(o)
}
