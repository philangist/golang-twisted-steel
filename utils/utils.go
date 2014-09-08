package utils

import (
  "testing"
  "reflect"
  "runtime"
  "path/filepath"
)

func Check(err error){
  if err != nil{
    panic(err)
  }
}

//stole this from bitly's nsq equal implementation
//https://github.com/bitly/nsq/blob/master/nsqd/test/cluster_test.go#L16
func Equal(t *testing.T, actual, expectation interface{}){
  if !reflect.DeepEqual(actual, expectation){
    _, file, line, _ := runtime.Caller(1)
    t.Logf("\033[31m%s:%d:\n\n\texpected: %#v\n\n\tactual: %#v\033[39m\n\n",
           filepath.Base(file), line, expectation, actual)
    t.FailNow()
  }
}
