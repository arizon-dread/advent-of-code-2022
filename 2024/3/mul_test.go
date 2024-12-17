package main

import (
	"reflect"
	"testing"
)

func Test_sumFoundMuls(t *testing.T) {
	type args struct {
		line string
	}
	test1_arg := args{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1 works", args: test1_arg, want: 161},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumFoundMuls(tt.args.line); got != tt.want {
				t.Errorf("sumFoundMuls() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDonts(t *testing.T) {
	type args struct {
		input []byte
	}
	test1_arg := args{[]byte(`do()xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
do()don't()asdmul(222,444)do()mul(2,3)`)}
	test1_want := []byte(`do()xmul(2,4)&mul[3,7]!^?mul(8,5))
do()mul(2,3)`)
	test2_arg := args{[]byte(`@why(692,996)[&}}^where(81,407)mul(247,89):[&[{<mul(980,958),?mul(529,895)!<#~!$&~when()+mul(519,986)what())#mul(710,934)%??*'!<mul(813,338)! +$what()<don't(){^mul(396,693)mul(337,541)}what()*<](@?~mul(64,644)[where()who()~,))mul(528,450)!' -do()who()#]where():(mul(909,368)mul(259,743)''when()^?from()]select()#;mul(227,252)<mul(118,202)-&!(when(806,911)~]who(58,451)- mul(135,37);mul(75,773)?~when()where()]mul(93,321)where(),where()!when(769,449)where(616,323)@&mul(489,237)&;do()<mul(803,622)mul(616,264)!from()why()~@],@from()mul(499,593)#<?/&(when()':mul(237,54)&-],+mul(856,447)})select()mul(540,283)don't()'@how()@mul(701,900)?!['mul(958,898)mul(403,891)+*why()&-)mul(877,695)where()&}{<when()]:[(mul(70,638)<who()[mul(219,485)[why()+% +what()*who()>do())[when()![mul(681,521)`)}
	test2_want := []byte(`@why(692,996)[&}}^where(81,407)mul(247,89):[&[{<mul(980,958),?mul(529,895)!<#~!$&~when()+mul(519,986)what())#mul(710,934)%??*'!<mul(813,338)! +$what()<who()#]where():(mul(909,368)mul(259,743)''when()^?from()]select()#;mul(227,252)<mul(118,202)-&!(when(806,911)~]who(58,451)- mul(135,37);mul(75,773)?~when()where()]mul(93,321)where(),where()!when(769,449)where(616,323)@&mul(489,237)&;do()<mul(803,622)mul(616,264)!from()why()~@],@from()mul(499,593)#<?/&(when()':mul(237,54)&-],+mul(856,447)})select()mul(540,283))[when()![mul(681,521)`)
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"test1", test1_arg, test1_want},
		{"test2", test2_arg, test2_want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDonts(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDonts() = %v,\n\n want %v", string(got), string(tt.want))
			}
		})
	}
}
