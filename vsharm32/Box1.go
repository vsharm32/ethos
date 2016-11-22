package main
import(
	"io"
	"math"
	"reflect"
	"unsafe"
	"ethos/llrb"
	"os"
	"bytes"
	"encoding/hex"
	"time"
	"fmt"
	
	"ethos/syscall"
	
	"ethos/goodmiddleman"
	
	"log"
	
)

// FIXME: this is a place holder for bytes. bytes is only used when there is an any type
var xxx = bytes.MinRead
var yyy hex.InvalidHexCharError
var sunday = time.Sunday // this needs to be remove after we fixed the sleep in Ipc and IpcWrite.


func Ipc(hostname string, serviceName string) (*Encoder, *Decoder, syscall.Status) {
     serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
     if status != syscall.StatusOk {
                  log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		  return nil, nil, status
     }

     sName := serviceName
     for i:=len(serviceName)-1; i>=0; i-- {
         if (serviceName[i]=='/') {
      	   sName = serviceName[i+1:]
      	}
     }


     netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
     if status != syscall.StatusOk {
                log.Fatalf ("Error calling Ipc: %v\n", status)
	    return nil, nil, status	
     }
     writer := goodmiddleman.NewWriter(netFd)
     reader := goodmiddleman.NewReader(netFd)
     goodmiddleman.Close(serviceFd)
     e := NewEncoder(writer)
     d := NewDecoder(reader)
     return e, d, syscall.StatusOk
}

func Advertise(serviceName string) (listeningFd syscall.Fd, status syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	listeningFd, status = goodmiddleman.Advertise (serviceFd, serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling Advertise: %v\n", status)
		return 0, status     
	}

	return listeningFd, status
}

func Import(listeningFd syscall.Fd) (user []byte, e *Encoder, d *Decoder, status syscall.Status) {
     user, netFd, status := goodmiddleman.Import (listeningFd)
     if status != syscall.StatusOk {
                log.Fatalf ("Error calling Import: %v\n", status)
		return []byte{}, nil, nil, status
     }
     writer := goodmiddleman.NewWriter(netFd)
     reader := goodmiddleman.NewReader(netFd)
     goodmiddleman.Close(listeningFd)
     e = NewEncoder(writer)
     d = NewDecoder(reader)
     return user, e, d, syscall.StatusOk
}

type EncoderDecoderFd struct {
     E *Encoder
     D *Decoder
     Fd syscall.Fd
}

func IpcEncoderDecoder(hostname string, serviceName string) (EncoderDecoderFd, syscall.Status) {
     var ed EncoderDecoderFd
     serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
     if status != syscall.StatusOk {
		  return ed, status
     }

     sName := serviceName
     for i:=len(serviceName)-1; i>=0; i-- {
         if (serviceName[i]=='/') {
      	   sName = serviceName[i+1:]
      	}
     }


     netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
     if status != syscall.StatusOk {
	    return ed, status	
     }
     writer := goodmiddleman.NewWriter(netFd)
     reader := goodmiddleman.NewReader(netFd)
     goodmiddleman.Close(serviceFd)
     e := NewEncoder(writer)
     d := NewDecoder(reader)
     ed.E = e
     ed.D = d
     ed.Fd = netFd
     return ed, syscall.StatusOk     
}

func (ed *EncoderDecoderFd) Close() {
     goodmiddleman.Close(ed.Fd)
}


type Box1 struct {
     
     x1 int32
     
     y1 int32
     
     x2 int32
     
     y2 int32
     
}


type Box2 struct {
     
     x3 int32
     
     y3 int32
     
     x4 int32
     
     y4 int32
     
     slope int32
     
}

func (e *Encoder) box1(v *Box1) (err os.Error){
     
     if _, _, error := e.t.PointerCheck(unsafe.Pointer(v), "0b597097877469841d234bd0630c40a2a79061ed999d2c271da54b487688c92523deaaa2f954ab65295b001e73a2ccc84b625f27e5dfe0096d726592f0baf780", uint64(unsafe.Sizeof(*v))); error == nil {
     	
       err = e.int32(v.x1)
       if err != nil {
       	  return err
       }
       
       err = e.int32(v.y1)
       if err != nil {
       	  return err
       }
       
       err = e.int32(v.x2)
       if err != nil {
       	  return err
       }
       
       err = e.int32(v.y2)
       if err != nil {
       	  return err
       }
        
     } else {
       return error
     }   
     return nil
}

func (e *Encoder) Box1 (v *Box1) (err os.Error){
     return e.box1(v)
}

func (e *Encoder) box1Internal(v *Box1) (err os.Error){
     
     
     	     err = e.int32(v.x1)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.int32(v.y1)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.int32(v.x2)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.int32(v.y2)
	     if err != nil {
	     	return err
	     }
     
     return nil
}


func (d *Decoder) box1() (v *Box1, error os.Error){
     
      

     var valv Box1 
     d.indexToValue = append(d.indexToValue, &valv)
     
     	
     p0, err0 := d.int32()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv.x1 = *p0
     
     
     	
     p1, err1 := d.int32()
     if err1 != nil {
     	return &valv, err1
     }
     
     valv.y1 = *p1
     
     
     	
     p2, err2 := d.int32()
     if err2 != nil {
     	return &valv, err2
     }
     
     valv.x2 = *p2
     
     
     	
     p3, err3 := d.int32()
     if err3 != nil {
     	return &valv, err3
     }
     
     valv.y2 = *p3
     
     
     v = &valv
     return v, nil
}

func (d *Decoder) Box1() (v *Box1, error os.Error){
     return d.box1()
}

func (d *Decoder) box1Internal() (v *Box1, error os.Error){
     
      
     var valv Box1
     
          
     p0, err0 := d.int32()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv.x1 = *p0
     
     
          
     p1, err1 := d.int32()
     if err1 != nil {
     	return &valv, err1
     }
     
     valv.y1 = *p1
     
     
          
     p2, err2 := d.int32()
     if err2 != nil {
     	return &valv, err2
     }
     
     valv.x2 = *p2
     
     
          
     p3, err3 := d.int32()
     if err3 != nil {
     	return &valv, err3
     }
     
     valv.y2 = *p3
     
     
     v = &valv     
     return v, nil
}

func (e *Encoder) box2(v *Box2) (err os.Error){
     
     if _, _, error := e.t.PointerCheck(unsafe.Pointer(v), "f64056a437617f591da3009d3e72c42d963d5a42d7b45a8251bdbc69b0aefd7ae82a3d17ed69e86bc7d226fefb4e35caf70b8d7183691d1cf8aaf967fabcbe66", uint64(unsafe.Sizeof(*v))); error == nil {
     	
       err = e.int32(v.x3)
       if err != nil {
       	  return err
       }
       
       err = e.int32(v.y3)
       if err != nil {
       	  return err
       }
       
       err = e.int32(v.x4)
       if err != nil {
       	  return err
       }
       
       err = e.int32(v.y4)
       if err != nil {
       	  return err
       }
       
       err = e.int32(v.slope)
       if err != nil {
       	  return err
       }
        
     } else {
       return error
     }   
     return nil
}

func (e *Encoder) Box2 (v *Box2) (err os.Error){
     return e.box2(v)
}

func (e *Encoder) box2Internal(v *Box2) (err os.Error){
     
     
     	     err = e.int32(v.x3)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.int32(v.y3)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.int32(v.x4)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.int32(v.y4)
	     if err != nil {
	     	return err
	     }
     
     	     err = e.int32(v.slope)
	     if err != nil {
	     	return err
	     }
     
     return nil
}


func (d *Decoder) box2() (v *Box2, error os.Error){
     
      

     var valv Box2 
     d.indexToValue = append(d.indexToValue, &valv)
     
     	
     p0, err0 := d.int32()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv.x3 = *p0
     
     
     	
     p1, err1 := d.int32()
     if err1 != nil {
     	return &valv, err1
     }
     
     valv.y3 = *p1
     
     
     	
     p2, err2 := d.int32()
     if err2 != nil {
     	return &valv, err2
     }
     
     valv.x4 = *p2
     
     
     	
     p3, err3 := d.int32()
     if err3 != nil {
     	return &valv, err3
     }
     
     valv.y4 = *p3
     
     
     	
     p4, err4 := d.int32()
     if err4 != nil {
     	return &valv, err4
     }
     
     valv.slope = *p4
     
     
     v = &valv
     return v, nil
}

func (d *Decoder) Box2() (v *Box2, error os.Error){
     return d.box2()
}

func (d *Decoder) box2Internal() (v *Box2, error os.Error){
     
      
     var valv Box2
     
          
     p0, err0 := d.int32()
     if err0 != nil {
     	return &valv, err0
     }
     
     valv.x3 = *p0
     
     
          
     p1, err1 := d.int32()
     if err1 != nil {
     	return &valv, err1
     }
     
     valv.y3 = *p1
     
     
          
     p2, err2 := d.int32()
     if err2 != nil {
     	return &valv, err2
     }
     
     valv.x4 = *p2
     
     
          
     p3, err3 := d.int32()
     if err3 != nil {
     	return &valv, err3
     }
     
     valv.y4 = *p3
     
     
          
     p4, err4 := d.int32()
     if err4 != nil {
     	return &valv, err4
     }
     
     valv.slope = *p4
     
     
     v = &valv     
     return v, nil
}

