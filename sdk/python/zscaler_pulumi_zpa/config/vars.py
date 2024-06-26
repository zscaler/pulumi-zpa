# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('zpa')


class _ExportableConfig(types.ModuleType):
    @property
    def zpa_client_id(self) -> Optional[str]:
        """
        zpa client id
        """
        return __config__.get('zpaClientId') or _utilities.get_env('ZPA_CLIENT_ID')

    @property
    def zpa_client_secret(self) -> Optional[str]:
        """
        zpa client secret
        """
        return __config__.get('zpaClientSecret') or _utilities.get_env('ZPA_CLIENT_SECRET')

    @property
    def zpa_cloud(self) -> Optional[str]:
        """
        Cloud to use PRODUCTION, ZPATWO, BETA, GOV, GOVUS, PREVIEW, DEV, QA, QA2
        """
        return __config__.get('zpaCloud') or _utilities.get_env('ZPA_CLOUD')

    @property
    def zpa_customer_id(self) -> Optional[str]:
        """
        zpa customer id
        """
        return __config__.get('zpaCustomerId') or _utilities.get_env('ZPA_CUSTOMER_ID')

