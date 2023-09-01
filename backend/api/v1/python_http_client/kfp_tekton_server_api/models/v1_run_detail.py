# coding: utf-8

"""
    Kubeflow Pipelines API

    This file contains REST API specification for Kubeflow Pipelines. The file is autogenerated from the swagger definition.

    Contact: kubeflow-pipelines@google.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from kfp_tekton_server_api.configuration import Configuration


class V1RunDetail(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'run': 'V1Run',
        'pipeline_runtime': 'V1PipelineRuntime'
    }

    attribute_map = {
        'run': 'run',
        'pipeline_runtime': 'pipeline_runtime'
    }

    def __init__(self, run=None, pipeline_runtime=None, local_vars_configuration=None):  # noqa: E501
        """V1RunDetail - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._run = None
        self._pipeline_runtime = None
        self.discriminator = None

        if run is not None:
            self.run = run
        if pipeline_runtime is not None:
            self.pipeline_runtime = pipeline_runtime

    @property
    def run(self):
        """Gets the run of this V1RunDetail.  # noqa: E501


        :return: The run of this V1RunDetail.  # noqa: E501
        :rtype: V1Run
        """
        return self._run

    @run.setter
    def run(self, run):
        """Sets the run of this V1RunDetail.


        :param run: The run of this V1RunDetail.  # noqa: E501
        :type run: V1Run
        """

        self._run = run

    @property
    def pipeline_runtime(self):
        """Gets the pipeline_runtime of this V1RunDetail.  # noqa: E501


        :return: The pipeline_runtime of this V1RunDetail.  # noqa: E501
        :rtype: V1PipelineRuntime
        """
        return self._pipeline_runtime

    @pipeline_runtime.setter
    def pipeline_runtime(self, pipeline_runtime):
        """Sets the pipeline_runtime of this V1RunDetail.


        :param pipeline_runtime: The pipeline_runtime of this V1RunDetail.  # noqa: E501
        :type pipeline_runtime: V1PipelineRuntime
        """

        self._pipeline_runtime = pipeline_runtime

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1RunDetail):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1RunDetail):
            return True

        return self.to_dict() != other.to_dict()