func (e *Encoder) any(a Any) (err os.Error){
     switch a.Value.(type) {
     
     case Box1:
     	hashByte, err := hex.DecodeString("0b597097877469841d234bd0630c40a2a79061ed999d2c271da54b487688c92523deaaa2f954ab65295b001e73a2ccc84b625f27e5dfe0096d726592f0baf780")
	if err != nil {
	   return err
	}
     	// hashByte := []byte("0b597097877469841d234bd0630c40a2a79061ed999d2c271da54b487688c92523deaaa2f954ab65295b001e73a2ccc84b625f27e5dfe0096d726592f0baf780")
     	// e.string("0b597097877469841d234bd0630c40a2a79061ed999d2c271da54b487688c92523deaaa2f954ab65295b001e73a2ccc84b625f27e5dfe0096d726592f0baf780")
	err = e.SliceOfBytes(hashByte)
	if err != nil {
	   return err
	}
     	t := a.Value.(Box1)
	err = e.box1(&t)
	if err != nil {
	   return err
	}
     
     case Box2:
     	hashByte, err := hex.DecodeString("f64056a437617f591da3009d3e72c42d963d5a42d7b45a8251bdbc69b0aefd7ae82a3d17ed69e86bc7d226fefb4e35caf70b8d7183691d1cf8aaf967fabcbe66")
	if err != nil {
	   return err
	}
     	// hashByte := []byte("f64056a437617f591da3009d3e72c42d963d5a42d7b45a8251bdbc69b0aefd7ae82a3d17ed69e86bc7d226fefb4e35caf70b8d7183691d1cf8aaf967fabcbe66")
     	// e.string("f64056a437617f591da3009d3e72c42d963d5a42d7b45a8251bdbc69b0aefd7ae82a3d17ed69e86bc7d226fefb4e35caf70b8d7183691d1cf8aaf967fabcbe66")
	err = e.SliceOfBytes(hashByte)
	if err != nil {
	   return err
	}
     	t := a.Value.(Box2)
	err = e.box2(&t)
	if err != nil {
	   return err
	}
     
     case nil:
     	err = e.SliceOfBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
	if err != nil {
	   return err
	}
	err = e.uint8(pNIL)
	if err != nil {
	   return err
	}
     case int8:
     	err = e.SliceOfBytes([]byte{0xc2, 0x71, 0x6e, 0x3c, 0x34, 0x13, 0xc0, 0xbe, 0xb3, 0x9e, 0x2e, 0xbc, 0xc7, 0x99, 0x62, 0x3f, 0x28, 0xe9, 0x9d, 0x19, 0x71, 0x6a, 0x0, 0x5b, 0x69, 0x4b, 0x1c, 0xbb, 0x3d, 0x8d, 0xcc, 0x45, 0xbd, 0x51, 0xca, 0x50, 0xc7, 0x3f, 0x8a, 0x64, 0xc8, 0xa0, 0xf0, 0xb, 0x46, 0x9a, 0x87, 0x5, 0xe9, 0x3c, 0xb9, 0x27, 0x91, 0xdf, 0x95, 0x88, 0x1b, 0x2e, 0xfa, 0x9f, 0xc0, 0x1f, 0xcb, 0xad}) 
	if err != nil {
	   return err
	}
	t := a.Value.(int8)
	err = e.int8(t)
	if err != nil {
	   return err
	}
     case uint8:
     	err = e.SliceOfBytes([]byte{0x9, 0x66, 0x78, 0xbb, 0x4a, 0x86, 0x63, 0x8d, 0x8e, 0xc6, 0x58, 0x8e, 0xcc, 0x2d, 0x89, 0x5a, 0x5c, 0x17, 0xb4, 0x86, 0x37, 0x8c, 0x81, 0xc2, 0xf2, 0xac, 0xf2, 0x48, 0x67, 0x57, 0x82, 0x13, 0x3e, 0x2e, 0x7d, 0x80, 0xb6, 0x66, 0x8d, 0x84, 0xc4, 0x55, 0xf4, 0xc9, 0xe1, 0x33, 0xfc, 0x71, 0xb, 0x77, 0x43, 0x63, 0x6, 0x82, 0x76, 0x63, 0x7, 0xbc, 0xf1, 0x64, 0xdb, 0xd9, 0x5c, 0x93})
	if err != nil {
	   return err
	}
	t := a.Value.(uint8)
	err = e.uint8(t)
	if err != nil {
	   return err
	}
     case int16:
     	err = e.SliceOfBytes([]byte{0x99, 0x7f, 0x69, 0x9c, 0x17, 0xfd, 0x6, 0x74, 0x8c, 0x2d, 0xba, 0xc4, 0x61, 0x1, 0x35, 0x21, 0xf1, 0x97, 0xd8, 0x1e, 0x74, 0x3c, 0x2f, 0x96, 0x56, 0xe6, 0xdc, 0xfc, 0x14, 0x1e, 0xd, 0x83, 0x36, 0xdc, 0x73, 0x36, 0xb0, 0xf4, 0x9e, 0x40, 0x2b, 0xfe, 0x97, 0x6f, 0xfa, 0xa5, 0x27, 0xac, 0xe1, 0xa2, 0x57, 0x2a, 0xae, 0x6d, 0x18, 0x22, 0xe2, 0xdc, 0xd8, 0x79, 0xe0, 0xb6, 0xf6, 0x7e})
	if err != nil {
	   return err
	}
	t := a.Value.(int16)
	err = e.int16(t)
	if err != nil {
	   return err
	}
     case uint16:
     	err = e.SliceOfBytes([]byte{0xa9, 0x1f, 0x92, 0xf4, 0xe9, 0x96, 0xb6, 0xd4, 0xa4, 0xcb, 0x85, 0x8d, 0x11, 0x80, 0x20, 0x10, 0xd6, 0x29, 0xba, 0x29, 0xe1, 0x89, 0x50, 0x2c, 0xa0, 0xf2, 0xcb, 0x1d, 0x86, 0xb8, 0x27, 0x3b, 0x67, 0xf3, 0x35, 0x72, 0xd1, 0x78, 0x15, 0xcc, 0xb0, 0x94, 0x94, 0x6f, 0x2, 0xab, 0x2e, 0x46, 0xcd, 0x74, 0xea, 0xf5, 0x15, 0xa, 0x26, 0xdc, 0x4e, 0xf7, 0xd, 0x9f, 0x3a, 0x9c, 0x6e, 0x55})
	if err != nil {
	   return err
	}
	t := a.Value.(uint16)
	err = e.uint16(t)
	if err != nil {
	   return err
	}
     case int32:
     	err = e.SliceOfBytes([]byte{0xeb, 0xc1, 0x67, 0x8b, 0x6, 0x82, 0x70, 0x21, 0x38, 0xc2, 0xd9, 0x9e, 0x33, 0x22, 0xd1, 0xa8, 0xc7, 0x2e, 0x9b, 0x68, 0xe9, 0x41, 0x12, 0x0, 0x1e, 0x3e, 0x51, 0xa8, 0xf5, 0xd9, 0xfa, 0x34, 0xc, 0x44, 0x9c, 0x6, 0x6d, 0x9d, 0x4c, 0xe7, 0x2a, 0x6, 0xab, 0x75, 0x77, 0x5d, 0xdf, 0x28, 0x34, 0x88, 0x7c, 0x7e, 0x96, 0x97, 0xbb, 0x8a, 0x95, 0xfe, 0x7, 0x65, 0xf7, 0x7c, 0x7e, 0x4c})
	if err != nil {
	   return err
	}
	t := a.Value.(int32)
	err = e.int32(t)
	if err != nil {
	   return err
	}
     case uint32:
     	err = e.SliceOfBytes([]byte{0xce, 0xf4, 0x3a, 0x5, 0xae, 0x67, 0xd9, 0x73, 0xc2, 0xa2, 0x1d, 0xf8, 0xcd, 0xf9, 0xd2, 0xde, 0x69, 0x8d, 0xd, 0xb7, 0x61, 0xb9, 0x51, 0x22, 0x58, 0xed, 0x8f, 0xb1, 0x83, 0xf1, 0x5c, 0xff, 0x5b, 0x84, 0xe2, 0x14, 0xe, 0x10, 0x68, 0x3f, 0x7a, 0xd9, 0xa7, 0x8f, 0x5b, 0xe4, 0x9e, 0x4e, 0x0, 0x7d, 0xcb, 0xfb, 0xd1, 0x69, 0x59, 0x9d, 0xbf, 0x9b, 0x75, 0x65, 0x15, 0x9e, 0x8b, 0x82})
	if err != nil {
	   return err
	}
	t := a.Value.(uint32)
	err = e.uint32(t)
	if err != nil {
	   return err
	}
     case int64:
     	err = e.SliceOfBytes([]byte{0x56, 0x16, 0x70, 0x35, 0xd0, 0x9, 0x18, 0x69, 0xe, 0xae, 0xad, 0x60, 0xd1, 0xee, 0x39, 0xa8, 0x61, 0x45, 0x58, 0x5b, 0x99, 0x20, 0x94, 0x57, 0x1f, 0xb0, 0x48, 0xeb, 0xb2, 0xcf, 0x5c, 0xa5, 0x8d, 0xc7, 0x8e, 0x7e, 0x3c, 0x89, 0xcd, 0x2f, 0xdc, 0xf2, 0x1c, 0x2a, 0xe3, 0xd2, 0x7f, 0x98, 0xc2, 0xad, 0x1c, 0x3d, 0x4e, 0x62, 0xd9, 0xdb, 0xc8, 0xc8, 0x59, 0xc5, 0xd5, 0xc6, 0xed, 0x7a})
	if err != nil {
	   return err
	}
	t := a.Value.(int64)
	err = e.int64(t)
	if err != nil {
	   return err
	}
     case uint64:
     	err = e.SliceOfBytes([]byte{0x11, 0x26, 0xb3, 0xd, 0x51, 0x59, 0x87, 0x5e, 0xd, 0x5b, 0x93, 0xfc, 0x92, 0xf0, 0x78, 0xaa, 0x12, 0xac, 0x93, 0xb8, 0x30, 0x1f, 0x48, 0xe, 0x13, 0x4d, 0x8b, 0xfb, 0x4c, 0x58, 0xfa, 0x3a, 0x69, 0x6a, 0x81, 0x1, 0xc5, 0x47, 0xc1, 0x55, 0x43, 0x95, 0x41, 0xdf, 0x3c, 0x8e, 0xb6, 0x96, 0x4a, 0x3c, 0x88, 0xab, 0x3f, 0x88, 0xed, 0x37, 0x5f, 0x8, 0x4a, 0x41, 0x8e, 0xd5, 0xda, 0x1e})
	if err != nil {
	   return err
	}
	t := a.Value.(uint64)
	err = e.uint64(t)
	if err != nil {
	   return err
	}
     case float32:
     	err = e.SliceOfBytes([]byte{0x42, 0x36, 0xae, 0xd3, 0x62, 0xca, 0x34, 0x75, 0x94, 0x52, 0xf0, 0x5f, 0x44, 0x83, 0x61, 0x75, 0x69, 0x39, 0xcf, 0x69, 0x74, 0x91, 0xee, 0x8d, 0x35, 0x8c, 0xd7, 0xa1, 0x63, 0xf, 0x88, 0x86, 0x6b, 0x52, 0xdd, 0x6d, 0xe1, 0xb2, 0x26, 0xf4, 0x3a, 0x9c, 0x9e, 0xf1, 0x56, 0xd, 0xf1, 0x48, 0x7, 0x39, 0x46, 0xf8, 0xe9, 0xd3, 0xab, 0x86, 0xe0, 0x1c, 0x98, 0xd, 0x17, 0x6b, 0x2, 0x63})
	if err != nil {
	   return err
	}
	t := a.Value.(float32)
	err = e.float32(t)
	if err != nil {
	   return err
	}
     case float64:
     	err = e.SliceOfBytes([]byte{0x65, 0x9b, 0xb2, 0x59, 0x85, 0xe2, 0x60, 0xe7, 0x1e, 0x12, 0x17, 0x3f, 0xc3, 0x1f, 0x20, 0x45, 0x8, 0x9e, 0x7e, 0x11, 0x6b, 0xaa, 0xb3, 0x1e, 0x6d, 0x7d, 0x7a, 0x5b, 0xe3, 0x3d, 0x40, 0xb5, 0x40, 0x6, 0x52, 0x85, 0x37, 0x80, 0x2c, 0xd8, 0x7d, 0x48, 0x67, 0xe3, 0x9a, 0xdd, 0xc9, 0x13, 0x11, 0x2c, 0xa5, 0xcc, 0x5a, 0x33, 0xbc, 0x35, 0x6b, 0x3e, 0xa8, 0x75, 0x93, 0x84, 0xcf, 0x1b})
	if err != nil {
	   return err
	}
	t := a.Value.(float64)
	err = e.float64(t)
	if err != nil {
	   return err
	}
     case bool:
     	err = e.SliceOfBytes([]byte{0x3e, 0x76, 0x6, 0x81, 0x73, 0x27, 0x61, 0xed, 0x16, 0x8e, 0xa3, 0x8, 0xe1, 0x10, 0x11, 0x85, 0xe1, 0xbd, 0x39, 0x2, 0xee, 0x67, 0x60, 0x21, 0x6a, 0x59, 0xc4, 0x7, 0x5a, 0x99, 0xc1, 0x46, 0xb7, 0xcd, 0x98, 0x14, 0xce, 0x14, 0x47, 0xe, 0xb0, 0x80, 0x6d, 0x91, 0x66, 0x50, 0xb0, 0xe5, 0xe, 0x77, 0x6f, 0x53, 0xe5, 0xd1, 0x72, 0x28, 0x1d, 0xd0, 0xe1, 0x70, 0x43, 0xc8, 0x65, 0x39})
	if err != nil {
	   return err
	}
	t := a.Value.(bool)
	err = e.bool(t)
	if err != nil {
	   return err
	}
     case string:
     	err = e.SliceOfBytes([]byte{0x27, 0x57, 0xcb, 0x3c, 0xaf, 0xc3, 0x9a, 0xf4, 0x51, 0xab, 0xb2, 0x69, 0x7b, 0xe7, 0x9b, 0x4a, 0xb6, 0x1d, 0x63, 0xd7, 0x4d, 0x85, 0xb0, 0x41, 0x86, 0x29, 0xde, 0x8c, 0x26, 0x81, 0x1b, 0x52, 0x9f, 0x3f, 0x37, 0x80, 0xd0, 0x15, 0x0, 0x63, 0xff, 0x55, 0xa2, 0xbe, 0xee, 0x74, 0xc4, 0xec, 0x10, 0x2a, 0x2a, 0x27, 0x31, 0xa1, 0xf1, 0xf7, 0xf1, 0xd, 0x47, 0x3a, 0xd1, 0x8a, 0x6a, 0x87})
	if err != nil {
	   return err
	}
	t := a.Value.(string)
	err = e.string(t)
	if err != nil {
	   return err
	}
     default:
	v := reflect.ValueOf(a)
     	e := fmt.Errorf("Wrong type used as Any " + v.Type().Name())
	return e
     }
     return nil
}

