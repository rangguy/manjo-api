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

# Endpoint

## 1. Get All Transactions

### GET `/transactions`

### Headers

| Key | Value |
|------|--------|
| api-key | Your API Key |

### Success Response

```json
{
    "data": [
        {
            "merchantId": "EP27842182",
            "originalReferenceNo": "A6096942400",
            "originalPartnerReferenceNo": "DIRECT-API-NMS-whhq7gvx58",
            "transactionStatusDesc": "Success",
            "transactionDate": "2026-03-02T17:44:36.096942+07:00",
            "paidTime": "2026-03-03T06:25:00+07:00",
            "amount": {
                "value": "10000",
                "currency": "IDR"
            }
        },
        {
            "merchantId": "",
            "originalReferenceNo": "A7431853900",
            "originalPartnerReferenceNo": "DIRECT-API-NMS-whhq7gvx58",
            "transactionStatusDesc": "",
            "transactionDate": "2026-03-02T15:15:17.431853+07:00",
            "paidTime": null,
            "amount": {
                "value": "10000",
                "currency": "IDR"
            }
        }
    ],
    "message": "success"
}
```

> **Catatan:** `merchantId` dan `transactionStatusDesc` dapat bernilai string kosong `""` apabila transaksi belum diproses. `paidTime` dapat bernilai `null` apabila transaksi belum dibayar.

### Response Fields

| Field | Type | Keterangan |
|-------|------|------------|
| `merchantId` | `string` | Kode merchant. Bisa kosong jika belum diisi |
| `originalReferenceNo` | `string` | Nomor referensi transaksi dari sistem |
| `originalPartnerReferenceNo` | `string` | Nomor referensi dari merchant/partner |
| `transactionStatusDesc` | `string` | Status transaksi (`Success`, `Pending`, `Failed`, atau kosong) |
| `transactionDate` | `string (ISO 8601)` | Waktu transaksi dibuat |
| `paidTime` | `string (ISO 8601) or null` | Waktu pembayaran. `null` jika belum dibayar |
| `amount.value` | `string` | Nominal transaksi |
| `amount.currency` | `string` | Mata uang (contoh: `IDR`) |

### Error Responses

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

## 2. Generate QR

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

#### Amount <= 0

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

## 3. Payment Callback

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

# Authentication Rules

- Semua endpoint wajib menyertakan header `api-key`
- Jika tidak ada header -> `missing required headers`
- Jika api key salah -> `invalid api key`

---

# Validation Rules

- `amount.value` harus lebih dari 0
- `referenceNo` harus valid dan terdaftar