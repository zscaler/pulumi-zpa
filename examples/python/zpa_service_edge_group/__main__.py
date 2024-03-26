"""A Python Pulumi program"""

import zscaler_pulumi_zpa as zpa

serviceEdgeGroup = zpa.serviceedge.ServiceEdgeGroup("service-edge-group-example",
    name = "Pulumi Service Edge Group",
    description = "Pulumi Service Edge Group",
    enabled = True,
    is_public = True,
    country_code = "US",
    city_country = "San Jose, US",
    location = "San Jose, CA, USA",
    latitude = "37.3382082",
    longitude = "-121.8863286",
    upgrade_day = "SUNDAY",
    upgrade_time_in_secs = "66600",
    override_version_profile = True,
    version_profile_id = "2",
)