func (e *Encoder) anyInternal(a Any) (err os.Error){
     switch a.Value.(type) {
     
     case Box1:
     	hashByte, err := hex.DecodeString("0b597097877469841d234bd0630c40a2a79061ed999d2c271da54b487688c92523deaaa2f954ab65295b001e73a2ccc84b625f27e5dfe0096d726592f0baf780")
	if err != nil {
	   return err
	}
     	// hashByte := []byte("0b597097877469841d234bd0630c40a2a79061ed999d2c271da54b487688c92523deaaa2f954ab65295b001e73a2ccc84b625f27e5dfe0096d726592f0baf780")
     	// e.string("0b597097877469841d234bd0630c40a2a79061ed999d2c271da54b487688c92523deaaa2f954ab65295b001e73a2ccc84b625f27e5dfe0096d726592f0baf780")
	err = e.SliceOfBytes(hashByte)
	if err != nil {
	   return err
	}
     	t := a.Value.(Box1)
	err = e.box1Internal(&t)
	if err != nil {
	   return err
	}
     
     case Box2:
     	hashByte, err := hex.DecodeString("f64056a437617f591da3009d3e72c42d963d5a42d7b45a8251bdbc69b0aefd7ae82a3d17ed69e86bc7d226fefb4e35caf70b8d7183691d1cf8aaf967fabcbe66")
	if err != nil {
	   return err
	}
     	// hashByte := []byte("f64056a437617f591da3009d3e72c42d963d5a42d7b45a8251bdbc69b0aefd7ae82a3d17ed69e86bc7d226fefb4e35caf70b8d7183691d1cf8aaf967fabcbe66")
     	// e.string("f64056a437617f591da3009d3e72c42d963d5a42d7b45a8251bdbc69b0aefd7ae82a3d17ed69e86bc7d226fefb4e35caf70b8d7183691d1cf8aaf967fabcbe66")
	err = e.SliceOfBytes(hashByte)
	if err != nil {
	   return err
	}
     	t := a.Value.(Box2)
	err = e.box2Internal(&t)
	if err != nil {
	   return err
	}
     
     case nil:
     	err = e.SliceOfBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
	if err != nil {
	   return err
	}
	err = e.uint8(pNIL)
	if err != nil {
	   return err
	}
     case int8:
     	err = e.SliceOfBytes([]byte{0xc2, 0x71, 0x6e, 0x3c, 0x34, 0x13, 0xc0, 0xbe, 0xb3, 0x9e, 0x2e, 0xbc, 0xc7, 0x99, 0x62, 0x3f, 0x28, 0xe9, 0x9d, 0x19, 0x71, 0x6a, 0x0, 0x5b, 0x69, 0x4b, 0x1c, 0xbb, 0x3d, 0x8d, 0xcc, 0x45, 0xbd, 0x51, 0xca, 0x50, 0xc7, 0x3f, 0x8a, 0x64, 0xc8, 0xa0, 0xf0, 0xb, 0x46, 0x9a, 0x87, 0x5, 0xe9, 0x3c, 0xb9, 0x27, 0x91, 0xdf, 0x95, 0x88, 0x1b, 0x2e, 0xfa, 0x9f, 0xc0, 0x1f, 0xcb, 0xad}) 
	if err != nil {
	   return err
	}
	t := a.Value.(int8)
	err = e.int8(t)
	if err != nil {
	   return err
	}
     case uint8:
     	err = e.SliceOfBytes([]byte{0x9, 0x66, 0x78, 0xbb, 0x4a, 0x86, 0x63, 0x8d, 0x8e, 0xc6, 0x58, 0x8e, 0xcc, 0x2d, 0x89, 0x5a, 0x5c, 0x17, 0xb4, 0x86, 0x37, 0x8c, 0x81, 0xc2, 0xf2, 0xac, 0xf2, 0x48, 0x67, 0x57, 0x82, 0x13, 0x3e, 0x2e, 0x7d, 0x80, 0xb6, 0x66, 0x8d, 0x84, 0xc4, 0x55, 0xf4, 0xc9, 0xe1, 0x33, 0xfc, 0x71, 0xb, 0x77, 0x43, 0x63, 0x6, 0x82, 0x76, 0x63, 0x7, 0xbc, 0xf1, 0x64, 0xdb, 0xd9, 0x5c, 0x93})
	if err != nil {
	   return err
	}
	t := a.Value.(uint8)
	err = e.uint8(t)
	if err != nil {
	   return err
	}
     case int16:
     	err = e.SliceOfBytes([]byte{0x99, 0x7f, 0x69, 0x9c, 0x17, 0xfd, 0x6, 0x74, 0x8c, 0x2d, 0xba, 0xc4, 0x61, 0x1, 0x35, 0x21, 0xf1, 0x97, 0xd8, 0x1e, 0x74, 0x3c, 0x2f, 0x96, 0x56, 0xe6, 0xdc, 0xfc, 0x14, 0x1e, 0xd, 0x83, 0x36, 0xdc, 0x73, 0x36, 0xb0, 0xf4, 0x9e, 0x40, 0x2b, 0xfe, 0x97, 0x6f, 0xfa, 0xa5, 0x27, 0xac, 0xe1, 0xa2, 0x57, 0x2a, 0xae, 0x6d, 0x18, 0x22, 0xe2, 0xdc, 0xd8, 0x79, 0xe0, 0xb6, 0xf6, 0x7e})
	if err != nil {
	   return err
	}
	t := a.Value.(int16)
	err = e.int16(t)
	if err != nil {
	   return err
	}
     case uint16:
     	err = e.SliceOfBytes([]byte{0xa9, 0x1f, 0x92, 0xf4, 0xe9, 0x96, 0xb6, 0xd4, 0xa4, 0xcb, 0x85, 0x8d, 0x11, 0x80, 0x20, 0x10, 0xd6, 0x29, 0xba, 0x29, 0xe1, 0x89, 0x50, 0x2c, 0xa0, 0xf2, 0xcb, 0x1d, 0x86, 0xb8, 0x27, 0x3b, 0x67, 0xf3, 0x35, 0x72, 0xd1, 0x78, 0x15, 0xcc, 0xb0, 0x94, 0x94, 0x6f, 0x2, 0xab, 0x2e, 0x46, 0xcd, 0x74, 0xea, 0xf5, 0x15, 0xa, 0x26, 0xdc, 0x4e, 0xf7, 0xd, 0x9f, 0x3a, 0x9c, 0x6e, 0x55})
	if err != nil {
	   return err
	}
	t := a.Value.(uint16)
	err = e.uint16(t)
	if err != nil {
	   return err
	}
     case int32:
     	err = e.SliceOfBytes([]byte{0xeb, 0xc1, 0x67, 0x8b, 0x6, 0x82, 0x70, 0x21, 0x38, 0xc2, 0xd9, 0x9e, 0x33, 0x22, 0xd1, 0xa8, 0xc7, 0x2e, 0x9b, 0x68, 0xe9, 0x41, 0x12, 0x0, 0x1e, 0x3e, 0x51, 0xa8, 0xf5, 0xd9, 0xfa, 0x34, 0xc, 0x44, 0x9c, 0x6, 0x6d, 0x9d, 0x4c, 0xe7, 0x2a, 0x6, 0xab, 0x75, 0x77, 0x5d, 0xdf, 0x28, 0x34, 0x88, 0x7c, 0x7e, 0x96, 0x97, 0xbb, 0x8a, 0x95, 0xfe, 0x7, 0x65, 0xf7, 0x7c, 0x7e, 0x4c})
	if err != nil {
	   return err
	}
	t := a.Value.(int32)
	err = e.int32(t)
	if err != nil {
	   return err
	}
     case uint32:
     	err = e.SliceOfBytes([]byte{0xce, 0xf4, 0x3a, 0x5, 0xae, 0x67, 0xd9, 0x73, 0xc2, 0xa2, 0x1d, 0xf8, 0xcd, 0xf9, 0xd2, 0xde, 0x69, 0x8d, 0xd, 0xb7, 0x61, 0xb9, 0x51, 0x22, 0x58, 0xed, 0x8f, 0xb1, 0x83, 0xf1, 0x5c, 0xff, 0x5b, 0x84, 0xe2, 0x14, 0xe, 0x10, 0x68, 0x3f, 0x7a, 0xd9, 0xa7, 0x8f, 0x5b, 0xe4, 0x9e, 0x4e, 0x0, 0x7d, 0xcb, 0xfb, 0xd1, 0x69, 0x59, 0x9d, 0xbf, 0x9b, 0x75, 0x65, 0x15, 0x9e, 0x8b, 0x82})
	if err != nil {
	   return err
	}
	t := a.Value.(uint32)
	err = e.uint32(t)
	if err != nil {
	   return err
	}
     case int64:
     	err = e.SliceOfBytes([]byte{0x56, 0x16, 0x70, 0x35, 0xd0, 0x9, 0x18, 0x69, 0xe, 0xae, 0xad, 0x60, 0xd1, 0xee, 0x39, 0xa8, 0x61, 0x45, 0x58, 0x5b, 0x99, 0x20, 0x94, 0x57, 0x1f, 0xb0, 0x48, 0xeb, 0xb2, 0xcf, 0x5c, 0xa5, 0x8d, 0xc7, 0x8e, 0x7e, 0x3c, 0x89, 0xcd, 0x2f, 0xdc, 0xf2, 0x1c, 0x2a, 0xe3, 0xd2, 0x7f, 0x98, 0xc2, 0xad, 0x1c, 0x3d, 0x4e, 0x62, 0xd9, 0xdb, 0xc8, 0xc8, 0x59, 0xc5, 0xd5, 0xc6, 0xed, 0x7a})
	if err != nil {
	   return err
	}
	t := a.Value.(int64)
	err = e.int64(t)
	if err != nil {
	   return err
	}
     case uint64:
     	err = e.SliceOfBytes([]byte{0x11, 0x26, 0xb3, 0xd, 0x51, 0x59, 0x87, 0x5e, 0xd, 0x5b, 0x93, 0xfc, 0x92, 0xf0, 0x78, 0xaa, 0x12, 0xac, 0x93, 0xb8, 0x30, 0x1f, 0x48, 0xe, 0x13, 0x4d, 0x8b, 0xfb, 0x4c, 0x58, 0xfa, 0x3a, 0x69, 0x6a, 0x81, 0x1, 0xc5, 0x47, 0xc1, 0x55, 0x43, 0x95, 0x41, 0xdf, 0x3c, 0x8e, 0xb6, 0x96, 0x4a, 0x3c, 0x88, 0xab, 0x3f, 0x88, 0xed, 0x37, 0x5f, 0x8, 0x4a, 0x41, 0x8e, 0xd5, 0xda, 0x1e})
	if err != nil {
	   return err
	}
	t := a.Value.(uint64)
	err = e.uint64(t)
	if err != nil {
	   return err
	}
     case float32:
     	err = e.SliceOfBytes([]byte{0x42, 0x36, 0xae, 0xd3, 0x62, 0xca, 0x34, 0x75, 0x94, 0x52, 0xf0, 0x5f, 0x44, 0x83, 0x61, 0x75, 0x69, 0x39, 0xcf, 0x69, 0x74, 0x91, 0xee, 0x8d, 0x35, 0x8c, 0xd7, 0xa1, 0x63, 0xf, 0x88, 0x86, 0x6b, 0x52, 0xdd, 0x6d, 0xe1, 0xb2, 0x26, 0xf4, 0x3a, 0x9c, 0x9e, 0xf1, 0x56, 0xd, 0xf1, 0x48, 0x7, 0x39, 0x46, 0xf8, 0xe9, 0xd3, 0xab, 0x86, 0xe0, 0x1c, 0x98, 0xd, 0x17, 0x6b, 0x2, 0x63})
	if err != nil {
	   return err
	}
	t := a.Value.(float32)
	err = e.float32(t)
	if err != nil {
	   return err
	}
     case float64:
     	err = e.SliceOfBytes([]byte{0x65, 0x9b, 0xb2, 0x59, 0x85, 0xe2, 0x60, 0xe7, 0x1e, 0x12, 0x17, 0x3f, 0xc3, 0x1f, 0x20, 0x45, 0x8, 0x9e, 0x7e, 0x11, 0x6b, 0xaa, 0xb3, 0x1e, 0x6d, 0x7d, 0x7a, 0x5b, 0xe3, 0x3d, 0x40, 0xb5, 0x40, 0x6, 0x52, 0x85, 0x37, 0x80, 0x2c, 0xd8, 0x7d, 0x48, 0x67, 0xe3, 0x9a, 0xdd, 0xc9, 0x13, 0x11, 0x2c, 0xa5, 0xcc, 0x5a, 0x33, 0xbc, 0x35, 0x6b, 0x3e, 0xa8, 0x75, 0x93, 0x84, 0xcf, 0x1b})
	if err != nil {
	   return err
	}
	t := a.Value.(float64)
	err = e.float64(t)
	if err != nil {
	   return err
	}
     case bool:
     	err = e.SliceOfBytes([]byte{0x3e, 0x76, 0x6, 0x81, 0x73, 0x27, 0x61, 0xed, 0x16, 0x8e, 0xa3, 0x8, 0xe1, 0x10, 0x11, 0x85, 0xe1, 0xbd, 0x39, 0x2, 0xee, 0x67, 0x60, 0x21, 0x6a, 0x59, 0xc4, 0x7, 0x5a, 0x99, 0xc1, 0x46, 0xb7, 0xcd, 0x98, 0x14, 0xce, 0x14, 0x47, 0xe, 0xb0, 0x80, 0x6d, 0x91, 0x66, 0x50, 0xb0, 0xe5, 0xe, 0x77, 0x6f, 0x53, 0xe5, 0xd1, 0x72, 0x28, 0x1d, 0xd0, 0xe1, 0x70, 0x43, 0xc8, 0x65, 0x39})
	if err != nil {
	   return err
	}
	t := a.Value.(bool)
	err = e.bool(t)
	if err != nil {
	   return err
	}
     case string:
     	err = e.SliceOfBytes([]byte{0x27, 0x57, 0xcb, 0x3c, 0xaf, 0xc3, 0x9a, 0xf4, 0x51, 0xab, 0xb2, 0x69, 0x7b, 0xe7, 0x9b, 0x4a, 0xb6, 0x1d, 0x63, 0xd7, 0x4d, 0x85, 0xb0, 0x41, 0x86, 0x29, 0xde, 0x8c, 0x26, 0x81, 0x1b, 0x52, 0x9f, 0x3f, 0x37, 0x80, 0xd0, 0x15, 0x0, 0x63, 0xff, 0x55, 0xa2, 0xbe, 0xee, 0x74, 0xc4, 0xec, 0x10, 0x2a, 0x2a, 0x27, 0x31, 0xa1, 0xf1, 0xf7, 0xf1, 0xd, 0x47, 0x3a, 0xd1, 0x8a, 0x6a, 0x87})
	if err != nil {
	   return err
	}
	t := a.Value.(string)
	err = e.string(t)
	if err != nil {
	   return err
	}
     default:
	v := reflect.ValueOf(a)
	e := fmt.Errorf("Wrong type used as Any " + v.Type().Name())
	return e
     }
     return nil
}

func (e *Encoder) Any(a Any) (err os.Error){
     return e.any(a)
}

func (d *Decoder) any() (retValue *Any, error os.Error) {
     l, err := d.uint32()
     if err != nil {
     	return nil, err
     }
     hashValue := make([]byte, *l)
     err = d.SliceOfBytes(hashValue, *l)
     if err != nil {
     	return nil, err
     }
     d.indexToValue = append(d.indexToValue, retValue)	
     index := len(d.indexToValue) - 1
     
     encodedHash := []byte(hex.EncodeToString(hashValue))
     
     switch {
     
     case bytes.Equal(encodedHash, []byte("0b597097877469841d234bd0630c40a2a79061ed999d2c271da54b487688c92523deaaa2f954ab65295b001e73a2ccc84b625f27e5dfe0096d726592f0baf780")):
     	
	p, err := d.box1()	
	if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     
     case bytes.Equal(encodedHash, []byte("f64056a437617f591da3009d3e72c42d963d5a42d7b45a8251bdbc69b0aefd7ae82a3d17ed69e86bc7d226fefb4e35caf70b8d7183691d1cf8aaf967fabcbe66")):
     	
	p, err := d.box2()	
	if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     
     case bytes.Equal(hashValue, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}):
     	  _, err  := d.uint8()
	  if err != nil {
     	   return nil, err
     	   }
     	  retValue.Value = nil
     case bytes.Equal(hashValue, []byte{0xc2, 0x71, 0x6e, 0x3c, 0x34, 0x13, 0xc0, 0xbe, 0xb3, 0x9e, 0x2e, 0xbc, 0xc7, 0x99, 0x62, 0x3f, 0x28, 0xe9, 0x9d, 0x19, 0x71, 0x6a, 0x0, 0x5b, 0x69, 0x4b, 0x1c, 0xbb, 0x3d, 0x8d, 0xcc, 0x45, 0xbd, 0x51, 0xca, 0x50, 0xc7, 0x3f, 0x8a, 0x64, 0xc8, 0xa0, 0xf0, 0xb, 0x46, 0x9a, 0x87, 0x5, 0xe9, 0x3c, 0xb9, 0x27, 0x91, 0xdf, 0x95, 0x88, 0x1b, 0x2e, 0xfa, 0x9f, 0xc0, 0x1f, 0xcb, 0xad}):
     	  p, err := d.int8()
	  if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x9, 0x66, 0x78, 0xbb, 0x4a, 0x86, 0x63, 0x8d, 0x8e, 0xc6, 0x58, 0x8e, 0xcc, 0x2d, 0x89, 0x5a, 0x5c, 0x17, 0xb4, 0x86, 0x37, 0x8c, 0x81, 0xc2, 0xf2, 0xac, 0xf2, 0x48, 0x67, 0x57, 0x82, 0x13, 0x3e, 0x2e, 0x7d, 0x80, 0xb6, 0x66, 0x8d, 0x84, 0xc4, 0x55, 0xf4, 0xc9, 0xe1, 0x33, 0xfc, 0x71, 0xb, 0x77, 0x43, 0x63, 0x6, 0x82, 0x76, 0x63, 0x7, 0xbc, 0xf1, 0x64, 0xdb, 0xd9, 0x5c, 0x93}):
     	  p, err := d.uint8()
	  if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x99, 0x7f, 0x69, 0x9c, 0x17, 0xfd, 0x6, 0x74, 0x8c, 0x2d, 0xba, 0xc4, 0x61, 0x1, 0x35, 0x21, 0xf1, 0x97, 0xd8, 0x1e, 0x74, 0x3c, 0x2f, 0x96, 0x56, 0xe6, 0xdc, 0xfc, 0x14, 0x1e, 0xd, 0x83, 0x36, 0xdc, 0x73, 0x36, 0xb0, 0xf4, 0x9e, 0x40, 0x2b, 0xfe, 0x97, 0x6f, 0xfa, 0xa5, 0x27, 0xac, 0xe1, 0xa2, 0x57, 0x2a, 0xae, 0x6d, 0x18, 0x22, 0xe2, 0xdc, 0xd8, 0x79, 0xe0, 0xb6, 0xf6, 0x7e}):
     	  p, err := d.int16()
	  if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0xa9, 0x1f, 0x92, 0xf4, 0xe9, 0x96, 0xb6, 0xd4, 0xa4, 0xcb, 0x85, 0x8d, 0x11, 0x80, 0x20, 0x10, 0xd6, 0x29, 0xba, 0x29, 0xe1, 0x89, 0x50, 0x2c, 0xa0, 0xf2, 0xcb, 0x1d, 0x86, 0xb8, 0x27, 0x3b, 0x67, 0xf3, 0x35, 0x72, 0xd1, 0x78, 0x15, 0xcc, 0xb0, 0x94, 0x94, 0x6f, 0x2, 0xab, 0x2e, 0x46, 0xcd, 0x74, 0xea, 0xf5, 0x15, 0xa, 0x26, 0xdc, 0x4e, 0xf7, 0xd, 0x9f, 0x3a, 0x9c, 0x6e, 0x55}):
     	  p, err := d.uint16()
	  if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0xeb, 0xc1, 0x67, 0x8b, 0x6, 0x82, 0x70, 0x21, 0x38, 0xc2, 0xd9, 0x9e, 0x33, 0x22, 0xd1, 0xa8, 0xc7, 0x2e, 0x9b, 0x68, 0xe9, 0x41, 0x12, 0x0, 0x1e, 0x3e, 0x51, 0xa8, 0xf5, 0xd9, 0xfa, 0x34, 0xc, 0x44, 0x9c, 0x6, 0x6d, 0x9d, 0x4c, 0xe7, 0x2a, 0x6, 0xab, 0x75, 0x77, 0x5d, 0xdf, 0x28, 0x34, 0x88, 0x7c, 0x7e, 0x96, 0x97, 0xbb, 0x8a, 0x95, 0xfe, 0x7, 0x65, 0xf7, 0x7c, 0x7e, 0x4c}):
     	  p, err := d.int32()
	  if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0xce, 0xf4, 0x3a, 0x5, 0xae, 0x67, 0xd9, 0x73, 0xc2, 0xa2, 0x1d, 0xf8, 0xcd, 0xf9, 0xd2, 0xde, 0x69, 0x8d, 0xd, 0xb7, 0x61, 0xb9, 0x51, 0x22, 0x58, 0xed, 0x8f, 0xb1, 0x83, 0xf1, 0x5c, 0xff, 0x5b, 0x84, 0xe2, 0x14, 0xe, 0x10, 0x68, 0x3f, 0x7a, 0xd9, 0xa7, 0x8f, 0x5b, 0xe4, 0x9e, 0x4e, 0x0, 0x7d, 0xcb, 0xfb, 0xd1, 0x69, 0x59, 0x9d, 0xbf, 0x9b, 0x75, 0x65, 0x15, 0x9e, 0x8b, 0x82}):
     	  p, err := d.uint32()
	  if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x56, 0x16, 0x70, 0x35, 0xd0, 0x9, 0x18, 0x69, 0xe, 0xae, 0xad, 0x60, 0xd1, 0xee, 0x39, 0xa8, 0x61, 0x45, 0x58, 0x5b, 0x99, 0x20, 0x94, 0x57, 0x1f, 0xb0, 0x48, 0xeb, 0xb2, 0xcf, 0x5c, 0xa5, 0x8d, 0xc7, 0x8e, 0x7e, 0x3c, 0x89, 0xcd, 0x2f, 0xdc, 0xf2, 0x1c, 0x2a, 0xe3, 0xd2, 0x7f, 0x98, 0xc2, 0xad, 0x1c, 0x3d, 0x4e, 0x62, 0xd9, 0xdb, 0xc8, 0xc8, 0x59, 0xc5, 0xd5, 0xc6, 0xed, 0x7a}):
     	  p, err := d.int64()
	  if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x11, 0x26, 0xb3, 0xd, 0x51, 0x59, 0x87, 0x5e, 0xd, 0x5b, 0x93, 0xfc, 0x92, 0xf0, 0x78, 0xaa, 0x12, 0xac, 0x93, 0xb8, 0x30, 0x1f, 0x48, 0xe, 0x13, 0x4d, 0x8b, 0xfb, 0x4c, 0x58, 0xfa, 0x3a, 0x69, 0x6a, 0x81, 0x1, 0xc5, 0x47, 0xc1, 0x55, 0x43, 0x95, 0x41, 0xdf, 0x3c, 0x8e, 0xb6, 0x96, 0x4a, 0x3c, 0x88, 0xab, 0x3f, 0x88, 0xed, 0x37, 0x5f, 0x8, 0x4a, 0x41, 0x8e, 0xd5, 0xda, 0x1e}):
     	  p, err := d.uint64()
	  if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x42, 0x36, 0xae, 0xd3, 0x62, 0xca, 0x34, 0x75, 0x94, 0x52, 0xf0, 0x5f, 0x44, 0x83, 0x61, 0x75, 0x69, 0x39, 0xcf, 0x69, 0x74, 0x91, 0xee, 0x8d, 0x35, 0x8c, 0xd7, 0xa1, 0x63, 0xf, 0x88, 0x86, 0x6b, 0x52, 0xdd, 0x6d, 0xe1, 0xb2, 0x26, 0xf4, 0x3a, 0x9c, 0x9e, 0xf1, 0x56, 0xd, 0xf1, 0x48, 0x7, 0x39, 0x46, 0xf8, 0xe9, 0xd3, 0xab, 0x86, 0xe0, 0x1c, 0x98, 0xd, 0x17, 0x6b, 0x2, 0x63}):
     	  p, err := d.float32()
	  if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x65, 0x9b, 0xb2, 0x59, 0x85, 0xe2, 0x60, 0xe7, 0x1e, 0x12, 0x17, 0x3f, 0xc3, 0x1f, 0x20, 0x45, 0x8, 0x9e, 0x7e, 0x11, 0x6b, 0xaa, 0xb3, 0x1e, 0x6d, 0x7d, 0x7a, 0x5b, 0xe3, 0x3d, 0x40, 0xb5, 0x40, 0x6, 0x52, 0x85, 0x37, 0x80, 0x2c, 0xd8, 0x7d, 0x48, 0x67, 0xe3, 0x9a, 0xdd, 0xc9, 0x13, 0x11, 0x2c, 0xa5, 0xcc, 0x5a, 0x33, 0xbc, 0x35, 0x6b, 0x3e, 0xa8, 0x75, 0x93, 0x84, 0xcf, 0x1b}):
     	  p, err := d.float64()
	  if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x3e, 0x76, 0x6, 0x81, 0x73, 0x27, 0x61, 0xed, 0x16, 0x8e, 0xa3, 0x8, 0xe1, 0x10, 0x11, 0x85, 0xe1, 0xbd, 0x39, 0x2, 0xee, 0x67, 0x60, 0x21, 0x6a, 0x59, 0xc4, 0x7, 0x5a, 0x99, 0xc1, 0x46, 0xb7, 0xcd, 0x98, 0x14, 0xce, 0x14, 0x47, 0xe, 0xb0, 0x80, 0x6d, 0x91, 0x66, 0x50, 0xb0, 0xe5, 0xe, 0x77, 0x6f, 0x53, 0xe5, 0xd1, 0x72, 0x28, 0x1d, 0xd0, 0xe1, 0x70, 0x43, 0xc8, 0x65, 0x39}):
     	  p, err := d.bool()
	  if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x27, 0x57, 0xcb, 0x3c, 0xaf, 0xc3, 0x9a, 0xf4, 0x51, 0xab, 0xb2, 0x69, 0x7b, 0xe7, 0x9b, 0x4a, 0xb6, 0x1d, 0x63, 0xd7, 0x4d, 0x85, 0xb0, 0x41, 0x86, 0x29, 0xde, 0x8c, 0x26, 0x81, 0x1b, 0x52, 0x9f, 0x3f, 0x37, 0x80, 0xd0, 0x15, 0x0, 0x63, 0xff, 0x55, 0xa2, 0xbe, 0xee, 0x74, 0xc4, 0xec, 0x10, 0x2a, 0x2a, 0x27, 0x31, 0xa1, 0xf1, 0xf7, 0xf1, 0xd, 0x47, 0x3a, 0xd1, 0x8a, 0x6a, 0x87}):
     	  p, err := d.string()
	  if err != nil {
     	     return nil, err
     	  }
     	retValue.Value = *(p)
     default:
	e := fmt.Errorf("Wrong type used as Any")
	return retValue, e
     }
     d.indexToValue[index] = retValue		
     return retValue, nil
}

