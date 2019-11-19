import asyncio
import io
import os
import aioboto3
from shared import shared

NUM_WORKERS = 10

params = shared.get_params()
bucket_name = params['bucket_name']
input_file = params['input_file']
output_dir = params['output_dir']
num_workers = int(params['workers']) if params['workers'] else NUM_WORKERS

chunk_size = 69 * 1024

files = shared.read_filenames(input_file)
loop = asyncio.get_event_loop()

async def save(full_path: str, s3_ob):
    async with io.FileIO(full_path, 'w') as file:
        async with s3_ob["Body"] as stream:
            file_data = await stream.read(chunk_size)
            while file_data:
                await file.write(file_data)
                file_data = await stream.read(chunk_size)


async def worker():
    async with aioboto3.client("s3") as s3:
        while len(files) > 0:
            filename = files.pop()
            s3_ob = await s3.get_object(Bucket=bucket_name, Key=filename)
            await save(os.path.join(output_dir, filename, s3_ob)

@asyncio.coroutine
async def main():
    workers = [worker() for i in range(0, num_workers)]
    await asyncio.wait(workers, return_when=asyncio.ALL_COMPLETED)


if __name__ == '__main__':
    loop.run_until_complete(main())
