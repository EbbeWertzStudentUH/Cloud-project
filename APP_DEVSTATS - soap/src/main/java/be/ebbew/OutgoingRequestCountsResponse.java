package be.ebbew;

import java.util.Map;

public class OutgoingRequestCountsResponse {
    private Map<String, Integer> counts;

    public OutgoingRequestCountsResponse() {
    }

    public OutgoingRequestCountsResponse(Map<String, Integer> counts) {
        this.counts = counts;
    }

    public Map<String, Integer> getCounts() {
        return counts;
    }

    public void setCounts(Map<String, Integer> counts) {
        this.counts = counts;
    }
}