package usecase

import (
	"context"
	"database/sql"
	"errors"
	db "esim-service/db/sqlc"
	"esim-service/domain"
	"esim-service/domain/order"
	"esim-service/util"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sesv2"
)

const (
	// Replace sender@example.com with your "From" address.
	// This address must be verified with Amazon SES.
	Sender = "no-reply@beliesim.com"

	// Specify a configuration set. To use a configuration
	// set, comment the next line and line 92.
	//ConfigurationSet = "ConfigSet"

	// The subject line for the email.
	Subject = "Esim Purchase"

	// // The HTML body for the email.
	// HtmlBody = "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
	// 	"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
	// 	"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

	// //The email body for recipients with non-HTML email clients.
	// TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."

	// The character encoding for the email.
	CharSet = "UTF-8"
)

func generateHTMLBody(
	order db.GetOrderByTopupIdRow,
	esim db.Esim,
	iccid string,
	topupId string,
	smdp string,
	activationCode string,
	qrCodeHTML string,
) string {
	productName := fmt.Sprintf(`Paket eSIM %s %d%s/%dDay(s) [%s]`, esim.CountryCode, esim.DataAmount, esim.DataUnit, esim.DurationInDays, esim.PlanOption)

	return fmt.Sprintf(`
	<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//ID" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
	<html lang="en" xmlns="http://www.w3.org/1999/xhtml">
		<head>
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<meta name="x-apple-disable-message-reformatting" />
			<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
			<meta name="color-scheme" content="light dark" />
			<meta name="supported-color-schemes" content="light dark" />
			<title></title>
			<style type="text/css" rel="stylesheet" media="all">
			/* Base ------------------------------ */
			
			@import url("https://fonts.googleapis.com/css?family=Nunito+Sans:400,700&display=swap");
			body {
				width: 100% !important;
				height: 100%;
				margin: 0;
				-webkit-text-size-adjust: none;
			}
			
			a {
				color: #3869D4;
			}
			
			a img {
				border: none;
			}
			
			td {
				word-break: break-word;
			}
			
			.preheader {
				display: none !important;
				visibility: hidden;
				mso-hide: all;
				font-size: 1px;
				line-height: 1px;
				max-height: 0;
				max-width: 0;
				opacity: 0;
				overflow: hidden;
			}
			/* Type ------------------------------ */
			
			body,
			td,
			th {
				font-family: "Nunito Sans", Helvetica, Arial, sans-serif;
			}
			
			h1 {
				margin-top: 0;
				color: #333333;
				font-size: 22px;
				font-weight: bold;
				text-align: left;
			}
			
			h2 {
				margin-top: 0;
				color: #333333;
				font-size: 16px;
				font-weight: bold;
				text-align: left;
			}
			
			h3 {
				margin-top: 0;
				color: #333333;
				font-size: 14px;
				font-weight: bold;
				text-align: left;
			}
			
			td,
			th {
				font-size: 16px;
			}
			
			p,
			ul,
			ol,
			blockquote {
				margin: .4em 0 1.1875em;
				font-size: 16px;
				line-height: 1.625;
			}
			
			p.sub {
				font-size: 13px;
			}
			/* Utilities ------------------------------ */
			
			.align-right {
				text-align: right;
			}
			
			.align-left {
				text-align: left;
			}
			
			.align-center {
				text-align: center;
			}
			
			.u-margin-bottom-none {
				margin-bottom: 0;
			}
			/* Buttons ------------------------------ */
			
			.button {
				background-color: #3869D4;
				border-top: 10px solid #3869D4;
				border-right: 18px solid #3869D4;
				border-bottom: 10px solid #3869D4;
				border-left: 18px solid #3869D4;
				display: inline-block;
				color: #FFF;
				text-decoration: none;
				border-radius: 3px;
				box-shadow: 0 2px 3px rgba(0, 0, 0, 0.16);
				-webkit-text-size-adjust: none;
				box-sizing: border-box;
			}
			
			.button--green {
				background-color: #22BC66;
				border-top: 10px solid #22BC66;
				border-right: 18px solid #22BC66;
				border-bottom: 10px solid #22BC66;
				border-left: 18px solid #22BC66;
			}
			
			.button--red {
				background-color: #FF6136;
				border-top: 10px solid #FF6136;
				border-right: 18px solid #FF6136;
				border-bottom: 10px solid #FF6136;
				border-left: 18px solid #FF6136;
			}
			
			@media only screen and (max-width: 500px) {
				.button {
					width: 100% !important;
					text-align: center !important;
				}
			}
			/* Attribute list ------------------------------ */
			
			.attributes {
				margin: 0 0 21px;
			}
			
			.attributes_content {
				background-color: #F4F4F7;
				padding: 16px;
			}
			
			.attributes_item {
				padding: 0;
			}
			/* Related Items ------------------------------ */
			
			.related {
				width: 100%;
				margin: 0;
				padding: 25px 0 0 0;
				-premailer-width: 100%;
				-premailer-cellpadding: 0;
				-premailer-cellspacing: 0;
			}
			
			.related_item {
				padding: 10px 0;
				color: #CBCCCF;
				font-size: 15px;
				line-height: 18px;
			}
			
			.related_item-title {
				display: block;
				margin: .5em 0 0;
			}
			
			.related_item-thumb {
				display: block;
				padding-bottom: 10px;
			}
			
			.related_heading {
				border-top: 1px solid #CBCCCF;
				text-align: center;
				padding: 25px 0 10px;
			}
			/* Discount Code ------------------------------ */
			
			.discount {
				width: 100%;
				margin: 0;
				padding: 24px;
				-premailer-width: 100%;
				-premailer-cellpadding: 0;
				-premailer-cellspacing: 0;
				background-color: #F4F4F7;
				border: 2px dashed #CBCCCF;
			}
			
			.discount_heading {
				text-align: center;
			}
			
			.discount_body {
				text-align: center;
				font-size: 15px;
			}
			/* Social Icons ------------------------------ */
			
			.social {
				width: auto;
			}
			
			.social td {
				padding: 0;
				width: auto;
			}
			
			.social_icon {
				height: 20px;
				margin: 0 8px 10px 8px;
				padding: 0;
			}
			/* Data table ------------------------------ */
			
			.purchase {
				width: 100%;
				margin: 0;
				padding: 35px 0;
				-premailer-width: 100%;
				-premailer-cellpadding: 0;
				-premailer-cellspacing: 0;
			}
			
			.purchase_content {
				width: 100%;
				margin: 0;
				padding: 25px 0 0 0;
				-premailer-width: 100%;
				-premailer-cellpadding: 0;
				-premailer-cellspacing: 0;
			}
			
			.purchase_item {
				padding: 10px 0;
				color: #51545E;
				font-size: 15px;
				line-height: 18px;
			}
			
			.purchase_heading {
				padding-bottom: 8px;
				border-bottom: 1px solid #EAEAEC;
			}
			
			.purchase_heading p {
				margin: 0;
				color: #85878E;
				font-size: 12px;
			}
			
			.purchase_footer {
				padding-top: 15px;
				border-top: 1px solid #EAEAEC;
			}
			
			.purchase_total {
				margin: 0;
				text-align: right;
				font-weight: bold;
				color: #333333;
			}
			
			.purchase_total--label {
				padding: 0 15px 0 0;
			}
			
			body {
				background-color: #F2F4F6;
				color: #51545E;
			}
			
			p {
				color: #51545E;
			}
			
			.email-wrapper {
				width: 100%;
				margin: 0;
				padding: 0;
				-premailer-width: 100%;
				-premailer-cellpadding: 0;
				-premailer-cellspacing: 0;
				background-color: #F2F4F6;
			}
			
			.email-content {
				width: 100%;
				margin: 0;
				padding: 0;
				-premailer-width: 100%;
				-premailer-cellpadding: 0;
				-premailer-cellspacing: 0;
			}
			/* Masthead ----------------------- */
			
			.email-masthead {
				padding: 25px 0;
				text-align: center;
			}
			
			.email-masthead_logo {
				width: 94px;
			}
			
			.email-masthead_name {
				font-size: 16px;
				font-weight: bold;
				color: #A8AAAF;
				text-decoration: none;
				text-shadow: 0 1px 0 white;
			}
			/* Body ------------------------------ */
			
			.email-body {
				width: 100%;
				margin: 0;
				padding: 0;
				-premailer-width: 100%;
				-premailer-cellpadding: 0;
				-premailer-cellspacing: 0;
			}
			
			.email-body_inner {
				width: 570px;
				margin: 0 auto;
				padding: 0;
				-premailer-width: 570px;
				-premailer-cellpadding: 0;
				-premailer-cellspacing: 0;
				background-color: #FFFFFF;
			}
			
			.email-footer {
				width: 570px;
				margin: 0 auto;
				padding: 0;
				-premailer-width: 570px;
				-premailer-cellpadding: 0;
				-premailer-cellspacing: 0;
				text-align: center;
			}
			
			.email-footer p {
				color: #A8AAAF;
			}
			
			.body-action {
				width: 100%;
				margin: 30px auto;
				padding: 0;
				-premailer-width: 100%;
				-premailer-cellpadding: 0;
				-premailer-cellspacing: 0;
				text-align: center;
			}
			
			.body-sub {
				margin-top: 25px;
				padding-top: 25px;
				border-top: 1px solid #EAEAEC;
			}
			
			.content-cell {
				padding: 45px;
			}
			/*Media Queries ------------------------------ */
			
			@media only screen and (max-width: 600px) {
				.email-body_inner,
				.email-footer {
					width: 100% !important;
				}
			}
			
			@media (prefers-color-scheme: dark) {
				body,
				.email-body,
				.email-body_inner,
				.email-content,
				.email-wrapper,
				.email-masthead,
				.email-footer {
					background-color: #333333 !important;
					color: #FFF !important;
				}
				p,
				ul,
				ol,
				blockquote,
				h1,
				h2,
				h3,
				span,
				.purchase_item {
					color: #FFF !important;
				}
				.attributes_content,
				.discount {
					background-color: #222 !important;
				}
				.email-masthead_name {
					text-shadow: none !important;
				}
			}
			
			:root {
				color-scheme: light dark;
				supported-color-schemes: light dark;
			}
			</style>
		</head>
		<body>
			<table class="email-wrapper" width="100%">
				<tr>
					<td align="center">
						<table class="email-content" width="100%" cellpadding="0" cellspacing="0" role="presentation">
							<tr>
								<td class="email-masthead">
									<a href="https://www.beliesim.com" class="f-fallback email-masthead_name">
									<img src="https://i.ibb.co/nkp2W8Q/logo-beliesimv2-2-1.png" alt="beliesim-logo">
								</a>
								</td>
							</tr>
							<!-- Email Body -->
							<tr>
								<td class="email-body" width="570" cellpadding="0" cellspacing="0">
									<table class="email-body_inner" align="center" width="570" cellpadding="0" cellspacing="0" role="presentation">
										<!-- Body content -->
										<tr>
											<td class="content-cell">
												<div class="f-fallback">
													<h1>Hallo Sobat,</h1>
													<p>Selamat! eSIM <b>%[8]s</b> sudah siap!</p>
													<p>Pastikan untuk mengaktifkannya dalam waktu 60 hari ya!<br>Baca panduan aktivasi eSIM dibawah ini.</p>
													<!-- Discount -->
													<table class="discount" align="center" width="100%" cellpadding="0" cellspacing="0" role="presentation">
														<tr>
															<td align="center">
																<h1 class="f-fallback discount_heading">Let’s Get Started!</h1>
																<p class="f-fallback discount_body"><b>Opsi Pertama - Scan QR Code</b> </p>
																%[2]s
																<p class="f-fallback discount_body">ICCID: %[3]s </p>
																<p class="f-fallback discount_body"><b>Opsi Kedua - Input SM-DP+Address & Activation Code</b> </p>
																<p class="f-fallback discount_body">SM-DP_Address: %[4]s</p>
																<p class="f-fallback discount_body">Activation Code:  %[5]s</p>
															</td>
														</tr>
													</table>
													<table class="purchase" width="100%" cellpadding="0" cellspacing="0" role="presentation">
														<tr>
															<td>
																<h3>No. Order&colon; %[6]d</h3></td>
															<td>
																<h3 class="align-right">%[1]s</h3></td>
														</tr>
														<tr>
															<td colspan="2">
																<table class="purchase_content" width="100%">
																	<tr>
																		<th class="purchase_heading" align="left">
																			<p class="f-fallback">Deskripsi</p>
																		</th>
																		<th class="purchase_heading" align="right">
																			<p class="f-fallback">Harga</p>
																		</th>
																	</tr>
																	<!-- {{#each receipt_details}} -->
																	<tr>
																		<td width="80%" class="purchase_item"><span class="f-fallback">%[8]s<span></td>
																		<td class="align-right" width="20%" class="purchase_item"><span class="f-fallback">Rp%[7]d,00</span></td>
																	</tr>
																	<!-- {{/each}} -->
																	<tr>
																		<td width="80%" class="purchase_footer" valign="middle">
																			<p class="f-fallback purchase_total purchase_total--label">Total</p>
																		</td>
																		<td width="20%" class="purchase_footer" valign="middle">
																			<p class="f-fallback purchase_total">Rp%[7]d,00</p>
																		</td>
																	</tr>
																</table>
															</td>
														</tr>
													</table>
													<h1>Catatan</h1>
													<ul>
														<li><b>Penting!</b> Silakan pindai kode QR di pengaturan eSIM perangkatmu untuk mengunduh profil. Jangan gunakan Kamera atau APLIKASI lain untuk memindai kode, itu harus dilakukan melalui pengaturan perangkatmu.</li>
														<li>Validitas kode QR eSIM adalah 60 hari, eSIMmu telah ditambahkan dan diaktifkan di negara cakupan paket data dalam waktu 60 hari, jika tidak maka akan habis masa berlakunya.</li>
														<li>Jika terjadi masalah selama pengaturan dan aktivasi, harap hubungi <a href="mailto:info@beliesim.com?subject=Permasalahan Aktivasi">Admin BELIESIM</a> dengan kode pesanan eSIM, UID eSIM, dan tangkapan layar kode QR eSIMmu.</li>
													</ul>
													<h2>Lihat panduan kami untuk mempelajari cara mengatur dan mengaktifkan eSIMmu.</h2>
													<br>
													<h1>Panduan Pengguna iPhone</h1>
													<b>Langkah Pengaturan:</b>
													<p>Untuk menambahkan eSIM ke perangkatmu yang mendukung eSIM.</p>
													<ol type="1">
														<li>Buka Pengaturan -> Seluler/Seluler -> Tambahkan Paket Seluler/Seluler</li>
														<li>Pindai Kode QR yang kamu terima. Jika kamu melihat pop up kecil "Paket Seluler Bersertifikat", klik saja OK</li>
														<li>Selesai! Secara default, eSIMmu akan diberi label Sekunder. Di sini kamu juga dapat mengganti nama profil eSIM baru ini, kami menyarankan kamu mengganti nama labelnya menjadi "eSIM-BELIESIM" agar mudah dikenali.</li>
													</ol>
													<b>Langkah Aktivasi:</b>
													<p>Untuk mengaktifkan eSIMmu di negara cakupan paket data.</p>
													<ol type="1">
														<li>Sesampainya di tempat tujuan, silahkan aktifkan eSIM "eSIM-BELIESIM"mu</li>
														<li>Gunakan eSIM "eSIM-BELIESIM"mu untuk data seluler</li>
														<li>Aktifkan Roaming data</li>
														<li>Untuk mengakses internet, kamu mungkin perlu mengatur Access Point Name (APN) seperti detail di atas.</li>
													</ol>
													<br>
													<h1>Panduan Pengguna Android</h1>
													<b>Langkah Pengaturan:</b>
													<p>Untuk menambahkan eSIM ke perangkatmu yang mendukung eSIM.</p>
													<ol type="1">
														<li>Buka Pengaturan-> Jaringan & internet -> Jaringan seluler -> Lanjutan -> Operator -> Tambahkan operator</li>
														<li>Tambahkan operator dengan kode QR yang diterima. Jika kamu melihat sedikit pop up "Nomor telepon ditambahkan", klik saja Selesai.</li>
														<li>Selesai! Secara default, eSIMmu akan diaktifkan. Jika status operatormu "tidak pernah diaktifkan", ikuti langkah aktivasi di bawah ini.</li>
													</ol>
													<b>Langkah Aktivasi:</b>
													<p>Untuk mengaktifkan eSIMmu di negara cakupan paket data.</p>
													<ol type="1">
														<li>Sesampainya di tempat tujuan, harap aktifkan eSIM "eSIM-BELIESIM"mu.</li>
														<li>Gunakan eSIM "eSIM-BELESIM"mu untuk data seluler.</li>
													</ol>
													<table class="body-sub" role="presentation">
														<tr>
															<td>
																<p class="f-fallback sub">Jika kamu memiliki pertanyaan, silahkan hubungi <a href="mailto:info@beliesim.com?subject=Pertanyaan Seputar Receipt">Admin BELIESIM</a> untuk mendapatkan bantuan.</p>
													<p class="f-fallback sub">Happy Traveling,
														<br>
														<br>BELIESIM TEAM</p>
															</td>
														</tr>
													</table>
												</div>
											</td>
										</tr>
									</table>
								</td>
							</tr>
							<tr>
								<td>
									<table class="email-footer" align="center" width="570" cellpadding="0" cellspacing="0" role="presentation">
										<tr>
											<td class="content-cell" align="center">
												<p class="f-fallback sub align-center">
													PT. CAKEPLABS GLOBAL TEKNOLOGI
													<br>Bali, Indonesia.
												</p>
											</td>
										</tr>
									</table>
								</td>
							</tr>
						</table>
					</td>
				</tr>
			</table>
		</body>
	</html>
		`,
		order.PaidAt.Time.Format("2 January 2006"),
		qrCodeHTML,
		iccid,
		smdp,
		activationCode,
		order.ID,
		order.IdrPrice,
		productName,
	)
}

