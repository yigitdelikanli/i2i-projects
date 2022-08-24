package JUnitDemo.JUnit;

import java.util.Scanner;

public class BusinessLogic {
	public static void main(String[] args) {
		Scanner scanner = new Scanner(System.in);
		System.out.println("Enter the number which will be control");
		int number = scanner.nextInt();
		if (isPrime(number)) {
			System.out.println(number + "is prime number");
		} else {
			System.out.println(number + "is not prime number");
		}
		scanner.close();
	}

	public static boolean isPrime(int number) {
		boolean flag = true;
		if (number <= 1)
			return false;
		for (int i = 2; i <= number / 2; i++) {
			if (number % i == 0) {
				flag = false;
				break;
			}
		}
		return flag;
	}

}