import * as zpa from "@bdzscaler/pulumi-zpa";

// Create a Application Server
const applicationServer = new zpa.ApplicationServer("application-server-example", {
    description: "Pulumi Application Server",
    enabled: true,
    address: "192.168.100.1"
 })

 export const server = applicationServer.id