type orderUsecase struct {
	server *domain.ConcreteServer
	ss     domain.SmtpService
	xhr    order.XenditHttpRepository
	uhr    order.UsimsaHttpRepository
}

func NewOrderUsecase(server *domain.ConcreteServer, ss domain.SmtpService, xhr order.XenditHttpRepository, uhr order.UsimsaHttpRepository) order.Usecase {
	return &orderUsecase{
		ss:     ss,
		server: server,
		xhr:    xhr,
		uhr:    uhr,
	}
}

func (ou *orderUsecase) CreatePendingOrder(c context.Context, esimId int64, quantity int32, email string) (*order.CreateOrderResp, error) {
	ctx, cancel := context.WithTimeout(c, ou.server.Timeout)
	defer cancel()

	// Get the current UTC time
	currentTime := time.Now().UTC()

	// Convert the current UTC time to a Unix timestamp
	timestamp := currentTime.Unix()

	var res *order.CreateOrderResp

	err := ou.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		esim, err := q.GetEsim(ctx, esimId)

		if err != nil {
			return fmt.Errorf("(usecase/order.go, CreatePendingOrder) GetEsim error : %w", err)
		}

		arg := db.CreateOrderParams{
			CustomerEmail:  email,
			EsimID:         sql.NullInt64{Int64: esimId, Valid: true},
			CountryCode:    esim.CountryCode,
			PlanOption:     esim.PlanOption,
			DataAmount:     esim.DataAmount,
			DataUnit:       esim.DataUnit,
			DurationInDays: esim.DurationInDays,
			OptionID:       esim.OptionID,
			IdrPrice:       esim.IdrPrice,
			Quantity:       quantity,
		}

		createdOrder, err := q.CreateOrder(ctx, arg)

		if err != nil {
			return fmt.Errorf("(usecase/order.go, CreatePendingOrder) CreateOrder error : %w", err)
		}

		uniqueOrderId := fmt.Sprintf("beliesim-%d-%d-%d", timestamp, esimId, createdOrder.ID)

		invoiceReq := order.CreateOrderInvoiceReq{
			// ExternalID: strconv.FormatInt(createdOrder.ID, 10),
			ExternalID:      uniqueOrderId,
			Amount:          int(esim.IdrPrice) * int(quantity),
			Currency:        "IDR",
			Description:     "Esim purchase",
			InvoiceDuration: 3600,
			Items: []domain.InvoiceItem{
				{
					Name:     esim.OptionID,
					Quantity: int(quantity),
					Price:    int(esim.IdrPrice),
				},
			},
			Customer: domain.InvoiceCustomer{
				Email: email,
			},
			SuccessRedirectURL: ou.server.Config.XenditInvoiceSuccessRedirectUrl,
			FailureRedirectURL: ou.server.Config.XenditInvoiceFailureRedirectUrl,
		}

		invoice, err := ou.xhr.CreateInvoice(invoiceReq)

		if err != nil {
			return fmt.Errorf("(usecase/order.go, CreatePendingOrder) CreateInvoice error : %w", err)
		}

		addOrderInvoiceArg := db.AddOrderInvoiceParams{
			ID:              createdOrder.ID,
			UniqueOrderID:   sql.NullString{String: invoice.ExternalID, Valid: true},
			XenditInvoiceID: sql.NullString{String: invoice.ID, Valid: true},
			PaymentStatus:   db.NullPaymentStatus{PaymentStatus: db.PaymentStatus(invoice.Status), Valid: true},
		}

		_, err = q.AddOrderInvoice(ctx, addOrderInvoiceArg)

		if err != nil {
			return fmt.Errorf("(usecase/order.go, CreatePendingOrder) AddOrderInvoice error : %w", err)
		}

		res = &order.CreateOrderResp{
			InvoiceUrl: invoice.InvoiceURL,
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	return res, err
}

func (ou *orderUsecase) UpdateInvoicePaymentAndOrderUsimsa(c context.Context, callback order.InvoicePaymentCallback) (*db.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.server.Timeout)
	defer cancel()

	// if err != nil {
	// 	fmt.Printf("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) orderId parseInt error : %v\n", err)
	// 	return nil, err
	// }

	utcTime := time.Now().UTC()

	var paidAt sql.NullTime

	var subscribedUsimsaOrder *order.SubscribeUsimsaOrderResp

	var err error

	if db.PaymentStatus(callback.Status) == db.PaymentStatusPAID {
		paidAt.Time = utcTime
		paidAt.Valid = true

		var products []domain.UsimsaOrderItem

		for _, item := range callback.Items {
			products = append(products, domain.UsimsaOrderItem{
				OptionId: item.Name,
				Quantity: int32(item.Quantity),
			})
		}

		subscribeOrderReq := order.SubscribeUsimsaOrderReq{
			UniqueOrderId: callback.ExternalID,
			Products:      products,
		}

		subscribedUsimsaOrder, err = ou.uhr.SubscribeOrder(subscribeOrderReq)

		if subscribedUsimsaOrder.Code != "0000" {
			err = errors.New(subscribedUsimsaOrder.Message)
			fmt.Printf("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) Subscribe Order error with code %s : %s\n", subscribedUsimsaOrder.Code, err)
			return nil, err
		}

		if err != nil {
			fmt.Printf("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) Subscribe Order error : %v\n", err)
			return nil, err
		}

	}

	var order db.Order

	err = ou.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		if q == nil {
			err := errors.New("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) q from ExecTx is nil")
			fmt.Print(err)
			return err
		}

		arg := db.UpdateOrderParams{
			UniqueOrderID:   sql.NullString{String: callback.ExternalID, Valid: true},
			XenditInvoiceID: sql.NullString{String: callback.ID, Valid: true},
			PaymentStatus:   db.NullPaymentStatus{PaymentStatus: db.PaymentStatus(callback.Status), Valid: true},
			PaidAt:          paidAt,
		}

		order, err = q.UpdateOrder(ctx, arg)

		if err != nil {
			fmt.Printf("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) UpdateOrder error : %v\n", err)
			return err
		}

		//Execute add topup ids data to db if the invoice is paid
		if db.PaymentStatus(callback.Status) == db.PaymentStatusPAID && subscribedUsimsaOrder != nil {
			for _, orderedItem := range subscribedUsimsaOrder.Products {
				orderIdStr, err := util.GetLastSegmentAfterDash(callback.ExternalID)

				if err != nil {
					return err
				}

				orderId, err := strconv.ParseInt(orderIdStr, 10, 64)

				if err != nil {
					fmt.Printf("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) error parsing orderId : %v\n", err)
					return err
				}

				arg := db.AddOrderTopUpIdParams{
					OrderID:       orderId,
					UsimsaTopupID: orderedItem.TopupId,
				}
				_, err = q.AddOrderTopUpId(ctx, arg)

				if err != nil {
					fmt.Printf("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) AddOrderTopUpId error : %v\n", err)

					return err
				}
			}
		}

		return err
	})

	if err != nil {
		fmt.Printf("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) ExecTx error : %v\n", err)
		return nil, err
	}

	if order == (db.Order{}) {
		err := errors.New("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) order is nil and can't be returned")
		fmt.Print(err)
		return nil, err
	}

	return &order, err

}

