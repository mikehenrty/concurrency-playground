import os
import sys
import boto3

bucket_name = sys.argv[1]
input_file = sys.argv[2]
output_dir = sys.argv[3]

session = boto3.Session()
resource = session.resource('s3')

with open(input_file, 'r') as data:
    lines = [line.strip() for line in data.readlines()]

while(len(lines) > 0):
    filename = lines.pop()
    bucket = resource.Bucket(bucket_name)
    output_filename = os.path.join('./', output_dir, filename)
    bucket.download_file(filename, output_filename)
