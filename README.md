# Manjo QR API

API untuk generate QR dan update status pembayaran.

Base URL:

```
http://localhost:8080/api/v1/qr
```

Authentication menggunakan header:

```
api-key: <your-api-key>
```

---

# 📌 Endpoint

## 1️⃣ Generate QR

### POST `/generate`

### Headers

| Key | Value |
|------|--------|
| api-key | Your API Key |
| Content-Type | application/json |

### Request Body

```json
{
  "partnerReferenceNo": "DIRECT-API-NMS-whhq7gvx58",
  "amount": {
    "value": "10000.00",
    "currency": "IDR"
  },
  "merchantId": "EP27842182"
}
```

### Success Response

```json
{
  "qrContent": "string",
  "referenceNo": "A6096942400",
  "responseCode": "2004700",
  "responseMessage": "Successful"
}
```

### Error Responses

#### Amount ≤ 0

```json
{
  "message": "amount must be greater than zero",
  "status": "error"
}
```

#### Missing Header

```json
{
  "message": "missing required headers",
  "status": "error"
}
```

#### Invalid API Key

```json
{
  "message": "invalid api key",
  "status": "error"
}
```

---

## 2️⃣ Payment Callback

### POST `/payment`

### Headers

| Key | Value |
|------|--------|
| api-key | Your API Key |
| Content-Type | application/json |

### Request Body

```json
{
  "originalReferenceNo": "A3889922900",
  "originalPartnerReferenceNo": "DIRECT-API-NMS-whhq7gvx58",
  "transactionStatusDesc": "Success",
  "paidTime": "2025-09-21T09:25:00+07:00",
  "amount": {
    "value": "10000.00",
    "currency": "IDR"
  }
}
```

### Success Response

```json
{
  "responseCode": "2005100",
  "responseMessage": "Successful",
  "transactionStatusDesc": "Success"
}
```

### Error Responses

#### Transaction Not Found

```json
{
  "message": "transaction not found",
  "status": "error"
}
```

#### Missing Header

```json
{
  "message": "missing required headers",
  "status": "error"
}
```

#### Invalid API Key

```json
{
  "message": "invalid api key",
  "status": "error"
}
```

---

# 🔐 Authentication Rules

- Semua endpoint wajib menyertakan header `api-key`
- Jika tidak ada header → `missing required headers`
- Jika api key salah → `invalid api key`

---

# ⚠️ Validation Rules

- `amount.value` harus lebih dari 0
- `referenceNo` harus valid dan terdaftar

---