// func (ou *orderUsecase) UpdateInvoicePaymentAndOrderUsimsa(c context.Context, callback order.InvoicePaymentCallback) (*db.Order, error) {
// 	ctx, cancel := context.WithTimeout(c, ou.server.Timeout)
// 	defer cancel()

// 	// if err != nil {
// 	// 	fmt.Printf("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) orderId parseInt error : %v\n", err)
// 	// 	return nil, err
// 	// }

// 	utcTime := time.Now().UTC()

// 	var paidAt sql.NullTime

// 	var subscribedUsimsaOrder *order.SubscribeUsimsaOrderResp

// 	var err error

// 	if db.PaymentStatus(callback.Status) == db.PaymentStatusPAID {
// 		paidAt.Time = utcTime
// 		paidAt.Valid = true

// 		var products []domain.UsimsaOrderItem

// 		for _, item := range callback.Items {
// 			products = append(products, domain.UsimsaOrderItem{
// 				OptionId: item.Name,
// 				Quantity: int32(item.Quantity),
// 			})
// 		}

// 		subscribeOrderReq := order.SubscribeUsimsaOrderReq{
// 			UniqueOrderId: callback.ExternalID,
// 			Products:      products,
// 		}

// 		subscribedUsimsaOrder, err = ou.uhr.SubscribeOrder(subscribeOrderReq)

