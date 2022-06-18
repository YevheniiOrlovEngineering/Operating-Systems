package var_3.launcher;

import var_3.utils.A;
import var_3.utils.B;
import var_3.utils.CoinExchanger;


public class CoinExchangerLauncher {
    public static void main(String[] args) {
        CoinExchanger coinExchanger = new CoinExchanger();

        A t1 = new A(coinExchanger);
        B t2 = new B(coinExchanger);

        t1.start();
        t2.start();
    }
}
