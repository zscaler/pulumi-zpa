using System.Collections.Generic;
using Pulumi;
using Zpa = zscaler.PulumiPackage.Zpa;

await Deployment.RunAsync(() =>

 {
     // Create a App Connector Group
     var appConnectorGroup = new Zpa.ZPAAppConnectorGroup("app-connector-group-example", new()
     {
         Name = "Pulumi App Connector Group",
         Description = "Pulumi App Connector Group",
         Enabled = true,
         CityCountry = "San Jose, CA",
         CountryCode = "US",
         Latitude = "37.338",
         Longitude = "-121.8863",
         Location = "San Jose, CA, US",
         OverrideVersionProfile = true,
         UpgradeDay = "SUNDAY",
         UpgradeTimeInSecs = "66600",
         VersionProfileId = "2",
         DnsQueryType = "IPV4_IPV6",
     });

     // Create a Server Group resource with Dynamic Discovery Enabled
     var serverGroup = new Zpa.ZPAServerGroup("server-group-example", new()
     {
         Name = "Pulumi Server Group",
         Description = "Pulumi Server Group",
         Enabled = true,
         DynamicDiscovery = true,
         AppConnectorGroups = new[]
         {
             new Zpa.Inputs.ZPAServerGroupAppConnectorGroupArgs
             {
                 Ids = new[]
                 {
                     appConnectorGroup.Id,
                 },
             },
         },
     }, new CustomResourceOptions
     {
         DependsOn = new[]
         {
             appConnectorGroup,
         },
     });

 });