import * as zpa from "@bdzscaler/pulumi-zpa";

// Create a App Connector Group
const appConnectorGroup = new zpa.ConnectorGroup("app-connector-group-example", {
    name: "Pulumi App Connector Group",
    description: "Pulumi App Connector Group",
    enabled: true,
    cityCountry: "San Jose, CA",
    countryCode: "US",
    latitude: "37.3382082",
    longitude: "-121.8863",
    location: "San Jose, CA, US",
    overrideVersionProfile: true,
    upgradeDay: "SUNDAY",
    upgradeTimeInSecs: "66600",
    versionProfileId: "2",
    dnsQueryType: "IPV4_IPV6",
 })

 export const groupId = appConnectorGroup.id