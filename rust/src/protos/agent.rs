// This file is generated by rust-protobuf 2.23.0. Do not edit
// @generated

// https://github.com/rust-lang/rust-clippy/issues/702
#![allow(unknown_lints)]
#![allow(clippy::all)]

#![allow(unused_attributes)]
#![cfg_attr(rustfmt, rustfmt::skip)]

#![allow(box_pointers)]
#![allow(dead_code)]
#![allow(missing_docs)]
#![allow(non_camel_case_types)]
#![allow(non_snake_case)]
#![allow(non_upper_case_globals)]
#![allow(trivial_casts)]
#![allow(unused_imports)]
#![allow(unused_results)]
//! Generated file from `agent.proto`

/// Generated files are compatible only with the same version
/// of protobuf runtime.
// const _PROTOBUF_VERSION_CHECK: () = ::protobuf::VERSION_2_23_0;

#[derive(PartialEq,Clone,Default)]
pub struct CollectCPUTimeRequest {
    // message fields
    pub timestamp_list: ::std::vec::Vec<u64>,
    pub cpu_time_ms_list: ::std::vec::Vec<u32>,
    pub resource_tag: ::std::vec::Vec<u8>,
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a CollectCPUTimeRequest {
    fn default() -> &'a CollectCPUTimeRequest {
        <CollectCPUTimeRequest as ::protobuf::Message>::default_instance()
    }
}

impl CollectCPUTimeRequest {
    pub fn new() -> CollectCPUTimeRequest {
        ::std::default::Default::default()
    }

    // repeated uint64 timestamp_list = 1;


    pub fn get_timestamp_list(&self) -> &[u64] {
        &self.timestamp_list
    }
    pub fn clear_timestamp_list(&mut self) {
        self.timestamp_list.clear();
    }

    // Param is passed by value, moved
    pub fn set_timestamp_list(&mut self, v: ::std::vec::Vec<u64>) {
        self.timestamp_list = v;
    }

    // Mutable pointer to the field.
    pub fn mut_timestamp_list(&mut self) -> &mut ::std::vec::Vec<u64> {
        &mut self.timestamp_list
    }

    // Take field
    pub fn take_timestamp_list(&mut self) -> ::std::vec::Vec<u64> {
        ::std::mem::replace(&mut self.timestamp_list, ::std::vec::Vec::new())
    }

    // repeated uint32 cpu_time_ms_list = 2;


    pub fn get_cpu_time_ms_list(&self) -> &[u32] {
        &self.cpu_time_ms_list
    }
    pub fn clear_cpu_time_ms_list(&mut self) {
        self.cpu_time_ms_list.clear();
    }

    // Param is passed by value, moved
    pub fn set_cpu_time_ms_list(&mut self, v: ::std::vec::Vec<u32>) {
        self.cpu_time_ms_list = v;
    }

    // Mutable pointer to the field.
    pub fn mut_cpu_time_ms_list(&mut self) -> &mut ::std::vec::Vec<u32> {
        &mut self.cpu_time_ms_list
    }

    // Take field
    pub fn take_cpu_time_ms_list(&mut self) -> ::std::vec::Vec<u32> {
        ::std::mem::replace(&mut self.cpu_time_ms_list, ::std::vec::Vec::new())
    }

    // bytes resource_tag = 3;


    pub fn get_resource_tag(&self) -> &[u8] {
        &self.resource_tag
    }
    pub fn clear_resource_tag(&mut self) {
        self.resource_tag.clear();
    }

    // Param is passed by value, moved
    pub fn set_resource_tag(&mut self, v: ::std::vec::Vec<u8>) {
        self.resource_tag = v;
    }

    // Mutable pointer to the field.
    // If field is not initialized, it is initialized with default value first.
    pub fn mut_resource_tag(&mut self) -> &mut ::std::vec::Vec<u8> {
        &mut self.resource_tag
    }

