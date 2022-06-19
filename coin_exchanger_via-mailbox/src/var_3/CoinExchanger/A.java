package var_3.CoinExchanger;

public class A  extends Thread {
    CoinExchanger coinExchanger;

    public A(CoinExchanger coinExchanger) { this.coinExchanger = coinExchanger; }

    public void run() {
        while (true) {
            try {
                this.coinExchanger.InputCoinToBeExchanged();
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        }
    }
}
