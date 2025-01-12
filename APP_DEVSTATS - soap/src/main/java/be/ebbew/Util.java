package be.ebbew;

import java.time.Instant;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class Util {

    public static Map<String, Integer> calculateCounts(ConcurrentHashMap<String, List<Long>> requestTimestamps){
        Map<String, Integer> countsMap = requestTimestamps.entrySet().stream()
                .collect(ConcurrentHashMap::new,
                        (map, entry) -> map.put(entry.getKey(), entry.getValue().size()),
                        ConcurrentHashMap::putAll);
        return countsMap;
    }

    public static Map<String, Integer> convertToAvgTimes(ConcurrentHashMap<String, AvarageTime> avarageTimesMap){
        Map<String, Integer> timeMap = avarageTimesMap.entrySet().stream()
                .collect(ConcurrentHashMap::new,
                        (map, entry) -> map.put(entry.getKey(), entry.getValue().getAverageTime()),
                        ConcurrentHashMap::putAll);
        return timeMap;
    }

    public static Map<String, Integer> calculateRMP(ConcurrentHashMap<String, List<Long>> requestTimestamps){
        Map<String, Integer> rpmMap = new HashMap<>();
        long currentTime = Instant.now().toEpochMilli();
        long oneMinuteMillis = 60000L;
        double decayFactor = 0.9;
        for (Map.Entry<String, List<Long>> entry : requestTimestamps.entrySet()) {
            String key = entry.getKey();
            List<Long> timestamps = entry.getValue();
            double rpm = 0.0;
            for (Long timestamp : timestamps) {
                long ageMillis = currentTime - timestamp;
                if (ageMillis <= oneMinuteMillis) {
                    double fractionalContribution = (oneMinuteMillis - ageMillis) / (double) oneMinuteMillis;
                    rpm += fractionalContribution;
                } else {
                    double decayContribution = Math.pow(decayFactor, (ageMillis - oneMinuteMillis) / 1000.0);
                    rpm += decayContribution;
                }
            }
            rpmMap.put(key, (int) Math.round(rpm)); // Store the calculated RPM for the key
        }
        return rpmMap;
    }
}