    // Take field
    pub fn take_resource_tag(&mut self) -> ::std::vec::Vec<u8> {
        ::std::mem::replace(&mut self.resource_tag, ::std::vec::Vec::new())
    }
}

impl ::protobuf::Message for CollectCPUTimeRequest {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                1 => {
                    ::protobuf::rt::read_repeated_uint64_into(wire_type, is, &mut self.timestamp_list)?;
                },
                2 => {
                    ::protobuf::rt::read_repeated_uint32_into(wire_type, is, &mut self.cpu_time_ms_list)?;
                },
                3 => {
                    ::protobuf::rt::read_singular_proto3_bytes_into(wire_type, is, &mut self.resource_tag)?;
                },
                _ => {
                    ::protobuf::rt::read_unknown_or_skip_group(field_number, wire_type, is, self.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u32 {
        let mut my_size = 0;
        for value in &self.timestamp_list {
            my_size += ::protobuf::rt::value_size(1, *value, ::protobuf::wire_format::WireTypeVarint);
        };
        for value in &self.cpu_time_ms_list {
            my_size += ::protobuf::rt::value_size(2, *value, ::protobuf::wire_format::WireTypeVarint);
        };
        if !self.resource_tag.is_empty() {
            my_size += ::protobuf::rt::bytes_size(3, &self.resource_tag);
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        for v in &self.timestamp_list {
            os.write_uint64(1, *v)?;
        };
        for v in &self.cpu_time_ms_list {
            os.write_uint32(2, *v)?;
        };
        if !self.resource_tag.is_empty() {
            os.write_bytes(3, &self.resource_tag)?;
        }
        os.write_unknown_fields(self.get_unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn get_cached_size(&self) -> u32 {
        self.cached_size.get()
    }

    fn get_unknown_fields(&self) -> &::protobuf::UnknownFields {
        &self.unknown_fields
    }

    fn mut_unknown_fields(&mut self) -> &mut ::protobuf::UnknownFields {
        &mut self.unknown_fields
    }

    fn as_any(&self) -> &dyn (::std::any::Any) {
        self as &dyn (::std::any::Any)
    }
    fn as_any_mut(&mut self) -> &mut dyn (::std::any::Any) {
        self as &mut dyn (::std::any::Any)
    }
    fn into_any(self: ::std::boxed::Box<Self>) -> ::std::boxed::Box<dyn (::std::any::Any)> {
        self
    }

    fn descriptor(&self) -> &'static ::protobuf::reflect::MessageDescriptor {
        Self::descriptor_static()
    }

    fn new() -> CollectCPUTimeRequest {
        CollectCPUTimeRequest::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let mut fields = ::std::vec::Vec::new();
            fields.push(::protobuf::reflect::accessor::make_vec_accessor::<_, ::protobuf::types::ProtobufTypeUint64>(
                "timestamp_list",
                |m: &CollectCPUTimeRequest| { &m.timestamp_list },
                |m: &mut CollectCPUTimeRequest| { &mut m.timestamp_list },
            ));
            fields.push(::protobuf::reflect::accessor::make_vec_accessor::<_, ::protobuf::types::ProtobufTypeUint32>(
                "cpu_time_ms_list",
                |m: &CollectCPUTimeRequest| { &m.cpu_time_ms_list },
                |m: &mut CollectCPUTimeRequest| { &mut m.cpu_time_ms_list },
            ));
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeBytes>(
                "resource_tag",
                |m: &CollectCPUTimeRequest| { &m.resource_tag },
                |m: &mut CollectCPUTimeRequest| { &mut m.resource_tag },
            ));
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<CollectCPUTimeRequest>(
                "CollectCPUTimeRequest",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static CollectCPUTimeRequest {
        static instance: ::protobuf::rt::LazyV2<CollectCPUTimeRequest> = ::protobuf::rt::LazyV2::INIT;
        instance.get(CollectCPUTimeRequest::new)
    }
}

impl ::protobuf::Clear for CollectCPUTimeRequest {
    fn clear(&mut self) {
        self.timestamp_list.clear();
        self.cpu_time_ms_list.clear();
        self.resource_tag.clear();
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for CollectCPUTimeRequest {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for CollectCPUTimeRequest {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

#[derive(PartialEq,Clone,Default)]
pub struct CollectCPUTimeResponse {
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a CollectCPUTimeResponse {
    fn default() -> &'a CollectCPUTimeResponse {
        <CollectCPUTimeResponse as ::protobuf::Message>::default_instance()
    }
}

impl CollectCPUTimeResponse {
    pub fn new() -> CollectCPUTimeResponse {
        ::std::default::Default::default()
    }
}

impl ::protobuf::Message for CollectCPUTimeResponse {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                _ => {
                    ::protobuf::rt::read_unknown_or_skip_group(field_number, wire_type, is, self.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u32 {
        let mut my_size = 0;
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        os.write_unknown_fields(self.get_unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn get_cached_size(&self) -> u32 {
        self.cached_size.get()
    }

    fn get_unknown_fields(&self) -> &::protobuf::UnknownFields {
        &self.unknown_fields
    }

    fn mut_unknown_fields(&mut self) -> &mut ::protobuf::UnknownFields {
        &mut self.unknown_fields
    }

    fn as_any(&self) -> &dyn (::std::any::Any) {
        self as &dyn (::std::any::Any)
    }
    fn as_any_mut(&mut self) -> &mut dyn (::std::any::Any) {
        self as &mut dyn (::std::any::Any)
    }
    fn into_any(self: ::std::boxed::Box<Self>) -> ::std::boxed::Box<dyn (::std::any::Any)> {
        self
    }

    fn descriptor(&self) -> &'static ::protobuf::reflect::MessageDescriptor {
        Self::descriptor_static()
    }

    fn new() -> CollectCPUTimeResponse {
        CollectCPUTimeResponse::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let fields = ::std::vec::Vec::new();
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<CollectCPUTimeResponse>(
                "CollectCPUTimeResponse",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static CollectCPUTimeResponse {
        static instance: ::protobuf::rt::LazyV2<CollectCPUTimeResponse> = ::protobuf::rt::LazyV2::INIT;
        instance.get(CollectCPUTimeResponse::new)
    }
}

impl ::protobuf::Clear for CollectCPUTimeResponse {
    fn clear(&mut self) {
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for CollectCPUTimeResponse {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for CollectCPUTimeResponse {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

static file_descriptor_proto_data: &'static [u8] = b"\
    \n\x0bagent.proto\"\x8a\x01\n\x15CollectCPUTimeRequest\x12%\n\x0etimesta\
    mp_list\x18\x01\x20\x03(\x04R\rtimestampList\x12'\n\x10cpu_time_ms_list\
    \x18\x02\x20\x03(\rR\rcpuTimeMsList\x12!\n\x0cresource_tag\x18\x03\x20\
    \x01(\x0cR\x0bresourceTag\"\x18\n\x16CollectCPUTimeResponse2[\n\x12Resou\
    rceUsageAgent\x12E\n\x0eCollectCPUTime\x12\x16.CollectCPUTimeRequest\x1a\
    \x17.CollectCPUTimeResponse\"\0(\x01b\x06proto3\
";

static file_descriptor_proto_lazy: ::protobuf::rt::LazyV2<::protobuf::descriptor::FileDescriptorProto> = ::protobuf::rt::LazyV2::INIT;

fn parse_descriptor_proto() -> ::protobuf::descriptor::FileDescriptorProto {
    ::protobuf::Message::parse_from_bytes(file_descriptor_proto_data).unwrap()
}

pub fn file_descriptor_proto() -> &'static ::protobuf::descriptor::FileDescriptorProto {
    file_descriptor_proto_lazy.get(|| {
        parse_descriptor_proto()
    })
}
