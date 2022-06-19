package var_3.CoinExchanger;

import java.util.HashMap;
import java.util.InputMismatchException;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.Scanner;
import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class CoinExchanger {
    private static final List<Integer> values = List.of(1, 2, 5, 10, 25, 50, 100);
    private final Map<Integer, Integer> coinsPool = new HashMap<>();
    private static final Map<String, String> exchangerMessages = new HashMap<>() {
        {
            this.put("success",
                                            """
                                            [INFO] Exchanging done successfully.
                                            [INFO] Money: [%d] coins of (%d) cent
                                            [INFO] Change: [1] coin of (%d) cent
                                            """
            );
            this.put("success no change",
                                            """
                                            [INFO] Exchanging done successfully.
                                            [INFO] Money: [%d] coins of (%d) cent
                                            """
            );
            this.put("money and change error",
                                            """
                                            [ERROR] Exchanging done with error.
                                            [ERROR] The coins amount [%d] of desired value (%d)
                                            [ERROR] and coins amount [%d] of change (%d) coins
                                            [ERROR] is not enough to exchange (%d) cent coin
                                            """);
            this.put("money error",
                                            """
                                            [ERROR] Exchanging done with error.
                                            [ERROR] The coins amount [%d] of desired value (%d) in not enough to exchange (%d) cent coin
                                            """);
            this.put("change error",
                                            """
                                            [ERROR] Exchanging done with error.
                                            [ERROR] The coins amount [%d] of change (%d) coins
                                            [ERROR] is not enough to exchange (%d) cent coin
                                            """);
        }
    };

    private int coinToExchange;
    private int coinValueToBeExchangedBy;

    private boolean coinPassed = false;
    private boolean exchangerWorking = false;

    private final Lock lock = new ReentrantLock();
    private final Condition PassedCoin;
    private final Condition FinishExchanging;

    public CoinExchanger() {
        this.PassedCoin = this.lock.newCondition();
        this.FinishExchanging = this.lock.newCondition();
        this.generateInitialCoinPool();
    }

    private static int generateNumber(int min, int max) {
        return (int)Math.floor(Math.random() * (max - min + 1) + min);
    }

    private void printHelpMessage() {
        String supportedValues = values.toString();
        System.out.println("[INFO]  Entered coin for exchanging: " + this.coinToExchange);
        System.out.println("[INFO]  Supported coin values for exchanging: " + supportedValues);
    }

    private void generateInitialCoinPool() {
        for (int nominal : values) {
            this.coinsPool.put(nominal, generateNumber(5, 30));
        }
    }

    private static void clearScreen() {
        System.out.print("\u001b[H\u001b[2J");
        System.out.flush();
    }

    public void InputCoinToBeExchanged() throws InterruptedException {
        this.lock.lock();

        try {
            while(this.exchangerWorking) {
                this.FinishExchanging.await();
            }

            int nominalIdx = generateNumber(0, 6);
            this.coinToExchange = values.get(nominalIdx);
            this.coinsPool.put(this.coinToExchange, this.coinsPool.get(this.coinToExchange) + 1);

            clearScreen();
            System.out.print("\n[INFO] Entered coin for exchanging: " + this.coinToExchange);

            this.coinPassed = true;
            this.exchangerWorking = true;

            this.PassedCoin.signal();
        } finally {
            this.lock.unlock();
        }

    }

    private Optional<String> validatedCoinValue(Integer desiredNominal) {
        if (!values.contains(desiredNominal)) {
            return Optional.of(
                    "[ERROR] Entered coin value for exchanging " +
                     this.coinValueToBeExchangedBy +
                     " is not supported. Try again.");
        } else {
            return this.coinToExchange < desiredNominal ? Optional.of(
                    "[ERROR] Entered wanted value for exchanging is bigger than passed coin") : Optional.empty();
        }
    }

    private void inputCoinValueToBeExchangedBy() {
        while(true) {
            try {
                Scanner scanner = new Scanner(System.in);

                System.out.println("\n[INPUT] Enter wanted coin value to exchange inputted coin");
                this.coinValueToBeExchangedBy = scanner.nextInt();

                Optional<String> validationError = this.validatedCoinValue(this.coinValueToBeExchangedBy);
                if (validationError.isEmpty()) {
                    return;
                }

                System.out.println(validationError.get());
                this.printHelpMessage();
            } catch (InputMismatchException var3) {
                System.out.println("[ERROR] Entered value is not a coin value. Mismatch error. Try again.");
                this.printHelpMessage();
            }
        }
    }

    public void Exchange() throws InterruptedException {
        this.lock.lock();

        try {
            while(!this.coinPassed) {
                this.PassedCoin.await();
            }

            boolean isSuccessfulChange = false;

            this.inputCoinValueToBeExchangedBy();

            int coinsNum = this.coinToExchange / this.coinValueToBeExchangedBy;
            int change = this.coinToExchange % this.coinValueToBeExchangedBy;

            if (this.coinsPool.get(this.coinValueToBeExchangedBy) >= coinsNum) {
                if (change == 0) {
                    System.out.printf(
                            exchangerMessages.get("success no change"),
                            coinsNum, this.coinValueToBeExchangedBy);
                    isSuccessfulChange = true;
                } else if (this.coinsPool.get(change) < 1) {
                    System.out.printf(
                            exchangerMessages.get("change error"),
                            this.coinsPool.get(this.coinValueToBeExchangedBy),
                            this.coinValueToBeExchangedBy, this.coinToExchange);
                } else {
                    System.out.printf(
                            exchangerMessages.get("success"),
                            coinsNum, this.coinValueToBeExchangedBy, change);
                    isSuccessfulChange = true;
                }
            } else if (change == 0) {
                System.out.printf(
                        exchangerMessages.get("money error"),
                        this.coinsPool.get(this.coinValueToBeExchangedBy),
                        this.coinValueToBeExchangedBy, this.coinToExchange);
            } else if (this.coinsPool.get(change) < 1) {
                System.out.printf(
                        exchangerMessages.get("money and change error"),
                        this.coinsPool.get(this.coinValueToBeExchangedBy), this.coinValueToBeExchangedBy,
                        this.coinsPool.get(change), change, this.coinToExchange);
            }

            if (!isSuccessfulChange) {
                this.coinsPool.put(this.coinToExchange, this.coinsPool.get(this.coinToExchange) - 1);
                System.out.printf("[ERROR] Returning back your coin %d\n", this.coinToExchange);
            } else {
                this.coinsPool.put(this.coinValueToBeExchangedBy,
                        this.coinsPool.get(this.coinValueToBeExchangedBy) - coinsNum);
                if (change != 0) {
                    this.coinsPool.put(change, this.coinsPool.get(change) - 1);
                }
            }

            this.exchangerWorking = false;
            this.coinPassed = false;

            this.FinishExchanging.signal();
        } finally {
            this.lock.unlock();
        }
    }
}
