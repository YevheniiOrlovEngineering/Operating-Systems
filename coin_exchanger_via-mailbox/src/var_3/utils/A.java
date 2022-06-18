package var_3.utils;

public class A extends Thread {
        CoinExchanger coinExchanger;

        // Thread A
        public A(CoinExchanger coinExchanger)
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
                                this.coinExchanger.InputCoinToBeExchanged();
                        } catch (InterruptedException e) {
                                throw new RuntimeException(e);
                        }
                }

        }
}
