package be.ebbew;

public class OutgoingStats {
    private String serviceType;
    private String identifier;
    private String serviceName;
    private int count;
    private int rpm; // avarage requests per minute
    private int requestTime;

    public OutgoingStats() {
    }

    public OutgoingStats(String serviceType, String identifier, int count, String serviceName, int rpm, int requestTime) {
        this.serviceType = serviceType;
        this.identifier = identifier;
        this.count = count;
        this.serviceName = serviceName;
        this.rpm = rpm;
        this.requestTime = requestTime;
    }

    public int getRequestTime() {
        return requestTime;
    }

    public void setRequestTime(int requestTime) {
        this.requestTime = requestTime;
    }

    public String getServiceType() {
        return serviceType;
    }

    public int getRpm() {
        return rpm;
    }

    public void setRpm(int rpm) {
        this.rpm = rpm;
    }

    public String getServiceName() {
        return serviceName;
    }

    public void setServiceName(String serviceName) {
        this.serviceName = serviceName;
    }

    public void setServiceType(String serviceType) {
        this.serviceType = serviceType;
    }

    public String getIdentifier() {
        return identifier;
    }

    public void setIdentifier(String identifier) {
        this.identifier = identifier;
    }

    public int getCount() {
        return count;
    }

    public void setCount(int count) {
        this.count = count;
    }
}
