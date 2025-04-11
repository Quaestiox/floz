package floz

type MiddleWare struct {
	list []ReqHandler
}

func NewMW(mws ...ReqHandler) *MiddleWare {
	list := make([]ReqHandler, 0)
	for _, mw := range mws {
		list = append(list, mw)
	}
	return &MiddleWare{
		list: list,
	}
}

func (mw *MiddleWare) addMW(handlers ...ReqHandler) {
	mw.list = append(mw.list, handlers...)
}

func (mw *MiddleWare) mwNum() int {
	return len(mw.list)
}
