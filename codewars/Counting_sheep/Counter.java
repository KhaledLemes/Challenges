public class Counter {
    public static int countSheeps(Boolean[] arrayOfSheeps) {
      int count = 0;
      for (Boolean present : arrayOfSheeps) {
        if (present != null && present) {
          count++;
        }
      }
      return count;
    }
}