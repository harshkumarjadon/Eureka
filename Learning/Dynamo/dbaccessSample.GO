/*
    OneTable Overview -- in JavaScript

    If using typescript, see the 'typescript' overview sample.
    This sample runs its own local dynamodb instance on port 4567.
 */

import {DynamoDBClient} from '@aws-sdk/client-dynamodb'
import DynamoDbLocal from 'dynamo-db-local'

// import { Table } from 'dynamodb-onetable'

//  Use local source when debugging
import {Table} from '../../../dist/mjs/index.js'

/*
    Import the OneTable schema
*/
import Schema from './schema.js'

//  Local DynamoDB connection port
const PORT = 4567

//  Create a client using the AWS V3 helper
const client = new DynamoDBClient({
    region: 'local',
    endpoint: `http://localhost:${PORT}`,
})

//  Crypto setup for to add additional encryption layer of email addresses
const Crypto = {
    primary: {
        cipher: 'aes-256-gcm',
        password: '1a22a-d27c9-12342-5f7bc-1a716-fc73e',
    },
}

/*
    Single-table schema and setup. This is used for general access and by `createTable`
 */
const table = new Table({
    client: client,
    crypto: Crypto,
    logger: true,
    name: 'TestOverview',
    partial: false,
    schema: Schema,
})

let dynamodb = null

async function start() {
    //  Start the dynamodb instance and then short wait for it to open a listening port.
    dynamodb = DynamoDbLocal.spawn({port: PORT})
    await delay(1000)
}

async function stop() {
    //  Stop the local test dynamodb instance
    process.kill(dynamodb.pid)
}

async function test() {
    /*
        For this sample, create the table. Normally would be created separately.
    */
    await table.createTable()

    /*
        Get OneTable entities. You can use these models for create/read/write etc. Or you can use table.*('Account').
    */
    const Account = table.getModel('Account')
    const Invoice = table.getModel('Invoice')
    const Product = table.getModel('Product')
    const User = table.getModel('User')

    /*
        Create account. This will allocate an account ID (ULID) and create item in primary and
        secondary index.
    */
    let account = await Account.create({name: 'Acme Rockets'})

    /*
        Add account ID to context. This will be blended with all API properties.
    */
    table.setContext({accountId: account.id})

    /*
        Create user. This will allocate a user ID, get the accountId from the context.
     */
    let user = await User.create({name: 'Road Runner', email: 'roadrunner@acme.com', address: {}})

    user = await User.get({email: 'roadrunner@acme.com'})

    /*
        Fetch user via name using the GSI. Follow will automatically follow the GSI to fetch the full item from the primary index.
    */
    user = await User.get({name: 'Road Runner'}, {index: 'gs1', follow: true})

    /*
        Fetch users for account. Uses accountId from context
    */
    let users = await User.find({})

    /*
        Add a filter expression to find only those users that have an admin role.
    */
    let adminUsers = await User.find({role: 'admin'})

    /*
        Simple update
    */
    user = await User.update({email: 'roadrunner@acme.com', balance: 0, role: 'admin'})

    /*
        Update a nested field
    */
    user = await User.update({email: 'roadrunner@acme.com'}, {set: {'address.zip': '{"98034"}'}})

    /*
        Different ways to update properties. Add will atomically add 10 to the balance as will the `set` expression.
    */
    user = await User.update({email: 'roadrunner@acme.com', balance: 110})
    user = await User.update({email: 'roadrunner@acme.com'}, {add: {balance: 10}})
    user = await User.update({email: 'roadrunner@acme.com'}, {set: {balance: '${balance} - {2}'}})

    //  Find users with a balance over $100
    users = await User.find(
        {accountId: account.id},
        {
            where: '${balance} > {100}',
        }
    )

    /*
        Get a collection of items in the account. (See below table.fetch is simpler)
     */
    let collection = await table.fetch(['Account', 'User', 'Invoice'], {pk: `account#${account.id}`})

    /*
        Create many users via batch
     */
    let batch = {}
    let i = 0,
        count = 0
    while (i++ < 200) {
        User.create({name: `user${i}`, email: `user${i}@acme.com`}, {batch})
        if (++count >= 25) {
            await table.batchWrite(batch)
            batch = {}
            count = 0
        }
    }

    /*
        Get a list of user email addresses. Need _type to know how to parse results.
    */
    let items = (await User.find({}, {fields: ['email', '_type']})).map((i) => i.email)

    /*
        Read a page of users in groups of 25 at a time
     */
    let start
    do {
        users = await User.find({}, {start, limit: 25})
        start = users.start
    } while (users.start)

    /*
        Create a product, not tied to an account.
    */
    let product = await Product.create({name: 'rocket', price: 10.99})

    /*
        Transaction to atomically create an invoice and update the user and account balance
    */
    let transaction = {}
    Invoice.create({product: product.name, count: 1, total: product.price}, {transaction})
    User.update({email: 'roadrunner@acme.com'}, {add: {balance: product.price}, transaction})
    Account.update({id: account.id}, {add: {balance: product.price}, transaction})
    let result = await table.transact('write', transaction)

    /*
        Fetch item collection of entities in the account
    */
    collection = await table.fetch(['Account', 'User', 'Invoice'], {pk: `account#${account.id}`})

    /*
        Get invoices for the account this month. The sk is of the form invoice#iso-date#id
        So we take advantage of the fact that ISO dates sort.
    */
    let from = new Date()
    from.setMonth(from.getMonth() - 1)
    let invoices = await Invoice.find(
        {
            gs1sk: {
                between: [`invoice#${from.toISOString()}`, `invoice#${new Date().toISOString()}`],
            },
        },
        {index: 'gs1', follow: true}
    )

    /*
        For maintenance, useful to be able to query by entity type. This is not a costly scan.
    */
    let accounts = await Account.find({}, {index: 'gs1'})
    users = await User.find({}, {index: 'gs1'})
    invoices = await Invoice.find({}, {index: 'gs1'})

    /*
        Cleanup
    */
    await table.deleteTable('DeleteTableForever')
}