func (d *Decoder) anyInternal() (retValue *Any, error os.Error) {
     // hashValue := d.string()
     l, err := d.uint32()
     if err != nil {
     	return nil, err
     }
     hashValue := make([]byte, *l)
     err = d.SliceOfBytes(hashValue, *l)
     if err != nil {
     	return nil, err
     }

     
     encodedHash := []byte(hex.EncodeToString(hashValue))     
     

     switch {
     
     case bytes.Equal(encodedHash, []byte("0b597097877469841d234bd0630c40a2a79061ed999d2c271da54b487688c92523deaaa2f954ab65295b001e73a2ccc84b625f27e5dfe0096d726592f0baf780")):
     	
	p, err := d.box1Internal()
	if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     
     case bytes.Equal(encodedHash, []byte("f64056a437617f591da3009d3e72c42d963d5a42d7b45a8251bdbc69b0aefd7ae82a3d17ed69e86bc7d226fefb4e35caf70b8d7183691d1cf8aaf967fabcbe66")):
     	
	p, err := d.box2Internal()
	if err != nil {
     	   return nil, err
     	}
     	retValue.Value = *(p)
     
     case bytes.Equal(hashValue, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}):
     _, err := d.uint8()
     if err != nil {
     	return nil, err
     }
     retValue.Value = nil
     case bytes.Equal(hashValue, []byte{0xc2, 0x71, 0x6e, 0x3c, 0x34, 0x13, 0xc0, 0xbe, 0xb3, 0x9e, 0x2e, 0xbc, 0xc7, 0x99, 0x62, 0x3f, 0x28, 0xe9, 0x9d, 0x19, 0x71, 0x6a, 0x0, 0x5b, 0x69, 0x4b, 0x1c, 0xbb, 0x3d, 0x8d, 0xcc, 0x45, 0xbd, 0x51, 0xca, 0x50, 0xc7, 0x3f, 0x8a, 0x64, 0xc8, 0xa0, 0xf0, 0xb, 0x46, 0x9a, 0x87, 0x5, 0xe9, 0x3c, 0xb9, 0x27, 0x91, 0xdf, 0x95, 0x88, 0x1b, 0x2e, 0xfa, 0x9f, 0xc0, 0x1f, 0xcb, 0xad}):
     	  p, err := d.int8()
	  if err != nil {
     	     return nil, err
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x9, 0x66, 0x78, 0xbb, 0x4a, 0x86, 0x63, 0x8d, 0x8e, 0xc6, 0x58, 0x8e, 0xcc, 0x2d, 0x89, 0x5a, 0x5c, 0x17, 0xb4, 0x86, 0x37, 0x8c, 0x81, 0xc2, 0xf2, 0xac, 0xf2, 0x48, 0x67, 0x57, 0x82, 0x13, 0x3e, 0x2e, 0x7d, 0x80, 0xb6, 0x66, 0x8d, 0x84, 0xc4, 0x55, 0xf4, 0xc9, 0xe1, 0x33, 0xfc, 0x71, 0xb, 0x77, 0x43, 0x63, 0x6, 0x82, 0x76, 0x63, 0x7, 0xbc, 0xf1, 0x64, 0xdb, 0xd9, 0x5c, 0x93}):
     	  p, err := d.uint8()
	  if err != nil {
     	     return nil, err
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x99, 0x7f, 0x69, 0x9c, 0x17, 0xfd, 0x6, 0x74, 0x8c, 0x2d, 0xba, 0xc4, 0x61, 0x1, 0x35, 0x21, 0xf1, 0x97, 0xd8, 0x1e, 0x74, 0x3c, 0x2f, 0x96, 0x56, 0xe6, 0xdc, 0xfc, 0x14, 0x1e, 0xd, 0x83, 0x36, 0xdc, 0x73, 0x36, 0xb0, 0xf4, 0x9e, 0x40, 0x2b, 0xfe, 0x97, 0x6f, 0xfa, 0xa5, 0x27, 0xac, 0xe1, 0xa2, 0x57, 0x2a, 0xae, 0x6d, 0x18, 0x22, 0xe2, 0xdc, 0xd8, 0x79, 0xe0, 0xb6, 0xf6, 0x7e}):
     	  p, err := d.int16()
	  if err != nil {
     	     return nil, err
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0xa9, 0x1f, 0x92, 0xf4, 0xe9, 0x96, 0xb6, 0xd4, 0xa4, 0xcb, 0x85, 0x8d, 0x11, 0x80, 0x20, 0x10, 0xd6, 0x29, 0xba, 0x29, 0xe1, 0x89, 0x50, 0x2c, 0xa0, 0xf2, 0xcb, 0x1d, 0x86, 0xb8, 0x27, 0x3b, 0x67, 0xf3, 0x35, 0x72, 0xd1, 0x78, 0x15, 0xcc, 0xb0, 0x94, 0x94, 0x6f, 0x2, 0xab, 0x2e, 0x46, 0xcd, 0x74, 0xea, 0xf5, 0x15, 0xa, 0x26, 0xdc, 0x4e, 0xf7, 0xd, 0x9f, 0x3a, 0x9c, 0x6e, 0x55}):
     	  p, err := d.uint16()
	  if err != nil {
     	     return nil, err
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0xeb, 0xc1, 0x67, 0x8b, 0x6, 0x82, 0x70, 0x21, 0x38, 0xc2, 0xd9, 0x9e, 0x33, 0x22, 0xd1, 0xa8, 0xc7, 0x2e, 0x9b, 0x68, 0xe9, 0x41, 0x12, 0x0, 0x1e, 0x3e, 0x51, 0xa8, 0xf5, 0xd9, 0xfa, 0x34, 0xc, 0x44, 0x9c, 0x6, 0x6d, 0x9d, 0x4c, 0xe7, 0x2a, 0x6, 0xab, 0x75, 0x77, 0x5d, 0xdf, 0x28, 0x34, 0x88, 0x7c, 0x7e, 0x96, 0x97, 0xbb, 0x8a, 0x95, 0xfe, 0x7, 0x65, 0xf7, 0x7c, 0x7e, 0x4c}):
     	  p, err := d.int32()
	  if err != nil {
     	     return nil, err
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0xce, 0xf4, 0x3a, 0x5, 0xae, 0x67, 0xd9, 0x73, 0xc2, 0xa2, 0x1d, 0xf8, 0xcd, 0xf9, 0xd2, 0xde, 0x69, 0x8d, 0xd, 0xb7, 0x61, 0xb9, 0x51, 0x22, 0x58, 0xed, 0x8f, 0xb1, 0x83, 0xf1, 0x5c, 0xff, 0x5b, 0x84, 0xe2, 0x14, 0xe, 0x10, 0x68, 0x3f, 0x7a, 0xd9, 0xa7, 0x8f, 0x5b, 0xe4, 0x9e, 0x4e, 0x0, 0x7d, 0xcb, 0xfb, 0xd1, 0x69, 0x59, 0x9d, 0xbf, 0x9b, 0x75, 0x65, 0x15, 0x9e, 0x8b, 0x82}):
     	  p, err := d.uint32()
	  if err != nil {
     	     return nil, err
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x56, 0x16, 0x70, 0x35, 0xd0, 0x9, 0x18, 0x69, 0xe, 0xae, 0xad, 0x60, 0xd1, 0xee, 0x39, 0xa8, 0x61, 0x45, 0x58, 0x5b, 0x99, 0x20, 0x94, 0x57, 0x1f, 0xb0, 0x48, 0xeb, 0xb2, 0xcf, 0x5c, 0xa5, 0x8d, 0xc7, 0x8e, 0x7e, 0x3c, 0x89, 0xcd, 0x2f, 0xdc, 0xf2, 0x1c, 0x2a, 0xe3, 0xd2, 0x7f, 0x98, 0xc2, 0xad, 0x1c, 0x3d, 0x4e, 0x62, 0xd9, 0xdb, 0xc8, 0xc8, 0x59, 0xc5, 0xd5, 0xc6, 0xed, 0x7a}):
     	  p, err := d.int64()
	  if err != nil {
     	     return nil, err
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x11, 0x26, 0xb3, 0xd, 0x51, 0x59, 0x87, 0x5e, 0xd, 0x5b, 0x93, 0xfc, 0x92, 0xf0, 0x78, 0xaa, 0x12, 0xac, 0x93, 0xb8, 0x30, 0x1f, 0x48, 0xe, 0x13, 0x4d, 0x8b, 0xfb, 0x4c, 0x58, 0xfa, 0x3a, 0x69, 0x6a, 0x81, 0x1, 0xc5, 0x47, 0xc1, 0x55, 0x43, 0x95, 0x41, 0xdf, 0x3c, 0x8e, 0xb6, 0x96, 0x4a, 0x3c, 0x88, 0xab, 0x3f, 0x88, 0xed, 0x37, 0x5f, 0x8, 0x4a, 0x41, 0x8e, 0xd5, 0xda, 0x1e}):
     	  p, err := d.uint64()
	  if err != nil {
     	     return nil, err
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x42, 0x36, 0xae, 0xd3, 0x62, 0xca, 0x34, 0x75, 0x94, 0x52, 0xf0, 0x5f, 0x44, 0x83, 0x61, 0x75, 0x69, 0x39, 0xcf, 0x69, 0x74, 0x91, 0xee, 0x8d, 0x35, 0x8c, 0xd7, 0xa1, 0x63, 0xf, 0x88, 0x86, 0x6b, 0x52, 0xdd, 0x6d, 0xe1, 0xb2, 0x26, 0xf4, 0x3a, 0x9c, 0x9e, 0xf1, 0x56, 0xd, 0xf1, 0x48, 0x7, 0x39, 0x46, 0xf8, 0xe9, 0xd3, 0xab, 0x86, 0xe0, 0x1c, 0x98, 0xd, 0x17, 0x6b, 0x2, 0x63}):
     	  p, err := d.float32()
	  if err != nil {
     	     return nil, err
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x65, 0x9b, 0xb2, 0x59, 0x85, 0xe2, 0x60, 0xe7, 0x1e, 0x12, 0x17, 0x3f, 0xc3, 0x1f, 0x20, 0x45, 0x8, 0x9e, 0x7e, 0x11, 0x6b, 0xaa, 0xb3, 0x1e, 0x6d, 0x7d, 0x7a, 0x5b, 0xe3, 0x3d, 0x40, 0xb5, 0x40, 0x6, 0x52, 0x85, 0x37, 0x80, 0x2c, 0xd8, 0x7d, 0x48, 0x67, 0xe3, 0x9a, 0xdd, 0xc9, 0x13, 0x11, 0x2c, 0xa5, 0xcc, 0x5a, 0x33, 0xbc, 0x35, 0x6b, 0x3e, 0xa8, 0x75, 0x93, 0x84, 0xcf, 0x1b}):
     	  p, err := d.float64()
	  if err != nil {
     	     return nil, err
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x3e, 0x76, 0x6, 0x81, 0x73, 0x27, 0x61, 0xed, 0x16, 0x8e, 0xa3, 0x8, 0xe1, 0x10, 0x11, 0x85, 0xe1, 0xbd, 0x39, 0x2, 0xee, 0x67, 0x60, 0x21, 0x6a, 0x59, 0xc4, 0x7, 0x5a, 0x99, 0xc1, 0x46, 0xb7, 0xcd, 0x98, 0x14, 0xce, 0x14, 0x47, 0xe, 0xb0, 0x80, 0x6d, 0x91, 0x66, 0x50, 0xb0, 0xe5, 0xe, 0x77, 0x6f, 0x53, 0xe5, 0xd1, 0x72, 0x28, 0x1d, 0xd0, 0xe1, 0x70, 0x43, 0xc8, 0x65, 0x39}):
     	  p, err := d.bool()
	  if err != nil {
     	     return nil, err
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x27, 0x57, 0xcb, 0x3c, 0xaf, 0xc3, 0x9a, 0xf4, 0x51, 0xab, 0xb2, 0x69, 0x7b, 0xe7, 0x9b, 0x4a, 0xb6, 0x1d, 0x63, 0xd7, 0x4d, 0x85, 0xb0, 0x41, 0x86, 0x29, 0xde, 0x8c, 0x26, 0x81, 0x1b, 0x52, 0x9f, 0x3f, 0x37, 0x80, 0xd0, 0x15, 0x0, 0x63, 0xff, 0x55, 0xa2, 0xbe, 0xee, 0x74, 0xc4, 0xec, 0x10, 0x2a, 0x2a, 0x27, 0x31, 0xa1, 0xf1, 0xf7, 0xf1, 0xd, 0x47, 0x3a, 0xd1, 0x8a, 0x6a, 0x87}):
     	  p, err := d.string()
	  if err != nil {
     	     return nil, err
	  }
     	retValue.Value = *(p)
     default:
	e := fmt.Errorf("Wrong type used as Any")
	return retValue, e
     }
     return retValue, nil
}

