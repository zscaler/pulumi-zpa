"""A Python Pulumi program"""

import zscaler_pulumi_zpa as zpa

segmentGroup = zpa.segmentgroup.SegmentGroup("segment-group-example",
    name = "Pulumi Segment Group Python",
    description = "Pulumi Segment Group Python",
    enabled = True,
    tcp_keep_alive_enabled = "1"
)