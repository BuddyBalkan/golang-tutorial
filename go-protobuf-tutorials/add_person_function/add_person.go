package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	// pb "github.com/protocolbuffers/protobuf/examples/go/tutorialpb"
	pb "kk.com/go-protobuf-tutorial/examples"
	"google.golang.org/protobuf/proto"
)

func promptForAddress(r io.Reader)(*pb.Person, error){
	// A procol buffer can be created like any struct.
	// 一个protobuf的对象可以被任意创建
	p := &pb.Person{}

	rd := bufio.NewReader(r)
	fmt.Print("Enter person ID number:")
	// An int32 field in the .proto file is represented as an int32 field in the generated Go struct.
	// .proto文件中的int32类型即go语言中的int32类型
	// 通过交互键入id值
	if _, err := fmt.Fscanf(rd, "%d\n", &p.Id); err != nil {
		return p, err
	}

	// 交互式键入name值
	fmt.Print("Enter name:")
	name, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}

	// A string field in .proto file results in a string field in go.
	// We trim the whitespace because rd.ReadString includes the trailing newline character in its output.
	p.Name = strings.TrimSpace(name)


	// 交互式键入email值
	fmt.Print("Enter email address (blank for none): ")
	email, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}
	p.Email = strings.TrimSpace(email)

	for {
		fmt.Print("Enter a phone number (or leave blank to finish): ")
		phone, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}
		phone = strings.TrimSpace(phone)
		if phone == ""{
			break
		}

		// The PhoneNumber message type is nested within the Person message in the .proto file. This results in a Go struct named using the name of the parent
		// prefixed to the name of the nested message. Just as with pb.Person, it can be created like any other struct.
		pn := &pb.Person_PhoneNumber{
			Number: phone,
		}

		fmt.Print("Is this a mobile, home, or work phone?")
		ptype, err := rd.ReadString('\n')
		if err != nil{
			return p, err
		}
		ptype = strings.TrimSpace(ptype)

		// A  proto enum results in a Go constant for each enum value.
		switch ptype {
		case "mobile":
			pn.Type = pb.Person_MOBILE
		case "home":
			pn.Type = pb.Person_HOME
		case "work":
			pn.Type = pb.Person_WORK
		default:
			fmt.Printf("Unknown phone type %q. Using default. \n", ptype)
		}

		// A repeated proto field maps to a slice field in Go. We can append to it like any other slice.
		p.Phones = append(p.Phones, pn)
	}
	return p, nil
}


// Main reads the entire address book from a file, adds one person based on user input, the writes it back out to the same file.
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	// Read the exiting address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err){
			fmt.Printf("%s: File Not Found. Creating new file. \n", fname)
		}else {
			log.Fatalln("Error reading file: ", err)
		}
	}

	// [Start marshal_proto]
	book := &pb.AddressBook{}
	// [Start_Exclude]
	if err := proto.Unmarshal(in, book); err != nil{
		log.Fatalln("Failed to parse address book: ", err)
	} 

	// Add an address.
	addr, err := promptForAddress(os.Stdin)
	if err != nil {
		log.Fatalln("Error with address: ", err)
	}
	book.People = append(book.People, addr)
	// [End_Exclude]

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book: ", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book: ", err)
	}
	// [End marshal_proto]
}