func (d *Decoder) Any() (*Any, os.Error) {
     return d.any()
}


type Any struct{Value interface{}}

const(
	pNIL = 0
	pIDX = 1
	pVAL = 2
	BufSize = 1024 * 50
)
type Encoder struct {
	w io.Writer
	buf []byte
	t *TypeTree
	m []interface{} // this is for storing maps
	curPos int
	bufSpace uint32
	bufStart uint32
	bufLen uint32
	count uint32
}
type Decoder struct {
	r io.Reader
	buf []byte
	indexToValue []interface{}
	m []interface{} // this is for storing maps
	curPos uint32
	bufStart uint32
	bufLen uint32
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w:w, buf:make([]byte, BufSize), t: NewTypeTree(), curPos:4, bufSpace:BufSize, bufStart:0, bufLen:0, count:0}
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r:r, buf:make([]byte, BufSize), curPos:0, bufStart:0, bufLen:0}
}

func (e *Encoder) MapCheck(t interface{}) (int, bool) {
	for index, entry := range e.m {
		if entry == t {
			return index, true
		}
	}
	return -1, false
}

func (e *Encoder) isEnoughSpace(length uint32) bool {
     if length <= e.bufSpace {
     	return true
     }    
     return false
}

func (d *Decoder) isEnoughData(length uint32) bool {
     if length <= d.bufLen {
     	return true
     }
     return false
}


func (d *Decoder) readAtLeast(length uint32) os.Error{
     if d.bufLen > 0 {
     	copy(d.buf, d.buf[d.bufStart:(d.bufStart + d.bufLen)])
     }
     n, err := io.ReadAtLeast(d.r, d.buf[d.bufLen:], int(length))
     if err != nil {
	return err
     }
     d.bufLen += uint32(n)
     d.bufStart = 0
     return nil
}

func (e *Encoder) byte(b byte) (os.Error){
     return e.uint8(uint8(b))
     
}

func (e *Encoder) Byte(b byte) (os.Error){
     return e.byte(b)
}

func (d *Decoder) byte() (b *byte, error os.Error) {
     v, err := d.uint8()
     if err != nil {
        return nil, err
     }
     value := byte(*v)
     return &value, err
}

func (d *Decoder) Byte() (b *byte, err os.Error) {
     return d.byte()
}

func (e *Encoder) uint8(u uint8) os.Error{
     	// e.write([]byte{byte(u)})
	// e.buf[e.curPos] = byte(u)
	// e.curPos++
	if !e.isEnoughSpace(1) {
	   err := e.Flush() 
	   if err != nil {
	      return err
	   }
	}
	e.buf[e.bufStart] = byte(u)
	e.bufStart = (e.bufStart + 1) % BufSize
	e.bufLen++
	e.bufSpace--
	return nil
}

func (e *Encoder) Uint8(u uint8) (os.Error){
     return e.uint8(u)
}

func (e *Encoder) int8(u int8) (os.Error){
	return e.uint8(uint8(u))
}

func (e *Encoder) Int8(u int8) (os.Error){
     return e.int8(u)
}

func (d *Decoder) uint8() (w *uint8, err os.Error) {
	if !d.isEnoughData(1) {
	   err = d.readAtLeast(1)	   
	   if err != nil {
	      return nil, err
	   }
	}
	v := uint8(d.buf[d.bufStart])
	d.bufStart = (d.bufStart + 1) % BufSize
	d.bufLen -= 1
	return &v, nil
}

func (d *Decoder) Uint8() (w *uint8, err os.Error) {
     return d.uint8()
}

func (d *Decoder) int8() (w *int8, error os.Error) {
     	v, err := d.uint8()
	if err != nil {
		return nil, err
	}
	r := int8(*v)
	return &r, err
}

func (d *Decoder) Int8() (w *int8, err os.Error) {
     return d.int8()
}

func (e *Encoder) uint16(u uint16) os.Error{
	if !e.isEnoughSpace(2) {
	   err := e.Flush() 
	   if err != nil {
	      return err
	   }
	}
	e.buf[e.bufStart] = byte(u)
	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 8)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.bufLen += 2
	e.bufSpace -= 2
	return nil
}

func (e *Encoder) Uint16(u uint16) (os.Error){
     return e.uint16(u)
}

func (e *Encoder) int16(u int16) (os.Error){

	return e.uint16(uint16(u))
}

func (e *Encoder) Int16(u int16) (os.Error){
     return e.int16(u)
}

func (d *Decoder) uint16() (w *uint16, err os.Error) {
	if !d.isEnoughData(2) {
	   err = d.readAtLeast(2)	   
	   if err != nil {
	      return nil, err
	   }
	}
	v := uint16(d.buf[d.bufStart]) | uint16(d.buf[d.bufStart + 1]) << 8
	d.bufStart = (d.bufStart + 2) % BufSize
	d.bufLen -= 2
	return &v, nil
}

func (d *Decoder) Uint16() (w *uint16, err os.Error) {
     return d.uint16()
}

func (d *Decoder) int16() (w *int16, error os.Error) {
     	v, err := d.uint16()
	if err != nil {
		return nil, err
	}
	r := int16(*v)
	return &r, err
}

func (d *Decoder) Int16() (w *int16, err os.Error) {
     return d.int16()
}

func (e *Encoder) uint32(u uint32) os.Error{
	if !e.isEnoughSpace(4) {
	   err := e.Flush()   		      
	   if err != nil {
	      return err
	   }
	}
	e.buf[e.bufStart] = byte(u)
	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 8)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 16)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 24)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.bufLen += 4
	e.bufSpace -= 4
	return nil
}

func (e *Encoder) Uint32(u uint32) (os.Error){
     return e.uint32(u)
}

func (e *Encoder) int32(u int32) (os.Error){
	return e.uint32(uint32(u))
}

func (e *Encoder) Int32(u int32) (os.Error){
     return e.int32(u)
}

func (d *Decoder) uint32() (w *uint32, err os.Error) {
	if !d.isEnoughData(4) {
	   err = d.readAtLeast(4)	   
	   if err != nil {
	      return nil, err
	   }
	}
	v := uint32(d.buf[d.bufStart]) | uint32(d.buf[d.bufStart + 1]) << 8 | uint32(d.buf[d.bufStart + 2]) << 16 | uint32(d.buf[d.bufStart + 3]) << 24
	d.bufStart = (d.bufStart + 4) % BufSize
	d.bufLen -= 4
	return &v, nil
}

func (d *Decoder) Uint32() (w *uint32, err os.Error) {
     return d.uint32()
}

func (d *Decoder) int32() (w *int32, error os.Error) {
     	v, err := d.uint32()
	r := int32(*v)
	return &r, err
}

func (d *Decoder) Int32() (w *int32, err os.Error) {
     return d.int32()
}

func (e *Encoder) uint64(u uint64) os.Error{
	if !e.isEnoughSpace(8) {
	   	err := e.Flush()    
		if err != nil {
	      	   return err
	      	}  
	}
	e.buf[e.bufStart] = byte(u)
	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 8)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 16)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 24)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 32)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 40)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 48)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 56)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.bufLen += 8
	e.bufSpace -= 8
	return nil
}

func (e *Encoder) Uint64(u uint64) (os.Error){
     return e.uint64(u)
}

func (e *Encoder) int64(u int64) (os.Error){
	return e.uint64(uint64(u))
}

func (e *Encoder) Int64(u int64) (os.Error){
     return e.int64(u)
}

func (d *Decoder) uint64() (w *uint64, err os.Error) {
	if !d.isEnoughData(8) {
	   err = d.readAtLeast(8)	   
	   if err != nil {
	      return nil, err
	   }
	}
	v := uint64(d.buf[d.bufStart]) | uint64(d.buf[d.bufStart + 1]) << 8 | uint64(d.buf[d.bufStart + 2]) << 16 | uint64(d.buf[d.bufStart + 3]) << 24 | uint64(d.buf[d.bufStart + 4]) << 32 | uint64(d.buf[d.bufStart + 5]) << 40 | uint64(d.buf[d.bufStart + 6]) << 48 | uint64(d.buf[d.bufStart + 7]) << 56
	d.bufStart = (d.bufStart + 8) % BufSize
	d.bufLen -= 8
	return &v, nil
}

func (d *Decoder) Uint64() (w *uint64, err os.Error) {
     return d.uint64()
}

func (d *Decoder) int64() (w *int64, error os.Error) {
	v, err := d.uint64()
	if err != nil {
		return nil, err
	}
	r := int64(*v)
	return &r, err
}

func (d *Decoder) Int64() (w *int64, err os.Error) {
     return d.int64()
}

func (e *Encoder) float32(u float32) (os.Error){
	return e.uint32(math.Float32bits(u))
}

func (e *Encoder) Float32(u float32) (os.Error){
     return e.float32(u)
}

func (d *Decoder) float32() (w *float32, error os.Error) {
     	v, err := d.uint32()
	if err != nil {
		return nil, err
	}
	r := math.Float32frombits(*v)
	return &r, err
}

func (d *Decoder) Float32() (w *float32, err os.Error) {
     return d.float32()
}

func (e *Encoder) float64(u float64) (os.Error){
	return e.uint64(math.Float64bits(u))
}

func (e *Encoder) Float64(u float64) (os.Error){
     return e.float64(u)
}

func (d *Decoder) float64() (w *float64, error os.Error) {
	v, err := d.uint64()
	if err != nil {
		return nil, err
	}
	r := math.Float64frombits(*v)
	return &r, err
}

func (d *Decoder) Float64() (w *float64, err os.Error) {
     return d.float64()
}

func (e *Encoder) bool(u bool) (err os.Error){
	if u {
		err = e.uint8(1)
	} else {
		err = e.uint8(0)
	}
	return err
}

func (e *Encoder) Bool(u bool) (os.Error){
     return e.bool(u)
}

func (d *Decoder) bool() (w *bool, error os.Error) {
	v, err := d.uint8()
	if err != nil {
		return nil, err
	}
	var u bool
	if *v == 1 {
		u = true
	} else {
		u = false
	}
	return &u, err
}

func (d *Decoder) Bool() (w *bool, err os.Error) {
     return d.bool()
}

func (e *Encoder) SliceOfBytes(u []byte) (err os.Error){
     err = e.length(uint32(len(u)))
     if err != nil {
     	return err
     }
     sliceStartPos := uint32(0)
     for ;!e.isEnoughSpace(uint32(len(u[sliceStartPos:]))); {
	copy(e.buf[e.bufStart:], u[sliceStartPos:(sliceStartPos + e.bufSpace)])
	sliceStartPos += e.bufSpace
	e.bufLen += e.bufSpace
	e.bufSpace = 0
	if e.bufLen != 0 {
	   err := e.Flush()
	   if err != nil {
	      return err
	   }
	}
     } 
     if len(u[sliceStartPos:]) > 0 {
       copy(e.buf[e.bufStart:], u[sliceStartPos:])
       e.bufStart += uint32(len(u[sliceStartPos:]))	
       e.bufLen += uint32(len(u[sliceStartPos:]))
       e.bufSpace -= uint32(len(u[sliceStartPos:]))
     }
     return nil
     
}

func (d *Decoder) SliceOfBytes(v []byte, length uint32) (err os.Error){
     if length > d.bufLen {
     	copy(v, d.buf[d.bufStart:(d.bufStart + d.bufLen)])
	io.ReadFull(d.r, v[d.bufLen:])
	d.bufStart = 0
	d.bufLen = 0
	return
     }
     if !d.isEnoughData(length) {
	   err = d.readAtLeast(length)
	   if err != nil {
	      return err
	   }
	}
     copy(v, d.buf[d.bufStart: (d.bufStart + length)])
     d.bufStart = (d.bufStart + length) % BufSize
     d.bufLen -= length
     return nil
}

