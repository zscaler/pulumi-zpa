import * as zpa from "@bdzscaler/pulumi-zpa";

// Create a App Connector Group
const serviceEdgeGroup = new zpa.ServiceEdgeGroup("service-edge-group-example", {
    name: "Pulumi Service Edge Group",
    description: "Pulumi Service Edge Group",
    enabled: true,
    isPublic: true,
    cityCountry: "San Jose, CA",
    countryCode: "US",
    latitude: "37.3382082",
    longitude: "-121.8863",
    location: "San Jose, CA, US",
    overrideVersionProfile: true,
    upgradeDay: "SUNDAY",
    upgradeTimeInSecs: "66600",
    versionProfileId: "2",
 })

 export const group = serviceEdgeGroup.id