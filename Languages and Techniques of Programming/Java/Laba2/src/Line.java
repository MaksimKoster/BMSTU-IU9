public class Line {
    private Point[] lineFractures;
    public Line (Point[] lineFractures){
        this.lineFractures = lineFractures;
    }
    public double getLength (){
        double length = 0;
        for (int i = 1; i < lineFractures.length; i++){
            double tempX = Math.pow(lineFractures[i].getX() - lineFractures[i - 1].getX(), 2);
            double tempY = Math.pow(lineFractures[i].getY() - lineFractures[i - 1].getY(), 2);
            length += Math.sqrt(tempX + tempY);
        }
        return length;Integer
    }

    public int getSize (){
        return lineFractures.length;
    }

    public Point getPoint (int i){
        return lineFractures[i];
    }

    public String toString(){
        /*String points = "";
        for (Point lineFracture : lineFractures) {
            points += "x: " + lineFracture.getX() + "; y: " + lineFracture.getY() + "\n";
        }*/

        StringBuilder points = new StringBuilder();
        for (Point lineFracture : lineFractures) {
            points.append("x: ").append(lineFracture.getX()).append("; y: ").append(lineFracture.getY()).append("\n");
        }

        return "Line Fractures: " + points;
    }
}
