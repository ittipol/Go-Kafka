import api from "@/apps/axios"
import { BankAccount } from "@/models/model"
import { useEffect, useState } from "react"

type BankAccountType = BankAccount[] | null

export async function getStaticProps() {

  // const { Kafka } = require('kafkajs')

  // const kafka = new Kafka({
  //   clientId: 'my-app',
  //   brokers: ['localhost:9092']
  // })

  // const consumer = kafka.consumer({ groupId: 'log' })

  // await consumer.connect()
  // await consumer.subscribe({ topic: 'mtopic', fromBeginning: true })

  // const run = async () => {
  
  //   // Consuming
  //   await consumer.connect()
  //   await consumer.subscribe({ topic: 'test-topic', fromBeginning: true })
  
  //   await consumer.run({
  //     eachMessage: async ({ topic, partition, message }:{ topic:any, partition:any, message:any }) => {
  //       console.table({
  //         partition,
  //         offset: message.offset,
  //         value: message.value.toString(),
  //       })
  //     },
  //   })
  // }

  // run().catch(console.error)

  // const res = await fetch('http://localhost:5000/getAccounts')
  // const bankAccounts:BankAccount[] = await res.json()

  return {
    props: {      
    },
  }
}

export default function Home() {

  const [bankAccounts, setBankAccounts] = useState<BankAccountType>(null)

  useEffect(() => {

    const fetchData = async () => {
      
      try{
        const res = await api.get<BankAccount[]>('getAccounts')
        const data:BankAccountType = res.data

        setBankAccounts(data)

      } catch(ex) {

      }
    }

    fetchData()
    

  }, [])

  return (
    <div className="w-[700px] mx-auto flex flex-col">
      <h1 className="text-2xl justify-center flex my-5">Bank Accounts</h1>
      <div>
        {
          bankAccounts 
          ?
          <div>
            <table className="table-fixed w-full">
              <thead>
                <tr className="border-b-2 border-gray-600">
                  <th className="px-6 py-3">ID</th>
                  <th className="px-6 py-3">Account Holder</th>
                  <th className="px-6 py-3">Account Type</th>
                  <th className="px-6 py-3">Balance</th>
                </tr>
              </thead>
              <tbody>
                {bankAccounts.map((value, index) => (
                  <tr key={index} className="text-center">
                    <td className="px-6 py-3">{value.ID}</td>
                    <td className="px-6 py-3">{value.AccountHolder}</td>
                    <td className="px-6 py-3">{value.AccountType}</td>
                    <td className="px-6 py-3">{value.Balance}</td>
                  </tr>
                ))}
              </tbody>
            </table>
            <hr className="mt-10 mb-9" />
            <div className="flex flex-col">
              <h2 className="flex justify-center mb-5 text-xl">Test managing account by Curl commands</h2>
              <div className="mb-4">
                <h3 className="pl-4 mb-1 text-md">Add bank accounts</h3>
                <div className="text-white bg-gray-600 rounded-lg">
                  <div className="p-4">
                  {
                    `curl --location --request POST 'http://127.0.0.1:4000/openAccount' \
                    --header 'Content-Type: application/json' \
                    --data-raw '{"AccountHolder": "Test user1","AccountType": 1,"OpeningBalance": 1000000.00}'`
                  }
                  </div>
                </div>
              </div>

              <div className="mb-4">
                <h3 className="pl-4 mb-1 text-md">Deposit</h3>
                <div className="text-white bg-gray-600 rounded-lg">
                  <div className="p-4">
                  {
                    `curl --location --request POST 'http://127.0.0.1:4000/depositFund' \
                    --header 'Content-Type: application/json' \
                    --data-raw '{"ID":"Input_ID","Amount":20000}'`
                  }
                  </div>
                </div>
              </div>

              <div className="mb-4">
                <h3 className="pl-4 mb-1 text-md">Withdraw</h3>
                <div className="text-white bg-gray-600 rounded-lg">
                  <div className="p-4">
                  {
                    `curl --location --request POST 'http://127.0.0.1:4000/withdrawFund' \
                    --header 'Content-Type: application/json' \
                    --data-raw '{"ID":"Input_ID","Amount":20000}'`
                  }
                  </div>
                </div>
              </div>

              <div className="mb-4">
                <h3 className="pl-4 mb-1 text-md">Close account</h3>
                <div className="text-white bg-gray-600 rounded-lg">
                  <div className="p-4">
                  {
                    `curl --location --request POST 'http://127.0.0.1:4000/closeAccount' \
                    --header 'Content-Type: application/json' \
                    --data-raw '{"ID":"Input_ID"}'`
                  }
                  </div>
                </div>
              </div>
            </div>
          </div>
        :
        <div className="flex flex-col">
          <div className="text-center">No Bank Account</div>            

          <div className="flex flex-col mt-6 ">
            <h2 className="flex justify-center mb-2">Use Curl command to add new user</h2>
            <div className="text-white bg-gray-600 rounded-lg">
                <div className="p-4">
                {
                  `curl --location --request POST 'http://127.0.0.1:4000/openAccount' \
                  --header 'Content-Type: application/json' \
                  --data-raw '{"AccountHolder": "Test user1","AccountType": 1,"OpeningBalance": 1000000.00}'`
                }
                </div>
            </div>
          </div>
        </div>
        }
      </div>
    </div>
  )
}
