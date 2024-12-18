package util

var KeyCodes = map[string] uint16 {
	"A": 0x41, "B": 0x42, "C": 0x43, "D": 0x44, "E": 0x45, "F": 0x46, "G": 0x47,
	"H": 0x48, "I": 0x49, "J": 0x4A, "K": 0x4B, "L": 0x4C, "M": 0x4D, "N": 0x4E,
	"O": 0x4F, "P": 0x50, "Q": 0x51, "R": 0x52, "S": 0x53, "T": 0x54, "U": 0x55,
	"V": 0x56, "W": 0x57, "X": 0x58, "Y": 0x59, "Z": 0x5A,

	"1": 0x31, "2": 0x32, "3": 0x33, "4": 0x34, "5": 0x35, "6": 0x36, "7": 0x37,
	"8": 0x38, "9": 0x39, "0": 0x30,

	"Minus": 0xBD, "Equals": 0xBB, "Backspace": 0x08, "Tab": 0x09, "Semicolon": 0xBA,
	"Apostrophe": 0xDE, "Comma": 0xBC, "Period": 0xBE, "Slash": 0xBF, "Backslash": 0xDC,
	"LeftBrace": 0xDB, "RightBrace": 0xDD,

	"F1": 0x70, "F2": 0x71, "F3": 0x72, "F4": 0x73, "F5": 0x74, "F6": 0x75,
	"F7": 0x76, "F8": 0x77, "F9": 0x78, "F10": 0x79, "F11": 0x7A, "F12": 0x7B,

	"Esc": 0x1B, " [Enter]": 0x0D, "Space": 0x20, "LeftCtrl": 0xA2, "RightCtrl": 0xA3,
	"LeftShift": 0xA0, "RightShift": 0xA1, "LeftAlt": 0xA4, "RightAlt": 0xA5,
	"LeftWin": 0x5B, "RightWin": 0x5C, "ContextMenu": 0x5D,

	"ArrowUp": 0x26, "ArrowDown": 0x28, "ArrowLeft": 0x25, "ArrowRight": 0x27,

	"NumLock": 0x90, "Numpad0": 0x60, "Numpad1": 0x61, "Numpad2": 0x62, "Numpad3": 0x63,
	"Numpad4": 0x64, "Numpad5": 0x65, "Numpad6": 0x66, "Numpad7": 0x67, "Numpad8": 0x68,
	"Numpad9": 0x69, "NumpadMultiply": 0x6A, "NumpadAdd": 0x6B, "NumpadSubtract": 0x6D,
	"NumpadDecimal": 0x6E, "NumpadDivide": 0x6F, "NumPadEnter": 0x9C,

	"Insert": 0x2D, "Delete": 0x2E, "Home": 0x24, "End": 0x23, "PageUp": 0x21,
	"PageDown": 0x22,

	"CapsLock": 0x14, "ScrollLock": 0x91, "Pause": 0x13, "PrintScreen": 0x2C,
}

var SpecialKeys = map[string] uint16 {
	"Tab": 0x09,
	"Esc": 0x1B,
	"LeftCtrl": 0xA2,
	"RightCtrl": 0xA3,
	"LeftShift": 0xA0,
	"RightShift": 0xA1,
	"LeftAlt": 0xA4,
	"RightAlt": 0xA5,
	"LeftWin": 0x5B,
	"RightWin": 0x5C,
	"ContextMenu": 0x5D,
	"Insert": 0x2D,
	"Delete": 0x2E,
	"Home": 0x24,
	"End": 0x23,
	"PageUp": 0x21,
	"PageDown": 0x22,
	"CapsLock": 0x14,
}