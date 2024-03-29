// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Install resource for Cilium. This is equivalent to cilium cli: `cilium install`, `cilium upgrade` and `cilium uninstall`: It manages cilium helm chart
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as cilium from "@littlejo/cilium";
 * import * as kind from "@pulumi/kind";
 *
 * const examplekind_cluster = new kind.index.Kind_cluster("examplekind_cluster", {
 *     name: "test-cluster",
 *     kindConfig: [{
 *         kind: "Cluster",
 *         apiVersion: "kind.x-k8s.io/v1alpha4",
 *         node: [
 *             {
 *                 role: "control-plane",
 *             },
 *             {
 *                 role: "worker",
 *             },
 *         ],
 *         networking: [{
 *             disableDefaultCni: true,
 *         }],
 *     }],
 * });
 * const exampleInstall = new cilium.Install("exampleInstall", {
 *     sets: [
 *         "ipam.mode=kubernetes",
 *         "ipam.operator.replicas=1",
 *         "tunnel=vxlan",
 *     ],
 *     version: "1.14.5",
 * });
 * ```
 */
export class Install extends pulumi.CustomResource {
    /**
     * Get an existing Install resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstallState, opts?: pulumi.CustomResourceOptions): Install {
        return new Install(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cilium:index/install:Install';

    /**
     * Returns true if the given object is an instance of Install.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Install {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Install.__pulumiType;
    }

    /**
     * Datapath mode to use { tunnel | native | aws-eni | gke | azure | aks-byocni } (Default: `autodetected`).
     */
    public readonly dataPath!: pulumi.Output<string>;
    /**
     * Namespace in which to install (Default: `kube-system`).
     */
    public readonly namespace!: pulumi.Output<string>;
    /**
     * Helm chart repository to download Cilium charts from (Default: `https://helm.cilium.io`).
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * When upgrading, reset the helm values to the ones built into the chart (Default: `false`).
     */
    public readonly reset!: pulumi.Output<boolean>;
    /**
     * When upgrading, reuse the helm values from the latest release unless any overrides from are set from other flags. This option takes precedence over HelmResetValues (Default: `true`).
     */
    public readonly reuse!: pulumi.Output<boolean>;
    /**
     * Set helm values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2 (Default: `[]`).
     */
    public readonly sets!: pulumi.Output<string[]>;
    /**
     * values in raw yaml to pass to helm. (Default: `empty`).
     */
    public readonly values!: pulumi.Output<string>;
    /**
     * Version of Cilium (Default: `v1.14.5`).
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * Wait for Cilium status is ok (Default: `true`).
     */
    public readonly wait!: pulumi.Output<boolean>;

    /**
     * Create a Install resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstallArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstallArgs | InstallState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstallState | undefined;
            resourceInputs["dataPath"] = state ? state.dataPath : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["reset"] = state ? state.reset : undefined;
            resourceInputs["reuse"] = state ? state.reuse : undefined;
            resourceInputs["sets"] = state ? state.sets : undefined;
            resourceInputs["values"] = state ? state.values : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["wait"] = state ? state.wait : undefined;
        } else {
            const args = argsOrState as InstallArgs | undefined;
            resourceInputs["dataPath"] = args ? args.dataPath : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["reset"] = args ? args.reset : undefined;
            resourceInputs["reuse"] = args ? args.reuse : undefined;
            resourceInputs["sets"] = args ? args.sets : undefined;
            resourceInputs["values"] = args ? args.values : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["wait"] = args ? args.wait : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Install.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Install resources.
 */
export interface InstallState {
    /**
     * Datapath mode to use { tunnel | native | aws-eni | gke | azure | aks-byocni } (Default: `autodetected`).
     */
    dataPath?: pulumi.Input<string>;
    /**
     * Namespace in which to install (Default: `kube-system`).
     */
    namespace?: pulumi.Input<string>;
    /**
     * Helm chart repository to download Cilium charts from (Default: `https://helm.cilium.io`).
     */
    repository?: pulumi.Input<string>;
    /**
     * When upgrading, reset the helm values to the ones built into the chart (Default: `false`).
     */
    reset?: pulumi.Input<boolean>;
    /**
     * When upgrading, reuse the helm values from the latest release unless any overrides from are set from other flags. This option takes precedence over HelmResetValues (Default: `true`).
     */
    reuse?: pulumi.Input<boolean>;
    /**
     * Set helm values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2 (Default: `[]`).
     */
    sets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * values in raw yaml to pass to helm. (Default: `empty`).
     */
    values?: pulumi.Input<string>;
    /**
     * Version of Cilium (Default: `v1.14.5`).
     */
    version?: pulumi.Input<string>;
    /**
     * Wait for Cilium status is ok (Default: `true`).
     */
    wait?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Install resource.
 */
export interface InstallArgs {
    /**
     * Datapath mode to use { tunnel | native | aws-eni | gke | azure | aks-byocni } (Default: `autodetected`).
     */
    dataPath?: pulumi.Input<string>;
    /**
     * Namespace in which to install (Default: `kube-system`).
     */
    namespace?: pulumi.Input<string>;
    /**
     * Helm chart repository to download Cilium charts from (Default: `https://helm.cilium.io`).
     */
    repository?: pulumi.Input<string>;
    /**
     * When upgrading, reset the helm values to the ones built into the chart (Default: `false`).
     */
    reset?: pulumi.Input<boolean>;
    /**
     * When upgrading, reuse the helm values from the latest release unless any overrides from are set from other flags. This option takes precedence over HelmResetValues (Default: `true`).
     */
    reuse?: pulumi.Input<boolean>;
    /**
     * Set helm values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2 (Default: `[]`).
     */
    sets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * values in raw yaml to pass to helm. (Default: `empty`).
     */
    values?: pulumi.Input<string>;
    /**
     * Version of Cilium (Default: `v1.14.5`).
     */
    version?: pulumi.Input<string>;
    /**
     * Wait for Cilium status is ok (Default: `true`).
     */
    wait?: pulumi.Input<boolean>;
}
