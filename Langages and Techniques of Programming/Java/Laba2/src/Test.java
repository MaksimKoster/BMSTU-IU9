public class Test {
    public static void main(String[] args){
        Point[] lineFractures = new Point[] {
                new Point (1.0, 2.0),
                new Point (5.0, 2.1),
                new Point (1.0,10.0),
                new Point (6.0, 1.0)};
        Line l1 = new Line (lineFractures);
        System.out.println("l1 length: " + l1.getLength());
        System.out.println(l1);
        System.out.println(l1.getSize());
        System.out.println(l1.getPoint(3));
    }
}
