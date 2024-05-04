# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetEnrollmentCertResult',
    'AwaitableGetEnrollmentCertResult',
    'get_enrollment_cert',
    'get_enrollment_cert_output',
]

@pulumi.output_type
class GetEnrollmentCertResult:
    """
    A collection of values returned by getEnrollmentCert.
    """
    def __init__(__self__, allow_signing=None, certificate=None, client_cert_type=None, cname=None, creation_time=None, csr=None, description=None, id=None, issued_by=None, issued_to=None, microtenant_id=None, modified_by=None, modified_time=None, name=None, parent_cert_id=None, parent_cert_name=None, private_key=None, private_key_present=None, serial_no=None, valid_from_in_epoch_sec=None, valid_to_in_epoch_sec=None, zrsa_encrypted_private_key=None, zrsa_encrypted_session_key=None):
        if allow_signing and not isinstance(allow_signing, bool):
            raise TypeError("Expected argument 'allow_signing' to be a bool")
        pulumi.set(__self__, "allow_signing", allow_signing)
        if certificate and not isinstance(certificate, str):
            raise TypeError("Expected argument 'certificate' to be a str")
        pulumi.set(__self__, "certificate", certificate)
        if client_cert_type and not isinstance(client_cert_type, str):
            raise TypeError("Expected argument 'client_cert_type' to be a str")
        pulumi.set(__self__, "client_cert_type", client_cert_type)
        if cname and not isinstance(cname, str):
            raise TypeError("Expected argument 'cname' to be a str")
        pulumi.set(__self__, "cname", cname)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if csr and not isinstance(csr, str):
            raise TypeError("Expected argument 'csr' to be a str")
        pulumi.set(__self__, "csr", csr)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if issued_by and not isinstance(issued_by, str):
            raise TypeError("Expected argument 'issued_by' to be a str")
        pulumi.set(__self__, "issued_by", issued_by)
        if issued_to and not isinstance(issued_to, str):
            raise TypeError("Expected argument 'issued_to' to be a str")
        pulumi.set(__self__, "issued_to", issued_to)
        if microtenant_id and not isinstance(microtenant_id, str):
            raise TypeError("Expected argument 'microtenant_id' to be a str")
        pulumi.set(__self__, "microtenant_id", microtenant_id)
        if modified_by and not isinstance(modified_by, str):
            raise TypeError("Expected argument 'modified_by' to be a str")
        pulumi.set(__self__, "modified_by", modified_by)
        if modified_time and not isinstance(modified_time, str):
            raise TypeError("Expected argument 'modified_time' to be a str")
        pulumi.set(__self__, "modified_time", modified_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parent_cert_id and not isinstance(parent_cert_id, str):
            raise TypeError("Expected argument 'parent_cert_id' to be a str")
        pulumi.set(__self__, "parent_cert_id", parent_cert_id)
        if parent_cert_name and not isinstance(parent_cert_name, str):
            raise TypeError("Expected argument 'parent_cert_name' to be a str")
        pulumi.set(__self__, "parent_cert_name", parent_cert_name)
        if private_key and not isinstance(private_key, str):
            raise TypeError("Expected argument 'private_key' to be a str")
        pulumi.set(__self__, "private_key", private_key)
        if private_key_present and not isinstance(private_key_present, bool):
            raise TypeError("Expected argument 'private_key_present' to be a bool")
        pulumi.set(__self__, "private_key_present", private_key_present)
        if serial_no and not isinstance(serial_no, str):
            raise TypeError("Expected argument 'serial_no' to be a str")
        pulumi.set(__self__, "serial_no", serial_no)
        if valid_from_in_epoch_sec and not isinstance(valid_from_in_epoch_sec, str):
            raise TypeError("Expected argument 'valid_from_in_epoch_sec' to be a str")
        pulumi.set(__self__, "valid_from_in_epoch_sec", valid_from_in_epoch_sec)
        if valid_to_in_epoch_sec and not isinstance(valid_to_in_epoch_sec, str):
            raise TypeError("Expected argument 'valid_to_in_epoch_sec' to be a str")
        pulumi.set(__self__, "valid_to_in_epoch_sec", valid_to_in_epoch_sec)
        if zrsa_encrypted_private_key and not isinstance(zrsa_encrypted_private_key, str):
            raise TypeError("Expected argument 'zrsa_encrypted_private_key' to be a str")
        pulumi.set(__self__, "zrsa_encrypted_private_key", zrsa_encrypted_private_key)
        if zrsa_encrypted_session_key and not isinstance(zrsa_encrypted_session_key, str):
            raise TypeError("Expected argument 'zrsa_encrypted_session_key' to be a str")
        pulumi.set(__self__, "zrsa_encrypted_session_key", zrsa_encrypted_session_key)

    @property
    @pulumi.getter(name="allowSigning")
    def allow_signing(self) -> bool:
        return pulumi.get(self, "allow_signing")

    @property
    @pulumi.getter
    def certificate(self) -> str:
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="clientCertType")
    def client_cert_type(self) -> str:
        return pulumi.get(self, "client_cert_type")

    @property
    @pulumi.getter
    def cname(self) -> str:
        return pulumi.get(self, "cname")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def csr(self) -> str:
        return pulumi.get(self, "csr")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="issuedBy")
    def issued_by(self) -> str:
        return pulumi.get(self, "issued_by")

    @property
    @pulumi.getter(name="issuedTo")
    def issued_to(self) -> str:
        return pulumi.get(self, "issued_to")

    @property
    @pulumi.getter(name="microtenantId")
    def microtenant_id(self) -> Optional[str]:
        return pulumi.get(self, "microtenant_id")

    @property
    @pulumi.getter(name="modifiedBy")
    def modified_by(self) -> str:
        return pulumi.get(self, "modified_by")

    @property
    @pulumi.getter(name="modifiedTime")
    def modified_time(self) -> str:
        return pulumi.get(self, "modified_time")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentCertId")
    def parent_cert_id(self) -> str:
        return pulumi.get(self, "parent_cert_id")

    @property
    @pulumi.getter(name="parentCertName")
    def parent_cert_name(self) -> str:
        return pulumi.get(self, "parent_cert_name")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> str:
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="privateKeyPresent")
    def private_key_present(self) -> bool:
        return pulumi.get(self, "private_key_present")

    @property
    @pulumi.getter(name="serialNo")
    def serial_no(self) -> str:
        return pulumi.get(self, "serial_no")

    @property
    @pulumi.getter(name="validFromInEpochSec")
    def valid_from_in_epoch_sec(self) -> str:
        return pulumi.get(self, "valid_from_in_epoch_sec")

    @property
    @pulumi.getter(name="validToInEpochSec")
    def valid_to_in_epoch_sec(self) -> str:
        return pulumi.get(self, "valid_to_in_epoch_sec")

    @property
    @pulumi.getter(name="zrsaEncryptedPrivateKey")
    def zrsa_encrypted_private_key(self) -> str:
        return pulumi.get(self, "zrsa_encrypted_private_key")

    @property
    @pulumi.getter(name="zrsaEncryptedSessionKey")
    def zrsa_encrypted_session_key(self) -> str:
        return pulumi.get(self, "zrsa_encrypted_session_key")


