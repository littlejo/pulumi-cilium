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
    'InstallCa',
]

@pulumi.output_type
class InstallCa(dict):
    def __init__(__self__, *,
                 crt: str,
                 key: str):
        pulumi.set(__self__, "crt", crt)
        pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def crt(self) -> str:
        return pulumi.get(self, "crt")

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")