// 		fmt.Printf("(usecase/order.go,  UpdateInvoicePaymentAndOrderUsimsa) cek subscribedUsimsaOrder : %+v\n", subscribedUsimsaOrder)

// 		if subscribedUsimsaOrder.Code != "0000" {
// 			err = errors.New(subscribedUsimsaOrder.Message)
// 			fmt.Printf("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) Subscribe Order error with code %s : %s\n", subscribedUsimsaOrder.Code, err)
// 			return nil, err
// 		}

// 		if err != nil {
// 			fmt.Printf("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) Subscribe Order error : %v\n", err)
// 			return nil, err
// 		}

// 	}

// 	var order db.Order

// 	// Create a channel to signal errors from goroutines
// 	errCh := make(chan error, 1)

// 	// Create a WaitGroup to wait for all goroutines to finish
// 	var wg sync.WaitGroup

// 	// Start a goroutine to listen for errors and cancel signals
// 	go func() {
// 		defer close(errCh)

// 		select {
// 		case <-ctx.Done():
// 			// Context canceled, no need to process errors
// 			return
// 		case err := <-errCh:
// 			// An error occurred in one of the goroutines
// 			fmt.Printf("Received error: %v\n", err)
// 			cancel() // Cancel the context to signal other goroutines to stop
// 		}
// 	}()

