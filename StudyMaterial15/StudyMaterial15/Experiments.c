
#include <stdio.h>
#include <limits.h>

#define SIZE 10

//_____________________________________________________________________

void playWithCommandLineArgs(int argc, char *argv[]) {
	for ( int i = 0 ; i < argc ; i++ ) {
		printf("\n Argument Value At %d = %s", i, argv[i] );
	}
}

//_____________________________________________________________________
// DESIGN 01 : BAD CODE

int sumOld( int x, int y ) {
	return x + y;
}

void playWithSum() {
	int a, b, result = 0;

	a = 2147483647;
	b = 1;

	result = sumOld( a, b );
	printf("\nResult : %d", result ); // 2147483648

	a = -2147483648;
	b = -10;

	result = sumOld( a, b );
	printf("\nResult : %d", result ); // -2147483658
}

// Function : playWithSum
// Result : -2147483648
// Result : 2147483638

//_____________________________________________________________________

// GOOD DESIGN
//		DRIVEN BY FUNAMENTALS
int summation(int x, int y) {
	int result = 0;

	if ( ( y > 0 && x > INT_MAX - y ) || ( y < 0 && x < INT_MIN - y ) ) {
		printf("Cant't Calcualte Sum");
	} else {
		result = x + y;
		return result;
	}
}

//_____________________________________________________________________

void playWithTypes() {
	// char b = 900;
	
	char b = 90;
	printf("\nValue %d", b);

	// In C By Default
	//		int Is Signed By Default
	//			[ -32768, 32767 ]
	//		usinged int Is UnSigned int
	//			[ 0, 65535 ]
	
	int i = 100;
	unsigned ui = 900;
}

//_____________________________________________________________________

// In C
//		Strings Are NULL \0 Terminated 
void playWithString() {
	// In This Case NULL Character Is Stored As Last Character
	//		Size Of String Will Be Number Of Characters In String + 1
	//		+1 Is For Storing \0 Character
	// Compiler Will Calculate Size Of char Array With Above Logic
	char greeting[] = "Good Morning!";
	char username[] = "amarjit@icicibank.com";
	char password[] = "Good Morning!";

	// In The Following Case NULL Character Is NOT Stored As Last Character
	//		Because I Am Fixing Size Of String As Number Of Characters In String
	char dingdong[9] = "Ding Dong";  
	// char dingdong1[9] = "Ding Dong\0";
	char dingdongAgain[10] = "Ding Dong\0";
	char *dingdongAgain1 = "Ding Dong\0";
	
	printf("\nGreeting : %s", greeting);
	printf("\nDing Dong: %s", dingdong);
	printf("\nDing Dong Again: %s", dingdongAgain);
	printf("\nDing Dong Again 01: %s", dingdongAgain1);

// Function : playWithString
// Greeting : Good Morning!
// Ding Dong: Ding DongGood Morning!
}

//_____________________________________________________________________

void playWithPointers() {
	int a = 20;
	printf("\nValue Of a: %d", a);

	// Pointer Variable
	//		Is A Variable Used To Address Of Memory Location
	// Address Of a Given By &a
	// Address Of a Storing In ptr Pointer Variable 
	int * ptr = &a; 

	// *ptr Will Defer Value At Address Stored In ptr
	printf("\nAddress Of a: %x", ptr );
	printf("\nValue Of Address %x =  %d",ptr,  *ptr );
}

//_____________________________________________________________________

void doChange(int a[], int size) {
	for (int i = 0 ; i < size ; i++ ) {
		a[i] = 99;
	}
}

// In C/C++
//		Arrays Are Pass By Reference
void printArray(int a[], int size) {
	for (int i = 0 ; i < size ; i++ ) {
		printf("\n\tAt Index: %d Value : %d", i, a[i] );
	}
}

//#define SIZE 10
void playWithArray() {
	int numbers[ SIZE ] = {10, 20, 30, 40, 50, 60, 70, 80, 90, 100};
	
	printf("\nArray Before doChange Called...");
	printArray( numbers, SIZE );
	doChange( numbers, SIZE );

	printf("\nArray After doChange Called...");
	printArray( numbers, SIZE );
}

