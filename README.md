## DAPR Example Project for Go
This repo created to understand basics of DAPR with GoLang.

The project has 3 service now. Services simply implies State Management, Pub-Sub and Dynamic Secret Configuration with Kubernetes with Dapr.

<table>
<thead>
<th>Service Name</th>
<th>Endpoints</th>

</tr>
<td>Order Service</td>
<td>
POST /listen which is listening events from pub-service.
</td>
</tr>

</tr>
<td>Publish Service</td>
<td>
POST / Send predefined event to order service from redis.
</br>
GET /secret Dynamically fetch Kubernetes Secret.
</td>
</tr>

</tr>
<td>User Service</td>
<td>
POST /add Create state in redis store
</br>
GET / Read data from Redis Store. You could change with query string to get value of desired key.

</td>
</tr>
</table>