package classfile

type parsingContext struct {
	content []byte
	err     error
	offset  int
}

type parseFunc func([]byte, int) (int, error)

func (content *parsingContext) parse(parser parseFunc) {
	if content.err == nil {
		content.offset, content.err = parser(content.content, content.offset)

		if content.err != nil {
			logger.Errorf("error encountered while parsing %v", content.err)
		}
	}
}