//_____________________________________________________________________

void playWithArrayAndPointers() {
	// Array Name Is Pointer To Starting Element Of Array 
	//		numbers Is Pointer To First Element Of Continuous Array
	int numbers[ SIZE ] = {10, 20, 30, 40, 50, 60, 70, 80, 90, 100};
	int * pointNumbers;

	printf("\nPrinting Values Using Array Index\n");
	for ( int i = 0 ; i < SIZE ; i++ ) {
		printf("\t%d", numbers[i] );
	}

	pointNumbers = numbers;
	printf("\nPrinting Values Using Pointer To Array\n");
	for ( int i = 0 ; i < SIZE ; i++ ) {
		printf("\t%d", *( pointNumbers + i ) );
	}
// Function : playWithArrayAndPointers
// 	Printing Values Using Array Index
// 	10	20	30	40	50	60	70	80	90	100
// 	Printing Values Using Pointer To Array
// 	10	20	30	40	50	60	70	80	90	100
}

//_____________________________________________________________________

//	Internally Function Name
//		Is Pointer To Starting Address Of Function Body
int multiply(int x, int y) { return x * y; }
int division(int x, int y) { return x / y; }

int multiply3(int x, int y, int z) { return x * y; }

void playWithPointersAgain() {
	int value = 40, result;
	int * pointerToValue = &value;

	printf("\nValue %d and Value Again %d", value, *pointerToValue );

	// Function Pointer: Pointer To Function
	//		doSomething Is Pointer To Function Which Returns int
	//			And Takes Two int Arguments
	int (* doSomething )( int, int );

	doSomething = multiply;

	result = multiply( 11, 10 );
	printf("\nResult : %d", result );

	result = doSomething( 11, 10 );
	printf("\nResult : %d", result );

	doSomething = division;

	result = division( 11, 10 );
	printf("\nResult : %d", result );

	result = doSomething( 11, 10 );
	printf("\nResult : %d", result );

	//assignment to ‘int (*)(int,  int)’ from incompatible pointer type 
	// ‘int (*)(int,  int,  int)’
	// doSomething = multiply3;
}

// Function : playWithPointersAgain
// Value 40 and Value Again 40
// Result : 110
// Result : 110

//_____________________________________________________________________

// addition, substraction and multiplication Have 
// 		Function Type (int, int) -> int
// Function Type
//		(int, int) -> int
int addition(int x, int y) 			{ return x + y; }
int substraction(int x, int y) 		{ return x - y; }
int multiplication(int x, int y ) 	{ return x * y; }

//	Higher Order Functions
//		Functions Which Takes Functions As Arguments
//			And/Or
//		Can Returns Function

// Can I Say calculator Function Is Polymorphic?
// Polymoprhism
//		Using Mechanism By Passing Function/Behviour To Function/Behaviour

int calculator( int x, int y, int (*operation)(int, int) ) {
	return operation(x, y);
}

void playWithCalculator() {
	int a = 40, b = 10, result = 0;

	result = calculator(a, b, addition);
	printf("\nResult : %d", result);

	result = calculator(a, b, substraction);
	printf("\nResult : %d", result);

	result = calculator(a, b, multiplication);
	printf("\nResult : %d", result);
}


//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________

int main(int argc, char *argv[]) {
	printf("\n\nFunction : playWithCommandLineArgs");
	playWithCommandLineArgs( argc, argv );

	printf("\n\nFunction : playWithSum");
	playWithSum();

	printf("\n\nFunction : playWithTypes");
	playWithTypes();

	printf("\n\nFunction : playWithString");
	playWithString();

	printf("\n\nFunction : playWithPointers");
	playWithPointers();

	printf("\n\nFunction : playWithArray");
	playWithArray();

	printf("\n\nFunction : playWithArrayAndPointers");
	playWithArrayAndPointers();

	printf("\n\nFunction : playWithPointersAgain");
	playWithPointersAgain();

	printf("\n\nFunction : playWithCalculator");
	playWithCalculator();

	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");
	// printf("\n\nFunction : ");	
}
