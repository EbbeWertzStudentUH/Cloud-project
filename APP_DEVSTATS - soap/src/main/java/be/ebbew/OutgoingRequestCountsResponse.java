package be.ebbew;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

public class OutgoingRequestCountsResponse {
    private List<OutgoingRequestCount> counts;

    public OutgoingRequestCountsResponse() {
        this.counts = new ArrayList<>();
    }

    public OutgoingRequestCountsResponse(Map<String, Integer> countsMap, Map<String, Integer> avgRequestsPerMinute) {
        this.counts = new ArrayList<>();
        for(String key : countsMap.keySet()) {
            String[] parts = key.split(":", 3);
            String serviceName = parts[0];
            String serviceType = parts[1];
            String identifier = parts[2];
            int count = countsMap.get(key);
            int rpm = avgRequestsPerMinute.get(key);
            this.counts.add(new OutgoingRequestCount(serviceType, identifier, count, serviceName, rpm));
        }
    }

    public List<OutgoingRequestCount> getCounts() {
        return counts;
    }

    public void setCounts(List<OutgoingRequestCount> counts) {
        this.counts = counts;
    }
}
