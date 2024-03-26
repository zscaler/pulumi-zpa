using System.Collections.Generic;
using Pulumi;
using Zpa = zscaler.PulumiPackage.Zpa;

await Deployment.RunAsync(() =>
{
	var serviceEdge = new Zpa.ZPAServiceEdgeGroup("service-edge-group-example", new()
	{
      Name = "Pulumi Service Edge Group",
      Description = "Pulumi Service Edge Group",
      Enabled = true,
      IsPublic = true,
      CountryCode = "US",
      CityCountry = "San Jose, US",
      Location = "San Jose, CA, USA",
      Latitude = "37.3382082",
      Longitude = "-121.8863286",
      UpgradeDay = "SUNDAY",
      UpgradeTimeInSecs = "66600",
      OverrideVersionProfile = true,
      VersionProfileId = "2",
	});
});