package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraImageCaptureControl struct {
	QMediaControl
}

type QCameraImageCaptureControl_ITF interface {
	QMediaControl_ITF
	QCameraImageCaptureControl_PTR() *QCameraImageCaptureControl
}

func PointerFromQCameraImageCaptureControl(ptr QCameraImageCaptureControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraImageCaptureControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraImageCaptureControlFromPointer(ptr unsafe.Pointer) *QCameraImageCaptureControl {
	var n = new(QCameraImageCaptureControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraImageCaptureControl_") {
		n.SetObjectName("QCameraImageCaptureControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraImageCaptureControl) QCameraImageCaptureControl_PTR() *QCameraImageCaptureControl {
	return ptr
}

func (ptr *QCameraImageCaptureControl) CancelCapture() {
	defer qt.Recovering("QCameraImageCaptureControl::cancelCapture")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_CancelCapture(ptr.Pointer())
	}
}

func (ptr *QCameraImageCaptureControl) Capture(fileName string) int {
	defer qt.Recovering("QCameraImageCaptureControl::capture")

	if ptr.Pointer() != nil {
		return int(C.QCameraImageCaptureControl_Capture(ptr.Pointer(), C.CString(fileName)))
	}
	return 0
}

func (ptr *QCameraImageCaptureControl) DriveMode() QCameraImageCapture__DriveMode {
	defer qt.Recovering("QCameraImageCaptureControl::driveMode")

	if ptr.Pointer() != nil {
		return QCameraImageCapture__DriveMode(C.QCameraImageCaptureControl_DriveMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageCaptureControl) ConnectError(f func(id int, error int, errorString string)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::error")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectError() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::error")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQCameraImageCaptureControlError
func callbackQCameraImageCaptureControlError(ptr unsafe.Pointer, ptrName *C.char, id C.int, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QCameraImageCaptureControl::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(int, int, string))(int(id), int(error), C.GoString(errorString))
	}

}

func (ptr *QCameraImageCaptureControl) Error(id int, error int, errorString string) {
	defer qt.Recovering("QCameraImageCaptureControl::error")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_Error(ptr.Pointer(), C.int(id), C.int(error), C.CString(errorString))
	}
}

func (ptr *QCameraImageCaptureControl) ConnectImageExposed(f func(requestId int)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::imageExposed")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectImageExposed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageExposed", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectImageExposed() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::imageExposed")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectImageExposed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageExposed")
	}
}

//export callbackQCameraImageCaptureControlImageExposed
func callbackQCameraImageCaptureControlImageExposed(ptr unsafe.Pointer, ptrName *C.char, requestId C.int) {
	defer qt.Recovering("callback QCameraImageCaptureControl::imageExposed")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageExposed"); signal != nil {
		signal.(func(int))(int(requestId))
	}

}

func (ptr *QCameraImageCaptureControl) ImageExposed(requestId int) {
	defer qt.Recovering("QCameraImageCaptureControl::imageExposed")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ImageExposed(ptr.Pointer(), C.int(requestId))
	}
}

func (ptr *QCameraImageCaptureControl) ConnectImageMetadataAvailable(f func(id int, key string, value *core.QVariant)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::imageMetadataAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectImageMetadataAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageMetadataAvailable", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectImageMetadataAvailable() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::imageMetadataAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectImageMetadataAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageMetadataAvailable")
	}
}

//export callbackQCameraImageCaptureControlImageMetadataAvailable
func callbackQCameraImageCaptureControlImageMetadataAvailable(ptr unsafe.Pointer, ptrName *C.char, id C.int, key *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCaptureControl::imageMetadataAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageMetadataAvailable"); signal != nil {
		signal.(func(int, string, *core.QVariant))(int(id), C.GoString(key), core.NewQVariantFromPointer(value))
	}

}

func (ptr *QCameraImageCaptureControl) ImageMetadataAvailable(id int, key string, value core.QVariant_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::imageMetadataAvailable")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ImageMetadataAvailable(ptr.Pointer(), C.int(id), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QCameraImageCaptureControl) ConnectImageSaved(f func(requestId int, fileName string)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::imageSaved")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectImageSaved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "imageSaved", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectImageSaved() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::imageSaved")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectImageSaved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "imageSaved")
	}
}

//export callbackQCameraImageCaptureControlImageSaved
func callbackQCameraImageCaptureControlImageSaved(ptr unsafe.Pointer, ptrName *C.char, requestId C.int, fileName *C.char) {
	defer qt.Recovering("callback QCameraImageCaptureControl::imageSaved")

	if signal := qt.GetSignal(C.GoString(ptrName), "imageSaved"); signal != nil {
		signal.(func(int, string))(int(requestId), C.GoString(fileName))
	}

}

func (ptr *QCameraImageCaptureControl) ImageSaved(requestId int, fileName string) {
	defer qt.Recovering("QCameraImageCaptureControl::imageSaved")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ImageSaved(ptr.Pointer(), C.int(requestId), C.CString(fileName))
	}
}

func (ptr *QCameraImageCaptureControl) IsReadyForCapture() bool {
	defer qt.Recovering("QCameraImageCaptureControl::isReadyForCapture")

	if ptr.Pointer() != nil {
		return C.QCameraImageCaptureControl_IsReadyForCapture(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraImageCaptureControl) ConnectReadyForCaptureChanged(f func(ready bool)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::readyForCaptureChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ConnectReadyForCaptureChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readyForCaptureChanged", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectReadyForCaptureChanged() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::readyForCaptureChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DisconnectReadyForCaptureChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readyForCaptureChanged")
	}
}

//export callbackQCameraImageCaptureControlReadyForCaptureChanged
func callbackQCameraImageCaptureControlReadyForCaptureChanged(ptr unsafe.Pointer, ptrName *C.char, ready C.int) {
	defer qt.Recovering("callback QCameraImageCaptureControl::readyForCaptureChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "readyForCaptureChanged"); signal != nil {
		signal.(func(bool))(int(ready) != 0)
	}

}

func (ptr *QCameraImageCaptureControl) ReadyForCaptureChanged(ready bool) {
	defer qt.Recovering("QCameraImageCaptureControl::readyForCaptureChanged")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ReadyForCaptureChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(ready)))
	}
}

func (ptr *QCameraImageCaptureControl) SetDriveMode(mode QCameraImageCapture__DriveMode) {
	defer qt.Recovering("QCameraImageCaptureControl::setDriveMode")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_SetDriveMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraImageCaptureControl) DestroyQCameraImageCaptureControl() {
	defer qt.Recovering("QCameraImageCaptureControl::~QCameraImageCaptureControl")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_DestroyQCameraImageCaptureControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraImageCaptureControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraImageCaptureControlTimerEvent
func callbackQCameraImageCaptureControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCaptureControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraImageCaptureControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraImageCaptureControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraImageCaptureControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraImageCaptureControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraImageCaptureControlChildEvent
func callbackQCameraImageCaptureControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCaptureControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraImageCaptureControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraImageCaptureControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraImageCaptureControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraImageCaptureControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraImageCaptureControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraImageCaptureControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraImageCaptureControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraImageCaptureControlCustomEvent
func callbackQCameraImageCaptureControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraImageCaptureControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraImageCaptureControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraImageCaptureControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraImageCaptureControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraImageCaptureControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraImageCaptureControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}