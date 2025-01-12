package be.ebbew;

import jakarta.jws.WebMethod;
import jakarta.jws.WebParam;
import jakarta.jws.WebService;
import jakarta.xml.ws.Endpoint;

import java.time.Instant;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

import io.github.cdimascio.dotenv.Dotenv;


@WebService
public class Main {

    private final ConcurrentHashMap<String, List<Long>> requestTimestamps = new ConcurrentHashMap<>();


    @WebMethod
    public void registerOutgoingRequest(@WebParam(name = "serviceType") String serviceType,
                                        @WebParam(name = "identifier") String identifier,
                                        @WebParam(name = "serviceName") String serviceName) {
        String key = serviceName + ":" + serviceType + ":" + identifier;
        requestTimestamps.computeIfAbsent(key, k -> Collections.synchronizedList(new ArrayList<>()))
                .add(Instant.now().toEpochMilli());
    }

    @WebMethod
    public OutgoingRequestCountsResponse getOutgoingRequestCounts() {
        Map<String, Integer> countsMap = Util.calculateCounts(requestTimestamps);
        Map<String, Integer> rpmMap = Util.calculateRMP(requestTimestamps);
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
