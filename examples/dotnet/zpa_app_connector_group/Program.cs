using System.Collections.Generic;
using Pulumi;
using Zpa = zscaler.PulumiPackage.Zpa;

await Deployment.RunAsync(() =>
{
	var group = new Zpa.ZPAAppConnectorGroup("app-connector-group-example", new()
	{
      Name = "Pulumi App Connector Group",
      Description = "Pulumi App Connector Group",
      CityCountry = "San Jose, CA",
      CountryCode = "US",
      DnsQueryType = "IPV4_IPV6",
      Enabled = true,
      Latitude = "37.338",
      Longitude = "-121.8863",
      Location = "San Jose, CA, US",
      OverrideVersionProfile = true,
      UpgradeDay = "SUNDAY",
      UpgradeTimeInSecs = "66600",
      VersionProfileId = "2",
	});
});