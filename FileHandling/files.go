package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	
	// lets start with reading a file in go here is a multiple way to do it but we will see the most common one

	// the first step is going to use the os package which is used to open the file in go...

	// f, err := os.Open("example.txt") // returns two values one is the file object and the other is the error if any

	// // in go we have a pattern to handling errors we dont just throe the error but we will return that error and handle it accordingly

	// if err != nil {
	// 	// log the error
	// 	panic(err) // abhi kuch nhi ho sakta hain panic mode 
	// }

	// fileInfo, err := f.Stat() // this will give us the information about the file like size , name , mode etc

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("File Name: ", fileInfo.Name())
	// fmt.Println("File Size: ", fileInfo.Size())
	// fmt.Println("File Mode: ", fileInfo.Mode())
	// fmt.Println("File ModTime: ", fileInfo.ModTime())
	// fmt.Println("Is Directory: ", fileInfo.IsDir())

	// read file---------------------------------------------------------2

	// f,err := os.Open("example.txt")

	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// // now look we need to close the file after we are done with it so we will use defer to close the file but why we are using here here is the response because if there is any error will came through out the program then programm will not reach to the close statement so to avoid that we are using defer here or jitna jaldi use kar sko kar dena

	// defer f.Close() // this will close the file after the main function is done executing

	// we need to store the data like we need to store that data inside a buffer so that we can read that 
    // buffer is nothing but a temporary space in a memory where we can store the data
	// buffer is a array of bytes

	// we need to use here file size by using stack to make it dynamic but lets says 100 for now

	// buff := make([]byte, 100) // this will create a buffer of 100 bytes and we are using make function to create an array

	// d, err := f.Read(buff) // this will read the file and store the data in the buffer and it will return the number of bytes read and error if any
	
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Data: ", string(buff)) // converting byte array to string
	// fmt.Println("Bytes Read: ", d) // number of bytes read

	// read file------------------------the simplest one but not efficient(for big files)---------------------------------3

	// data , err := os.ReadFile("example.txt") // this will read the file and return the data and error if any
	
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Data: ", string(data)) // converting byte array to string

 // read folder -------------------4---------------------------------4

	// dir, err := os.Open("../")
	
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close() // closing the directory after we are done with it. 


	// fileInfos, err := dir.ReadDir(-1) // retuerns a slice of FileInfo and error if any

	// for _, fileInfo := range fileInfos {
	// 	fmt.Println("File Name: ", fileInfo.Name())
	// 	fmt.Println("Is Directory: ", fileInfo.IsDir())
	// 	fmt.Println("File Type: ", fileInfo.Type())
	// 	fmt.Println("---------------------")
	// }

	// how can we create a file in go lang---------------------5---------------------------------5

	// file, err := os.Create("newFile.txt") // this will create a new file in the current directory

	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close() // closing the file after we are done with it.

	// // file.WriteString("Hello goLang!!!") // this will write the string to the file
	// // file.WriteString("\nThis is a new file created using go lang") // this will write the string to the file // append mode
	 
	
	// fmt.Println("File created successfully")

	// // file is nothing but a binary data so we can write anything in it like images , videos etc but we need to convert that data into byte array

	//  bytes := []byte("Hello goLang!!!\nThis is a new file created using go lang") // converting string to byte array

	//  file.Write(bytes) // this will write the byte array to the file

	//--------------------------------6---------------------------------6 now we are going to read data from one file and then insert that data into another file lets see how to do that: through streams

	// sourcefile,err := os.Open("example.txt") // opening the source file

	// if err != nil {
    //  panic(err)
	// }
	// defer sourcefile.Close() // closing the source file after we are done with it

	// destinationfile, err := os.Create("destination.txt") // creating the destination file
	
	// if err != nil {
	// 	panic(err)
	// }

	// defer destinationfile.Close() // closing the destination file after we are done with it

	// reader := bufio.NewReader(sourcefile) // creating a new reader
	// writer := bufio.NewWriter(destinationfile) // creating a new writer

	// for {

	// 	b, err := reader.ReadByte() // reading one byte at a time
		
	// 	if err != nil {
	// 		if err.Error() != "EOF" {    // this error will come when we reach the end of the file if there is nothing to read  
	// 			panic(err)
	// 		}
	// 		break // breaking the loop if we reach the end of the file
	// 	}

	// 	e:= writer.WriteByte(b) // writing the byte to the destination file
		 
	// 	if err != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush() // flushing the writer to write the data to the file

	// fmt.Println("File copied successfully")


	// // Lets see how can we delete the file in go.       ---------------------7---------------------------------7

	sourcefile,err := os.Open("example.txt") // opening the source file

	if err != nil {
     panic(err)
	}
	defer sourcefile.Close() // closing the source file after we are done with it


	er:= os.Remove("example.txt") // this will delete the file

	if er != nil {
		panic(er)
	}

	fmt.Println("File deleted successfully")

/// this is simple isn't it...
// bro please read doc as well and revision is neccessary....

}


// "use terminal bro go run go and cd FileHandling" to run the code
