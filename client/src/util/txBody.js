export default {
  createUser: (id, name) => {
    return {
      type: 'create_user',
      entity: {
        _id: id,
        name: name
      }
    }
  },
  sendFunds: (to, amount) => {
    return {
      type: 'send_funds',
      entity: {
        to: to,
        amount: amount
      }
    }
  }
}
