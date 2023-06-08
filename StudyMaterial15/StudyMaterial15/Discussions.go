_______________________________________________________________________________________

RECOMMENDED ENVIRONMENT SETUP
_______________________________________________________________________________________
	
	1. Git Repository Access 
	2. Ubuntu 22.04 64 Bit Desktop In Virtualisation Environment (Guest OS)
	3. Or Get Installed Go In Host Machine

_______________________________________________________________________________________

INSTALLING GO COMPILER
_______________________________________________________________________________________
	
	Reference Link: 
		https://go.dev/dl/

	To Check Go Installed Or Not?
	
	go version

_______________________________________________________________________________________

HELLO WORLD IN GO
_______________________________________________________________________________________

1. Write Following Code In Hello.go File

	package main

	import "fmt"

	func main() {
		fmt.Println("Hello World!!!!")
	}

2. Save Code In CoforgeGo
	
	YOUR_USER_DIRECTOR/Documents/CoforgeGo


2. Running Go Program
	Change Your Directory To Code Directory

	cd YOUR_USER_DIRECTOR/Documents/CoforgeGo

	go run Hello.go

	OR

	go build Hello.go
	./Hello

_______________________________________________________________________________________
_______________________________________________________________________________________

CODING EXCERCISE C1

Write sum Function With Following Given Signature In C
	int sum( int x, int y )

// DESIGN 01 : BAD CODE

	int sum( int x, int y ) {
		int total;
		total = x + y;
		return total
	}


	int sum( int x, int y ) {
		return x + y;
	}

_______________________________________________________________________________________
_______________________________________________________________________________________


DESIGN 02 : BAD CODE IT WILL NOT WORK

  	int range= 4294967295 // eg of range of int data type
    
    if(sum <= range){
        return sum;
        }
    else{
        printf("Can't calculate sum as it's exceeding the range of int");
        return -1;
    }


Upadhaya Omprakash, Ajay, Shashank
 DESIGN 03 : GOOD DIRECTION

	int sum(int x, int y){
	    if((y>0&&x>INT_MAX-y)||(y<0&&x<INT_MIN-y)){
	        println("cant calculate given sum")
	    }
	    println(x+y)
	}


DESIGN 04 : BAD CODE IT WILL NOT WORK

[12:51 PM] Kundan Kumar Sahani
    int sum(int x, int y) {
     int arithmeticSum = sum(x, y);
    
    if(arithmeticSum < sizeof(int)){
        printf("The arithmetic sum from %d to %d is %d\n", x, y, arithmeticSum);
        retrn x+y;
    }else{
            printf("Can't calculate sum of given values");
            return 0;
}


DESIGN 05 : BAD CODE IT WILL NOT WORK
[12:57 PM] Yatendra Singh
    int sum( int x, int y) {
    	 if (x >= INT_MIN && x <=INT_MAX && y >= INT_MIN && y <=INT_MAX) 
    	{ 
    		return x+ y 

        } else {
        	 printf("invalid Values"); return; }
         }


_______________________________________________________________________________________

CODING EXCERCISE C2


Write sum Function With Following Given Signature In C
	// Preconditions Or Premises
	It Should Return Valid Arithematic Sum
				OR
	Print Can;t Calcualte Sum Of Given Values

	int sum( int x, int y )

BAD DESIGN CHOICES

	1. Design Based On sizeof Operator It's Bad Design Will Not Work

	2.  MIN <= x, y <= MAX  // ALWAYS TRUE CONDITION
		  result = x + y

		MIN <= result <= MAX // ALWAYS TRUE CONDITION

	3.  x <= MIN || x >= MAX && y <= MIN || y >= MAX // ALWAYS FALSE CONDITION
		  result = x + y

		result <= MIN || result >= MAX // ALWAYS FALSE CONDITION


GOOD DESIGN

	int sum(int x, int y) {
		int result = 0;
		// Type Deriven Design : Type Safe Code
		//		Respect Type Definition
		if ( ( y > 0 && x > INT_MAX - y ) || ( y < 0 && x < INT_MIN - y ) ) {
			printf("Cant't Calcualte Sum");
		} else {
			result = x + y;
			return result;
		}
	}

_______________________________________________________________________________________
_______________________________________________________________________________________


WHAT IS DATA TYPE???

_______________________________________________________________________________________
_______________________________________________________________________________________


_____________________________________________________________________________

CREATING GO MODULES
_____________________________________________________________________________

1. Create Hello Directory
2. cd Hello
3. Create main.go File In Hello Directory
4. Write Following Code In main.go

    1
    2 package main
    3
    4 //  Package
    5 import "fmt"
    6 import "rsc.io/quote"
    7
    8 func Hello() string {
    9     return quote.Hello()
   10 }
   11
   12 func main() {
   13     fmt.Println("Hello World!!!!")
   14     Hello()
   15 }
   16

5. To Make Directory As go  Module Run Following Command
	 go mod init example.com/hello

6. It Will Create go.mod File

7. To Resolve Dependencies Run Following Command
	 
   go mod tidy

8. To Run Code
	 go run .
	 go run main.go

_____________________________________________________________________________
_____________________________________________________________________________
_____________________________________________________________________________

