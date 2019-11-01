from shared import shared

params = shared.get_params()
bucket_name = params['bucket_name']
input_file = params['input_file']
output_dir = params['output_dir']

lines = shared.read_filenames(input_file)


def main():
    while len(lines) > 0:
        filename = lines.pop()
        shared.download_file(bucket_name, filename, output_dir)


if __name__ == '__main__':
    main()
