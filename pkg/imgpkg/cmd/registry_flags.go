package cmd

import (
	"os"

	ctlimg "github.com/btwiuse/dkg/pkg/imgpkg/image"
	"github.com/spf13/cobra"
)

type RegistryFlags struct {
	CACertPaths []string
	VerifyCerts bool

	Username string
	Password string
	Token    string
	Anon     bool
}

func (s *RegistryFlags) Set(cmd *cobra.Command) {
	cmd.Flags().StringSliceVar(&s.CACertPaths, "registry-ca-cert-path", nil, "Add CA certificates for registry API (format: /tmp/foo) (can be specified multiple times)")
	cmd.Flags().BoolVar(&s.VerifyCerts, "registry-verify-certs", true, "Set whether to verify server's certificate chain and host name")

	cmd.Flags().StringVar(&s.Username, "registry-username", "", "Set username for auth ($IMGPKG_USERNAME)")
	cmd.Flags().StringVar(&s.Password, "registry-password", "", "Set password for auth ($IMGPKG_PASSWORD)")
	cmd.Flags().StringVar(&s.Token, "registry-token", "", "Set token for auth ($IMGPKG_TOKEN)")
	cmd.Flags().BoolVar(&s.Anon, "registry-anon", false, "Set anonymous auth ($IMGPKG_ANON)")
}

func (s *RegistryFlags) AsRegistryOpts() ctlimg.RegistryOpts {
	opts := ctlimg.RegistryOpts{
		CACertPaths: s.CACertPaths,
		VerifyCerts: s.VerifyCerts,

		Username: s.Username,
		Password: s.Password,
		Token:    s.Token,
		Anon:     s.Anon,
	}

	if len(opts.Username) == 0 {
		opts.Username = os.Getenv("IMGPKG_USERNAME")
	}
	if len(opts.Password) == 0 {
		opts.Password = os.Getenv("IMGPKG_PASSWORD")
	}
	if len(opts.Token) == 0 {
		opts.Token = os.Getenv("IMGPKG_TOKEN")
	}
	if os.Getenv("IMGPKG_ANON") == "true" {
		opts.Anon = true
	}

	return opts
}
