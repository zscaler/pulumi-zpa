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

 // Create a Server Group resource with Dynamic Discovery Enabled
 const serverGroup = new zpa.ServerGroup("server-group-example", {
    name: "Pulumi Server Group",
    description: "Pulumi Server Group",
    enabled: true,
    dynamicDiscovery: true,
    appConnectorGroups: [{
         ids: [appConnectorGroup.id],
     }],
 }, {
     dependsOn: [appConnectorGroup],
 });

 export const serverGroupId = serverGroup.id