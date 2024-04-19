import os
import subprocess
import json
import argparse



def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--stream', dest='stream', type=str, help='Fedora stream', required=True)
    args = parser.parse_args()

    builds = getBuildsForStream(args.stream)
    for build in builds:
        build_id=build['id']
        arches=build['arches']
        for arch in arches:
            print(f"The build is {build_id}")
            buildFetch(args.stream, build_id, arch)
            meta = open('builds/'+build_id+'/'+arch+'/meta.json')
            data = json.load(meta)

            # Delete this when actually running. Just here while I make this script
            data ={"amis":[{
                "name": "us-east-1",
                "hvm": "ami-0016d5df3041499f9",
                "snapshot": "snap-0c1ca4850fcd5e573"
            }]}
            amis = data['amis']
            for ami in amis:
                checkAndAddTag(ami["hvm"], ami["name"])
                checkAndAddTag(ami["snapshot"], ami["name"])
            return

def checkAndAddTag(resourceId, region):
    tagExists = checkTag(resourceId)
    if tagExists:
        print(f"{resourceId} already tagged with FedoraUser=coreos tag")
    else:
        addTag(resourceId, region)
        print(f"FedoraUser=coreos tag successfully added to {resourceId}")    

def checkTag(resourceId):
    checkTagCmd = f'aws ec2 describe-tags --filters Name=resource-id,Values={resourceId} Name=value,Values=coreos'
    try:
        tagCheck=subprocess.run([checkTagCmd], shell=True, capture_output=True, text=True)
        if "FedoraUser" and "coreos" in tagCheck.stdout:
            return True
        return False
    except subprocess.CalledProcessError as e:
        return(e.output)

def addTag(resourceId, region):
    UpdateTagCmd = f'aws ec2 create-tags --resource {resourceId} --tags Key="FedoraUser",Value="coreos" --region {region}'
    try:
        subprocess.run([UpdateTagCmd], shell=True)
    except subprocess.CalledProcessError as e:
        return(e.output)
    
def getBuildsForStream(stream):
    buildFetch = 'cosa buildfetch --stream='+ stream + ' --arch=all'
    try:
        subprocess.call(['/bin/bash', '-i', '-c', buildFetch])
    except subprocess.CalledProcessError as e:
        return(e.output)

    f = open('builds/builds.json')
    data = json.load(f)
    return data['builds']

def buildFetch(stream, build, arch):
    buildFetchCmd = 'cosa buildfetch --stream='+ stream + ' --build=' + build + ' --arch=' + arch
    try:
        subprocess.call(['/bin/bash', '-i', '-c', buildFetchCmd])
    except subprocess.CalledProcessError as e:
        return(e.output)

if __name__ == '__main__':
    main()
