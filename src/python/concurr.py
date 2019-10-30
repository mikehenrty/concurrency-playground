import queue
import threading
from shared import shared

NUM_WORKERS = 300

params = shared.get_params()
bucket_name = params['bucket_name']
input_file = params['input_file']
output_dir = params['output_dir']

lines = shared.read_filenames(input_file)

q = queue.Queue()


def worker():
    while True:
        filename = q.get()
        shared.download_file(bucket_name, filename, output_dir)
        q.task_done()


def main():
    for i in range(len(lines)):
        q.put_nowait(lines[i])

    workers = []
    for i in range(NUM_WORKERS):
        t = threading.Thread(target=worker)
        t.daemon = True
        t.start()
        workers.append(t)

    q.join()


if __name__ == '__main__':
    main()
