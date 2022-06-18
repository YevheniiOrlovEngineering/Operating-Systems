package var_3.CoinExchanger;

public class B extends Thread {
    CoinExchanger coinExchanger;

    public B(CoinExchanger coinExchanger) { this.coinExchanger = coinExchanger; }

    public void run() {
        while (true) {
            try {
                this.coinExchanger.Exchange();
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        }
    }
}
