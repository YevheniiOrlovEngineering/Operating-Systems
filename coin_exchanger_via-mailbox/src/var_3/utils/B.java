package var_3.utils;

public class B extends Thread {
    CoinExchanger coinExchanger;

    // Thread B
    public B(CoinExchanger coinExchanger)
    {
        this.coinExchanger = coinExchanger;
    }

    // run() method for this thread invoked as
    // start() method is called in the CoinExchangerLauncher() method
    public void run()
    {
        while (true)
        {
            try {
                this.coinExchanger.Exchange();
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        }
    }
}
