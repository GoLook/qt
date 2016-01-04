package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractUriResolver struct {
	core.QObject
}

type QAbstractUriResolver_ITF interface {
	core.QObject_ITF
	QAbstractUriResolver_PTR() *QAbstractUriResolver
}

func PointerFromQAbstractUriResolver(ptr QAbstractUriResolver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractUriResolver_PTR().Pointer()
	}
	return nil
}

func NewQAbstractUriResolverFromPointer(ptr unsafe.Pointer) *QAbstractUriResolver {
	var n = new(QAbstractUriResolver)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractUriResolver_") {
		n.SetObjectName("QAbstractUriResolver_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractUriResolver) QAbstractUriResolver_PTR() *QAbstractUriResolver {
	return ptr
}

func (ptr *QAbstractUriResolver) Resolve(relative core.QUrl_ITF, baseURI core.QUrl_ITF) *core.QUrl {
	defer qt.Recovering("QAbstractUriResolver::resolve")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QAbstractUriResolver_Resolve(ptr.Pointer(), core.PointerFromQUrl(relative), core.PointerFromQUrl(baseURI)))
	}
	return nil
}

func (ptr *QAbstractUriResolver) DestroyQAbstractUriResolver() {
	defer qt.Recovering("QAbstractUriResolver::~QAbstractUriResolver")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_DestroyQAbstractUriResolver(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractUriResolver) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractUriResolver::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractUriResolver::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractUriResolverTimerEvent
func callbackQAbstractUriResolverTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractUriResolver::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractUriResolver) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractUriResolver) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractUriResolver::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractUriResolver::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractUriResolverChildEvent
func callbackQAbstractUriResolverChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractUriResolver::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractUriResolver) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractUriResolver) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractUriResolver::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractUriResolver) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractUriResolver::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractUriResolverCustomEvent
func callbackQAbstractUriResolverCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractUriResolver::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractUriResolverFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractUriResolver) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractUriResolver) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractUriResolver::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractUriResolver_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}