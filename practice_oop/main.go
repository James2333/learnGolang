package practice_oop

type User_Info struct {
	Name string
	Age int
	Height int
	School string
	Hobby []string
	MoreInfo interface{}
}
func NewUser_Info(name string,age int,height int ,school string,hobby []string,moreinfo interface{}) *User_Info {

	return &User_Info{
		Name:     name,
		Age:      age,
		Height:   height,
		School:   school,
		Hobby:    hobby,
		MoreInfo: moreinfo,
	}
	
}

