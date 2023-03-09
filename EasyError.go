package EasyError

import "fmt"

type HandContinue struct {
	Error error
}

type Hand func()

func HandError(err error) *HandContinue {
	if err != nil {
		return &HandContinue{Error: err}
	}
	return &HandContinue{Error: nil}
}

func (h *HandContinue) ThenGo(hand Hand) *HandContinue {
	if h.Error != nil {
		hand()
		return h
	}
	return h
}

func (h *HandContinue) PrintError(log ...any) *HandContinue {
	if h.Error != nil {

		var logstr string
		for i := 0; i < len(log); i++ {
			logstr = fmt.Sprint(logstr, log[i])
		}
		fmt.Printf("%s%s", logstr, h.Error)
		return h

	}
	return h
}

func (h *HandContinue) PrintlnError(log ...any) *HandContinue {
	if h.Error != nil {
		var logstr string
		for i := 0; i < len(log); i++ {
			logstr = fmt.Sprint(logstr, log[i])
		}
		fmt.Printf("%s%s\n", logstr, h.Error)
		return h
	}
	return h
}

func (h *HandContinue) Panic() *HandContinue {
	if h.Error != nil {
		panic(h.Error)
	}
	return h
}

func (h *HandContinue) Finaly(head Hand) {
	head()
}
