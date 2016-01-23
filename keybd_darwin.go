package keybd_event
/*
 #cgo CFLAGS: -x objective-c
 #cgo LDFLAGS: -framework Cocoa
 #import <Cocoa/Cocoa.h>
 CGEventRef Create(int k){
	CGEventRef event = CGEventCreateKeyboardEvent (NULL, (CGKeyCode)k, true);
	return event;
 }
 void KeyTap(CGEventRef event){
	CGEventPost(kCGAnnotatedSessionEventTap, event);
	CFRelease(event);
 }
 void AddActionKey(CGEventFlags type,CGEventRef event){
 	CGEventSetFlags(event, type);
 }
*/
import "C"
import (

	"time"
)
const (
	_AShift = C.kCGEventFlagMaskAlphaShift
	_VK_SHIFT = C.kCGEventFlagMaskShift
	_VK_CTRL = C.kCGEventFlagMaskControl
	_VK_ALT = C.kCGEventFlagMaskAlternate
	_VK_CMD = C.kCGEventFlagMaskCommand
	_Help = C.kCGEventFlagMaskHelp
	_VK_FN = C.kCGEventFlagMaskSecondaryFn
	_NumPad = C.kCGEventFlagMaskNumericPad
	_Coalesced = C.kCGEventFlagMaskNonCoalesced
	_VK_Control = 0x3B
	_VK_RightShift = 0x3C
	_VK_RightControl = 0x3E
	_VK_Command = 0x37
	_VK_Shift = 0x38
)
func initKeyBD() error { return nil}
// Launch key bounding
func (k *KeyBounding) Launching() error {

	for _, key := range keys {
		tapKey(key)
	}
}
func shift(event uintptr){
	C.AddActionKey(VK_SHIF,event)
}
func ctrl(event uintptr){
	C.AddActionKey(VK_CTRL,event)
}
func alt(event uintptr){
	C.AddActionKey(VK_ALT,event)
}
func cmd(event uintptr){
	C.AddActionKey(VK_CMD,event)
}
func tapKey(key int) {
	event := C.Create(int)
	if k.hasALT{
		C.AddActionKey(_VK_ALT,event)
	}
	if k.hasCTRL{
		C.AddActionKey(_VK_CTRL,event)
	}
	if k.hasSHIFT{
		C.AddActionKey(_VK_SHIFT,event)
	}
	if k.hasRCTRL{
		C.AddActionKey(_VK_CTRL,event)
	}
	if k.hasRSHIFT {
		C.AddActionKey(_VK_SHIFT,event)
	}
	C.KeyTap(event)
}
const (
	VK_A = 0x00
	VK_S = 0x01
	VK_D = 0x02
	VK_F = 0x03
	VK_H = 0x04
	VK_G = 0x05
	VK_Z = 0x06
	VK_X = 0x07
	VK_C = 0x08
	VK_V = 0x09
	VK_B = 0x0B
	VK_Q = 0x0C
	VK_W = 0x0D
	VK_E = 0x0E
	VK_R = 0x0F
	VK_Y = 0x10
	VK_T = 0x11
	VK_1 = 0x12
	VK_2 = 0x13
	VK_3 = 0x14
	VK_4 = 0x15
	VK_6 = 0x16
	VK_5 = 0x17
	VK_Equal = 0x18
	VK_9 = 0x19
	VK_7 = 0x1A
	VK_Minus = 0x1B
	VK_8 = 0x1C
	VK_0 = 0x1D
	VK_RightBracket = 0x1E
	VK_O = 0x1F
	VK_U = 0x20
	VK_LeftBracket = 0x21
	VK_I = 0x22
	VK_P = 0x23
	VK_L = 0x25
	VK_J = 0x26
	VK_Quote = 0x27
	VK_K = 0x28
	VK_Semicolon = 0x29
	VK_Backslash = 0x2A
	VK_Comma = 0x2B
	VK_Slash = 0x2C
	VK_N = 0x2D
	VK_M = 0x2E
	VK_Period = 0x2F
	VK_Grave = 0x32
	VK_KeypadDecimal = 0x41
	VK_KeypadMultiply = 0x43
	VK_KeypadPlus = 0x45
	VK_KeypadClear = 0x47
	VK_KeypadDivide = 0x4B
	VK_KeypadEnter = 0x4C
	VK_KeypadMinus = 0x4E
	VK_KeypadEquals = 0x51
	VK_Keypad0 = 0x52
	VK_Keypad1 = 0x53
	VK_Keypad2 = 0x54
	VK_Keypad3 = 0x55
	VK_Keypad4 = 0x56
	VK_Keypad5 = 0x57
	VK_Keypad6 = 0x58
	VK_Keypad7 = 0x59
	VK_Keypad8 = 0x5B
	VK_Keypad9 = 0x5C

	VK_Return = 0x24
	VK_Tab = 0x30
	VK_Space = 0x31
	VK_Delete = 0x33
	VK_Escape = 0x35
	VK_CapsLock = 0x39
	VK_Option = 0x3A
	VK_RightOption = 0x3D
	VK_Function = 0x3F
	VK_F17 = 0x40
	VK_VolumeUp = 0x48
	VK_VolumeDown = 0x49
	VK_Mute = 0x4A
	VK_F18 = 0x4F
	VK_F19 = 0x50
	VK_F20 = 0x5A
	VK_F5 = 0x60
	VK_F6 = 0x61
	VK_F7 = 0x62
	VK_F3 = 0x63
	VK_F8 = 0x64
	VK_F9 = 0x65
	VK_F11 = 0x67
	VK_F13 = 0x69
	VK_F16 = 0x6A
	VK_F14 = 0x6B
	VK_F10 = 0x6D
	VK_F12 = 0x6F
	VK_F15 = 0x71
	VK_Help = 0x72
	VK_Home = 0x73
	VK_PageUp = 0x74
	VK_ForwardDelete = 0x75
	VK_F4 = 0x76
	VK_End = 0x77
	VK_F2 = 0x78
	VK_PageDown = 0x79
	VK_F1 = 0x7A
	VK_LeftArrow = 0x7B
	VK_RightArrow = 0x7C
	VK_DownArrow = 0x7D
	VK_UpArrow = 0x7E
)