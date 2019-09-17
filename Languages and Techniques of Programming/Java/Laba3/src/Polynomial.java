
public class Polynomial implements Comparable<Polynomial> {
    private int[] cf;
    //private int degree;
    public Polynomial(int ... cf){
        //setDegree(cf.length - 1);
        setNumberOfCf(cf.length);
        this.cf = cf;
    }
//    private void setDegree(int degree) {
//        this.degree = degree;
//    }
    private void setNumberOfCf(int value) {
        cf = new int[value];
    }

    public int getNumberOfRoots(){
        int numberOfRoots = 0;
        for (int i = 0; i <= 10; i++) {
            int res = 0;
            for (int j = cf.length - 1; j > 0; j--) {
                res += Math.pow(i, j) * cf[j];
            }
            if (res == 0) numberOfRoots++;
        }
        return numberOfRoots;
    }

    public String toString(){
        int i;
        StringBuilder res = new StringBuilder();
        for(i = cf.length - 1; i >= 0; i--){
            if(cf[i] != 0){
                res.append(cf[i]).append("*x^").append(i);
                break;
            }
        }
        for(i-- ; i >= 0; i--) {
            if (cf[i] != 0)
                res.append(" + ").append(cf[i]).append("*x^").append(i);
        }
        return res.toString();
    }

    public int compareTo(Polynomial obj){
        if (getNumberOfRoots() == obj.getNumberOfRoots()) return 0;
        else if (getNumberOfRoots() > obj.getNumberOfRoots()) return 1;
        else return -1;
    }
}