func (e *Encoder) string(u string) (err os.Error){
	err = e.length(uint32(len(u)))
	if err != nil {
	   return err
	}
	stringStartPos := uint32(0)
	for ;!e.isEnoughSpace(uint32(len(u[stringStartPos:]))); {
	   copy(e.buf[e.bufStart:], u[stringStartPos:(stringStartPos + e.bufSpace)])
	   stringStartPos += e.bufSpace
	   e.bufLen += e.bufSpace
	   e.bufSpace = 0
	   if e.bufSpace == 0 {
	      err = e.Flush()
	      if err != nil {
	      	 return err
	      }
	   }	  
	} 
	if len(u[stringStartPos:]) > 0 {
		copy(e.buf[e.bufStart:], u[stringStartPos:])
		e.bufStart += uint32(len(u[stringStartPos:]))
		e.bufLen += uint32(len(u[stringStartPos:]))
		e.bufSpace -= uint32(len(u[stringStartPos:]))
	}
	return nil
	
}

func (e *Encoder) String(u string) (os.Error){
     return e.string(u)
}

func (d *Decoder) string() (w *string, err os.Error) {
	len, err := d.length()
	if err != nil {
	   return nil, err
	}

	if len > d.bufLen {
 	   b := make([]byte, len) 
	   copy(b[0:], d.buf[d.bufStart:(d.bufStart + d.bufLen)])
	   _, err = io.ReadFull(d.r, b[d.bufLen:])
	   if err != nil {
	      return nil, err
	   }
	   d.bufStart = 0
	   d.bufLen = 0	   
	   str := string(b)
	   return &str, nil
	}
	if !d.isEnoughData(uint32(len)) {
	  	   err = d.readAtLeast(uint32(len))
		   if err != nil {
	      	      return nil, err
		   }  
	}
	b := d.buf[d.bufStart:(d.bufStart + len)]
	d.bufStart = (d.bufStart + len) % BufSize
	d.bufLen -= len
	str := string(b)
	return &str, nil
}

func (d *Decoder) String() (w *string, err os.Error) {
     return d.string()
}

func (e *Encoder) length(l uint32) (os.Error){
	// e.buf[e.curPos] = byte(l)
	// e.curPos++
	// e.buf[e.curPos] = byte(l >> 8)
	// e.curPos++
	// e.buf[e.curPos] = byte(l >> 16)
	// e.curPos++
	// e.buf[e.curPos] = byte(l >> 24)
	// e.curPos++
	// return
	return e.uint32(l)
}

func (d *Decoder) length() (l uint32, error os.Error) {
     	v, err := d.uint32()
	if err != nil {
		return 0, err
	}
	return *v, err
}

func (e *Encoder) Flush() os.Error{
	// e.bufSize = uint32(e.curPos) - 4
	// e.buf[0] = byte(e.bufSize)
	// e.buf[1] = byte(e.bufSize >> 8)
	// e.buf[2] = byte(e.bufSize >> 16)
	// e.buf[3] = byte(e.bufSize >> 24)
	// e.write(e.buf[:e.curPos])
	// e.curPos = 4
	
        // e.count++
	// if e.count %40 != 0 {
	//     return 
	// }
	if e.bufLen == 0 {
	   return nil
	}
	if _, err := e.w.Write(e.buf[:e.bufLen]); err != nil {
	   return err
	}
	e.bufStart = 0
	e.bufLen = 0
	e.bufSpace = BufSize
	e.reset()
	return nil
}

func (d *Decoder) ReadAll() {
     /* FIXME: Temporarily place reset here. We also reset in RPC calls. This 
      *        should probably go in one place: each type decoder, but we need
      *        to be able to differentiate between an application-called decode
      *        and a decode of internal elements. The latter case should not 
      *        call reset. */
     d.reset() 
}

func (e *Encoder) reset() {
	//e.t = NewTypeTree()
	e.t.Reset()
	e.m = make([]interface{}, 0)
}

func (d *Decoder) reset() {
	d.indexToValue = make([]interface{}, 0)
	d.m = make([]interface{}, 0)
	d.curPos = 0
}


func Hash(v interface{}) reflect.Type{
     return reflect.ValueOf(v).Type()
}

func Sizeof(v interface{}) uint64 {
     return uint64(unsafe.Sizeof(v))
}

// type Node struct {
// 	ptr unsafe.Pointer
// 	eleType string
// 	eleSize uint64
// 	index uint32
// }

var emptyPtrNode *llrb.Item

type TypeTree struct{
	tree *llrb.Tree
	Index uint32
	reusedItem llrb.Item
	min *llrb.Item
	max *llrb.Item
}

func lessPtr(a, b llrb.Item) bool {
	return uintptr(a.Ptr) < uintptr(b.Ptr)
}

func NewTypeTree() *TypeTree{
	t := TypeTree{}
	t.tree = llrb.New(lessPtr)
	t.Index = 0
	emptyPtrNode = &llrb.Item{0, "", 0, 0}
	return &t
}

func (t *TypeTree) Reset() {
     t.tree.Reset()
}

func (t *TypeTree) closestPtr(ptr uint64, typ string, size uint64, index uint32) (*llrb.Item, *llrb.Item){
	t.reusedItem.Ptr = ptr 
	t.reusedItem.EleType = typ
	t.reusedItem.EleSize = size
	t.reusedItem.Index = index
	minItem, maxItem := t.tree.FindAdjacentNodes(t.reusedItem)
	//minItem, maxItem = t.tree.FindAdjacentNodesLog(t.reusedItem)
	return minItem, maxItem
	
}

func (t *TypeTree) addToTree(elePtr uint64, eleType string, eleSize uint64, eleIndex uint32) {
	t.reusedItem.Ptr = elePtr 
	t.reusedItem.EleType = eleType
	t.reusedItem.EleSize = eleSize
	t.reusedItem.Index = eleIndex
	t.tree.InsertNoReplace(&t.reusedItem)
}

func (t *TypeTree) PointerCheck(ptr_unsafe unsafe.Pointer, typ string, size uint64) (index uint32, encoded bool, err os.Error) {
       	
	ptr := uint64(uintptr(ptr_unsafe))
	t.reusedItem.Ptr = ptr 
	t.reusedItem.EleType = typ
	t.reusedItem.EleSize = size
	t.reusedItem.Index = 0
	sameItem := t.tree.Get(t.reusedItem)
	t.min, t.max = t.closestPtr(ptr, sameItem.EleType, sameItem.EleSize, sameItem.Index) 
	switch {
	case !sameItem.Equal(emptyPtrNode)  && sameItem.EleType == typ:
		// already in the tree
		return sameItem.Index, true, nil
	case !sameItem.Equal(emptyPtrNode)  && sameItem.EleType != typ:
		t.addToTree(ptr, typ, size, t.Index)
		t.Index++
		return t.Index, false, nil
	case (t.min.Equal(emptyPtrNode) && !t.max.Equal(emptyPtrNode) && (ptr + size) <= t.max.Ptr) ||
	     (!t.min.Equal(emptyPtrNode) && t.max.Equal(emptyPtrNode) && ptr >= (t.min.Ptr + t.min.EleSize)) ||
	     (!t.min.Equal(emptyPtrNode) && !t.max.Equal(emptyPtrNode) && (ptr + size) <= t.max.Ptr && ptr >= (t.min.Ptr + t.min.EleSize)) ||
	     (t.min.Equal(emptyPtrNode) && t.max.Equal(emptyPtrNode)):
		t.addToTree(ptr, typ, size, t.Index)
		t.Index++
		return t.Index, false, nil
	default:
		e := fmt.Errorf("Illegal pointer")
		return 0, false, e
	}
	e := fmt.Errorf("Illegal pointer")
	return 0, false, e
}


/*type BufferedIO struct{
	writeBuf []byte
	readBuf []byte
	readBufCount int
	head int
	hasReadOnce bool
	conn net.Conn
	writeBufStart int
	writeBufLen int
}
var ReadBufSize = 1024*40

func NewBufferedIO(conn net.Conn) (bufIO *BufferedIO, error os.Error){
	//buff := make([]byte, 4096*1024)
	bufIO = &BufferedIO{}
	if conn == nil{
		return nil, os.EINVAL
	}
	bufIO.conn = conn
	bufIO.readBuf = make([]byte, ReadBufSize)
	bufIO.readBufCount = 0
	bufIO.head = 0
	bufIO.hasReadOnce = false
	bufIO.writeBuf = make([]byte, ReadBufSize)
	bufIO.writeBufStart = 0
	bufIO.writeBufLen = 0
	return bufIO, error
}

func (bo *BufferedIO) GetConn() net.Conn{
	return bo.conn
}

func (bo *BufferedIO) Write(b []byte) (n int, err os.Error){
	copy(bo.writeBuf[bo.writeBufStart:], b)
	bo.writeBufLen += len(b)
	bo.writeBufStart = bo.writeBufLen
	return bo.writeBufLen, nil
}

func (bo *BufferedIO) Flush() os.Error{

	if bo.writeBuf != nil{
		_, error := bo.conn.Write(bo.writeBuf[:bo.writeBufLen])
		bo.writeBufStart = 0
		bo.writeBufLen = 0
		return error
	}

	return nil
}

func DoMove (dest []byte, src *BufferedIO, size int) {
	copy(dest, src.readBuf[src.head:src.head+size])
	src.readBufCount -= size
	src.head += size
}

func (bo *BufferedIO) Read(b []byte) (n int, err os.Error){
	// read the smaller of the size of b, or what's left in bo
	//MyPrint("Read is called")
	copied :=  len(b)
	if bo.readBufCount < len(b) {
		copied = bo.readBufCount
	}
	remainder := len (b) - copied
	DoMove (b, bo, copied)
	//MyPrint ("READ", bo.readBuf[:copied])

	if remainder > ReadBufSize { // big read, do it directly
		n, err = bo.conn.Read (b[copied:])
		if err != nil {
			return n, err
		}
		//MyPrint("REAL READ", b[copied:])
	} else if remainder != 0 { // small read, buffer it
		n, err = io.ReadAtLeast (bo.conn, bo.readBuf, remainder)
		bo.readBufCount = n
		bo.head = 0
		//MyPrint("REAL READ", bo.readBuf)
		if err != nil {
			return n, err
		}
		DoMove (b[copied:], bo, remainder)
		//MyPrint("READ", b[copied:])
		//b = append (b, bo.readBuf[:remainder])
	}
	return len(b), err
}

func (bo *BufferedIO) Close() os.Error{
	var error os.Error

	if bo.writeBuf != nil{
		//MyPrint("Flush buffer before close", bo, bo.conn, bo.writeBuf)
		_, error = bo.conn.Write(bo.writeBuf[:bo.writeBufLen])
		bo.writeBufStart = 0
		bo.writeBufLen = 0
	}else{
		//MyPrint("No data to flush before close")
	}
	//MyPrint("CLOSE", bo.buf[:len(bo.buf)])
	if error != nil{
		bo.conn.Close()
		return error
	}
	if bo.conn != nil{
		return bo.conn.Close()
	}
	return nil
	//return bo.conn.Close()
}*/
func (t *Box1) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteBox1: %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.Box1(t)
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *Box1) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.Box1(t)
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Box1) WriteVar(path string) (syscall.Status) {
	fd, name, status := goodmiddleman.OpenPathLastDirectory(path)
	if status != syscall.StatusOk {
		return status
	}

	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.Box1(t)
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Box1) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.Box1()
	if err != nil {
		return syscall.StatusFail
	}

	t = value
	return syscall.StatusOk
}


func (t *Box1) ReadVar(path string) (syscall.Status) {
	fd, name, status := goodmiddleman.OpenPathLastDirectory(path)
	if status != syscall.StatusOk {
		return status
	}

	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.Box1()
	if err != nil {
		return syscall.StatusFail
	}

	t = value
	return syscall.StatusOk
}



func (t *Box1) CreateDirectory(fd syscall.Fd, name string, label string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x0b,0x59,0x70,0x97,0x87,0x74,0x69,0x84,0x1d,0x23,0x4b,0xd0,0x63,0x0c,0x40,0xa2,0xa7,0x90,0x61,0xed,0x99,0x9d,0x2c,0x27,0x1d,0xa5,0x4b,0x48,0x76,0x88,0xc9,0x25,0x23,0xde,0xaa,0xa2,0xf9,0x54,0xab,0x65,0x29,0x5b,0x00,0x1e,0x73,0xa2,0xcc,0xc8,0x4b,0x62,0x5f,0x27,0xe5,0xdf,0xe0,0x09,0x6d,0x72,0x65,0x92,0xf0,0xba,0xf7,0x80, }  
       return goodmiddleman.CreateDirectory(fd, name, label, hash)
}

func (t *Box1) CreateDirectoryPath(path string, label string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x0b,0x59,0x70,0x97,0x87,0x74,0x69,0x84,0x1d,0x23,0x4b,0xd0,0x63,0x0c,0x40,0xa2,0xa7,0x90,0x61,0xed,0x99,0x9d,0x2c,0x27,0x1d,0xa5,0x4b,0x48,0x76,0x88,0xc9,0x25,0x23,0xde,0xaa,0xa2,0xf9,0x54,0xab,0x65,0x29,0x5b,0x00,0x1e,0x73,0xa2,0xcc,0xc8,0x4b,0x62,0x5f,0x27,0xe5,0xdf,0xe0,0x09,0x6d,0x72,0x65,0x92,0xf0,0xba,0xf7,0x80, }
       return goodmiddleman.CreateDirectoryPath(path, label, hash)
}
func (t *Box2) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteBox2: %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.Box2(t)
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *Box2) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.Box2(t)
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Box2) WriteVar(path string) (syscall.Status) {
	fd, name, status := goodmiddleman.OpenPathLastDirectory(path)
	if status != syscall.StatusOk {
		return status
	}

	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.Box2(t)
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Box2) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.Box2()
	if err != nil {
		return syscall.StatusFail
	}

	t = value
	return syscall.StatusOk
}


func (t *Box2) ReadVar(path string) (syscall.Status) {
	fd, name, status := goodmiddleman.OpenPathLastDirectory(path)
	if status != syscall.StatusOk {
		return status
	}

	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.Box2()
	if err != nil {
		return syscall.StatusFail
	}

	t = value
	return syscall.StatusOk
}



func (t *Box2) CreateDirectory(fd syscall.Fd, name string, label string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0xf6,0x40,0x56,0xa4,0x37,0x61,0x7f,0x59,0x1d,0xa3,0x00,0x9d,0x3e,0x72,0xc4,0x2d,0x96,0x3d,0x5a,0x42,0xd7,0xb4,0x5a,0x82,0x51,0xbd,0xbc,0x69,0xb0,0xae,0xfd,0x7a,0xe8,0x2a,0x3d,0x17,0xed,0x69,0xe8,0x6b,0xc7,0xd2,0x26,0xfe,0xfb,0x4e,0x35,0xca,0xf7,0x0b,0x8d,0x71,0x83,0x69,0x1d,0x1c,0xf8,0xaa,0xf9,0x67,0xfa,0xbc,0xbe,0x66, }  
       return goodmiddleman.CreateDirectory(fd, name, label, hash)
}

