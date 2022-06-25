// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'

type Data = {
  representation: string,
  kana: string,
}

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data[]>
) {
  const data: Data[] = await fetch('localhost:8000/kana', {
    method: 'GET'
  }).then(res => res.json())
  res.status(200).json(data)
}
