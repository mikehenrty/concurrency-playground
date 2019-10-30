from gevent import monkey

import gevent

# patches stdlib (including socket and ssl modules) to cooperate with other
monkey.patch_all()

from shared import shared


NUM_WORKERS = 10

params = shared.get_params()
bucket_name = params['bucket_name']
input_file = params['input_file']
output_dir = params['output_dir']
num_workers = int(params['workers']) if params['workers'] else NUM_WORKERS

files = shared.read_filenames(input_file)


def worker():
    while len(files) > 0:
        filename = files.pop()
        shared.download_file(bucket_name, filename, output_dir)


def main():
    workers = []
    for i in range(0, num_workers):
        w = gevent.spawn(worker)
        workers.append(w)

    gevent.wait(workers)


if __name__ == '__main__':
    main()
