public class Multiples {

    public static int FunctionDivisibleByNumber(int n, int target){
        int sum = 0;
        int samples = target/n;
        int sumation_samples = (int)(samples * (samples + 1) * n * 0.5);
        return sumation_samples;
    }
    public static int lcm(int a, int b){
        int num = 1;
        while(true){
            if (num % a == 0 && num % b == 0) {
                return num;
            }
            num += 1;
        }
    }
    public static void main(String[] args) {
        int a = 3;
        int b = 5;
        int target = 999;
        int x = FunctionDivisibleByNumber(a, target);
        int y = FunctionDivisibleByNumber(b, target);
        int lcm_val = lcm(a, b);
        int lcm_multiples = FunctionDivisibleByNumber(lcm_val, target);
        int output = x + y - lcm_multiples;
        System.out.println(output);
    }
}
