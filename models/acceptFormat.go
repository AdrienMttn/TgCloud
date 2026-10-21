package models

type AcceptFormat struct {
	Supports []string `json:"supports"`
}

func NewAcceptFormat() AcceptFormat{
	return AcceptFormat{Supports : []string{"image/jpeg", "image/png","image/jpg","image/gif", "image/webp"}}
}

func (AcceptFormatList *AcceptFormat) TestIsSupport(format string) bool{
	for _, f := range AcceptFormatList.Supports {
		if (f == format){
			return true
		}
	}
	return false
}