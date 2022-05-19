package zenity

const (
	_CW_USEDEFAULT = 0x80000000

	_WS_TABSTOP      = 0x00010000
	_WS_GROUP        = 0x00020000
	_WS_SYSMENU      = 0x00080000
	_WS_VSCROLL      = 0x00200000
	_WS_DLGFRAME     = 0x00400000
	_WS_BORDER       = 0x00800000
	_WS_CLIPSIBLINGS = 0x04000000
	_WS_DISABLED     = 0x08000000
	_WS_VISIBLE      = 0x10000000
	_WS_CHILD        = 0x40000000
	_WS_POPUP        = 0x80000000
	_WS_POPUPWINDOW  = _WS_POPUP | _WS_BORDER | _WS_SYSMENU

	_WS_EX_DLGMODALFRAME = 0x00000001
	_WS_EX_WINDOWEDGE    = 0x00000100
	_WS_EX_CLIENTEDGE    = 0x00000200
	_WS_EX_CONTROLPARENT = 0x00010000

	_BS_DEFPUSHBUTTON = 0x0001
	_ES_AUTOHSCROLL   = 0x0080
	_ES_PASSWORD      = 0x0020
	_SS_NOPREFIX      = 0x0080
	_SS_EDITCONTROL   = 0x2000
	_SS_WORDELLIPSIS  = 0xc000
	_LBS_EXTENDEDSEL  = 0x0800
	_MCS_NOTODAY      = 0x0010
	_PBS_SMOOTH       = 0x0001
	_PBS_MARQUEE      = 0x0008

	_SW_NORMAL = 1

	_SWP_NOMOVE   = 0x0002
	_SWP_NOZORDER = 0x0004

	_MB_OKCANCEL        = 0x00000001
	_MB_YESNOCANCEL     = 0x00000003
	_MB_YESNO           = 0x00000004
	_MB_ICONERROR       = 0x00000010
	_MB_ICONQUESTION    = 0x00000020
	_MB_ICONWARNING     = 0x00000030
	_MB_ICONINFORMATION = 0x00000040
	_MB_DEFBUTTON1      = 0x00000000
	_MB_DEFBUTTON2      = 0x00000100
	_MB_DEFBUTTON3      = 0x00000200

	_IDOK     = 1
	_IDCANCEL = 2
	_IDYES    = 6
	_IDNO     = 7

	_SC_CLOSE = 0xf060

	_WM_DESTROY     = 0x0002
	_WM_CLOSE       = 0x0010
	_WM_SETFONT     = 0x0030
	_WM_SETICON     = 0x0080
	_WM_NCCREATE    = 0x0081
	_WM_NCDESTROY   = 0x0082
	_WM_COMMAND     = 0x0111
	_WM_SYSCOMMAND  = 0x0112
	_WM_DPICHANGED  = 0x02e0
	_WM_USER        = 0x0400
	_EM_SETSEL      = 0x00b1
	_LB_ADDSTRING   = 0x0180
	_LB_GETCURSEL   = 0x0188
	_LB_GETSELCOUNT = 0x0190
	_LB_GETSELITEMS = 0x0191
	_MCM_GETCURSEL  = 0x1001
	_MCM_SETCURSEL  = 0x1002
	_PBM_SETPOS     = _WM_USER + 2
	_PBM_SETRANGE32 = _WM_USER + 6
	_PBM_SETMARQUEE = _WM_USER + 10
	_STM_SETICON    = 0x0170

	_GWL_STYLE = -16

	_PROGRESS_CLASS = "msctls_progress32"
	_MONTHCAL_CLASS = "SysMonthCal32"
)