func (t *Box2) CreateDirectoryPath(path string, label string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0xf6,0x40,0x56,0xa4,0x37,0x61,0x7f,0x59,0x1d,0xa3,0x00,0x9d,0x3e,0x72,0xc4,0x2d,0x96,0x3d,0x5a,0x42,0xd7,0xb4,0x5a,0x82,0x51,0xbd,0xbc,0x69,0xb0,0xae,0xfd,0x7a,0xe8,0x2a,0x3d,0x17,0xed,0x69,0xe8,0x6b,0xc7,0xd2,0x26,0xfe,0xfb,0x4e,0x35,0xca,0xf7,0x0b,0x8d,0x71,0x83,0x69,0x1d,0x1c,0xf8,0xaa,0xf9,0x67,0xfa,0xbc,0xbe,0x66, }
       return goodmiddleman.CreateDirectoryPath(path, label, hash)
}
type Int8 int8

func (t *Int8) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteInt8: %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.Int8(int8(*t))
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *Int8) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.Int8(int8(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Int8) WriteVar(fd syscall.Fd, name string) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.Int8(int8(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Int8) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.Int8()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Int8(*value)
	return syscall.StatusOk
}


func (t *Int8) ReadVar(fd syscall.Fd, name string) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.Int8()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Int8(*value)
	return syscall.StatusOk
}



func (t *Int8) CreateDirectory(fd syscall.Fd, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0xc2,0x71,0x6e,0x3c,0x34,0x13,0xc0,0xbe,0xb3,0x9e,0x2e,0xbc,0xc7,0x99,0x62,0x3f,0x28,0xe9,0x9d,0x19,0x71,0x6a,0x00,0x5b,0x69,0x4b,0x1c,0xbb,0x3d,0x8d,0xcc,0x45,0xbd,0x51,0xca,0x50,0xc7,0x3f,0x8a,0x64,0xc8,0xa0,0xf0,0x0b,0x46,0x9a,0x87,0x05,0xe9,0x3c,0xb9,0x27,0x91,0xdf,0x95,0x88,0x1b,0x2e,0xfa,0x9f,0xc0,0x1f,0xcb,0xad, }  
       return goodmiddleman.CreateDirectory(fd, name, "", hash)
}

func (t *Int8) CreateDirectoryPath(path string, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0xc2,0x71,0x6e,0x3c,0x34,0x13,0xc0,0xbe,0xb3,0x9e,0x2e,0xbc,0xc7,0x99,0x62,0x3f,0x28,0xe9,0x9d,0x19,0x71,0x6a,0x00,0x5b,0x69,0x4b,0x1c,0xbb,0x3d,0x8d,0xcc,0x45,0xbd,0x51,0xca,0x50,0xc7,0x3f,0x8a,0x64,0xc8,0xa0,0xf0,0x0b,0x46,0x9a,0x87,0x05,0xe9,0x3c,0xb9,0x27,0x91,0xdf,0x95,0x88,0x1b,0x2e,0xfa,0x9f,0xc0,0x1f,0xcb,0xad, }
       return goodmiddleman.CreateDirectoryPath(path,"", hash)
}
type Uint8 uint8

func (t *Uint8) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteUint8: %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.Uint8(uint8(*t))
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *Uint8) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.Uint8(uint8(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Uint8) WriteVar(fd syscall.Fd, name string) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.Uint8(uint8(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Uint8) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.Uint8()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Uint8(*value)
	return syscall.StatusOk
}


func (t *Uint8) ReadVar(fd syscall.Fd, name string) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.Uint8()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Uint8(*value)
	return syscall.StatusOk
}



func (t *Uint8) CreateDirectory(fd syscall.Fd, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x09,0x66,0x78,0xbb,0x4a,0x86,0x63,0x8d,0x8e,0xc6,0x58,0x8e,0xcc,0x2d,0x89,0x5a,0x5c,0x17,0xb4,0x86,0x37,0x8c,0x81,0xc2,0xf2,0xac,0xf2,0x48,0x67,0x57,0x82,0x13,0x3e,0x2e,0x7d,0x80,0xb6,0x66,0x8d,0x84,0xc4,0x55,0xf4,0xc9,0xe1,0x33,0xfc,0x71,0x0b,0x77,0x43,0x63,0x06,0x82,0x76,0x63,0x07,0xbc,0xf1,0x64,0xdb,0xd9,0x5c,0x93, }  
       return goodmiddleman.CreateDirectory(fd, name, "", hash)
}

func (t *Uint8) CreateDirectoryPath(path string, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x09,0x66,0x78,0xbb,0x4a,0x86,0x63,0x8d,0x8e,0xc6,0x58,0x8e,0xcc,0x2d,0x89,0x5a,0x5c,0x17,0xb4,0x86,0x37,0x8c,0x81,0xc2,0xf2,0xac,0xf2,0x48,0x67,0x57,0x82,0x13,0x3e,0x2e,0x7d,0x80,0xb6,0x66,0x8d,0x84,0xc4,0x55,0xf4,0xc9,0xe1,0x33,0xfc,0x71,0x0b,0x77,0x43,0x63,0x06,0x82,0x76,0x63,0x07,0xbc,0xf1,0x64,0xdb,0xd9,0x5c,0x93, }
       return goodmiddleman.CreateDirectoryPath(path,"", hash)
}
type Bool bool

func (t *Bool) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteBool: %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.Bool(bool(*t))
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *Bool) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.Bool(bool(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Bool) WriteVar(fd syscall.Fd, name string) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.Bool(bool(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Bool) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.Bool()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Bool(*value)
	return syscall.StatusOk
}


func (t *Bool) ReadVar(fd syscall.Fd, name string) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.Bool()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Bool(*value)
	return syscall.StatusOk
}



func (t *Bool) CreateDirectory(fd syscall.Fd, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x3e,0x76,0x06,0x81,0x73,0x27,0x61,0xed,0x16,0x8e,0xa3,0x08,0xe1,0x10,0x11,0x85,0xe1,0xbd,0x39,0x02,0xee,0x67,0x60,0x21,0x6a,0x59,0xc4,0x07,0x5a,0x99,0xc1,0x46,0xb7,0xcd,0x98,0x14,0xce,0x14,0x47,0x0e,0xb0,0x80,0x6d,0x91,0x66,0x50,0xb0,0xe5,0x0e,0x77,0x6f,0x53,0xe5,0xd1,0x72,0x28,0x1d,0xd0,0xe1,0x70,0x43,0xc8,0x65,0x39, }  
       return goodmiddleman.CreateDirectory(fd, name, "", hash)
}

func (t *Bool) CreateDirectoryPath(path string, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x3e,0x76,0x06,0x81,0x73,0x27,0x61,0xed,0x16,0x8e,0xa3,0x08,0xe1,0x10,0x11,0x85,0xe1,0xbd,0x39,0x02,0xee,0x67,0x60,0x21,0x6a,0x59,0xc4,0x07,0x5a,0x99,0xc1,0x46,0xb7,0xcd,0x98,0x14,0xce,0x14,0x47,0x0e,0xb0,0x80,0x6d,0x91,0x66,0x50,0xb0,0xe5,0x0e,0x77,0x6f,0x53,0xe5,0xd1,0x72,0x28,0x1d,0xd0,0xe1,0x70,0x43,0xc8,0x65,0x39, }
       return goodmiddleman.CreateDirectoryPath(path,"", hash)
}
type Int16 int16

func (t *Int16) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteInt16: %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.Int16(int16(*t))
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *Int16) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.Int16(int16(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Int16) WriteVar(fd syscall.Fd, name string) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.Int16(int16(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Int16) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.Int16()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Int16(*value)
	return syscall.StatusOk
}


func (t *Int16) ReadVar(fd syscall.Fd, name string) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.Int16()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Int16(*value)
	return syscall.StatusOk
}



func (t *Int16) CreateDirectory(fd syscall.Fd, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x99,0x7f,0x69,0x9c,0x17,0xfd,0x06,0x74,0x8c,0x2d,0xba,0xc4,0x61,0x01,0x35,0x21,0xf1,0x97,0xd8,0x1e,0x74,0x3c,0x2f,0x96,0x56,0xe6,0xdc,0xfc,0x14,0x1e,0x0d,0x83,0x36,0xdc,0x73,0x36,0xb0,0xf4,0x9e,0x40,0x2b,0xfe,0x97,0x6f,0xfa,0xa5,0x27,0xac,0xe1,0xa2,0x57,0x2a,0xae,0x6d,0x18,0x22,0xe2,0xdc,0xd8,0x79,0xe0,0xb6,0xf6,0x7e, }  
       return goodmiddleman.CreateDirectory(fd, name, "", hash)
}

func (t *Int16) CreateDirectoryPath(path string, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x99,0x7f,0x69,0x9c,0x17,0xfd,0x06,0x74,0x8c,0x2d,0xba,0xc4,0x61,0x01,0x35,0x21,0xf1,0x97,0xd8,0x1e,0x74,0x3c,0x2f,0x96,0x56,0xe6,0xdc,0xfc,0x14,0x1e,0x0d,0x83,0x36,0xdc,0x73,0x36,0xb0,0xf4,0x9e,0x40,0x2b,0xfe,0x97,0x6f,0xfa,0xa5,0x27,0xac,0xe1,0xa2,0x57,0x2a,0xae,0x6d,0x18,0x22,0xe2,0xdc,0xd8,0x79,0xe0,0xb6,0xf6,0x7e, }
       return goodmiddleman.CreateDirectoryPath(path,"", hash)
}
type Uint16 uint16

func (t *Uint16) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteUint16: %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.Uint16(uint16(*t))
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *Uint16) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.Uint16(uint16(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Uint16) WriteVar(fd syscall.Fd, name string) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.Uint16(uint16(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Uint16) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.Uint16()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Uint16(*value)
	return syscall.StatusOk
}


func (t *Uint16) ReadVar(fd syscall.Fd, name string) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.Uint16()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Uint16(*value)
	return syscall.StatusOk
}



func (t *Uint16) CreateDirectory(fd syscall.Fd, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0xa9,0x1f,0x92,0xf4,0xe9,0x96,0xb6,0xd4,0xa4,0xcb,0x85,0x8d,0x11,0x80,0x20,0x10,0xd6,0x29,0xba,0x29,0xe1,0x89,0x50,0x2c,0xa0,0xf2,0xcb,0x1d,0x86,0xb8,0x27,0x3b,0x67,0xf3,0x35,0x72,0xd1,0x78,0x15,0xcc,0xb0,0x94,0x94,0x6f,0x02,0xab,0x2e,0x46,0xcd,0x74,0xea,0xf5,0x15,0x0a,0x26,0xdc,0x4e,0xf7,0x0d,0x9f,0x3a,0x9c,0x6e,0x55, }  
       return goodmiddleman.CreateDirectory(fd, name, "", hash)
}

func (t *Uint16) CreateDirectoryPath(path string, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0xa9,0x1f,0x92,0xf4,0xe9,0x96,0xb6,0xd4,0xa4,0xcb,0x85,0x8d,0x11,0x80,0x20,0x10,0xd6,0x29,0xba,0x29,0xe1,0x89,0x50,0x2c,0xa0,0xf2,0xcb,0x1d,0x86,0xb8,0x27,0x3b,0x67,0xf3,0x35,0x72,0xd1,0x78,0x15,0xcc,0xb0,0x94,0x94,0x6f,0x02,0xab,0x2e,0x46,0xcd,0x74,0xea,0xf5,0x15,0x0a,0x26,0xdc,0x4e,0xf7,0x0d,0x9f,0x3a,0x9c,0x6e,0x55, }
       return goodmiddleman.CreateDirectoryPath(path,"", hash)
}
type Int32 int32

func (t *Int32) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteInt32: %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.Int32(int32(*t))
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *Int32) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.Int32(int32(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Int32) WriteVar(fd syscall.Fd, name string) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.Int32(int32(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Int32) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.Int32()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Int32(*value)
	return syscall.StatusOk
}


func (t *Int32) ReadVar(fd syscall.Fd, name string) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.Int32()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Int32(*value)
	return syscall.StatusOk
}



func (t *Int32) CreateDirectory(fd syscall.Fd, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0xeb,0xc1,0x67,0x8b,0x06,0x82,0x70,0x21,0x38,0xc2,0xd9,0x9e,0x33,0x22,0xd1,0xa8,0xc7,0x2e,0x9b,0x68,0xe9,0x41,0x12,0x00,0x1e,0x3e,0x51,0xa8,0xf5,0xd9,0xfa,0x34,0x0c,0x44,0x9c,0x06,0x6d,0x9d,0x4c,0xe7,0x2a,0x06,0xab,0x75,0x77,0x5d,0xdf,0x28,0x34,0x88,0x7c,0x7e,0x96,0x97,0xbb,0x8a,0x95,0xfe,0x07,0x65,0xf7,0x7c,0x7e,0x4c, }  
       return goodmiddleman.CreateDirectory(fd, name, "", hash)
}

func (t *Int32) CreateDirectoryPath(path string, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0xeb,0xc1,0x67,0x8b,0x06,0x82,0x70,0x21,0x38,0xc2,0xd9,0x9e,0x33,0x22,0xd1,0xa8,0xc7,0x2e,0x9b,0x68,0xe9,0x41,0x12,0x00,0x1e,0x3e,0x51,0xa8,0xf5,0xd9,0xfa,0x34,0x0c,0x44,0x9c,0x06,0x6d,0x9d,0x4c,0xe7,0x2a,0x06,0xab,0x75,0x77,0x5d,0xdf,0x28,0x34,0x88,0x7c,0x7e,0x96,0x97,0xbb,0x8a,0x95,0xfe,0x07,0x65,0xf7,0x7c,0x7e,0x4c, }
       return goodmiddleman.CreateDirectoryPath(path,"", hash)
}
type Uint32 uint32

func (t *Uint32) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteUint32: %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.Uint32(uint32(*t))
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *Uint32) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.Uint32(uint32(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Uint32) WriteVar(fd syscall.Fd, name string) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.Uint32(uint32(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Uint32) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.Uint32()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Uint32(*value)
	return syscall.StatusOk
}


func (t *Uint32) ReadVar(fd syscall.Fd, name string) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.Uint32()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Uint32(*value)
	return syscall.StatusOk
}



func (t *Uint32) CreateDirectory(fd syscall.Fd, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0xce,0xf4,0x3a,0x05,0xae,0x67,0xd9,0x73,0xc2,0xa2,0x1d,0xf8,0xcd,0xf9,0xd2,0xde,0x69,0x8d,0x0d,0xb7,0x61,0xb9,0x51,0x22,0x58,0xed,0x8f,0xb1,0x83,0xf1,0x5c,0xff,0x5b,0x84,0xe2,0x14,0x0e,0x10,0x68,0x3f,0x7a,0xd9,0xa7,0x8f,0x5b,0xe4,0x9e,0x4e,0x00,0x7d,0xcb,0xfb,0xd1,0x69,0x59,0x9d,0xbf,0x9b,0x75,0x65,0x15,0x9e,0x8b,0x82, }  
       return goodmiddleman.CreateDirectory(fd, name, "", hash)
}

func (t *Uint32) CreateDirectoryPath(path string, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0xce,0xf4,0x3a,0x05,0xae,0x67,0xd9,0x73,0xc2,0xa2,0x1d,0xf8,0xcd,0xf9,0xd2,0xde,0x69,0x8d,0x0d,0xb7,0x61,0xb9,0x51,0x22,0x58,0xed,0x8f,0xb1,0x83,0xf1,0x5c,0xff,0x5b,0x84,0xe2,0x14,0x0e,0x10,0x68,0x3f,0x7a,0xd9,0xa7,0x8f,0x5b,0xe4,0x9e,0x4e,0x00,0x7d,0xcb,0xfb,0xd1,0x69,0x59,0x9d,0xbf,0x9b,0x75,0x65,0x15,0x9e,0x8b,0x82, }
       return goodmiddleman.CreateDirectoryPath(path,"", hash)
}
type Int64 int64

