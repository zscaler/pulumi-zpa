using System.Collections.Generic;
using Pulumi;
using Zpa = zscaler.PulumiPackage.Zpa;

await Deployment.RunAsync(() =>
{
	var group = new Zpa.ZPASegmentGroup("segment-group-example", new()
	{
		  Name = "Pulumi Segment Group",
        Description = "Pulumi Segment Group",
        Enabled = true,
        TcpKeepAliveEnabled = "1",
	});
});