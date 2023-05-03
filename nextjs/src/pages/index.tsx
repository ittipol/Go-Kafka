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

  return {
    props: {}, // will be passed to the page component as props
  }
}

export default function Home() {

  console.log('Start')

  return (
    <div>index</div>
  )
}
