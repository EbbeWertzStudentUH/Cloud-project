package be.ebbew;

import jakarta.jws.WebMethod;
import jakarta.jws.WebParam;
import jakarta.jws.WebService;
import jakarta.xml.ws.Endpoint;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicInteger;
import io.github.cdimascio.dotenv.Dotenv;


@WebService
public class Main {

    private final ConcurrentHashMap<String, AtomicInteger> outgoingRequests = new ConcurrentHashMap<>();

    @WebMethod
    public void registerOutgoingRequest(@WebParam(name = "serviceType") String serviceType,
                                        @WebParam(name = "identifier") String identifier) {
        String key = serviceType + ":" + identifier;
        outgoingRequests.computeIfAbsent(key, k -> new AtomicInteger(0)).incrementAndGet();
    }

    @WebMethod
    public OutgoingRequestCountsResponse getOutgoingRequestCounts() {
        Map<String, Integer> counts = outgoingRequests.entrySet().stream()
                .collect(ConcurrentHashMap::new,
                        (map, entry) -> map.put(entry.getKey(), entry.getValue().get()),
                        ConcurrentHashMap::putAll);

        return new OutgoingRequestCountsResponse(counts);
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
