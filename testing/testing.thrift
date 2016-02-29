namespace go testinglion

struct TFoo {
  1: optional string one,
  2: optional i32 two,
  3: optional string string_field,
  4: optional i32 int32_field,
  5: optional TBar bar,
}

struct TBar {
  1: optional string one,
  2: optional string two,
  3: optional string string_field,
  4: optional i32 int32_field,
}

struct TBan {
  1: optional string string_field,
  2: optional i32 int32_field,
}

struct TBat {
  1: optional TBan ban,
}

struct TBaz {
  1: optional TBat bat,
}

struct TEmpty {}

struct TNoStdJson {
  1: optional map<i64, string> one,
}