// 	err = ou.server.Store.ExecTx(ctx, func(q *db.Queries) error {

// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()

// 			// Check for cancellation signal
// 			select {
// 			case <-ctx.Done():
// 				// Context canceled, no need to process errors
// 				return
// 			default:

// 				arg := db.UpdateOrderParams{
// 					UniqueOrderID:   sql.NullString{String: callback.ExternalID, Valid: true},
// 					XenditInvoiceID: sql.NullString{String: callback.ID, Valid: true},
// 					PaymentStatus:   db.NullPaymentStatus{PaymentStatus: db.PaymentStatus(callback.Status), Valid: true},
// 					PaidAt:          paidAt,
// 				}

// 				order, err = q.UpdateOrder(ctx, arg)

// 				if err != nil {
// 					fmt.Printf("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) UpdateOrder error : %v\n", err)
// 					errCh <- err // Send the error to the channel
// 					return
// 				}
// 			}

// 		}()

// 		for _, orderedItem := range subscribedUsimsaOrder.Products {
// 			wg.Add(1)
// 			go func(item domain.UsimsaOrderedItem) {
// 				defer wg.Done()

// 				select {
// 				case <-ctx.Done():
// 					// Context canceled, no need to process errors
// 					return
// 				default:

// 					orderIdStr := utils.GetLastSegmentAfterDash(callback.ExternalID)

