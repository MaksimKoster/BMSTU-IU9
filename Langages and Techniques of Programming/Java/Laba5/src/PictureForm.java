import javax.swing.*;
import javax.swing.event.ChangeEvent;
import javax.swing.event.ChangeListener;
import java.awt.event.*;

public class PictureForm {
    private JPanel mainPanel;
    private JSpinner value2;
    private JSpinner value1;
    private JSpinner value3;
    private CanvasPanel canvasPanel;
    private JSlider red;
    private JSlider green;
    private JSlider blue;
    private JComboBox<String> colorSwitch;
    private JButton vertexAUpDownButton;

    public PictureForm() {
        value1.addChangeListener(e -> {
            int a = (int)value1.getValue();
            canvasPanel.setA(a);
        });
        value1.setValue(50);
        value2.addChangeListener(e -> {
            int b = (int)value2.getValue();
            canvasPanel.setB(b);
        });
        value2.setValue(50);
        value3.addChangeListener(e -> {
            int c = (int)value3.getValue();
            canvasPanel.setC(c);
        });
        value3.setValue(50);

        colorSwitch.addActionListener(e -> {
            String color = (String)colorSwitch.getSelectedItem();
            canvasPanel.setColor(color);
            if (colorSwitch.getSelectedItem().equals("Red")){
                red.setValue(255);
                green.setValue(0);
                blue.setValue(0);
            }else if (colorSwitch.getSelectedItem().equals("Green")){
                red.setValue(0);
                green.setValue(255);
                blue.setValue(0);
            }else if (colorSwitch.getSelectedItem().equals("Blue")){
                red.setValue(0);
                green.setValue(0);
                blue.setValue(255);
            }if (colorSwitch.getSelectedItem().equals("Black")){
                red.setValue(0);
                green.setValue(0);
                blue.setValue(0);
            }
        });
        colorSwitch.addItem("RGB");
        colorSwitch.addItem("Red");
        colorSwitch.addItem("Blue");
        colorSwitch.addItem("Green");
        colorSwitch.addItem("Black");

        red.addChangeListener(e -> {
            int r = red.getValue();
            canvasPanel.setRed(r);
        });
        red.setValue(20);
        green.addChangeListener(e -> {
            int g = green.getValue();
            canvasPanel.setGreen(g);

        });
        green.setValue(40);
        blue.addChangeListener(e -> {
            int b = blue.getValue();
            canvasPanel.setBlue(b);
        });
        blue.setValue(30);
        canvasPanel.addKeyListener(new KeyAdapter() {

            public void keyReleased(KeyEvent e) {}
            public void keyTyped(KeyEvent e) {}
            @Override
            public void keyPressed(KeyEvent e) {
                super.keyTyped(e);
                switch (e.getKeyCode()){
                    case KeyEvent.VK_DOWN:
                        canvasPanel.setZeroY(canvasPanel.getZeroY() + 1);
                        break;
                    case KeyEvent.VK_UP:
                        canvasPanel.setZeroY(canvasPanel.getZeroY() - 1);
                        break;
                    case KeyEvent.VK_LEFT:
                        canvasPanel.setZeroX(canvasPanel.getZeroX() - 1);
                        break;
                    case KeyEvent.VK_RIGHT:
                        canvasPanel.setZeroX(canvasPanel.getZeroX() + 1);
                        break;
                }
            }
        });
        canvasPanel.addMouseListener(new MouseAdapter() {
            @Override
            public void mousePressed(MouseEvent e) {
                canvasPanel.requestFocusInWindow();
            }
        });
        vertexAUpDownButton.addActionListener(e -> {
            canvasPanel.reverse();
        });
    }

    public static void main(String[] args) {
        JFrame frame = new JFrame("Triangle");
        frame.setContentPane(new PictureForm().mainPanel);
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.pack();
        frame.setVisible(true);
    }
}
