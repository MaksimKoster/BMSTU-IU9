import java.util.Iterator;

public class IntSequence implements Iterable<Integer> {
    private int[] seq;
    public IntSequence(int ... seq){
        this.seq = seq;
    }
    public Iterator<Integer> iterator() {
        return new GCDIterator();
    }
    public int getGCD(int a, int b){
        while (b !=0) {
            int tmp = a%b;
            a = b;
            b = tmp;
        }
        return a;
    }
    public void changeElem(int pos, int x){
        seq[pos] = x;
    }
    private class GCDIterator implements Iterator<Integer>{
        private int pos;

        public GCDIterator(){
            pos = 0;
        }

        public boolean hasNext(){
            return pos < seq.length - 1;
        }
        public Integer next(){
            return getGCD(seq[pos], seq[++pos]);
        }
    }

    public String toString(){
        StringBuilder res = new StringBuilder();
        for (int i = 0; i < seq.length; i++) {
            res.append(seq[i]).append(" ");
        }
        return res.toString();
    }
}