// 					orderId, err := strconv.ParseInt(orderIdStr, 10, 64)

// 					if err != nil {
// 						fmt.Printf("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) error parsing orderId : %v\n", err)
// 						errCh <- err
// 						return
// 					}

// 					arg := db.AddOrderTopUpIdParams{
// 						OrderID:       orderId,
// 						UsimsaTopupID: item.TopupId,
// 					}
// 					_, err = q.AddOrderTopUpId(ctx, arg)

// 					if err != nil {
// 						fmt.Printf("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) AddOrderTopUpId error : %v\n", err)
// 						errCh <- err
// 						return
// 					}
// 				}

// 			}(orderedItem)
// 		}

// 		wg.Wait() // Wait for all goroutines to finish
// 		close(errCh)

// 		return err
// 	})

// 	if err != nil {
// 		fmt.Printf("(usecase/order.go, UpdateInvoicePaymentAndOrderUsimsa) ExecTx error : %v\n", err)
// 		return nil, err
// 	}

// 	return &order, err

// }

func (ou *orderUsecase) SendPurchasedEsimEmail(c context.Context, topupId string, iccid string, smdp string, activationCode string, qrCodeImgUrl string) (*sesv2.SendEmailOutput, error) {
	ctx, cancel := context.WithTimeout(c, ou.server.Timeout)
	defer cancel()

	// topupIdArg := sql.NullString{
	// 	String: topupId,
	// 	Valid:  true,
	// }

	order, err := ou.server.Store.GetOrderByTopupId(ctx, topupId)
	if err != nil {
		fmt.Printf("(usecase/order.go, SendPurchasedEsimEmail) error on GetOrderByTopupId : %v\n", err)
		return nil, err
	}

	esim, err := ou.server.Store.GetEsim(ctx, order.EsimID.Int64)
	if err != nil {
		fmt.Printf("(usecase/order.go, SendPurchasedEsimEmail) error on GetEsim : %v\n", err)
		return nil, err
	}

	senderEmail := Sender

	qrCodeHTML := fmt.Sprintf(`<img src="%s" alt="QR Code">`, qrCodeImgUrl)
	htmlBody := generateHTMLBody(order, esim, iccid, topupId, smdp, activationCode, qrCodeHTML)

	// Assemble the email.
	input := &sesv2.SendEmailInput{
		FromEmailAddress: &senderEmail,
		Destination: &sesv2.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(order.CustomerEmail),
			},
		},
		Content: &sesv2.EmailContent{
			Simple: &sesv2.Message{
				Body: &sesv2.Body{
					Html: &sesv2.Content{
						Charset: aws.String(CharSet),
						Data:    aws.String(htmlBody),
					},
					// Text: &sesv2.Content{
					// 	Charset: aws.String(CharSet),
					// 	Data:    aws.String(textBody),
					// },
				},
				Subject: &sesv2.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(Subject),
				},
			},
		},
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	emailResp, err := ou.ss.SendEmail(input)

	if err != nil {
		fmt.Printf("(usecase/order.go, SendPurchasedEsimEmail) error on SendEmail : %v\n", err)
		return nil, err
	}

	return emailResp, err
}

func (ou *orderUsecase) Get(c context.Context, id int64) (*db.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.server.Timeout)
	defer cancel()

	res, err := ou.server.Store.GetOrder(ctx, id)

	return &res, err
}
