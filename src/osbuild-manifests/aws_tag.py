#!/usr/bin/python3

# Script to go through all builds and add the FedoraGroup=coreos
# tag to all AMIs and snapshots that we know about.

import subprocess
import json
import argparse

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--stream', dest='stream', type=str, help='Fedora stream', required=True)
    parser.add_argument('--dry-run', dest='dry_run', help='Check if the resources have tags but not add them', action='store_true')
    args = parser.parse_args()

    builds = getBuildsForStream(args.stream)
    for build in builds:
        build_id=build['id']
        arches=build['arches']
        for arch in arches:
            print(f"Parsing AMIs for {build_id} for {arch}")
            buildFetch(args.stream, build_id, arch)
            meta = open(f'builds/{build_id}/{arch}/meta.json')
            data = json.load(meta)

            if 'amis' in data.keys():
                amis = data['amis']
            else:
                print(f"{build_id} does not have any AMIs for {arch} in meta.json")
                continue
            for ami in amis:
                region = ami["name"]
                checkAndAddTag(ami["hvm"], region, args.dry_run)
                checkAndAddTag(ami["snapshot"], region, args.dry_run)

def checkAndAddTag(resourceId, region, dry_run):
    describeTagsCmd = f'aws ec2 describe-tags --filters Name=resource-id,Values={resourceId} --region {region} --output=json'
    tagCheck=subprocess.run(describeTagsCmd.split(' '), capture_output=True, text=True)
    if tagCheck.stdout == None or tagCheck.stdout == '':
        print(f"\tNo tags detected for {resourceId}; assuming it doesn't exist")
        return
    tagCheck=json.loads(tagCheck.stdout)

    if any((tag['Key'] == 'FedoraGroup' and tag['Value'] == 'coreos') for tag in tagCheck['Tags']):
        print(f"\t{resourceId} in {region} already tagged with FedoraGroup=coreos tag")
    else:
        addTag(resourceId, region, dry_run)

def addTag(resourceId, region, dry_run):
    if dry_run:
        print(f"\tWould add tag 'FedoraGroup=coreos' to {resourceId} in region {region}")
    else:
        UpdateTagCmd = f'aws ec2 create-tags --resource {resourceId} --tags Key="FedoraGroup",Value="coreos" --region {region}'
        subprocess.run(UpdateTagCmd.split(' '))
        print(f"\t'FedoraGroup=coreos' tag successfully added to {resourceId} in {region}")

def getBuildsForStream(stream):
    buildFetchCmd = 'cosa buildfetch --stream='+ stream + ' --arch=all'
    subprocess.check_output(buildFetchCmd.split(' '))

    f = open(f'builds/builds.json')
    data = json.load(f)
    return data['builds']

def buildFetch(stream, build, arch):
    buildFetchCmd = 'cosa buildfetch --stream='+ stream + ' --build=' + build + ' --arch=' + arch
    subprocess.check_output(buildFetchCmd.split(' '))

if __name__ == '__main__':
    main()
