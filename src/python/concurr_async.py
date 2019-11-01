import os
import asyncio
import boto3
import concurrent.futures
from shared import shared


NUM_WORKERS = 10

params = shared.get_params()
bucket_name = params['bucket_name']
input_file = params['input_file']
output_dir = params['output_dir']
num_workers = int(params['workers']) if params['workers'] else NUM_WORKERS
num_threads = max(num_workers, min(32, os.cpu_count() + 4))  # min(32, os.cpu_count() + 4) is the default


def downloader(bucket, key):
    """ Download a file from S3 and save it into output_dir """
    output_filename = os.path.join(output_dir, key)
    bucket.download_file(key, output_filename)

    return


async def worker(queue, executor):
    """ Worker """
    # Create boto3 objects in each worker
    session = boto3.Session()
    resource = session.resource('s3')
    bucket = resource.Bucket(bucket_name)

    while not queue.empty():
        item = await queue.get()

        # Download file in the executor (ThreadPoolExecutor)
        await loop.run_in_executor(executor, downloader, bucket, item)

        # Decrement the item counter of the queue
        queue.task_done()


async def main(loop, executor):
    """ Main """
    # Read target files and put them into queue
    queue = asyncio.Queue(loop=loop)
    with open(input_file, 'r') as data:
        for line in data.readlines():
            await queue.put(line.strip())

    # Start workers
    for i in range(0, num_workers):
        asyncio.ensure_future(worker(queue, executor))

    # Wait until all items in the queue are processed
    await queue.join()

    return


if __name__ == '__main__':
    loop = asyncio.get_event_loop()

    with concurrent.futures.ThreadPoolExecutor(max_workers=num_threads) as executor:
        loop.run_until_complete(main(loop, executor))

    loop.close()
