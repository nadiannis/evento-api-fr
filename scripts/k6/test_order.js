import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 200,
  duration: '5s',
};

const customerID = 1;
const ticketID = 2;
const initialTicketQuantity = 100;

export function setup() {
  const baseUrl = 'http://localhost:8080/api';
  const headers = { 'Content-Type': 'application/json' };

  const getTicketQuantityUrl = `${baseUrl}/tickets/${ticketID}`;
  const ticketResponse = http.get(getTicketQuantityUrl);
  const ticketData = JSON.parse(ticketResponse.body);

  if (ticketData.data.quantity === 0) {
    const addTicketQuantityUrl = `${baseUrl}/tickets/${ticketID}/quantities`;
    const ticketQuantityPayload = JSON.stringify({
      action: 'add',
      quantity: 100,
    });
    const addTicketResponse = http.patch(
      addTicketQuantityUrl,
      ticketQuantityPayload,
      { headers }
    );

    check(addTicketResponse, {
      'added ticket quantity': (r) => r.status === 200,
    });
  }

  const getCustomerUrl = `${baseUrl}/customers/${customerID}`;
  const customerResponse = http.get(getCustomerUrl);
  let customerData = JSON.parse(customerResponse.body);

  if (customerData.data.balance === 0) {
    const addBalanceUrl = `${baseUrl}/customers/${customerID}/balances`;
    const balancePayload = JSON.stringify({ action: 'add', balance: 100000 });
    const balanceResponse = http.patch(addBalanceUrl, balancePayload, {
      headers,
    });
    customerData = JSON.parse(balanceResponse.body);

    check(balanceResponse, {
      'added customer balance': (r) => r.status === 200,
    });
  }

  const deleteOrdersUrl = `${baseUrl}/orders`;
  const deleteOrdersResponse = http.del(deleteOrdersUrl);

  check(deleteOrdersResponse, {
    'deleted all orders': (r) => r.status === 200,
  });

  return {
    ticketId: ticketID,
    customerId: customerID,
    initialTicketQuantity: initialTicketQuantity,
    initialCustomerBalance: customerData.data.balance,
  };
}

export default function () {
  const url = 'http://localhost:8080/api/orders';
  const payload = JSON.stringify({
    customer_id: customerID,
    ticket_id: ticketID,
    quantity: 1,
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.post(url, payload, params);

  check(res, {
    'status is 201 or 400': (r) => r.status === 201 || r.status === 400,
  });

  sleep(1);
}

export function teardown(data) {
  console.log(data);

  // Check final state
  const ticketUrl = `http://localhost:8080/api/tickets/${ticketID}`;
  const ticketRes = http.get(ticketUrl);

  check(ticketRes, {
    'final ticket quantity is 0': (r) => JSON.parse(r.body).data.quantity === 0,
  });

  const customerUrl = `http://localhost:8080/api/customers/${customerID}`;
  const customerRes = http.get(customerUrl);
  const totalPrice = data.initialTicketQuantity * 250;
  const finalBalance = data.initialCustomerBalance - totalPrice;

  check(customerRes, {
    [`final customer balance is ${finalBalance}`]: (r) =>
      JSON.parse(r.body).data.balance === finalBalance,
  });

  const orderUrl = 'http://localhost:8080/api/orders';
  const orderRes = http.get(orderUrl);

  check(orderRes, {
    'order count matches initial ticket quantity': (r) =>
      JSON.parse(r.body).data.length === data.initialTicketQuantity,
  });
}