//  Short nap
async function delay(time) {
    return new Promise(function (resolve, reject) {
        setTimeout(() => resolve(true), time)
    })
}

async function main() {
    await start()
    try {
        await test()
    } catch (err) {
        console.error(err)
    }
    await stop()
}

//  Ah, if only for a top-level await
main()





/*
    Overview schema
*/

const Match = {}
//     ulid: /^[0123456789ABCDEFGHJKMNPQRSTVWXYZ]{26}$/,
//     email: /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
//     name: /^[a-z0-9 ,.'-]+$/i,
//     address: /[a-z0-9 ,.-]+$/,
//     zip: /^\d{5}(?:[-\s]\d{4})?$/,
//     phone: /^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$/,
// }

export default {
    format: 'onetable:1.1.0',
    version: '0.0.1',
    indexes: {
        primary: {hash: 'pk', sort: 'sk'},
        gs1: {hash: 'gs1pk', sort: 'gs1sk', project: ['gs1pk', 'gs1sk']},
    },
    models: {
        Account: {
            pk: {type: String, value: 'account#${id}'},
            sk: {type: String, value: 'account#'},
            id: {type: String, generate: 'ulid', validate: Match.ulid},
            name: {type: String, required: true, unique: true, validate: Match.name},
            balance: {type: Number, default: 0},

            //  Search by account name or by type
            gs1pk: {type: String, value: 'account#'},
            gs1sk: {type: String, value: 'account#${name}${id}'},
        },

        User: {
            pk: {type: String, value: 'account#${accountId}'},
            sk: {type: String, value: 'user#${email}'},
            accountId: {type: String, required: true},
            id: {type: String, generate: 'ulid', validate: Match.ulid},
            name: {type: String, required: true, validate: Match.name},
            email: {type: String, required: true, validate: Match.email, crypt: true},
            role: {type: String, default: 'user', enum: ['admin', 'user']},

            address: {
                type: Object,
                default: {},
                schema: {
                    street: {type: String},
                    city: {type: String},
                    zip: {type: String},
                },
            },

            status: {type: String, required: true, default: 'active', enum: ['active', 'inactive']},
            balance: {type: Number, default: 0},

            //  Search by user name or by type
            gs1pk: {type: String, value: 'user#'},
            gs1sk: {type: String, value: 'user#${name}#${id}'},
        },

        Product: {
            pk: {type: String, value: 'product#${id}'},
            sk: {type: String, value: 'product#'},
            id: {type: String, generate: 'ulid', validate: Match.ulid},
            name: {type: String, required: true},
            price: {type: Number, required: true},

            //  Search by product name or by type
            gs1pk: {type: String, value: 'product#'},
            gs1sk: {type: String, value: 'product#${name}#${id}'},
        },

        Invoice: {
            pk: {type: String, value: 'account#${accountId}'},
            sk: {type: String, value: 'invoice#${id}'},

            accountId: {type: String, required: true},
            date: {type: Date},
            id: {type: String, generate: 'ulid'},
            product: {type: String},
            count: {type: Number},
            total: {type: Number},

            //  Search by invoice date or by type
            gs1pk: {type: String, value: 'invoice#'},
            gs1sk: {type: String, value: 'invoice#${date}#${id}'},
        },
    },
    params: {
        isoDates: true,
        timestamps: true,
    },
