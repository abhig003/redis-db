package core
import "errors"
//encode the command and its arguments into RESP format
//decode the RESP format into command and its arguments

//length finder
//simple string encoder and decoder
//bulk string encoder and ecoder
//array encoder and decoder
//integer encoder and decoder
//error encoder and decoder


func readLength(data []byte) (int , int){
	
	var len int = 0
	for pos:=range data{
		b:=data[pos]
		if !(b>='0' && b<='9'){
			return len,pos+2
		}
		len=len*10 + int(b-'0')
	}
    return 0,0

}



func readSimpleSTring(data []byte) (string ,int, error){//it will return the string an dthe next postion after \r\n
   // we have need to find the locakltion of \r and the string will be between first charecter and \r
    pos:=1
	for ;data[pos]!='\r';pos++{
	}
	return string(data[1:pos]),pos+2,nil


}



func readError(data []byte) (string ,int, error){//it will return the string an dthe next postion after \r\n
    return readSimpleSTring(data)



}


func readInt64(data []byte)(int64,int,error){//it will return the value and the next postion after \r\n

	pos:=1
	var num int64=0

	for ;data[pos]!='\r';pos++{

		b:=data[pos]

		if(b<'0' || b>'9'){
			return 0,0,errors.New("invalid integer format")
		}

		num=num*10+int64(data[pos]-'0')

	}
	return num,pos+2,nil
	}




func readBUlkstring(data []byte)(string, int, error){
    //$4\r\nhello\r\n
	pos:=1;
	len1,nextpos:=readLength(data[pos:])

	if nextpos+int(len1)+2> len(data){
		return "",0,errors.New("invalid bulk string format")
	}


	str:=string(data[pos+nextpos:nextpos+pos+int(len1)])

	return str,nextpos+pos+int(len1)+2,nil

}


func readArray(data []byte)(interface{},int,error){ 

	//*length\r\n{anything in RESP format}*length\r\n{anything in RESP format}*length\r\n{anything in RESP format}

    pos:=1

	len1,nextpos:=readLength(data[pos:])

     pos+=nextpos
	if nextpos+int(len1)+2> len(data){
		return nil,0,errors.New("invalid array format")
	}


	var elements []interface{}= make([]interface{},len1)
    
	for val:=range elements{
		element,nextpos,err:=decodeResp(data[pos:])
		if err!=nil{
			return nil,0,err
		}
		elements[val]=element
		pos+=nextpos
	}

	return elements,pos,nil

}


func decodeResp(data []byte)(interface{},int,error){
   if len(data)==0{
	return nil,0,errors.New("empty data")
   }

   switch data[0]{

   case '+':
	return readSimpleSTring(data);

   case '-':
	return readError(data)	

   case ':':
	return readInt64(data)	
 
   case '$':
	return readBUlkstring(data)

   case '*':
	return readArray(data)

   default:
	return nil,0,errors.New("invalid RESP format")


   }
}

