
package learnJava;

public class Experiments {
	public static void printArray(int a[], int size) {
		for (int i = 0 ; i < size ; i++ ) {
			System.out.printf("\nAt Index %d Value %d", i, a[i]);
		}
	}

	// In Java
	//		Arrays Are Pass By Reference
	public static void doChange(int a[], int size) {
		for (int i = 0 ; i < size ; i++ ) {
			a[i] = 99;
		}
	}

	public static void main(String[] args) {
		int [] numbers = new int[]{ 10, 20, 30, 40, 50, 60, 70, 80, 90, 100 };
	
		System.out.printf("\nArray Before doChange Called");
		printArray( numbers, 10 );
		doChange(numbers, 10);
		System.out.printf("\nArray After doChange Called");
		printArray( numbers, 10 );
	}
}

