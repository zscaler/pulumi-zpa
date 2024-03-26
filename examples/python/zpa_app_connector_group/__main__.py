"""A Python Pulumi program"""

import zscaler_pulumi_zpa as zpa

appConnectorGroup = zpa.connector.ConnectorGroup("app-connector-group-example",
    name = "Pulumi App Connector Group",
    description = "Pulumi App Connector Group",
    enabled = True,
    country_code = "US",
    city_country = "San Jose, US",
    location = "San Jose, CA, USA",
    latitude = "37.3382082",
    longitude = "-121.8863286",
    upgrade_day = "SUNDAY",
    upgrade_time_in_secs = "66600",
    override_version_profile = True,
    version_profile_id = "2",
    dns_query_type = "IPV4_IPV6",
)