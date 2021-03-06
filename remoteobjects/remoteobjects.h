// +build !minimal

#pragma once

#ifndef GO_QTREMOTEOBJECTS_H
#define GO_QTREMOTEOBJECTS_H

#include <stdint.h>

#ifdef __cplusplus
int QRemoteObjectAbstractPersistedStore_QRemoteObjectAbstractPersistedStore_QRegisterMetaType();
int QRemoteObjectDynamicReplica_QRemoteObjectDynamicReplica_QRegisterMetaType();
int QRemoteObjectHost_QRemoteObjectHost_QRegisterMetaType();
int QRemoteObjectHostBase_QRemoteObjectHostBase_QRegisterMetaType();
int QRemoteObjectNode_QRemoteObjectNode_QRegisterMetaType();
int QRemoteObjectRegistry_QRemoteObjectRegistry_QRegisterMetaType();
int QRemoteObjectRegistryHost_QRemoteObjectRegistryHost_QRegisterMetaType();
int QRemoteObjectReplica_QRemoteObjectReplica_QRegisterMetaType();
extern "C" {
#endif

struct QtRemoteObjects_PackedString { char* data; long long len; };
struct QtRemoteObjects_PackedList { void* data; long long len; };
void* QRemoteObjectAbstractPersistedStore_NewQRemoteObjectAbstractPersistedStore(void* parent);
struct QtRemoteObjects_PackedString QRemoteObjectAbstractPersistedStore_QRemoteObjectAbstractPersistedStore_Tr(char* s, char* c, int n);
struct QtRemoteObjects_PackedString QRemoteObjectAbstractPersistedStore_QRemoteObjectAbstractPersistedStore_TrUtf8(char* s, char* c, int n);
struct QtRemoteObjects_PackedList QRemoteObjectAbstractPersistedStore_RestoreProperties(void* ptr, struct QtRemoteObjects_PackedString repName, void* repSig);
void QRemoteObjectAbstractPersistedStore_SaveProperties(void* ptr, struct QtRemoteObjects_PackedString repName, void* repSig, void* values);
void QRemoteObjectAbstractPersistedStore_DestroyQRemoteObjectAbstractPersistedStore(void* ptr);
void QRemoteObjectAbstractPersistedStore_DestroyQRemoteObjectAbstractPersistedStoreDefault(void* ptr);
void* QRemoteObjectAbstractPersistedStore_MetaObjectDefault(void* ptr);
void* QRemoteObjectAbstractPersistedStore___restoreProperties_atList(void* ptr, int i);
void QRemoteObjectAbstractPersistedStore___restoreProperties_setList(void* ptr, void* i);
void* QRemoteObjectAbstractPersistedStore___restoreProperties_newList(void* ptr);
void* QRemoteObjectAbstractPersistedStore___saveProperties_values_atList(void* ptr, int i);
void QRemoteObjectAbstractPersistedStore___saveProperties_values_setList(void* ptr, void* i);
void* QRemoteObjectAbstractPersistedStore___saveProperties_values_newList(void* ptr);
void* QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_atList(void* ptr, int i);
void QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_setList(void* ptr, void* i);
void* QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_newList(void* ptr);
void* QRemoteObjectAbstractPersistedStore___findChildren_atList2(void* ptr, int i);
void QRemoteObjectAbstractPersistedStore___findChildren_setList2(void* ptr, void* i);
void* QRemoteObjectAbstractPersistedStore___findChildren_newList2(void* ptr);
void* QRemoteObjectAbstractPersistedStore___findChildren_atList3(void* ptr, int i);
void QRemoteObjectAbstractPersistedStore___findChildren_setList3(void* ptr, void* i);
void* QRemoteObjectAbstractPersistedStore___findChildren_newList3(void* ptr);
void* QRemoteObjectAbstractPersistedStore___findChildren_atList(void* ptr, int i);
void QRemoteObjectAbstractPersistedStore___findChildren_setList(void* ptr, void* i);
void* QRemoteObjectAbstractPersistedStore___findChildren_newList(void* ptr);
void* QRemoteObjectAbstractPersistedStore___children_atList(void* ptr, int i);
void QRemoteObjectAbstractPersistedStore___children_setList(void* ptr, void* i);
void* QRemoteObjectAbstractPersistedStore___children_newList(void* ptr);
char QRemoteObjectAbstractPersistedStore_EventDefault(void* ptr, void* e);
char QRemoteObjectAbstractPersistedStore_EventFilterDefault(void* ptr, void* watched, void* event);
void QRemoteObjectAbstractPersistedStore_ChildEventDefault(void* ptr, void* event);
void QRemoteObjectAbstractPersistedStore_ConnectNotifyDefault(void* ptr, void* sign);
void QRemoteObjectAbstractPersistedStore_CustomEventDefault(void* ptr, void* event);
void QRemoteObjectAbstractPersistedStore_DeleteLaterDefault(void* ptr);
void QRemoteObjectAbstractPersistedStore_DisconnectNotifyDefault(void* ptr, void* sign);
void QRemoteObjectAbstractPersistedStore_TimerEventDefault(void* ptr, void* event);
void QRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplica(void* ptr);
void QRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplicaDefault(void* ptr);
void* QRemoteObjectHost_NewQRemoteObjectHost(void* parent);
void* QRemoteObjectHost_NewQRemoteObjectHost3(void* address, void* parent);
void* QRemoteObjectHost_NewQRemoteObjectHost2(void* address, void* registryAddress, void* parent);
char QRemoteObjectHost_SetHostUrl(void* ptr, void* hostAddress);
char QRemoteObjectHost_SetHostUrlDefault(void* ptr, void* hostAddress);
void QRemoteObjectHost_DestroyQRemoteObjectHost(void* ptr);
void QRemoteObjectHost_DestroyQRemoteObjectHostDefault(void* ptr);
void* QRemoteObjectHost_HostUrl(void* ptr);
void* QRemoteObjectHost_HostUrlDefault(void* ptr);
char QRemoteObjectHostBase_DisableRemoting(void* ptr, void* remoteObject);
char QRemoteObjectHostBase_EnableRemoting3(void* ptr, void* model, struct QtRemoteObjects_PackedString name, void* roles, void* selectionModel);
char QRemoteObjectHostBase_EnableRemoting2(void* ptr, void* object, struct QtRemoteObjects_PackedString name);
int QRemoteObjectHostBase___enableRemoting_roles_atList3(void* ptr, int i);
void QRemoteObjectHostBase___enableRemoting_roles_setList3(void* ptr, int i);
void* QRemoteObjectHostBase___enableRemoting_roles_newList3(void* ptr);
void* QRemoteObjectNode_AcquireDynamic(void* ptr, struct QtRemoteObjects_PackedString name);
void* QRemoteObjectNode_NewQRemoteObjectNode(void* parent);
void* QRemoteObjectNode_NewQRemoteObjectNode2(void* registryAddress, void* parent);
struct QtRemoteObjects_PackedString QRemoteObjectNode_QRemoteObjectNode_Tr(char* s, char* c, int n);
struct QtRemoteObjects_PackedString QRemoteObjectNode_QRemoteObjectNode_TrUtf8(char* s, char* c, int n);
char QRemoteObjectNode_ConnectToNode(void* ptr, void* address);
char QRemoteObjectNode_SetRegistryUrl(void* ptr, void* registryAddress);
char QRemoteObjectNode_SetRegistryUrlDefault(void* ptr, void* registryAddress);
char QRemoteObjectNode_WaitForRegistry(void* ptr, int timeout);
void QRemoteObjectNode_ConnectError(void* ptr);
void QRemoteObjectNode_DisconnectError(void* ptr);
void QRemoteObjectNode_Error(void* ptr, long long errorCode);
void QRemoteObjectNode_ConnectHeartbeatIntervalChanged(void* ptr);
void QRemoteObjectNode_DisconnectHeartbeatIntervalChanged(void* ptr);
void QRemoteObjectNode_HeartbeatIntervalChanged(void* ptr, int heartbeatInterval);
void QRemoteObjectNode_SetHeartbeatInterval(void* ptr, int interval);
void QRemoteObjectNode_SetName(void* ptr, struct QtRemoteObjects_PackedString name);
void QRemoteObjectNode_SetNameDefault(void* ptr, struct QtRemoteObjects_PackedString name);
void QRemoteObjectNode_SetPersistedStore(void* ptr, void* persistedStore);
void QRemoteObjectNode_TimerEventDefault(void* ptr, void* vqt);
void QRemoteObjectNode_DestroyQRemoteObjectNode(void* ptr);
void QRemoteObjectNode_DestroyQRemoteObjectNodeDefault(void* ptr);
void* QRemoteObjectNode_PersistedStore(void* ptr);
long long QRemoteObjectNode_LastError(void* ptr);
struct QtRemoteObjects_PackedString QRemoteObjectNode_Instances2(void* ptr, struct QtRemoteObjects_PackedString typeName);
void* QRemoteObjectNode_RegistryUrl(void* ptr);
void* QRemoteObjectNode_MetaObjectDefault(void* ptr);
void* QRemoteObjectNode_Registry(void* ptr);
int QRemoteObjectNode_HeartbeatInterval(void* ptr);
int QRemoteObjectNode___acquireModel_rolesHint_atList(void* ptr, int i);
void QRemoteObjectNode___acquireModel_rolesHint_setList(void* ptr, int i);
void* QRemoteObjectNode___acquireModel_rolesHint_newList(void* ptr);
void* QRemoteObjectNode___retrieveProperties_atList(void* ptr, int i);
void QRemoteObjectNode___retrieveProperties_setList(void* ptr, void* i);
void* QRemoteObjectNode___retrieveProperties_newList(void* ptr);
void* QRemoteObjectNode___persistProperties_props_atList(void* ptr, int i);
void QRemoteObjectNode___persistProperties_props_setList(void* ptr, void* i);
void* QRemoteObjectNode___persistProperties_props_newList(void* ptr);
void* QRemoteObjectNode___dynamicPropertyNames_atList(void* ptr, int i);
void QRemoteObjectNode___dynamicPropertyNames_setList(void* ptr, void* i);
void* QRemoteObjectNode___dynamicPropertyNames_newList(void* ptr);
void* QRemoteObjectNode___findChildren_atList2(void* ptr, int i);
void QRemoteObjectNode___findChildren_setList2(void* ptr, void* i);
void* QRemoteObjectNode___findChildren_newList2(void* ptr);
void* QRemoteObjectNode___findChildren_atList3(void* ptr, int i);
void QRemoteObjectNode___findChildren_setList3(void* ptr, void* i);
void* QRemoteObjectNode___findChildren_newList3(void* ptr);
void* QRemoteObjectNode___findChildren_atList(void* ptr, int i);
void QRemoteObjectNode___findChildren_setList(void* ptr, void* i);
void* QRemoteObjectNode___findChildren_newList(void* ptr);
void* QRemoteObjectNode___children_atList(void* ptr, int i);
void QRemoteObjectNode___children_setList(void* ptr, void* i);
void* QRemoteObjectNode___children_newList(void* ptr);
char QRemoteObjectNode_EventDefault(void* ptr, void* e);
char QRemoteObjectNode_EventFilterDefault(void* ptr, void* watched, void* event);
void QRemoteObjectNode_ChildEventDefault(void* ptr, void* event);
void QRemoteObjectNode_ConnectNotifyDefault(void* ptr, void* sign);
void QRemoteObjectNode_CustomEventDefault(void* ptr, void* event);
void QRemoteObjectNode_DeleteLaterDefault(void* ptr);
void QRemoteObjectNode_DisconnectNotifyDefault(void* ptr, void* sign);
void QRemoteObjectRegistry_QRemoteObjectRegistry_RegisterMetatypes();
void QRemoteObjectRegistry_DestroyQRemoteObjectRegistry(void* ptr);
void QRemoteObjectRegistry_DestroyQRemoteObjectRegistryDefault(void* ptr);
void* QRemoteObjectRegistryHost_NewQRemoteObjectRegistryHost(void* registryAddress, void* parent);
void QRemoteObjectRegistryHost_DestroyQRemoteObjectRegistryHost(void* ptr);
void QRemoteObjectRegistryHost_DestroyQRemoteObjectRegistryHostDefault(void* ptr);
struct QtRemoteObjects_PackedString QRemoteObjectReplica_QRemoteObjectReplica_Tr(char* s, char* c, int n);
struct QtRemoteObjects_PackedString QRemoteObjectReplica_QRemoteObjectReplica_TrUtf8(char* s, char* c, int n);
char QRemoteObjectReplica_WaitForSource(void* ptr, int timeout);
void QRemoteObjectReplica_ConnectInitialized(void* ptr);
void QRemoteObjectReplica_DisconnectInitialized(void* ptr);
void QRemoteObjectReplica_Initialized(void* ptr);
void QRemoteObjectReplica_SetNode(void* ptr, void* node);
void QRemoteObjectReplica_SetNodeDefault(void* ptr, void* node);
void QRemoteObjectReplica_ConnectStateChanged(void* ptr);
void QRemoteObjectReplica_DisconnectStateChanged(void* ptr);
void QRemoteObjectReplica_StateChanged(void* ptr, long long state, long long oldState);
void* QRemoteObjectReplica_Node(void* ptr);
long long QRemoteObjectReplica_State(void* ptr);
char QRemoteObjectReplica_IsInitialized(void* ptr);
char QRemoteObjectReplica_IsReplicaValid(void* ptr);
void* QRemoteObjectReplica_MetaObjectDefault(void* ptr);
void* QRemoteObjectReplica___sendWithReply_args_atList(void* ptr, int i);
void QRemoteObjectReplica___sendWithReply_args_setList(void* ptr, void* i);
void* QRemoteObjectReplica___sendWithReply_args_newList(void* ptr);
void* QRemoteObjectReplica___send_args_atList(void* ptr, int i);
void QRemoteObjectReplica___send_args_setList(void* ptr, void* i);
void* QRemoteObjectReplica___send_args_newList(void* ptr);
void* QRemoteObjectReplica___setProperties_properties_atList(void* ptr, int i);
void QRemoteObjectReplica___setProperties_properties_setList(void* ptr, void* i);
void* QRemoteObjectReplica___setProperties_properties_newList(void* ptr);
void* QRemoteObjectReplica___retrieveProperties_atList(void* ptr, int i);
void QRemoteObjectReplica___retrieveProperties_setList(void* ptr, void* i);
void* QRemoteObjectReplica___retrieveProperties_newList(void* ptr);
void* QRemoteObjectReplica___persistProperties_props_atList(void* ptr, int i);
void QRemoteObjectReplica___persistProperties_props_setList(void* ptr, void* i);
void* QRemoteObjectReplica___persistProperties_props_newList(void* ptr);
void* QRemoteObjectReplica___dynamicPropertyNames_atList(void* ptr, int i);
void QRemoteObjectReplica___dynamicPropertyNames_setList(void* ptr, void* i);
void* QRemoteObjectReplica___dynamicPropertyNames_newList(void* ptr);
void* QRemoteObjectReplica___findChildren_atList2(void* ptr, int i);
void QRemoteObjectReplica___findChildren_setList2(void* ptr, void* i);
void* QRemoteObjectReplica___findChildren_newList2(void* ptr);
void* QRemoteObjectReplica___findChildren_atList3(void* ptr, int i);
void QRemoteObjectReplica___findChildren_setList3(void* ptr, void* i);
void* QRemoteObjectReplica___findChildren_newList3(void* ptr);
void* QRemoteObjectReplica___findChildren_atList(void* ptr, int i);
void QRemoteObjectReplica___findChildren_setList(void* ptr, void* i);
void* QRemoteObjectReplica___findChildren_newList(void* ptr);
void* QRemoteObjectReplica___children_atList(void* ptr, int i);
void QRemoteObjectReplica___children_setList(void* ptr, void* i);
void* QRemoteObjectReplica___children_newList(void* ptr);
char QRemoteObjectReplica_EventDefault(void* ptr, void* e);
char QRemoteObjectReplica_EventFilterDefault(void* ptr, void* watched, void* event);
void QRemoteObjectReplica_ChildEventDefault(void* ptr, void* event);
void QRemoteObjectReplica_ConnectNotifyDefault(void* ptr, void* sign);
void QRemoteObjectReplica_CustomEventDefault(void* ptr, void* event);
void QRemoteObjectReplica_DeleteLaterDefault(void* ptr);
void QRemoteObjectReplica_DisconnectNotifyDefault(void* ptr, void* sign);
void QRemoteObjectReplica_TimerEventDefault(void* ptr, void* event);
struct QtRemoteObjects_PackedString QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_CLASS();
struct QtRemoteObjects_PackedString QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_MODEL();
struct QtRemoteObjects_PackedString QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_Local();
struct QtRemoteObjects_PackedString QRemoteObjectStringLiterals_QRemoteObjectStringLiterals_Tcp();
struct QtRemoteObjects_PackedString QtRemoteObjects_QtRemoteObjects_Qt_getEnumName(long long vqt);
void QtRemoteObjects_QtRemoteObjects_CopyStoredProperties3(void* mo, void* src, void* dst);
void QtRemoteObjects_QtRemoteObjects_CopyStoredProperties2(void* mo, void* src, void* dst);
void QtRemoteObjects_QtRemoteObjects_CopyStoredProperties(void* mo, void* src, void* dst);

#ifdef __cplusplus
}
#endif

#endif