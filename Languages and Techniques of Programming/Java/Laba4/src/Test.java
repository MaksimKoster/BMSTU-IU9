public class Test {
    public static void main(String[] args){
        int[] arr = new int[]{1, 2, 4, 5, 10};
        IntSequence sq = new IntSequence(arr);
        for (Integer a : sq) System.out.print(a + " ");
        System.out.println();
        sq.changeElem(1, 5);
        for (Integer a : sq) System.out.print(a + " ");
    }
}
