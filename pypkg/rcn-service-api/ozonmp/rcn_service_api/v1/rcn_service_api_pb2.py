# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ozonmp/rcn_service_api/v1/rcn_service_api.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from validate import validate_pb2 as validate_dot_validate__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n/ozonmp/rcn_service_api/v1/rcn_service_api.proto\x12\x19ozonmp.rcn_service_api.v1\x1a\x17validate/validate.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"b\n\x08Template\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12\x10\n\x03\x66oo\x18\x02 \x01(\x04R\x03\x66oo\x12\x34\n\x07\x63reated\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x07\x63reated\"E\n\x19\x44\x65scribeTemplateV1Request\x12(\n\x0btemplate_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\ntemplateId\"W\n\x1a\x44\x65scribeTemplateV1Response\x12\x39\n\x05value\x18\x01 \x01(\x0b\x32#.ozonmp.rcn_service_api.v1.TemplateR\x05value2\xc0\x01\n\x15OmpTemplateApiService\x12\xa6\x01\n\x12\x44\x65scribeTemplateV1\x12\x34.ozonmp.rcn_service_api.v1.DescribeTemplateV1Request\x1a\x35.ozonmp.rcn_service_api.v1.DescribeTemplateV1Response\"#\x82\xd3\xe4\x93\x02\x1d\x12\x1b/v1/templates/{template_id}BGZEgithub.com/ozonmp/rcn-service-api/pkg/rcn-service-api;rcn_service_apib\x06proto3')



_TEMPLATE = DESCRIPTOR.message_types_by_name['Template']
_DESCRIBETEMPLATEV1REQUEST = DESCRIPTOR.message_types_by_name['DescribeTemplateV1Request']
_DESCRIBETEMPLATEV1RESPONSE = DESCRIPTOR.message_types_by_name['DescribeTemplateV1Response']
Template = _reflection.GeneratedProtocolMessageType('Template', (_message.Message,), {
  'DESCRIPTOR' : _TEMPLATE,
  '__module__' : 'ozonmp.rcn_service_api.v1.rcn_service_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.rcn_service_api.v1.Template)
  })
_sym_db.RegisterMessage(Template)

DescribeTemplateV1Request = _reflection.GeneratedProtocolMessageType('DescribeTemplateV1Request', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBETEMPLATEV1REQUEST,
  '__module__' : 'ozonmp.rcn_service_api.v1.rcn_service_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.rcn_service_api.v1.DescribeTemplateV1Request)
  })
_sym_db.RegisterMessage(DescribeTemplateV1Request)

DescribeTemplateV1Response = _reflection.GeneratedProtocolMessageType('DescribeTemplateV1Response', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBETEMPLATEV1RESPONSE,
  '__module__' : 'ozonmp.rcn_service_api.v1.rcn_service_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.rcn_service_api.v1.DescribeTemplateV1Response)
  })
_sym_db.RegisterMessage(DescribeTemplateV1Response)

_OMPTEMPLATEAPISERVICE = DESCRIPTOR.services_by_name['OmpTemplateApiService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZEgithub.com/ozonmp/rcn-service-api/pkg/rcn-service-api;rcn_service_api'
  _DESCRIBETEMPLATEV1REQUEST.fields_by_name['template_id']._options = None
  _DESCRIBETEMPLATEV1REQUEST.fields_by_name['template_id']._serialized_options = b'\372B\0042\002 \000'
  _OMPTEMPLATEAPISERVICE.methods_by_name['DescribeTemplateV1']._options = None
  _OMPTEMPLATEAPISERVICE.methods_by_name['DescribeTemplateV1']._serialized_options = b'\202\323\344\223\002\035\022\033/v1/templates/{template_id}'
  _TEMPLATE._serialized_start=166
  _TEMPLATE._serialized_end=264
  _DESCRIBETEMPLATEV1REQUEST._serialized_start=266
  _DESCRIBETEMPLATEV1REQUEST._serialized_end=335
  _DESCRIBETEMPLATEV1RESPONSE._serialized_start=337
  _DESCRIBETEMPLATEV1RESPONSE._serialized_end=424
  _OMPTEMPLATEAPISERVICE._serialized_start=427
  _OMPTEMPLATEAPISERVICE._serialized_end=619
# @@protoc_insertion_point(module_scope)
