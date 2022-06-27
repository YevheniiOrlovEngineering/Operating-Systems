package var3;


import java.util.Scanner;
import java.util.regex.Pattern;

import static java.lang.System.exit;

public class Main {
    private static void printMenu(String[] options){
        for (String option : options){
            System.out.println(option);
        }
        System.out.print("Choose your option : ");
    }

    private static void generateTree(BTree<String> bTree, int filesAmount, int step) {
        for (int i = 1; i < filesAmount * step; i+=step) {
            String id;
            if (i < 10) {
                id = "00" + i;
            } else if (i < 100) {
                id = "0" + i;
            } else {
                id = String.valueOf(i);
            }
            bTree.add(id);
        }
    }

    public static void main(String[] args) {
        String[] options =
        {
                "1 - Add file",
                "2 - Remove file",
                "3 - Find file",
                "4 - Print file system",
                "5 - Exit"
        };

        BTree<String> bTree = new BTree<>(2);
        Scanner scanner = new Scanner(System.in);
        int option = 0;

        generateTree(bTree, 20, 3);

        while (option!=5) {
            printMenu(options);
            try {
                option = scanner.nextInt();

                switch (option) {
                    case 1 -> addFile(bTree);
                    case 2 -> removeFile(bTree);
                    case 3 -> findFile(bTree);
                    case 4 -> printFileSystem(bTree);
                    case 5 -> exit(0);
                    default -> throw new RuntimeException(
                            "Please enter an integer value between 1 and " + options.length);
                }
            }
            catch (RuntimeException e){
                System.out.println("\n" + e.getMessage() + "\n");
            }
        }
    }

    // Options
    private static void addFile(BTree<String> bTree) {
        while (true) {
            Scanner scanner = new Scanner(System.in);
            System.out.print("\nEnter file name: ");
            String fileName = scanner.next();
            // file name validation
            boolean isOnlyThreeDigits = Pattern.matches("\\d\\d\\d", fileName);
            if (isOnlyThreeDigits) {
                bTree.add(fileName);
                System.out.println();
                break;
            } else {
                System.out.println("Entered file name doesn't consist of only three digits. Try again.");
            }
        }
    }

    private static void removeFile(BTree<String> bTree) {
        while (true) {
            Scanner scanner = new Scanner(System.in);
            System.out.print("\nEnter file name: ");
            String fileName = scanner.next();
            boolean isFileExists = bTree.contains(fileName);
            if (isFileExists) {
                bTree.remove(fileName);
                System.out.println();
                break;
            } else {
                System.out.println("Entered file name doesn't exist in a file system. Try again.");
            }
        }
    }

    private static void findFile(BTree<String> bTree) {
        while (true) {
            Scanner scanner = new Scanner(System.in);
            System.out.print("\nEnter file name: ");
            String fileName = scanner.next();
            boolean isFileExists = bTree.contains(fileName);
            if (isFileExists) {
                String path = bTree.search(fileName);
                System.out.println(path);
                System.out.println();
                break;
            } else {
                System.out.println("Entered file name doesn't exist in a file system. Try again.");
            }
        }
    }

    private static void printFileSystem(BTree<String> bTree) {
        System.out.println("\n" + bTree);
    }
}
