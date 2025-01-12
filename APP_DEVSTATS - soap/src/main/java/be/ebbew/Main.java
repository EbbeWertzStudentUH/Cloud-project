package be.ebbew;

import jakarta.jws.WebMethod;
import jakarta.jws.WebParam;
import jakarta.jws.WebService;
import jakarta.xml.ws.Endpoint;

import java.time.Instant;
import java.util.*;
import java.util.concurrent.ConcurrentHashMap;

import io.github.cdimascio.dotenv.Dotenv;


@WebService
public class Main {

    // key = service info, val = timestamps
    private final ConcurrentHashMap<String, List<Long>> requestTimestamps = new ConcurrentHashMap<>();
    // key = request UUID, val = request info
    private final ConcurrentHashMap<String, String> uuidMap = new ConcurrentHashMap<>();
    // key = request UUID, val = timestamp
    private final ConcurrentHashMap<String, Long> requestStartTimes = new ConcurrentHashMap<>();
    // key = service info, val = stats
    private final ConcurrentHashMap<String, AvarageTime> requestAvarageTimes = new ConcurrentHashMap<>();


    @WebMethod
    public String registerOutgoingStart(@WebParam(name = "serviceType") String serviceType,
                                          @WebParam(name = "identifier") String identifier,
                                          @WebParam(name = "serviceName") String serviceName) {
        String key = serviceName + ":" + serviceType + ":" + identifier;
        String requestId = UUID.randomUUID().toString();
        long currentTime = Instant.now().toEpochMilli();
        requestTimestamps.computeIfAbsent(key, k -> Collections.synchronizedList(new ArrayList<>())).add(currentTime);
        requestStartTimes.put(requestId, currentTime);
        uuidMap.put(requestId, key);
        return requestId;
    }

    @WebMethod
    public void registerOutgoingEnd(@WebParam(name = "requestId") String requestId) {
        String key = uuidMap.get(requestId);
        Long startTime = requestStartTimes.remove(requestId);
        if (startTime == null) {
            throw new IllegalArgumentException("Invalid request ID: " + requestId);
        }
        long endTime = Instant.now().toEpochMilli();
        long requestDuration = endTime - startTime;
        requestAvarageTimes.computeIfAbsent(key, k -> new AvarageTime()).addRequestTime(requestDuration);
    }

    @WebMethod
    public OutgoingStatsResponse getOutgoingStats() {
        Map<String, Integer> countsMap = Util.calculateCounts(requestTimestamps);
        Map<String, Integer> rpmMap = Util.calculateRMP(requestTimestamps);
        Map<String, Integer> timeMap = Util.convertToAvgTimes(requestAvarageTimes);
        return new OutgoingStatsResponse(countsMap, rpmMap, timeMap);
    }


    public static void main(String[] args) {
        Dotenv dotenv = Dotenv.configure().load();
        String port = dotenv.get("LISTEN_PORT", "8080");
        String url = "http://0.0.0.0:" + port + "/devstats";
        System.out.println("Starting SOAP Service at " + url);
        Main service = new Main();
        Endpoint.publish(url, service);
    }
}
