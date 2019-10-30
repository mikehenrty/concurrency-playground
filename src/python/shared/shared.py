import boto3
import os
import sys

session = boto3.Session()
resource = session.resource('s3')


def get_params():
    params = {
        'bucket_name': sys.argv[1],
        'input_file': sys.argv[2],
        'output_dir': sys.argv[3],
    }

    if len(sys.argv) > 4:
        params['workers'] = sys.argv[4]

    return params


def read_filenames(input_file):
    with open(input_file, 'r') as data:
        lines = [line.strip() for line in data.readlines()]
    return lines


def download_file(bucket, key, output_dir):
    bucket = resource.Bucket(bucket)
    output_filename = os.path.join(output_dir, key)
    bucket.download_file(key, output_filename)
