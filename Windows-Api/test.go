package winapi

/*
#include <windows.h>
#include <winnetwk.h>
int datspecialnumber() {
    return ERROR_CONNECTED_OTHER_PASSWORD_DEFAULT;
}
*/
import "C"
import "fmt"

func Test(){
  text := fmt.Sprintf("%x", C.datspecialnumber())
  size := 8-len(text)
  needed := "0x"
  for i := 0; i < size; i++{
    needed += "0"
  }
  needed +=text
  fmt.Println(needed)
}
