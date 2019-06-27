package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/afero"
	"io/ioutil"
	"log"
	"os"
)


var AppFs = afero.NewOsFs()
var scanner = bufio.NewScanner(os.Stdin)
var f afero.File
var n int
var path = "/home/parulraich/go/src/afero_filesystem"

func main() {


	n=takeUserinput()
	Implementation(n)
}
func getDirectoryName() (Dirname string){
	fmt.Println("input directory name")
	_,_=fmt.Scanln(&Dirname)
	if !fileExists(Dirname){
		err:=AppFs.MkdirAll("./"+Dirname,0777)
		if err != nil{
			log.Fatal(err)
		}
	}

	return
}

func getFileName() (filename string){
	fmt.Println("Enter the filename")
	_, _ = fmt.Scanln(&filename)
   return
}

func getString() (str string){
	fmt.Println("Enter the string to be appended")
	scanner.Scan()
	str = scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from input: ", err)
	}
	return
}

func createFile(){
	var stat string
	fmt.Println("do you want to create the file in a directory? \n press Y if yes and N if no")
	_,_=fmt.Scanln(&stat)
	if stat == "Y" || stat == "y"{
		Dirname := getDirectoryName()
		Filename := getFileName()
		path = path + "/" + Dirname + "/" + Filename
		if !fileExists(path){
			_,err := AppFs.Create(path)
			if err !=nil{
				log.Fatal(err)
			}
		}

	}else {
		var filename string
		fmt.Println("Enter the filename")
		_, _ = fmt.Scanln(&filename)

		//create a file
		if !fileExists(filename) {
			_, err := AppFs.Create(filename)
			if err != nil {
				panic(err)
			}
		}
	}

}


