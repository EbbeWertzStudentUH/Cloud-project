package be.ebbew;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

public class OutgoingStatsResponse {
    private List<OutgoingStats> counts;

    public OutgoingStatsResponse() {
        this.counts = new ArrayList<>();
    }

    public OutgoingStatsResponse(Map<String, Integer> countsMap, Map<String, Integer> avgRequestsPerMinute, Map<String, Integer> timeMap) {
        this.counts = new ArrayList<>();
        for(String key : countsMap.keySet()) {
            String[] parts = key.split(":", 3);
            String serviceName = parts[0];
            String serviceType = parts[1];
            String identifier = parts[2];
            int count = countsMap.get(key);
            int rpm = avgRequestsPerMinute.get(key);
            int avgtime = timeMap.getOrDefault(key, -1);
            this.counts.add(new OutgoingStats(serviceType, identifier, count, serviceName, rpm, avgtime));
        }
    }

    public List<OutgoingStats> getCounts() {
        return counts;
    }

    public void setCounts(List<OutgoingStats> counts) {
        this.counts = counts;
    }
}
