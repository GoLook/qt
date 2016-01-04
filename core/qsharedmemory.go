package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSharedMemory struct {
	QObject
}

type QSharedMemory_ITF interface {
	QObject_ITF
	QSharedMemory_PTR() *QSharedMemory
}

func PointerFromQSharedMemory(ptr QSharedMemory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSharedMemory_PTR().Pointer()
	}
	return nil
}

func NewQSharedMemoryFromPointer(ptr unsafe.Pointer) *QSharedMemory {
	var n = new(QSharedMemory)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSharedMemory_") {
		n.SetObjectName("QSharedMemory_" + qt.Identifier())
	}
	return n
}

func (ptr *QSharedMemory) QSharedMemory_PTR() *QSharedMemory {
	return ptr
}

//QSharedMemory::AccessMode
type QSharedMemory__AccessMode int64

const (
	QSharedMemory__ReadOnly  = QSharedMemory__AccessMode(0)
	QSharedMemory__ReadWrite = QSharedMemory__AccessMode(1)
)

//QSharedMemory::SharedMemoryError
type QSharedMemory__SharedMemoryError int64

const (
	QSharedMemory__NoError          = QSharedMemory__SharedMemoryError(0)
	QSharedMemory__PermissionDenied = QSharedMemory__SharedMemoryError(1)
	QSharedMemory__InvalidSize      = QSharedMemory__SharedMemoryError(2)
	QSharedMemory__KeyError         = QSharedMemory__SharedMemoryError(3)
	QSharedMemory__AlreadyExists    = QSharedMemory__SharedMemoryError(4)
	QSharedMemory__NotFound         = QSharedMemory__SharedMemoryError(5)
	QSharedMemory__LockError        = QSharedMemory__SharedMemoryError(6)
	QSharedMemory__OutOfResources   = QSharedMemory__SharedMemoryError(7)
	QSharedMemory__UnknownError     = QSharedMemory__SharedMemoryError(8)
)

func NewQSharedMemory2(parent QObject_ITF) *QSharedMemory {
	defer qt.Recovering("QSharedMemory::QSharedMemory")

	return NewQSharedMemoryFromPointer(C.QSharedMemory_NewQSharedMemory2(PointerFromQObject(parent)))
}

func NewQSharedMemory(key string, parent QObject_ITF) *QSharedMemory {
	defer qt.Recovering("QSharedMemory::QSharedMemory")

	return NewQSharedMemoryFromPointer(C.QSharedMemory_NewQSharedMemory(C.CString(key), PointerFromQObject(parent)))
}

func (ptr *QSharedMemory) Attach(mode QSharedMemory__AccessMode) bool {
	defer qt.Recovering("QSharedMemory::attach")

	if ptr.Pointer() != nil {
		return C.QSharedMemory_Attach(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QSharedMemory) ConstData() unsafe.Pointer {
	defer qt.Recovering("QSharedMemory::constData")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSharedMemory_ConstData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSharedMemory) Create(size int, mode QSharedMemory__AccessMode) bool {
	defer qt.Recovering("QSharedMemory::create")

	if ptr.Pointer() != nil {
		return C.QSharedMemory_Create(ptr.Pointer(), C.int(size), C.int(mode)) != 0
	}
	return false
}

func (ptr *QSharedMemory) Data() unsafe.Pointer {
	defer qt.Recovering("QSharedMemory::data")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSharedMemory_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSharedMemory) Data2() unsafe.Pointer {
	defer qt.Recovering("QSharedMemory::data")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSharedMemory_Data2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSharedMemory) Detach() bool {
	defer qt.Recovering("QSharedMemory::detach")

	if ptr.Pointer() != nil {
		return C.QSharedMemory_Detach(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSharedMemory) Error() QSharedMemory__SharedMemoryError {
	defer qt.Recovering("QSharedMemory::error")

	if ptr.Pointer() != nil {
		return QSharedMemory__SharedMemoryError(C.QSharedMemory_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSharedMemory) ErrorString() string {
	defer qt.Recovering("QSharedMemory::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSharedMemory_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSharedMemory) IsAttached() bool {
	defer qt.Recovering("QSharedMemory::isAttached")

	if ptr.Pointer() != nil {
		return C.QSharedMemory_IsAttached(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSharedMemory) Key() string {
	defer qt.Recovering("QSharedMemory::key")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSharedMemory_Key(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSharedMemory) Lock() bool {
	defer qt.Recovering("QSharedMemory::lock")

	if ptr.Pointer() != nil {
		return C.QSharedMemory_Lock(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSharedMemory) NativeKey() string {
	defer qt.Recovering("QSharedMemory::nativeKey")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSharedMemory_NativeKey(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSharedMemory) SetKey(key string) {
	defer qt.Recovering("QSharedMemory::setKey")

	if ptr.Pointer() != nil {
		C.QSharedMemory_SetKey(ptr.Pointer(), C.CString(key))
	}
}

func (ptr *QSharedMemory) SetNativeKey(key string) {
	defer qt.Recovering("QSharedMemory::setNativeKey")

	if ptr.Pointer() != nil {
		C.QSharedMemory_SetNativeKey(ptr.Pointer(), C.CString(key))
	}
}

func (ptr *QSharedMemory) Size() int {
	defer qt.Recovering("QSharedMemory::size")

	if ptr.Pointer() != nil {
		return int(C.QSharedMemory_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSharedMemory) Unlock() bool {
	defer qt.Recovering("QSharedMemory::unlock")

	if ptr.Pointer() != nil {
		return C.QSharedMemory_Unlock(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSharedMemory) DestroyQSharedMemory() {
	defer qt.Recovering("QSharedMemory::~QSharedMemory")

	if ptr.Pointer() != nil {
		C.QSharedMemory_DestroyQSharedMemory(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSharedMemory) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QSharedMemory::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSharedMemory) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSharedMemory::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSharedMemoryTimerEvent
func callbackQSharedMemoryTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSharedMemory::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQSharedMemoryFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSharedMemory) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QSharedMemory::timerEvent")

	if ptr.Pointer() != nil {
		C.QSharedMemory_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QSharedMemory) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QSharedMemory::timerEvent")

	if ptr.Pointer() != nil {
		C.QSharedMemory_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QSharedMemory) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QSharedMemory::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSharedMemory) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSharedMemory::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSharedMemoryChildEvent
func callbackQSharedMemoryChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSharedMemory::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQSharedMemoryFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QSharedMemory) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QSharedMemory::childEvent")

	if ptr.Pointer() != nil {
		C.QSharedMemory_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QSharedMemory) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QSharedMemory::childEvent")

	if ptr.Pointer() != nil {
		C.QSharedMemory_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QSharedMemory) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QSharedMemory::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSharedMemory) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSharedMemory::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSharedMemoryCustomEvent
func callbackQSharedMemoryCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSharedMemory::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQSharedMemoryFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QSharedMemory) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QSharedMemory::customEvent")

	if ptr.Pointer() != nil {
		C.QSharedMemory_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QSharedMemory) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QSharedMemory::customEvent")

	if ptr.Pointer() != nil {
		C.QSharedMemory_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}