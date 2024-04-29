package zenity

import (
	"context"
	"image/color"
	"reflect"
	"testing"
	"time"

	"github.com/dcam0050/zenity/internal/zencmd"
)

func Test_applyOptions(t *testing.T) {
	t.Parallel()
	date := time.Date(2006, 1, 1, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name string
		args Option
		want options
	}{
		// General options
		{name: "Title", args: Title("Title"), want: options{title: ptr("Title")}},
		{name: "Width", args: Width(100), want: options{width: 100}},
		{name: "Height", args: Height(100), want: options{height: 100}},
		{name: "OKLabel", args: OKLabel("OK"), want: options{okLabel: ptr("OK")}},
		{name: "CancelLabel", args: CancelLabel("Cancel"), want: options{cancelLabel: ptr("Cancel")}},
		{name: "ExtraButton", args: ExtraButton("Extra"), want: options{extraButton: ptr("Extra")}},
		{name: "DefaultCancel", args: DefaultCancel(), want: options{defaultCancel: true}},
		{name: "WindowIcon", args: WindowIcon(ErrorIcon), want: options{windowIcon: ErrorIcon}},
		{name: "WindowIcon", args: WindowIcon("error"), want: options{windowIcon: "error"}},
		{name: "Icon", args: Icon(ErrorIcon), want: options{icon: ErrorIcon}},
		{name: "Icon", args: Icon("error"), want: options{icon: "error"}},
		{name: "Modal", args: Modal(), want: options{modal: true}},
		{name: "Attach", args: Attach(zencmd.ParseWindowId("12345")), want: options{attach: zencmd.ParseWindowId("12345")}},
		{name: "Display", args: Display(":1"), want: options{display: ":1"}},
		{name: "ClassHint", args: ClassHint("Name", "Class"), want: options{name: "Name", class: "Class"}},

		// Message options
		{name: "NoWrap", args: NoWrap(), want: options{noWrap: true}},
		{name: "Ellipsize", args: Ellipsize(), want: options{ellipsize: true}},

		// Entry options
		{name: "EntryText", args: EntryText("text"), want: options{entryText: "text"}},
		{name: "HideText", args: HideText(), want: options{hideText: true}},
		{name: "Username", args: Username(), want: options{username: true}},

		// List options
		{name: "CheckList", args: CheckList(), want: options{listKind: checkListKind}},
		{name: "RadioList", args: RadioList(), want: options{listKind: radioListKind}},
		{name: "MidSearch", args: MidSearch(), want: options{midSearch: true}},
		{name: "DisallowEmpty", args: DisallowEmpty(), want: options{disallowEmpty: true}},
		{name: "DefaultItems", args: DefaultItems("a", "b"), want: options{defaultItems: []string{"a", "b"}}},

		// Calendar options
		{name: "DefaultDate", args: DefaultDate(2006, time.January, 1), want: options{time: &date}},

		// File selection options
		{name: "Directory", args: Directory(), want: options{directory: true}},
		{name: "ConfirmOverwrite", args: ConfirmOverwrite(), want: options{confirmOverwrite: true}},
		{name: "ConfirmCreate", args: ConfirmCreate(), want: options{confirmCreate: true}},
		{name: "ShowHidden", args: ShowHidden(), want: options{showHidden: true}},
		{name: "Filename", args: Filename("file.go"), want: options{filename: "file.go"}},
		{name: "FileFilter", args: FileFilter{"All files", []string{"*.*"}, true}, want: options{
			fileFilters: FileFilters{{"All files", []string{"*.*"}, true}},
		}},
		{name: "FileFilters", args: FileFilters{{"Go files", []string{"*.go"}, true}}, want: options{
			fileFilters: FileFilters{{"Go files", []string{"*.go"}, true}},
		}},

		// Color selection options
		{name: "Color", args: Color(color.Black), want: options{color: color.Black}},
		{name: "ShowPalette", args: ShowPalette(), want: options{showPalette: true}},

		// Progress indication options
		{name: "MaxValue", args: MaxValue(100), want: options{maxValue: 100}},
		{name: "Pulsate", args: Pulsate(), want: options{maxValue: -1}},
		{name: "NoCancel", args: NoCancel(), want: options{noCancel: true}},
		{name: "TimeRemaining", args: TimeRemaining(), want: options{timeRemaining: true}},

		// Context for timeout
		{name: "Context", args: Context(context.TODO()), want: options{ctx: context.TODO()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyOptions([]Option{tt.args}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyOptions() = %v; want %v", got, tt.want)
			}
		})
	}
}
