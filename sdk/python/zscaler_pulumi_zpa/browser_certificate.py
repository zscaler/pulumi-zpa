# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BrowserCertificateArgs', 'BrowserCertificate']

@pulumi.input_type
class BrowserCertificateArgs:
    def __init__(__self__, *,
                 cert_blob: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BrowserCertificate resource.
        :param pulumi.Input[str] cert_blob: The description of the certificate
        :param pulumi.Input[str] description: The description of the certificate
        :param pulumi.Input[str] microtenant_id: The unique identifier of the Microtenant
        :param pulumi.Input[str] name: The name of the certificate.
        """
        if cert_blob is not None:
            pulumi.set(__self__, "cert_blob", cert_blob)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if microtenant_id is not None:
            pulumi.set(__self__, "microtenant_id", microtenant_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="certBlob")
    def cert_blob(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the certificate
        """
        return pulumi.get(self, "cert_blob")

    @cert_blob.setter
    def cert_blob(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_blob", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the certificate
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the Microtenant
        """
        return pulumi.get(self, "microtenant_id")

    @microtenant_id.setter
    def microtenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "microtenant_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the certificate.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _BrowserCertificateState:
    def __init__(__self__, *,
                 cert_blob: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BrowserCertificate resources.
        :param pulumi.Input[str] cert_blob: The description of the certificate
        :param pulumi.Input[str] certificate: The certificate text in PEM format
        :param pulumi.Input[str] description: The description of the certificate
        :param pulumi.Input[str] microtenant_id: The unique identifier of the Microtenant
        :param pulumi.Input[str] name: The name of the certificate.
        """
        if cert_blob is not None:
            pulumi.set(__self__, "cert_blob", cert_blob)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if microtenant_id is not None:
            pulumi.set(__self__, "microtenant_id", microtenant_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="certBlob")
    def cert_blob(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the certificate
        """
        return pulumi.get(self, "cert_blob")

    @cert_blob.setter
    def cert_blob(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_blob", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate text in PEM format
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the certificate
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the Microtenant
        """
        return pulumi.get(self, "microtenant_id")

    @microtenant_id.setter
    def microtenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "microtenant_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the certificate.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class BrowserCertificate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cert_blob: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        * [Official documentation](https://help.zscaler.com/zpa/about-web-server-certificates)
        * [API documentation](https://help.zscaler.com/zpa/configuring-certificates-using-api)

        Use the **zpa_ba_certificate** creates a browser access certificate with a private key in the Zscaler Private Access cloud. This resource is required when creating a browser access application segment resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_zpa as zpa

        foo = zpa.get_ba_certificate(name="example.acme.com")
        ```

        ######### PASSWORDS OR RELATED CREDENTIALS ATTRIBUTES IN THIS FILE #########\\
        ######### ARE FOR EXAMPLE ONLY AND NOT USED IN PRODUCTION SYSTEMS ##########
        ```python
        import pulumi
        import zscaler_pulumi_zpa as zpa

        # ZPA Browser Access resource
        this = zpa.BrowserCertificate("this",
            cert_blob=\"\"\"-----BEGIN PRIVATE KEY-----
        MIIDyzCCArOgA
        -----END PRIVATE KEY-----
        -----BEGIN CERTIFICATE-----
        MIIDyzCCArOgAwIBAgIUekBD+iu64583B3u5ew7Bqj2O5cQwDQYJKoZIhvcNAQEL
        -----END CERTIFICATE-----

        \"\"\",
            description="server.example.com")
        ```

        ## Let's Encrypt Certbot

        This example demonstrates generatoring a domain certificate with letsencrypt
        certbot https://letsencrypt.org/getting-started/

        Use letsencrypt's certbot to generate domain certificates in RSA output mode.
        The generator's output corresponds to `BrowserCertificate` fields in the
        following manner.

        Zscaler Field          | Certbot file
        --------------------|--------------
        `certblob`          | `cert.pem`
        `certblob`          | `privkey.pem`

        ## Import

        This resource does not support importing.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cert_blob: The description of the certificate
        :param pulumi.Input[str] description: The description of the certificate
        :param pulumi.Input[str] microtenant_id: The unique identifier of the Microtenant
        :param pulumi.Input[str] name: The name of the certificate.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[BrowserCertificateArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        * [Official documentation](https://help.zscaler.com/zpa/about-web-server-certificates)
        * [API documentation](https://help.zscaler.com/zpa/configuring-certificates-using-api)

        Use the **zpa_ba_certificate** creates a browser access certificate with a private key in the Zscaler Private Access cloud. This resource is required when creating a browser access application segment resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_zpa as zpa

        foo = zpa.get_ba_certificate(name="example.acme.com")
        ```

        ######### PASSWORDS OR RELATED CREDENTIALS ATTRIBUTES IN THIS FILE #########\\
        ######### ARE FOR EXAMPLE ONLY AND NOT USED IN PRODUCTION SYSTEMS ##########
        ```python
        import pulumi
        import zscaler_pulumi_zpa as zpa

        # ZPA Browser Access resource
        this = zpa.BrowserCertificate("this",
            cert_blob=\"\"\"-----BEGIN PRIVATE KEY-----
        MIIDyzCCArOgA
        -----END PRIVATE KEY-----
        -----BEGIN CERTIFICATE-----
        MIIDyzCCArOgAwIBAgIUekBD+iu64583B3u5ew7Bqj2O5cQwDQYJKoZIhvcNAQEL
        -----END CERTIFICATE-----

        \"\"\",
            description="server.example.com")
        ```

        ## Let's Encrypt Certbot

        This example demonstrates generatoring a domain certificate with letsencrypt
        certbot https://letsencrypt.org/getting-started/

        Use letsencrypt's certbot to generate domain certificates in RSA output mode.
        The generator's output corresponds to `BrowserCertificate` fields in the
        following manner.

        Zscaler Field          | Certbot file
        --------------------|--------------
        `certblob`          | `cert.pem`
        `certblob`          | `privkey.pem`

        ## Import

        This resource does not support importing.

        :param str resource_name: The name of the resource.
        :param BrowserCertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BrowserCertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cert_blob: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 microtenant_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BrowserCertificateArgs.__new__(BrowserCertificateArgs)

            __props__.__dict__["cert_blob"] = cert_blob
            __props__.__dict__["description"] = description
            __props__.__dict__["microtenant_id"] = microtenant_id
            __props__.__dict__["name"] = name
            __props__.__dict__["certificate"] = None
        super(BrowserCertificate, __self__).__init__(
            'zpa:index/browserCertificate:BrowserCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cert_blob: Optional[pulumi.Input[str]] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            microtenant_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'BrowserCertificate':
        """
        Get an existing BrowserCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cert_blob: The description of the certificate
        :param pulumi.Input[str] certificate: The certificate text in PEM format
        :param pulumi.Input[str] description: The description of the certificate
        :param pulumi.Input[str] microtenant_id: The unique identifier of the Microtenant
        :param pulumi.Input[str] name: The name of the certificate.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BrowserCertificateState.__new__(_BrowserCertificateState)

        __props__.__dict__["cert_blob"] = cert_blob
        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["description"] = description
        __props__.__dict__["microtenant_id"] = microtenant_id
        __props__.__dict__["name"] = name
        return BrowserCertificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certBlob")
    def cert_blob(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the certificate
        """
        return pulumi.get(self, "cert_blob")

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[str]:
        """
        The certificate text in PEM format
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the certificate
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> pulumi.Output[str]:
        """
        The unique identifier of the Microtenant
        """
        return pulumi.get(self, "microtenant_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the certificate.
        """
        return pulumi.get(self, "name")

