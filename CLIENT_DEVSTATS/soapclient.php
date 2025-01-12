<?php

$env = parse_ini_file('.env');
$wsdl_url = $env['SOAP_WSDL_URL'];


try {
    $client = new SoapClient($wsdl_url, ['trace' => true,'exceptions' => true,]);
    $response = $client->__soapCall("getOutgoingStats", []);

    $counts = [];
    if (isset($response->return->counts)) {
        $countsData = is_array($response->return->counts) 
            ? $response->return->counts 
            : [$response->return->counts];
    
        foreach ($countsData as $count) {
            $counts[] = [
                'serviceName' => (string)$count->serviceName,
                'serviceType' => (string)$count->serviceType,
                'identifier' => (string)$count->identifier,
                'count' => (int)$count->count,
                'requestTime' => (int)$count->requestTime,
                'rpm' => (float)$count->rpm
            ];
        }
    }

    header('Content-Type: application/json');
    echo json_encode($counts);

} catch (SoapFault $e) {
    header('Content-Type: application/json');
    echo json_encode(['error' => $e->getMessage()]);
}

?>
