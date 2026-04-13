package provider

import "context"

type Provider interface {
	ResolveCertificate(ctx context.Context, domain string, current *ObservedCertificate, options ResolveOptions) (*CertificateResolution, error)
}
