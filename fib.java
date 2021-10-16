import java.util.*;
import java.lang.*;
import java.io.*;

public class Main
{
	
    // define the golden ratio
    public static final double GOLDEN_RATIO = (1 + Math.sqrt(5)) / 2;

    public static double fib(int seq) {
        if (seq <= 0) return 0;
        return Math.round(Math.pow(GOLDEN_RATIO, seq) / Math.sqrt(5));
    }
    
	public static void main (String[] args) throws java.lang.Exception
	{
		for (int i = 1; i <= 10; i++) {
			System.out.println(fib(i));
		}
	}
}
