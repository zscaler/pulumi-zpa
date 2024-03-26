"""A Python Pulumi program"""

import zscaler_pulumi_zpa as zpa

appConnectorGroup01 = zpa.connector.ConnectorGroup("app-connector-group",
    name = "Pulumi App Connector Group",
    description = "Pulumi App Connector Group",
    enabled = True,
    country_code = "CA",
    city_country = "Langley, CA",
    location = "Langley City, BC, Canada",
    latitude = "49.1041779",
    longitude = "-122.6603519",
    upgrade_day = "SUNDAY",
    upgrade_time_in_secs = "66600",
    override_version_profile = True,
    version_profile_id = "2",
    dns_query_type = "IPV4_IPV6",
)

provisioningKey = zpa.provisioningkey.ProvisioningKey("provisioning-key",
    name = "Pulumi Provisioning Key",
    association_type = "CONNECTOR_GRP",
    max_usage        = "10",
    enrollment_cert_id = "6573",
    zcomponent_id = appConnectorGroup01.id,
)