func editFile(){
	var stat string
	fmt.Println("do you want to edit file in a directory? \n press Y if yes and N if no")
	_,_=fmt.Scanln(&stat)
	if stat == "Y" || stat == "y"{
		var str string
		Dirname := getDirectoryName()
		Filename := getFileName()
		path1 := "/home/parulraich/go/src/afero_filesystem"
		path1 = path1 + "/" + Dirname + "/" + Filename
		fmt.Println(path)
		f,err := AppFs.OpenFile(path1, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
		if err!=nil{
			panic(err)
		}
		defer f.Close()
		str = getString()
		_, _ = f.Write([]byte(str))
		fmt.Println("File Updated")

	}else {
		filename := getFileName()
		f, erro := AppFs.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
		if erro != nil {
			panic(erro)
		}
		defer f.Close()

		//edit a file

		//str, _ = reader.ReadString('\n')
		var str string
		str = getString()
		_, _ = f.Write([]byte(str))
		fmt.Println("File Updated")
	}

}

func moveFile(){
		Dirname := getDirectoryName()
		Filename := getFileName()
		path_initial := path + "/" + Dirname + "/" + Filename
		var Dir_new string
		fmt.Println("enter new directory name")
		_,_ =fmt.Scanln(&Dir_new)
		path_change :=path + "/" + Dir_new + "/" + Filename
		if fileExists(path_initial){
			err :=AppFs.Rename(path_initial,path_change)
			if err!=nil{
				fmt.Println("file cannot be moved")
			}
		}else {
			fmt.Println("File Doesn't Exists")
		}

}

func copyFile(){
	Dirname := getDirectoryName()
	Filename := getFileName()
	path1 := "/home/parulraich/go/src/afero_filesystem"
	path_initial := path1 + "/" + Dirname + "/" + Filename
	content, err := ioutil.ReadFile(path_initial)
	if err != nil {
		log.Fatal(err)
	}
	var stat string
	fmt.Println("do you want to copy the file in the other file in same folder? if yes then type Y ,if no then type N")
	_,_=fmt.Scanln(&stat)
	if stat == "Y" || stat == "y" {
		var newFile string
		fmt.Println("enter the file name in which data is to be copied")
		_, _ = fmt.Scanln(&newFile)
		path1 := "/home/parulraich/go/src/afero_filesystem"
		path_change := path1 + "/" + Dirname + "/" + newFile
		if !fileExists(path_change){
			_,err:=AppFs.Create(path_change)
			if err != nil {
				panic(err)
			}
			f, err := AppFs.OpenFile(path_change, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
		   if err != nil {
			panic(err)
		  }
		defer f.Close()
		_, _ = f.Write([]byte(content))

		}else if fileExists(path_change){
			f, err := AppFs.OpenFile(path_change, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			_, _ = f.Write([]byte(content))
		}

	}else {

		var Dir_new,New_file string
		fmt.Println("enter new directory name")
		_,_ =fmt.Scanln(&Dir_new)
		fmt.Println("enter the file name")
		_,_=fmt.Scanln(&New_file)
		path1 := "/home/parulraich/go/src/afero_filesystem"
		path_change:= path1 + "/" + Dir_new + "/" + New_file
		if fileExists(path_change){
			f,err := AppFs.OpenFile(path_change, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
			if err!=nil{
				panic(err)
			}
			defer f.Close()
			_, _ = f.Write([]byte(content))
		}
	}
}

func renameFile(){
	var stat string
	fmt.Println("do you want to rename file in a directory? \n press Y if yes and N if no")
	_,_=fmt.Scanln(&stat)
	if stat == "Y" || stat == "y" {
		var new string
		Dirname := getDirectoryName()
		Filename := getFileName()
		path_initial := path + "/" + Dirname + "/" + Filename
		fmt.Println("enter new file name")
		_,_=fmt.Scanln(&new)
		path_change :=path + "/" + Dirname + "/" + new
		if fileExists(path_initial){
			err :=AppFs.Rename(path_initial,path_change)
			if err!=nil{
				fmt.Println("file cannot be renamed")
			}
		}else {
			fmt.Println("File Doesn't Exists")
		}

	}else {
		var filename, new string
		fmt.Println("Enter Old file name")
		_, _ = fmt.Scanln(&filename)
		fmt.Println("Enter New file name")
		_, _ = fmt.Scanln(&new)
		if fileExists(filename) {
			AppFs.Rename(filename, new)
			fmt.Println("File Renamed")
		} else {
			fmt.Println("File Doesn't Exists")
		}
	}

}

func Fileinfo(){
	var stat string
	fmt.Println("do you want to rename file in a directory? \n press Y if yes and N if no")
	_,_=fmt.Scanln(&stat)
	if stat == "Y" || stat == "y"{
		Dirname := getDirectoryName()
		Filename := getFileName()
		path_i := path + "/" + Dirname + "/" + Filename
		if fileExists(path_i) {
			fmt.Println(AppFs.Stat(path_i))
		}else {
			fmt.Println("File Doesn't Exists")
		}

	}else {
		filename := getFileName()
		if fileExists(filename) {
			fmt.Println(AppFs.Stat(filename))
		} else {
			fmt.Println("File Doesn't Exists")
		}
	}
}

func deletefile(){
	var stat string
	fmt.Println("do you want to rename file in a directory? \n press Y if yes and N if no")
	_,_=fmt.Scanln(&stat)
	if stat == "Y" || stat == "y"{
		Dirname := getDirectoryName()
		Filename := getFileName()
		path_i := path + "/" + Dirname + "/" + Filename
		if fileExists(path_i) {
			AppFs.Remove(path_i)
			fmt.Println("File Deleted")
		}else {
			fmt.Println("File Doesn't Exists")
		}

	}
	filename := getFileName()
	if fileExists(filename) {
		AppFs.Remove(filename)
		fmt.Println("File Deleted")
	}else {
		fmt.Println("File Doesn't Exists")
	}
}

func exitServer(){
	os.Exit(0)
	fmt.Println("Server Ended")
}

func createDir(){
	path := "/home/parulraich/go/src/afero_filesystem"
	//err := AppFs.MkdirAll(path, 777)
	var Dirname string
	fmt.Println("enter directory name")
	_,_ = fmt.Scanln(&Dirname)
	path = path+ "/" + Dirname
	err := AppFs.MkdirAll(path, 0777)
	if err == nil{
		fmt.Println("path created successfully")
	}
	if err != nil {
		fmt.Println("path directory either exists or some error occured")
	}

}

func readFile(){
	var stat string
	fmt.Println("do you want to rename file in a directory? \n press Y if yes and N if no")
	_,_=fmt.Scanln(&stat)
	if stat == "Y" || stat == "y" {
		Dirname := getDirectoryName()
		Filename := getFileName()
		path_i := path + "/" + Dirname + "/" + Filename
		content, err := ioutil.ReadFile(path_i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Content of "+path_i+":\n", string(content))
	}else {
		filename := getFileName()
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Content of "+filename+":\n", string(content))
	}

}

func removeDir(){
	var Dirname string
	path1 := "/home/parulraich/go/src/afero_filesystem"
	fmt.Println("enter directory name")
	_,_ =fmt.Scanln(&Dirname)
	path1 = path1 + "/" + Dirname
	err := AppFs.RemoveAll(path)
	if err!=nil{
		log.Fatal(err)
	}

}

//checks the file is not there and is also not a directory
func fileExists(filename string) bool {
	info, err := AppFs.Stat(filename)
	if os.IsNotExist(err){
		return false
	}
	return  !info.IsDir()
}

func takeUserinput() (n int){
	fmt.Println("1:create directory\n2: create a file" +
		"\n3:edit a file a file\n4: rename file \n5: File info\n6: delete file\n7:read file\n8:remove directory\n9:move file \n10:exit server\n11:copy file\n\n Enter your choice")
	fmt.Scanln(&n)
	return n
}
func Implementation(n int) {

	switch n {

		case 1:
			createDir()
			n=takeUserinput()
			Implementation(n)

	    case 2:
            createFile()
			n=takeUserinput()
			Implementation(n)
		case 3:
            editFile()
			n=takeUserinput()
			Implementation(n)

		case 4:
			renameFile()
			n=takeUserinput()
			Implementation(n)
		case 5:
			Fileinfo()
			n=takeUserinput()
			Implementation(n)
		case 6:
			deletefile()
			n=takeUserinput()
			Implementation(n)


		case 7:
			readFile()
			n=takeUserinput()
			Implementation(n)

		case 8:
			removeDir()
			n=takeUserinput()
			Implementation(n)

	    case 9:
	    	moveFile()
			n=takeUserinput()
			Implementation(n)


	    case 10:
			exitServer()


	     case 11:
		copyFile()
		n=takeUserinput()
		Implementation(n)

		default:
			fmt.Println("Invalid choice")
			n=takeUserinput()
			Implementation(n)
	}

}