class AwaitableGetEnrollmentCertResult(GetEnrollmentCertResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnrollmentCertResult(
            allow_signing=self.allow_signing,
            certificate=self.certificate,
            client_cert_type=self.client_cert_type,
            cname=self.cname,
            creation_time=self.creation_time,
            csr=self.csr,
            description=self.description,
            id=self.id,
            issued_by=self.issued_by,
            issued_to=self.issued_to,
            microtenant_id=self.microtenant_id,
            modified_by=self.modified_by,
            modified_time=self.modified_time,
            name=self.name,
            parent_cert_id=self.parent_cert_id,
            parent_cert_name=self.parent_cert_name,
            private_key=self.private_key,
            private_key_present=self.private_key_present,
            serial_no=self.serial_no,
            valid_from_in_epoch_sec=self.valid_from_in_epoch_sec,
            valid_to_in_epoch_sec=self.valid_to_in_epoch_sec,
            zrsa_encrypted_private_key=self.zrsa_encrypted_private_key,
            zrsa_encrypted_session_key=self.zrsa_encrypted_session_key)


def get_enrollment_cert(id: Optional[str] = None,
                        microtenant_id: Optional[str] = None,
                        name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnrollmentCertResult:
    """
    * [Official documentation](https://help.zscaler.com/zpa/about-enrollment-ca-certificates)
    * [API documentation](https://help.zscaler.com/zpa/obtaining-enrollment-certificate-details-using-api)

    Use the **zpa_enrollment_cert** data source to get information about all configured enrollment certificate details created in the Zscaler Private Access cloud. This data source is required when creating provisioning key resources.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_zpa as zpa

    root = zpa.get_enrollment_cert(name="Root")
    client = zpa.get_enrollment_cert(name="Client")
    connector = zpa.get_enrollment_cert(name="Connector")
    service_edge = zpa.get_enrollment_cert(name="Service Edge")
    isolation_client = zpa.get_enrollment_cert(name="Isolation Client")
    ```
    <!--End PulumiCodeChooser -->
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['microtenantId'] = microtenant_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zpa:index/getEnrollmentCert:getEnrollmentCert', __args__, opts=opts, typ=GetEnrollmentCertResult).value

    return AwaitableGetEnrollmentCertResult(
        allow_signing=pulumi.get(__ret__, 'allow_signing'),
        certificate=pulumi.get(__ret__, 'certificate'),
        client_cert_type=pulumi.get(__ret__, 'client_cert_type'),
        cname=pulumi.get(__ret__, 'cname'),
        creation_time=pulumi.get(__ret__, 'creation_time'),
        csr=pulumi.get(__ret__, 'csr'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        issued_by=pulumi.get(__ret__, 'issued_by'),
        issued_to=pulumi.get(__ret__, 'issued_to'),
        microtenant_id=pulumi.get(__ret__, 'microtenant_id'),
        modified_by=pulumi.get(__ret__, 'modified_by'),
        modified_time=pulumi.get(__ret__, 'modified_time'),
        name=pulumi.get(__ret__, 'name'),
        parent_cert_id=pulumi.get(__ret__, 'parent_cert_id'),
        parent_cert_name=pulumi.get(__ret__, 'parent_cert_name'),
        private_key=pulumi.get(__ret__, 'private_key'),
        private_key_present=pulumi.get(__ret__, 'private_key_present'),
        serial_no=pulumi.get(__ret__, 'serial_no'),
        valid_from_in_epoch_sec=pulumi.get(__ret__, 'valid_from_in_epoch_sec'),
        valid_to_in_epoch_sec=pulumi.get(__ret__, 'valid_to_in_epoch_sec'),
        zrsa_encrypted_private_key=pulumi.get(__ret__, 'zrsa_encrypted_private_key'),
        zrsa_encrypted_session_key=pulumi.get(__ret__, 'zrsa_encrypted_session_key'))


@_utilities.lift_output_func(get_enrollment_cert)
def get_enrollment_cert_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                               microtenant_id: Optional[pulumi.Input[Optional[str]]] = None,
                               name: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEnrollmentCertResult]:
    """
    * [Official documentation](https://help.zscaler.com/zpa/about-enrollment-ca-certificates)
    * [API documentation](https://help.zscaler.com/zpa/obtaining-enrollment-certificate-details-using-api)

    Use the **zpa_enrollment_cert** data source to get information about all configured enrollment certificate details created in the Zscaler Private Access cloud. This data source is required when creating provisioning key resources.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_zpa as zpa

    root = zpa.get_enrollment_cert(name="Root")
    client = zpa.get_enrollment_cert(name="Client")
    connector = zpa.get_enrollment_cert(name="Connector")
    service_edge = zpa.get_enrollment_cert(name="Service Edge")
    isolation_client = zpa.get_enrollment_cert(name="Isolation Client")
    ```
    <!--End PulumiCodeChooser -->
    """
    ...
