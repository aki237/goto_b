package main

import "os"
import "os/exec"
import "fmt"
import "database/sql"
import _"github.com/mattn/go-sqlite3"
import "strings"

func register(nick,destination string) bool {
	//Init Database
	home := os.Getenv("HOME")
	gotodb,err := sql.Open("sqlite3",home+"/.config/goto/goto.db")
	var returnstate bool = false
	var err_n,err_d string
	if err != nil {
		fmt.Println("Unable to access db")
	}
	//Checking whether the nick is unique and inserting if yes.
	rows, err := gotodb.Query("select * from goto where  NICK= ?", nick)
	if err != nil{
		fmt.Println(err)
	}
	defer rows.Close()
	if (!rows.Next()){
		stmt, err := gotodb.Prepare("INSERT INTO goto(NICK,DESTINATION) VALUES(?,?)")
		if err != nil{
			fmt.Println("Unable to prepare database")
		}
		_,err = stmt.Exec(nick,destination)
		if err != nil {
			fmt.Println("Unable to record the nick.")
		}
		returnstate = true
	}else{
		if(rows.Next()){
			err = rows.Scan(&err_n, &err_d)
			fmt.Println("The given nick is already in use for another Destination : "+err_n,err_d)
		}
		returnstate = false
	}
	defer gotodb.Close()
	return returnstate
}

//
func listall(what string) {
	var nick ,dest string
	home := os.Getenv("HOME")
	gotodb,err := sql.Open("sqlite3",home+"/.config/goto/goto.db")
	if err != nil {
		fmt.Println("Unable to access db")
	}
	//Checking whether the nick is unique and inserting if yes.
	rows, err := gotodb.Query("select * from goto")
	if err != nil{
		fmt.Println(err)
	}
	defer rows.Close()
	for(rows.Next()){
		err = rows.Scan(&nick,&dest)
		if err!=nil{
			fmt.Println(err)
		}else{
			switch what{
			case "all":
				fmt.Println(nick + " - "+dest)
			case "nick":
				fmt.Println(nick)
			default:
				fmt.Println(nick + " - "+dest)
			}
		}
	}
}

//
func getbranch(nick string) {
	dir := gonick(nick)
	err := os.Chdir(dir)
	out,err := exec.Command("git", "branch").Output()
	if err != nil{
		fmt.Println("Not a git repository",err)
	}
	retbash := string(out)
	retbash = strings.Replace(retbash,"*","",1)
	retbash = strings.Replace(retbash,"\n"," ",-1)
	fmt.Println(retbash)
}
//
func gonick(nicks string) string {
	var nick ,dest string
	gotodb,err := sql.Open("sqlite3","/home/aki237/.config/goto/goto.db")
	if err != nil {
		fmt.Println("Unable to access db")
	}
	//Checking whether the nick is unique and inserting if yes.
	rows, err := gotodb.Query("select * from goto where NICK= ?" , nicks)
	if err != nil{
		fmt.Println(err)
	}
	defer rows.Close()
	if(rows.Next()){
		err = rows.Scan(&nick,&dest)
		if err!=nil{
			fmt.Println(err)
		}
	}else{
		rows, err = gotodb.Query("select * from goto")
		if err != nil{
			fmt.Println(err)
		}
		defer rows.Close()
		for(rows.Next()){
			err = rows.Scan(&nick,&dest)
			if err!=nil{
				fmt.Println(err)
			}
			if(strings.Contains(nick,nicks)){
				break
			}
		}
	}
	return dest
}

//
func usage(){
	fmt.Println("Usage : ")
	fmt.Println("\tgoto [command] { [nick] [target] }")
	fmt.Println("go [nick]           -     go to the directory assigned to a given nick")
	fmt.Println("list                -     list all the recorded nicks against the directory")
	fmt.Println("reg [nick] [target] -     register a nick against a given target (or) the current working directoy")
}

//
func main() {
	var command,nick,destination string
	destination,err := os.Getwd()
	if err != nil{
		fmt.Println("Something is wrong")
	}
	switch len(os.Args){
	case 1:
		usage()
		return
	case 2:
		command = os.Args[1]
	case 3:
		command = os.Args[1]
		nick = os.Args[2]
	case 4:
		command = os.Args[1]
		nick = os.Args[2]
		destination = os.Args[3]
	}
	switch command {
	case "list":
		listall("all")
		return
	case "listnick":
		listall("nick")
		return
	case "reg":
		register(nick, destination)
	case "go":
		printout := gonick(nick)
		fmt.Println(printout)
		return
	case "branch":
		getbranch(nick)
		return
	default:
		usage()
		return
	}
}
