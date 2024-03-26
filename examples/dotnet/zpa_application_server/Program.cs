using System.Collections.Generic;
using Pulumi;
using Zpa = zscaler.PulumiPackage.Zpa;

await Deployment.RunAsync(() =>
{
	var appServer = new Zpa.ZPAApplicationServer("application-server-example", new()
	{
		  Name = "Pulumi Application Server",
        Description = "Pulumi Application Server",
        Enabled = true,
        Address = "192.168.100.1",
	});
});