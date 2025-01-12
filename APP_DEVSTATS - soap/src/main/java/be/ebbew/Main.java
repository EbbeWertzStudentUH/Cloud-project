package be.ebbew;

import jakarta.jws.WebMethod;
import jakarta.jws.WebParam;
import jakarta.jws.WebService;
import jakarta.xml.ws.Endpoint;

import java.time.Instant;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Collections;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicInteger;
import io.github.cdimascio.dotenv.Dotenv;


@WebService
public class Main {

    private final ConcurrentHashMap<String, AtomicInteger> outgoingRequests = new ConcurrentHashMap<>();
    private final ConcurrentHashMap<String, List<Long>> requestTimestamps = new ConcurrentHashMap<>();


    @WebMethod
    public void registerOutgoingRequest(@WebParam(name = "serviceType") String serviceType,
                                        @WebParam(name = "identifier") String identifier,
                                        @WebParam(name = "serviceName") String serviceName) {
        String key = serviceName + ":" + serviceType + ":" + identifier;
        outgoingRequests.computeIfAbsent(key, k -> new AtomicInteger(0)).incrementAndGet();
        requestTimestamps.computeIfAbsent(key, k -> Collections.synchronizedList(new ArrayList<>()))
                .add(Instant.now().toEpochMilli());
    }

    @WebMethod
    public OutgoingRequestCountsResponse getOutgoingRequestCounts() {
        Map<String, Integer> countsMap = outgoingRequests.entrySet().stream()
                .collect(ConcurrentHashMap::new,
                        (map, entry) -> map.put(entry.getKey(), entry.getValue().get()),
                        ConcurrentHashMap::putAll);

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

        return new OutgoingRequestCountsResponse(countsMap, rpmMap);
    }




    public static void main(String[] args) {
        Dotenv dotenv = Dotenv.configure().load();
        String port = dotenv.get("LISTEN_PORT", "8080");
        String url = "http://localhost:" + port + "/devstats";
        System.out.println("Starting SOAP Service at " + url);
        Main service = new Main();
        Endpoint.publish(url, service);
    }
}
