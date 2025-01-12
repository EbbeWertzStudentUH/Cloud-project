package be.ebbew;

public class AvarageTime {
    private long totalRequestTime = 0;
    private int count = 0;

    public synchronized void addRequestTime(long time) {
        totalRequestTime += time;
        count++;
    }

    public synchronized int getAverageTime() {
        double avgDouble = count == 0 ? 0 : (double) totalRequestTime / count;
        return (int) Math.round(avgDouble);
    }
}
