import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class Test {
    public static void main(String[] args) {
        Map<Boolean, List<Integer>> res;
        Sentence s = new Sentence("a bb  cad ccac lal cawe vabnavba  ib");
        //false - четное true - нечетное

        System.out.println(s.getNumberString());
//        System.out.println(s.getNumber(false).collect(Collectors.joining()));
        System.out.println(s.getWords().get(0));

        res = s.getNumberStream()
                .collect(Collectors.partitioningBy(x -> x % 2 != 0));

        System.out.println(res + "\n false - even; true - odd");
    }
}
