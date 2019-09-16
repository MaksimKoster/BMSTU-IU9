import java.util.Arrays;

public class Test {
    public static void main (String[] args){
        Polynomial[] p = new Polynomial[]{
                new Polynomial(1, 2, 3, 4, 5, 7),
                new Polynomial(1, 2, 4),
                new Polynomial(5, 2, 0),
                new Polynomial(1, -2, 1),
                new Polynomial(5, 4, 3, 2, 1),
                new Polynomial(1),
                new Polynomial(2, 1, 0)
                };
        //for (Polynomial cf : p) System.out.println(cf + "\nNumber of roots:  " + cf.getNumberOfRoots());
        Arrays.sort(p);
        System.out.println("\nSorted: \n");
        for (Polynomial cf : p) System.out.println(cf + "\nNumber of roots:  " + cf.getNumberOfRoots());
    }
}
