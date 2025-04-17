//
//  Generated code. Do not modify.
//  source: proto/service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class GetMessagesRequest_Direction extends $pb.ProtobufEnum {
  static const GetMessagesRequest_Direction None = GetMessagesRequest_Direction._(0, _omitEnumNames ? '' : 'None');
  static const GetMessagesRequest_Direction NextPage = GetMessagesRequest_Direction._(1, _omitEnumNames ? '' : 'NextPage');
  static const GetMessagesRequest_Direction PrevPage = GetMessagesRequest_Direction._(2, _omitEnumNames ? '' : 'PrevPage');
  static const GetMessagesRequest_Direction BothPage = GetMessagesRequest_Direction._(3, _omitEnumNames ? '' : 'BothPage');
  static const GetMessagesRequest_Direction LastPage = GetMessagesRequest_Direction._(4, _omitEnumNames ? '' : 'LastPage');

  static const $core.List<GetMessagesRequest_Direction> values = <GetMessagesRequest_Direction> [
    None,
    NextPage,
    PrevPage,
    BothPage,
    LastPage,
  ];

  static final $core.Map<$core.int, GetMessagesRequest_Direction> _byValue = $pb.ProtobufEnum.initByValue(values);
  static GetMessagesRequest_Direction? valueOf($core.int value) => _byValue[value];

  const GetMessagesRequest_Direction._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
