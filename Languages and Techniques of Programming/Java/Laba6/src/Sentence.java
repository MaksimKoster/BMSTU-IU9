import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class Sentence {
    private List<String> splitted;
    private String full;

    public Sentence(String full) {
        this.full = full.replaceAll("\\s+", " ");
        this.splitted = splitString();
    }

    public List<String> splitString() {
        return Arrays.asList(full.split("\\s+"));
    }

    public Stream<Integer> getNumberStream(){
        return getStream()
                .map(x -> ((int) x.toLowerCase().charAt(0) - 96));
    }

    public String getNumberString(){
        return getStream()
                .map(x -> ((int) x.toLowerCase().charAt(0) - 96) + " ")
                .collect(Collectors.joining());
    }

    public Stream<String> getStream() {
        return splitted.stream();
    }

    public List<String> getWords(){
        return getStream()
                .filter(x -> x.charAt(0) - 96 == x.length())
                .collect(Collectors.toList());
    }
}
