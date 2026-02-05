package models

type ProductStats struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}

type DailyReport struct {
	TotalRevenue   int          `json:"total_revenue"`
	TotalTransaksi int          `json:"total_transaksi"`
	ProdukTerlaris ProductStats `json:"produk_terlaris"`
}
