# proto 

- [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
- [Language Guide (proto3)](https://developers.google.com/protocol-buffers/docs/proto3)
- [消息中字段的Unique Number是什么?](https://developers.google.com/protocol-buffers/docs/proto3#assigning_field_numbers)
    - 1-15仅取一个字节进行编码， 16-2047仅取两个字节进行编码
    - 1-15应该留给频繁出现的字段
    - 应该预留些数据给以后可能添加的频繁字段

```protobuf
// 消息
message Person {
  // 消息字段, 包括类型、名称和唯一号码
  string name = 1;
  int32 id = 2;
  
  // 数组
  repeated string phones = 3;

  // 非频繁字段
  string email = 16;
  google.protobuf.Timestamp last_updated = 17;
}
```