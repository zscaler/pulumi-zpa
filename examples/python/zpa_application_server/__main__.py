"""A Python Pulumi program"""

import zscaler_pulumi_zpa as zpa

applicationServer = zpa.server.ApplicationServer("application-server-example",
    name = "Pulumi Application Server",
    description = "Pulumi Application Server",
    enabled = True,
    address = "192.168.100.1"
)

