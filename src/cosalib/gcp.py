import os
import re
from cosalib.cmdlib import run_verbose
from tenacity import (
    retry,
    stop_after_attempt
)


# This is the naming rule used by GCP and is used to check image
# names during upload. See:
# https://cloud.google.com/compute/docs/reference/rest/v1/images/insert
GCP_NAMING_RE = r"[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?|[1-9][0-9]{0,19}"


@retry(reraise=True, stop=stop_after_attempt(3))
def remove_gcp_image(gcp_id, json_key, project):
    print(f"GCP: removing image {gcp_id}")
    try:
        run_verbose([
            'ore', 'gcloud', 'delete-images', gcp_id,
            '--json-key', json_key,
            '--project', project
        ])
    except SystemExit:
        raise Exception("Failed to remove image")


@retry(reraise=True, stop=stop_after_attempt(3))
def gcp_run_ore(build, args):
    """
    Execute ore to upload the tarball and register the image
    """
    arg_exp_str = "parameter '--{}' or envVar '{}' must be defined"
    if args.bucket is None:
        raise Exception(arg_exp_str.format("bucket", "GCP_BUCKET"))
    if args.json_key is None:
        raise Exception(arg_exp_str.format("json-key", "GCP_JSON_AUTH"))
    if args.project is None:
        raise Exception(arg_exp_str.format("project", "GCP_PROJECT"))

    gcp_name = re.sub(r'[_\.]', '-', build.image_name_base)
    if not re.fullmatch(GCP_NAMING_RE, gcp_name):
        raise Exception(f"{gcp_name} does match the naming rule: file a bug")
    urltmp = os.path.join(build.tmpdir, "gcp-url")

    ore_common_args = [
        'ore',
        'gcloud',
        '--project', args.project,
        '--json-key', args.json_key,

    ]
    if args.log_level == "DEBUG":
        ore_common_args.extend(['--log-level', "DEBUG"])

    ore_upload_cmd = ore_common_args + [
        'upload',
        '--basename', build.build_name,
        '--force',  # We want to support restarting the pipeline
        '--bucket', f'{args.bucket}',
        '--name', gcp_name,
        '--file', f"{build.image_path}",
        '--write-url', urltmp,
    ]
    if args.description:
        ore_upload_cmd.extend(['--description', args.description])
    if not args.create_image:
        ore_upload_cmd.extend(['--create-image=false'])
    run_verbose(ore_upload_cmd)

    # Run deprecate image to deprecate if requested
    if args.deprecated:
        ore_deprecate_cmd = ore_common_args + [
            'deprecate-image',
            '--image', gcp_name,
            '--state', 'DEPRECATED'
        ]
        run_verbose(ore_deprecate_cmd)

    # Run update-image to add to an image family if requested.
    # We run this as a separate API call because we want to run
    # it AFTER the deprecation if the user passed --deprecated
    if args.family:
        ore_update_cmd = ore_common_args + [
            'update-image',
            '--image', gcp_name,
            '--family', args.family
        ]
        run_verbose(ore_update_cmd)

    build.meta['gcp'] = {
        'image': gcp_name,
        'url': open(urltmp).read().strip()
    }
    build.meta_write()


def gcp_run_ore_replicate(*args, **kwargs):
    print("""
Google Cloud Compute Engine does not require regional
replication. This command is a place-holder only.
""")


# https://stackoverflow.com/questions/44561722/why-in-argparse-a-true-is-always-true
def boolean_string(s):
    if s.lower() not in {'false', 'true'}:
        raise ValueError('Not a valid boolean string')
    return s.lower() == 'true'


def gcp_cli(parser):
    """
    Extend a parser with the GCP options
    """
    parser.add_argument("--bucket",
                        help="Storage account to write image to",
                        default=os.environ.get("GCP_BUCKET"))
    parser.add_argument("--gce",
                        help="Use GCE as the platform ID instead of GCP",
                        action="store_true",
                        default=bool(
                            os.environ.get("GCP_GCE_PLATFORM_ID", False))
                        )
    parser.add_argument("--json-key",
                        help="GCP Service Account JSON Auth",
                        default=os.environ.get("GCP_JSON_AUTH"))
    parser.add_argument("--project",
                        help="GCP Project name",
                        default=os.environ.get("GCP_PROJECT_NAME"))
    parser.add_argument("--family",
                        help="GCP image family to attach image to",
                        default=None)
    parser.add_argument("--description",
                        help="The description that should be attached to the image",
                        default=None)
    parser.add_argument("--create-image",
                        type=boolean_string,
                        help="Whether or not to create an image in GCP after upload.",
                        default=True)
    parser.add_argument("--deprecated",
                        action="store_true",
                        default=False,
                        help="If the image should be marked as deprecated")
    return parser
