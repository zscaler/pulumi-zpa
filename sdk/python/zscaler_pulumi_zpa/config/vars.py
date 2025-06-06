# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

import types

__config__ = pulumi.Config('zpa')


class _ExportableConfig(types.ModuleType):
    @property
    def backoff(self) -> Optional[bool]:
        """
        Use exponential back off strategy for rate limits.
        """
        return __config__.get_bool('backoff')

    @property
    def client_id(self) -> Optional[str]:
        """
        zpa client id
        """
        return __config__.get('clientId') or _utilities.get_env('ZSCALER_CLIENT_ID')

    @property
    def client_secret(self) -> Optional[str]:
        """
        zpa client secret
        """
        return __config__.get('clientSecret') or _utilities.get_env('ZSCALER_CLIENT_SECRET')

    @property
    def customer_id(self) -> Optional[str]:
        """
        zpa customer id
        """
        return __config__.get('customerId') or _utilities.get_env('ZPA_CUSTOMER_ID')

    @property
    def http_proxy(self) -> Optional[str]:
        """
        Alternate HTTP proxy of scheme://hostname or scheme://hostname:port format
        """
        return __config__.get('httpProxy')

    @property
    def max_retries(self) -> Optional[int]:
        """
        maximum number of retries to attempt before erroring out.
        """
        return __config__.get_int('maxRetries')

    @property
    def max_wait_seconds(self) -> Optional[int]:
        """
        maximum seconds to wait when rate limit is hit. We use exponential backoffs when backoff is enabled.
        """
        return __config__.get_int('maxWaitSeconds')

    @property
    def microtenant_id(self) -> Optional[str]:
        """
        zpa microtenant ID
        """
        return __config__.get('microtenantId')

    @property
    def min_wait_seconds(self) -> Optional[int]:
        """
        minimum seconds to wait when rate limit is hit. We use exponential backoffs when backoff is enabled.
        """
        return __config__.get_int('minWaitSeconds')

    @property
    def parallelism(self) -> Optional[int]:
        """
        Number of concurrent requests to make within a resource where bulk operations are not possible. Take note of
        https://help.zscaler.com/zpa/understanding-rate-limiting.
        """
        return __config__.get_int('parallelism')

    @property
    def private_key(self) -> Optional[str]:
        """
        zpa private key
        """
        return __config__.get('privateKey') or _utilities.get_env('ZSCALER_PRIVATE_KEY')

    @property
    def request_timeout(self) -> Optional[int]:
        """
        Timeout for single request (in seconds) which is made to Zscaler, the default is `0` (means no limit is set). The
        maximum value can be `300`.
        """
        return __config__.get_int('requestTimeout')

    @property
    def use_legacy_client(self) -> Optional[bool]:
        """
        Enables interaction with the ZPA legacy API framework
        """
        return __config__.get_bool('useLegacyClient') or _utilities.get_env_bool('ZSCALER_USE_LEGACY_CLIENT')

    @property
    def vanity_domain(self) -> Optional[str]:
        """
        Zscaler Vanity Domain
        """
        return __config__.get('vanityDomain') or _utilities.get_env('ZSCALER_VANITY_DOMAIN')

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

    @property
    def zscaler_cloud(self) -> Optional[str]:
        """
        Zscaler Cloud Name
        """
        return __config__.get('zscalerCloud') or _utilities.get_env('ZSCALER_CLOUD')

