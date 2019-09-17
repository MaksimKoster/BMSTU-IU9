import javax.swing.*;
import java.awt.*;

public class CanvasPanel extends JPanel {
    private int a = 50, b = 50, c = 50;
    private int red = 0, green = 0, blue = 0;
//    private char f = 1;
    private String color = "RGB";
    private double gamma = (b * b + a * a - c * c) / (double)(2 * b * a); // cos gamma
    private double h = (b * Math.sqrt(1 - gamma * gamma)), oh = a - (int) Math.sqrt((b * b) - (h * h));
    private int zeroX = 50, zeroY = 50;

    public void setA(int a) {
        this.a = a;
        gamma = (b * b + a * a - c * c) / (double)(2 * b * a); // cos gamma
        if (h > 0)
            h = (b * Math.sqrt(1 - gamma * gamma));
        else h = - (b * Math.sqrt(1 - gamma * gamma));
        oh = a - (int) Math.sqrt((b * b) - (h * h));
        repaint();
    }

    public void setB(int b) {
        this.b = b;
        gamma = (b * b + a * a - c * c) / (double)(2 * b * a); // cos gamma
        if (h > 0)
            h = (b * Math.sqrt(1 - gamma * gamma));
        else h = - (b * Math.sqrt(1 - gamma * gamma));
        oh = a - (int) Math.sqrt((b * b) - (h * h));
        repaint();
    }

    public void setC(int c) {
        this.c = c;
        gamma = (b * b + a * a - c * c) / (double)(2 * b * a); // cos gamma
        if (h > 0)
            h = (b * Math.sqrt(1 - gamma * gamma));
        else h = - (b * Math.sqrt(1 - gamma * gamma));
        oh = a - (int) Math.sqrt((b * b) - (h * h));
        repaint();
    }

    public void setRed(int red) {
        this.red = red;
        repaint();
    }

    public void setGreen(int green) {
        this.green = green;
        repaint();
    }

    public void setBlue(int blue) {
        this.blue = blue;
        repaint();
    }

    public void setZeroX(int x){
        zeroX = x;
        repaint();
    }

    public void setZeroY(int y) {
        zeroY = y;
        repaint();
    }

    public int getZeroX(){ return zeroX; }
    public int getZeroY(){ return zeroY; }

    public void setColor(String color) {
        this.color = color;
        repaint();
    }

    public void reverse(){
        h *= -1;
        repaint();
    }

    protected void paintComponent(Graphics g) {
        super.paintComponent(g);
        try {
            if (a + b <= c || b + c <= a || a + c <= b){
                throw new RuntimeException("Sum of two different side lengths must be bigger then the third");
            }
            Color cl = new Color(red, green, blue);
            Polygon p = new Polygon(new int[]{zeroX, zeroX + (int) oh, zeroX + a}/*x's*/, new int[]{zeroY, zeroY + (int) h, zeroY}/*y's*/, 3);


            g.setColor(cl);

            g.drawPolygon(p);
            g.fillPolygon(p);
        }catch (RuntimeException e){
            g.setColor(Color.RED);
            //g.drawString("Sum of two different\n side lengths must be bigger\n then the third", 20,20 );
            g.drawString("Sum of two different", 20, 20);
            g.drawString("side lengths must be bigger", 20, 40);
            g.drawString("then the third", 20,60 );
        }
    }
}
