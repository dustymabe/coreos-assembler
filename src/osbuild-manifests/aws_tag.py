import os
import subprocess
import json
import argparse



def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--stream', dest='stream', type=str, help='Fedora stream', required=True)
    parser.add_argument('--arch', dest='arch', type=str, help='Architecture', default='x86_64')
    args = parser.parse_args()

    builds = getBuildsForStream(args.stream, args.arch)
    for build in builds:
        print("The build is "+build)
        buildFetch(args.stream, build, args.arch)
        meta = open('builds/'+build+'/'+args.arch+'/meta.json')
        data = json.load(meta)

        # Delete this when actually running. Just here while I make this script
        # data ={"amis":[{
        #     "name": "us-east-1",
        #     "hvm": "ami-0016d5df3041499f9",
        #     "snapshot": "snap-0c1ca4850fcd5e573"
        # }]}
        amis = data['amis']
        for ami in amis:
            UpdateTagCmd = 'aws ec2 create-tags --resource ' + ami['hvm'] + ' --tags '+ 'Key="FedoraUser",Value="coreos"' + ' --region=' + ami['name']
            try:
                subprocess.call(['/bin/bash', '-i', '-c', UpdateTagCmd])
            except subprocess.CalledProcessError as e:
                return(e.output)
        return

def getBuildsForStream(stream, arch):
    buildFetch = 'cosa buildfetch --stream='+ stream + ' --arch='+ arch
    try:
        subprocess.call(['/bin/bash', '-i', '-c', buildFetch])
    except subprocess.CalledProcessError as e:
        return(e.output)

    f = open('builds/builds.json')
    data = json.load(f)
    builds = []

    for i in data['builds']:
        builds.append(i['id'])
    return builds

def buildFetch(stream, build, arch):
    buildFetchCmd = 'cosa buildfetch --stream='+ stream + '--build=' + build + '--arch=' + arch
    try:
        subprocess.call(['/bin/bash', '-i', '-c', buildFetchCmd])
    except subprocess.CalledProcessError as e:
        return(e.output)

if __name__ == '__main__':
    main()
