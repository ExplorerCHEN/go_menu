package main
 
import "fmt"
 
/**************************************************************************************************/
/* Copyright (C) cpchen@ustc.edu, SSE@USTC, 2022-2023                                             */
/*                                                                                                */
/*  FILE NAME             :  menu.go                                                              */
/*  PRINCIPAL AUTHOR      :  Chencanpeng                                                          */
/*  SUBSYSTEM NAME        :  menu                                                                 */
/*  MODULE NAME           :  menu                                                                 */
/*  LANGUAGE              :  GO                                                                   */
/*  TARGET ENVIRONMENT    :  ANY                                                                  */
/*  DATE OF FIRST RELEASE :  2022/03/25                                                           */
/*  DESCRIPTION           :  This is a menu program                                               */
/**************************************************************************************************/


func main() {
	var cmd string
	for {
		fmt.Print("Please input a command: ")
		fmt.Scanln(&cmd)
		switch cmd {
		case "help":
			fmt.Println("This is help command.")
		case "quit":
			fmt.Println("Bye.")
			return
		default:
			fmt.Println("Wrong command!")
		}

	}
}
