package parse

type ParseContext struct {
	BitBuf_reader *BitBuf_reader
	BitBuf_writer *BitBuf_writer
	ParentContext *ParseContext
	//指针类型
	Instance any
}

func ToParseContext(instance any, parentContext *ParseContext) *ParseContext {
	return &ParseContext{
		BitBuf_reader: nil,
		BitBuf_writer: nil,
		ParentContext: parentContext,
		Instance:      instance,
	}
}

func GetBitBuf_reader(byteBuf *ByteBuf, parentParseContext *ParseContext) *BitBuf_reader {
	if parentParseContext == nil || parentParseContext.BitBuf_reader == nil {
		return ToBitBuf_reader(byteBuf)
	} else {
		return parentParseContext.BitBuf_reader
	}
}

func GetBitBuf_writer(byteBuf *ByteBuf, parentParseContext *ParseContext) *BitBuf_writer {
	if parentParseContext == nil || parentParseContext.BitBuf_writer == nil {
		return ToBitBuf_writer(byteBuf)
	} else {
		return parentParseContext.BitBuf_writer
	}
}
