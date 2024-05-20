package promptpayqr

import (
	"flag"
	"fmt"
	"image"
	"os"

	"github.com/divan/qrlogo"
	"github.com/skip2/go-qrcode"
)

// QRCodeGenerator is a struct that provides methods to generate QR code payloads.
type QRCodeGenerator struct{}

// NewQRCodeGenerator creates a new instance of QRCodeGenerator.
func NewQRCodeGenerator() *QRCodeGenerator {
	return &QRCodeGenerator{}
}

//change v1.0.1

// GeneratePayload generates a QR code payload for a target with an optional amount.
func (q *QRCodeGenerator) GeneratePayload(target string, amount *string) string {
	// Placeholder implementation; replace with actual payload generation logic.
	if amount != nil {
		return fmt.Sprintf("payload_for_%s_with_amount_%s", target, *amount)
	}
	return fmt.Sprintf("payload_for_%s", target)
}

// GenerateBillPaymentPayload generates a QR code payload for bill payment.
func (q *QRCodeGenerator) GenerateBillPaymentPayload(billerID, ref1, ref2 string, terminalID, amount *string) string {
	// Placeholder implementation; replace with actual bill payment payload generation logic.
	return fmt.Sprintf("bill_payment_payload_for_%s_%s_%s_%s_%s", billerID, ref1, ref2, *terminalID, *amount)
}

func QRForTargetWithAmount(target, amount string) (string, error) {
	qr := New()
	payload := qr.GeneratePayload(target, &amount)
	return payload, nil
}

func QRForBillPayment(billerID string, ref1 string, ref2 string, terminalID string, amount string) (*[]byte, error) {
	qr := New()
	payload := qr.GenerateBillPaymentPayload(billerID, ref1, ref2, &terminalID, &amount)

	png, err := qrcode.Encode(payload, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return &png, nil
}

func QRForTarget(target string) (*[]byte, error) {
	qr := New()
	payload := qr.GeneratePayload(target, nil)

	png, err := qrcode.Encode(payload, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return &png, nil
}

func QRWithPromptpayLogoForTargetWithAmount(target, amount string) (*[]byte, error) {
	var (
		input = flag.String("promptpay", "promptpay.png", "Prompt Pay Logo")
		size  = flag.Int("size", 256, "Image size in pixels")
	)
	flag.Parse()

	qr := New()
	payload := qr.GeneratePayload(target, &amount)

	file, err := os.Open(*input)
	if err != nil {
		fmt.Println("Failed to open logo:", err)
		os.Exit(1)
	}
	defer file.Close()

	logo, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Failed to decode PNG with logo:", err)
		os.Exit(1)
	}

	qrImage, err := qrlogo.Encode(payload, logo, *size)
	if err != nil {
		fmt.Println("Failed to encode QR:", err)
		os.Exit(1)
	}

	qrBytes := qrImage.Bytes()
	return &qrBytes, err
}
