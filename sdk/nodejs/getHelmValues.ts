// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Helm values of cilium
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cilium from "@pulumi/cilium";
 * import * as local from "@pulumi/local";
 *
 * const exampleHelmValues = cilium.getHelmValues({});
 * const exampleFile = new local.File("exampleFile", {
 *     content: exampleHelmValues.then(exampleHelmValues => exampleHelmValues.yaml),
 *     filename: `${path.module}/values.yaml`,
 * });
 * ```
 */
export function getHelmValues(args?: GetHelmValuesArgs, opts?: pulumi.InvokeOptions): Promise<GetHelmValuesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("cilium:index/getHelmValues:getHelmValues", {
        "namespace": args.namespace,
        "release": args.release,
    }, opts);
}

/**
 * A collection of arguments for invoking getHelmValues.
 */
export interface GetHelmValuesArgs {
    /**
     * Namespace of cilium (Default: `kube-system`).
     */
    namespace?: string;
    /**
     * Helm release (Default: `cilium.Install`).
     */
    release?: string;
}

/**
 * A collection of values returned by getHelmValues.
 */
export interface GetHelmValuesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Namespace of cilium (Default: `kube-system`).
     */
    readonly namespace?: string;
    /**
     * Helm release (Default: `cilium.Install`).
     */
    readonly release?: string;
    /**
     * Yaml output
     */
    readonly yaml: string;
}
/**
 * Helm values of cilium
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cilium from "@pulumi/cilium";
 * import * as local from "@pulumi/local";
 *
 * const exampleHelmValues = cilium.getHelmValues({});
 * const exampleFile = new local.File("exampleFile", {
 *     content: exampleHelmValues.then(exampleHelmValues => exampleHelmValues.yaml),
 *     filename: `${path.module}/values.yaml`,
 * });
 * ```
 */
export function getHelmValuesOutput(args?: GetHelmValuesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHelmValuesResult> {
    return pulumi.output(args).apply((a: any) => getHelmValues(a, opts))
}

/**
 * A collection of arguments for invoking getHelmValues.
 */
export interface GetHelmValuesOutputArgs {
    /**
     * Namespace of cilium (Default: `kube-system`).
     */
    namespace?: pulumi.Input<string>;
    /**
     * Helm release (Default: `cilium.Install`).
     */
    release?: pulumi.Input<string>;
}