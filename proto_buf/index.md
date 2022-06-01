# Protocol Buffers



Protocol buffers are a language-neutral, platform-neutral extensible mechanism for serializing structured data.
<!--more-->
<br />

## tl;dr

{{< admonition note "What are protocol buffers" >}}
_Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages._
{{< /admonition  >}}

How do I start?

1. [Download and install](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation) the protocol buffer compiler.
2. Read the [overview](https://developers.google.com/protocol-buffers/docs/overview).
3. Try the tutorial for your chosen language. ([python](https://developers.google.com/protocol-buffers/docs/pythontutorial))

## overview
> https://developers.google.com/protocol-buffers/docs/overview

{{< admonition note "protobuf" >}}
_It’s like JSON, except it's smaller and faster, and it generates native language bindings._

_Protocol buffers are a combination of the definition language (created in .proto files), the code that the proto compiler generates to interface with data, language-specific runtime libraries, and the **serialization** format for data that is written to a file (or sent across a network connection)._
{{< /admonition  >}}

These are protobuf's main components (2)

1. `Protoc` (compiler)
  - It is for data format compile
  - It compiles `.proto` files
2. `SDK`
  - each language support

- The proto compiler is invoked at build time on .proto files to generate code in various programming languages
- Each generated class contains simple accessors for each field and methods to serialize and parse the whole structure to and from raw bytes.

### languages

The following languages are supported directly in the protocol buffers compiler, protoc:

- C++
- C#
- Java
- Kotlin
- Objective-C
- PHP
- Python
- Ruby

The following languages are supported by Google, but the projects' source code resides in GitHub repositories. The protoc compiler uses plugins for these languages

- Dart
- Go

[for other languages](https://github.com/protocolbuffers/protobuf/blob/main/docs/third_party.md)


### Pros

- language/platform-neutral (low coupling, microsevice)
- Compact data storage
- Fast Parsing (compared to json?)
- Availability in many programming languages
- **Optimized functionality through automatically-generated classes**
- **You can update `Proto Definitions` without updating code.**
  - which refers you can control code(especially data schema) version compatibility.

### Cons

- Protocol buffers tend to assume that entire messages can be loaded into memory at once and are not larger than an object graph. For data that exceeds a few megabytes, consider a different solution; when working with larger data, you may effectively end up with several copies of the data due to serialized copies, which can cause surprising spikes in memory usage.
- When protocol buffers are serialized, the same data can have many different binary serializations. You cannot compare two messages for equality without fully parsing them.
- Messages are not compressed. While messages can be zipped or gzipped like any other file, special-purpose compression algorithms like the ones used by JPEG and PNG will produce much smaller files for data of the appropriate type.
- Protocol buffer messages are less than maximally efficient in both size and speed for many scientific and engineering uses that involve large, multi-dimensional arrays of floating point numbers. For these applications, `FITS` and similar formats have less overhead.
- Protocol buffers are not well supported in non-object-oriented languages popular in scientific computing, such as Fortran and IDL.
- Protocol buffer messages don't inherently self-describe their data, but they have a fully reflective schema that you can use to implement self-description. That is, you cannot fully interpret one without access to its corresponding `.proto` file.
- Protocol buffers are not a formal standard of any organization. This makes them unsuitable for use in environments with legal or other requirements to build on top of standards.

### Flow

![](/images/protocol-buffers-concepts.png)

### .proto definition syntax (3)

1. optionality(field rules)
  - `optional`
  - `repeated`
    - Repeated fields are represented as an object that acts like a Python sequence
  - `singular`(proto3, default, 단수형)
  - `required`(deprecated)
  - `reversed`
2. field type
  - `message`
    - you can nest parts of the definition, such as for repeating sets of data.
  - `enum`
    - set of values to choose from.
  - `oneof`
    - which you can use when a message has many optional fields and at most one field will be set at the same time.
  - `map`
3. field number
4. basic scalar type
5. [additional scalar type](https://developers.google.com/protocol-buffers/docs/overview#common-types)

{{< admonition warning "field number" >}}
_Field numbers cannot be repurposed or reused. If you delete a field, you should reserve its field number to prevent someone from accidentally reusing the number._

- [from microsoft description](https://docs.microsoft.com/ko-kr/dotnet/architecture/grpc-for-wcf-developers/protobuf-messages#field-numbers)

_필드 번호는 Protobuf의 중요한 부분입니다. 이진 인코딩된 데이터의 필드를 식별하는 데 사용됩니다. 즉, 서비스 버전에서 버전으로 변경할 수 없습니다. 장점은 이전 버전과의 호환성 및 앞으로 호환성이 가능하다는 것입니다. 클라이언트 및 서비스는 누락된 값의 가능성이 처리되는 한 모르는 필드 번호를 무시합니다._

_이진 형식에서 필드 번호는 형식 식별자와 결합됩니다. 1에서 15까지의 필드 번호는 해당 형식으로 단일 바이트로 인코딩할 수 있습니다. 16에서 2,047까지의 숫자는 2바이트를 사용합니다. 어떤 이유로든 메시지에 2,047개 이상의 필드가 필요한 경우 더 높아질 수 있습니다. 필드 번호 1에서 15까지의 싱글 바이트 식별자는 더 나은 성능을 제공하므로 가장 기본적으로 자주 사용되는 필드에 사용해야 합니다._
{{< /admonition  >}}

### Example

This is the `status.proto` file used by Google. ([refs](https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto))

```protobuf
// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

syntax = "proto3";

package google.rpc;

import "google/protobuf/any.proto";

option cc_enable_arenas = true;
option go_package = "google.golang.org/genproto/googleapis/rpc/status;status";
option java_multiple_files = true;
option java_outer_classname = "StatusProto";
option java_package = "com.google.rpc";
option objc_class_prefix = "RPC";

// The `Status` type defines a logical error model that is suitable for
// different programming environments, including REST APIs and RPC APIs. It is
// used by [gRPC](https://github.com/grpc). Each `Status` message contains
// three pieces of data: error code, error message, and error details.
//
// You can find out more about this error model and how to work with it in the
// [API Design Guide](https://cloud.google.com/apis/design/errors).
message Status {
  // The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].
  int32 code = 1;

  // A developer-facing error message, which should be in English. Any
  // user-facing error message should be localized and sent in the
  // [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
  string message = 2;

  // A list of messages that carry the error details.  There is a common set of
  // message types for APIs to use.
  repeated google.protobuf.Any details = 3;
}
```


### protobuf: python
> https://developers.google.com/protocol-buffers/docs/reference/python-generated#invocation


### protobuf vs ...

- protobuf(grpc) vs thrift
  - [grpc vs thrift](https://www.alluxio.io/blog/moving-from-apache-thrift-to-grpc-a-perspective-from-alluxio/)
- protobuf(grpc) vs graphql
  - https://github.com/google/rejoiner
  - 구글에서 graphql과 호환 lib 오픈소스화 시킴.
  - https://medium.com/@lvdbrink/graphql-meets-protocol-buffers-in-go-cdbf11090934

### refs

- [protobuf github](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation)

