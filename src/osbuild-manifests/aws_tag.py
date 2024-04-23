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
            # Delete this when actually running. Just here while I make this script
            # data ={"amis":[{
            #     "name": "us-east-1",
            #     "hvm": "ami-0016d5df3041499f9",
            #     "snapshot": "snap-0c1ca4850fcd5e573"
            # }]}
            # amis = data['amis']

            for ami in amis:
                region = ami["name"]
                checkAndAddTag(ami["hvm"], region, args.dry_run)
                checkAndAddTag(ami["snapshot"], region, args.dry_run)
    return
 
def checkAndAddTag(resourceId, region, dry_run):
    describeTagsCmd = f'aws ec2 describe-tags --filters Name=resource-id,Values={resourceId} --region {region} --output=json'
    tagCheck=subprocess.run([describeTagsCmd], shell=True, capture_output=True, text=True)
    if tagCheck.stdout == None or tagCheck.stdout == '':
        print(f"No tags detected for {resourceId}; assuming it doesn't exist")
        return
    tagCheck=json.loads(tagCheck.stdout)

    if any((tag['Key'] == 'FedoraGroup' and tag['Value'] == 'coreos') for tag in tagCheck['Tags']):
        print(f"{resourceId} already tagged with FedoraGroup=coreos tag")
        return
    else:
        if dry_run:
            print(f"Would add tag 'FedoraGroup=coreos' to {resourceId} in region {region}")
            return
        else:                                                                                                                      
            addTag(resourceId, region, dry_run)                                                                                    
                                                                                                                               
def addTag(resourceId, region, dry_run):                                                                                       
    if dry_run:                                                                                                                
        print(f"Would add tag 'FedoraGroup=coreos' to {resourceId} in region {region}")                                        
    else:                                                                                                                      
        UpdateTagCmd = f'aws ec2 create-tags --resource {resourceId} --tags Key="FedoraGroup",Value="coreos" --region {region}'
        subprocess.run([UpdateTagCmd], shell=True)                                                                             
        print(f"'FedoraGroup=coreos' tag successfully added to {resourceId}")

def getBuildsForStream(stream):
    buildFetchCmd = 'cosa buildfetch --stream='+ stream + ' --arch=all'
    subprocess.check_output(['/bin/bash', '-i', '-c', buildFetchCmd])
 
    f = open(f'builds/builds.json')
    data = json.load(f)
    return data['builds']
 
def buildFetch(stream, build, arch):
    buildFetchCmd = 'cosa buildfetch --stream='+ stream + ' --build=' + build + ' --arch=' + arch
    subprocess.check_output(['/bin/bash', '-i', '-c', buildFetchCmd])
 
if __name__ == '__main__':
    main()
