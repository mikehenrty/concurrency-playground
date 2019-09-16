from gevent import monkey

import gevent

# patches stdlib (including socket and ssl modules) to cooperate with other
monkey.patch_all()

from shared import get_params, read_filenames, download_file


NUM_WORKERS = 40

params = get_params()
bucket_name = params['bucket_name']
input_file = params['input_file']
output_dir = params['output_dir']

files = read_filenames(input_file)


def worker():
    while len(files) > 0:
        filename = files.pop()
        download_file(bucket_name, filename, output_dir)


def main():
    workers = []
    for i in range(0, NUM_WORKERS):
        w = gevent.spawn(worker)
        workers.append(w)

    gevent.wait(workers)
    # gevent.joinall(workers)


if __name__ == '__main__':
    main()