func (t *Int64) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteInt64: %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.Int64(int64(*t))
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *Int64) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.Int64(int64(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Int64) WriteVar(fd syscall.Fd, name string) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.Int64(int64(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Int64) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.Int64()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Int64(*value)
	return syscall.StatusOk
}


func (t *Int64) ReadVar(fd syscall.Fd, name string) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.Int64()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Int64(*value)
	return syscall.StatusOk
}



func (t *Int64) CreateDirectory(fd syscall.Fd, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x56,0x16,0x70,0x35,0xd0,0x09,0x18,0x69,0x0e,0xae,0xad,0x60,0xd1,0xee,0x39,0xa8,0x61,0x45,0x58,0x5b,0x99,0x20,0x94,0x57,0x1f,0xb0,0x48,0xeb,0xb2,0xcf,0x5c,0xa5,0x8d,0xc7,0x8e,0x7e,0x3c,0x89,0xcd,0x2f,0xdc,0xf2,0x1c,0x2a,0xe3,0xd2,0x7f,0x98,0xc2,0xad,0x1c,0x3d,0x4e,0x62,0xd9,0xdb,0xc8,0xc8,0x59,0xc5,0xd5,0xc6,0xed,0x7a, }  
       return goodmiddleman.CreateDirectory(fd, name, "", hash)
}

func (t *Int64) CreateDirectoryPath(path string, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x56,0x16,0x70,0x35,0xd0,0x09,0x18,0x69,0x0e,0xae,0xad,0x60,0xd1,0xee,0x39,0xa8,0x61,0x45,0x58,0x5b,0x99,0x20,0x94,0x57,0x1f,0xb0,0x48,0xeb,0xb2,0xcf,0x5c,0xa5,0x8d,0xc7,0x8e,0x7e,0x3c,0x89,0xcd,0x2f,0xdc,0xf2,0x1c,0x2a,0xe3,0xd2,0x7f,0x98,0xc2,0xad,0x1c,0x3d,0x4e,0x62,0xd9,0xdb,0xc8,0xc8,0x59,0xc5,0xd5,0xc6,0xed,0x7a, }
       return goodmiddleman.CreateDirectoryPath(path,"", hash)
}
type Uint64 uint64

func (t *Uint64) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteUint64: %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.Uint64(uint64(*t))
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *Uint64) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.Uint64(uint64(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Uint64) WriteVar(fd syscall.Fd, name string) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.Uint64(uint64(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Uint64) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.Uint64()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Uint64(*value)
	return syscall.StatusOk
}


func (t *Uint64) ReadVar(fd syscall.Fd, name string) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.Uint64()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Uint64(*value)
	return syscall.StatusOk
}



func (t *Uint64) CreateDirectory(fd syscall.Fd, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x11,0x26,0xb3,0x0d,0x51,0x59,0x87,0x5e,0x0d,0x5b,0x93,0xfc,0x92,0xf0,0x78,0xaa,0x12,0xac,0x93,0xb8,0x30,0x1f,0x48,0x0e,0x13,0x4d,0x8b,0xfb,0x4c,0x58,0xfa,0x3a,0x69,0x6a,0x81,0x01,0xc5,0x47,0xc1,0x55,0x43,0x95,0x41,0xdf,0x3c,0x8e,0xb6,0x96,0x4a,0x3c,0x88,0xab,0x3f,0x88,0xed,0x37,0x5f,0x08,0x4a,0x41,0x8e,0xd5,0xda,0x1e, }  
       return goodmiddleman.CreateDirectory(fd, name, "", hash)
}

func (t *Uint64) CreateDirectoryPath(path string, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x11,0x26,0xb3,0x0d,0x51,0x59,0x87,0x5e,0x0d,0x5b,0x93,0xfc,0x92,0xf0,0x78,0xaa,0x12,0xac,0x93,0xb8,0x30,0x1f,0x48,0x0e,0x13,0x4d,0x8b,0xfb,0x4c,0x58,0xfa,0x3a,0x69,0x6a,0x81,0x01,0xc5,0x47,0xc1,0x55,0x43,0x95,0x41,0xdf,0x3c,0x8e,0xb6,0x96,0x4a,0x3c,0x88,0xab,0x3f,0x88,0xed,0x37,0x5f,0x08,0x4a,0x41,0x8e,0xd5,0xda,0x1e, }
       return goodmiddleman.CreateDirectoryPath(path,"", hash)
}
type Float32 float32

func (t *Float32) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteFloat32: %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.Float32(float32(*t))
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *Float32) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.Float32(float32(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Float32) WriteVar(fd syscall.Fd, name string) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.Float32(float32(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Float32) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.Float32()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Float32(*value)
	return syscall.StatusOk
}


func (t *Float32) ReadVar(fd syscall.Fd, name string) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.Float32()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Float32(*value)
	return syscall.StatusOk
}



func (t *Float32) CreateDirectory(fd syscall.Fd, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x42,0x36,0xae,0xd3,0x62,0xca,0x34,0x75,0x94,0x52,0xf0,0x5f,0x44,0x83,0x61,0x75,0x69,0x39,0xcf,0x69,0x74,0x91,0xee,0x8d,0x35,0x8c,0xd7,0xa1,0x63,0x0f,0x88,0x86,0x6b,0x52,0xdd,0x6d,0xe1,0xb2,0x26,0xf4,0x3a,0x9c,0x9e,0xf1,0x56,0x0d,0xf1,0x48,0x07,0x39,0x46,0xf8,0xe9,0xd3,0xab,0x86,0xe0,0x1c,0x98,0x0d,0x17,0x6b,0x02,0x63, }  
       return goodmiddleman.CreateDirectory(fd, name, "", hash)
}

func (t *Float32) CreateDirectoryPath(path string, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x42,0x36,0xae,0xd3,0x62,0xca,0x34,0x75,0x94,0x52,0xf0,0x5f,0x44,0x83,0x61,0x75,0x69,0x39,0xcf,0x69,0x74,0x91,0xee,0x8d,0x35,0x8c,0xd7,0xa1,0x63,0x0f,0x88,0x86,0x6b,0x52,0xdd,0x6d,0xe1,0xb2,0x26,0xf4,0x3a,0x9c,0x9e,0xf1,0x56,0x0d,0xf1,0x48,0x07,0x39,0x46,0xf8,0xe9,0xd3,0xab,0x86,0xe0,0x1c,0x98,0x0d,0x17,0x6b,0x02,0x63, }
       return goodmiddleman.CreateDirectoryPath(path,"", hash)
}
type Float64 float64

func (t *Float64) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteFloat64: %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.Float64(float64(*t))
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *Float64) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.Float64(float64(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Float64) WriteVar(fd syscall.Fd, name string) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.Float64(float64(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Float64) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.Float64()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Float64(*value)
	return syscall.StatusOk
}


func (t *Float64) ReadVar(fd syscall.Fd, name string) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.Float64()
	if err != nil {
		return syscall.StatusFail
	}

	*t = Float64(*value)
	return syscall.StatusOk
}



func (t *Float64) CreateDirectory(fd syscall.Fd, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x65,0x9b,0xb2,0x59,0x85,0xe2,0x60,0xe7,0x1e,0x12,0x17,0x3f,0xc3,0x1f,0x20,0x45,0x08,0x9e,0x7e,0x11,0x6b,0xaa,0xb3,0x1e,0x6d,0x7d,0x7a,0x5b,0xe3,0x3d,0x40,0xb5,0x40,0x06,0x52,0x85,0x37,0x80,0x2c,0xd8,0x7d,0x48,0x67,0xe3,0x9a,0xdd,0xc9,0x13,0x11,0x2c,0xa5,0xcc,0x5a,0x33,0xbc,0x35,0x6b,0x3e,0xa8,0x75,0x93,0x84,0xcf,0x1b, }  
       return goodmiddleman.CreateDirectory(fd, name, "", hash)
}

func (t *Float64) CreateDirectoryPath(path string, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x65,0x9b,0xb2,0x59,0x85,0xe2,0x60,0xe7,0x1e,0x12,0x17,0x3f,0xc3,0x1f,0x20,0x45,0x08,0x9e,0x7e,0x11,0x6b,0xaa,0xb3,0x1e,0x6d,0x7d,0x7a,0x5b,0xe3,0x3d,0x40,0xb5,0x40,0x06,0x52,0x85,0x37,0x80,0x2c,0xd8,0x7d,0x48,0x67,0xe3,0x9a,0xdd,0xc9,0x13,0x11,0x2c,0xa5,0xcc,0x5a,0x33,0xbc,0x35,0x6b,0x3e,0xa8,0x75,0x93,0x84,0xcf,0x1b, }
       return goodmiddleman.CreateDirectoryPath(path,"", hash)
}
type String string

func (t *String) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteString: %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.String(string(*t))
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *String) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.String(string(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *String) WriteVar(fd syscall.Fd, name string) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.String(string(*t))
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *String) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.String()
	if err != nil {
		return syscall.StatusFail
	}

	*t = String(*value)
	return syscall.StatusOk
}


func (t *String) ReadVar(fd syscall.Fd, name string) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.String()
	if err != nil {
		return syscall.StatusFail
	}

	*t = String(*value)
	return syscall.StatusOk
}



func (t *String) CreateDirectory(fd syscall.Fd, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x27,0x57,0xcb,0x3c,0xaf,0xc3,0x9a,0xf4,0x51,0xab,0xb2,0x69,0x7b,0xe7,0x9b,0x4a,0xb6,0x1d,0x63,0xd7,0x4d,0x85,0xb0,0x41,0x86,0x29,0xde,0x8c,0x26,0x81,0x1b,0x52,0x9f,0x3f,0x37,0x80,0xd0,0x15,0x00,0x63,0xff,0x55,0xa2,0xbe,0xee,0x74,0xc4,0xec,0x10,0x2a,0x2a,0x27,0x31,0xa1,0xf1,0xf7,0xf1,0x0d,0x47,0x3a,0xd1,0x8a,0x6a,0x87, }  
       return goodmiddleman.CreateDirectory(fd, name, "", hash)
}

func (t *String) CreateDirectoryPath(path string, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{ 0x27,0x57,0xcb,0x3c,0xaf,0xc3,0x9a,0xf4,0x51,0xab,0xb2,0x69,0x7b,0xe7,0x9b,0x4a,0xb6,0x1d,0x63,0xd7,0x4d,0x85,0xb0,0x41,0x86,0x29,0xde,0x8c,0x26,0x81,0x1b,0x52,0x9f,0x3f,0x37,0x80,0xd0,0x15,0x00,0x63,0xff,0x55,0xa2,0xbe,0xee,0x74,0xc4,0xec,0x10,0x2a,0x2a,0x27,0x31,0xa1,0xf1,0xf7,0xf1,0x0d,0x47,0x3a,0xd1,0x8a,0x6a,0x87, }
       return goodmiddleman.CreateDirectoryPath(path,"", hash)
}
func (t *Any) IpcWrite(serviceName string, hostname string) (syscall.Fd, syscall.Status) {
	serviceFd, status := goodmiddleman.OpenDirectoryPath("/services/" + serviceName)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling OpenDirectory /services/%s: %v\n", serviceName, status)
		return 0, status
	}

	sName := serviceName
	for i:=len(serviceName)-1; i>=0; i-- {
		if (serviceName[i]=='/') {
			sName = serviceName[i+1:]
		}
	}


	netFd, status := goodmiddleman.IpcRepeat(serviceFd, sName, hostname, nil)
	if status != syscall.StatusOk {
		log.Fatalf ("Error calling IpcWriteAny %v\n", status)
		return 0, status
	}
	writer := goodmiddleman.NewWriter(netFd)
	goodmiddleman.Close(serviceFd)
	e := NewEncoder(writer)
	err := e.Any(*t)
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return 0, syscall.StatusWriteFail
	}

	return netFd, syscall.StatusOk
}



func (t *Any) Write(fd syscall.Fd) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewWriter(fd))
	err := e.Any(*t)
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Any) WriteVar(fd syscall.Fd, name string) (syscall.Status) {
	e := NewEncoder(goodmiddleman.NewVarWriter(fd, name))
	err := e.Any(*t)
	if err != nil {
		return syscall.StatusWriteFail
	}

	err = e.Flush()
	if err != nil {
		return syscall.StatusWriteFail
	}

	return syscall.StatusOk
}


func (t *Any) Read(fd syscall.Fd) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewReader(fd))
	value, err := d.Any()
	if err != nil {
		return syscall.StatusFail
	}

	t.Value = value.Value
	return syscall.StatusOk
}


func (t *Any) ReadVar(fd syscall.Fd, name string) (syscall.Status) {
	d := NewDecoder(goodmiddleman.NewVarReader(fd, name))
	value, err := d.Any()
	if err != nil {
		return syscall.StatusFail
	}

	t.Value = value.Value
	return syscall.StatusOk
}



func (t *Any) CreateDirectory(fd syscall.Fd, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{0x40, 0xd3, 0x80, 0xd9, 0x2f, 0x53, 0xad, 0x12, 0xcf, 0x21, 0x94, 0x59, 0x68, 0x74, 0xa9, 0x17, 0x9f, 0x1e, 0xe3, 0xf9, 0x2e, 0x8f, 0x7c, 0x99, 0x4b, 0xf9, 0x4d, 0xb3, 0x29, 0x1a, 0xbb, 0x89, 0xc3, 0xff, 0x35, 0x1e, 0xd2, 0xb9, 0x11, 0x30, 0x15, 0x7f, 0xc7, 0xd3, 0x2f, 0x84, 0x2c, 0xed, 0x4b, 0x99, 0x8a, 0x9d, 0xe6, 0xe0, 0xe0, 0x1d, 0x98, 0x7a, 0x28, 0xd9, 0x69, 0x34, 0xe6, 0xcc}  
       return goodmiddleman.CreateDirectory(fd, name, "", hash)
}

func (t *Any) CreateDirectoryPath(path string, name string) (syscall.Status) {
       hash := goodmiddleman.HashValue{0x40, 0xd3, 0x80, 0xd9, 0x2f, 0x53, 0xad, 0x12, 0xcf, 0x21, 0x94, 0x59, 0x68, 0x74, 0xa9, 0x17, 0x9f, 0x1e, 0xe3, 0xf9, 0x2e, 0x8f, 0x7c, 0x99, 0x4b, 0xf9, 0x4d, 0xb3, 0x29, 0x1a, 0xbb, 0x89, 0xc3, 0xff, 0x35, 0x1e, 0xd2, 0xb9, 0x11, 0x30, 0x15, 0x7f, 0xc7, 0xd3, 0x2f, 0x84, 0x2c, 0xed, 0x4b, 0x99, 0x8a, 0x9d, 0xe6, 0xe0, 0xe0, 0x1d, 0x98, 0x7a, 0x28, 0xd9, 0x69, 0x34, 0xe6, 0xcc}
       return goodmiddleman.CreateDirectoryPath(path,"", hash)
}
