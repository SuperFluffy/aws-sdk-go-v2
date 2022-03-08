// Code generated by smithy-go-codegen DO NOT EDIT.

// Package route53recoverycluster provides the API client, operations, and
// parameter types for Route53 Recovery Cluster.
//
// Welcome to the Routing Control (Recovery Cluster) API Reference Guide for Amazon
// Route 53 Application Recovery Controller. With Amazon Route 53 Application
// Recovery Controller, you can use routing control with extreme reliability to
// recover applications by rerouting traffic across Availability Zones or AWS
// Regions. Routing controls are simple on/off switches hosted on a highly
// available cluster in Application Recovery Controller. A cluster provides a set
// of five redundant Regional endpoints against which you can run API calls to get
// or update the state of routing controls. To implement failover, you set one
// routing control on and another one off, to reroute traffic from one Availability
// Zone or Amazon Web Services Region to another. Be aware that you must specify
// the Regional endpoints for a cluster when you work with API cluster operations
// to get or update routing control states in Application Recovery Controller. In
// addition, you must specify the US West (Oregon) Region for Application Recovery
// Controller API calls. For example, use the parameter region us-west-2 with AWS
// CLI commands. For more information, see  Get and update routing control states
// using the API
// (https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.update.api.html)
// in the Amazon Route 53 Application Recovery Controller Developer Guide. This API
// guide includes information about the API operations for how to get and update
// routing control states in Application Recovery Controller. You also must set up
// the structures to support routing controls: clusters and control panels. For
// more information about working with routing control in Application Recovery
// Controller, see the following:
//
// * To create clusters, routing controls, and
// control panels by using the control plane API for routing control, see the
// Recovery Control Configuration API Reference Guide for Amazon Route 53
// Application Recovery Controller
// (https://docs.aws.amazon.com/recovery-cluster/latest/api/).
//
// * Learn about the
// components in recovery control configuration, including clusters, routing
// controls, and control panels. For more information, see  Recovery control
// components
// (https://docs.aws.amazon.com/r53recovery/latest/dg/introduction-components.html#introduction-components-routing)
// in the Amazon Route 53 Application Recovery Controller Developer Guide.
//
// *
// Application Recovery Controller also provides readiness checks that run
// continually to help make sure that your applications are scaled and ready to
// handle failover traffic. For more information about the related API actions, see
// the Recovery Readiness API Reference Guide for Amazon Route 53 Application
// Recovery Controller
// (https://docs.aws.amazon.com/recovery-readiness/latest/api/).
//
// * For more
// information about creating resilient applications and preparing for recovery
// readiness with Application Recovery Controller, see the Amazon Route 53
// Application Recovery Controller Developer Guide
// (https://docs.aws.amazon.com/r53recovery/latest/dg/).
package route53recoverycluster
