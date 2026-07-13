
# Example for cross-origin httpOnly secure cookie

```shell
caddy run
# if blocked by browser
# firefox > Advanced > click the certificate, download the PEM (not chain)
# Settings > Privacy > Certificates > View Certificates > Import the pem above

# install autorecompile
go install github.com/air-verse/air

cd sts
air
```

open the https://console.local.benalu.dev afterward and click the button, inspect element

![image](https://github.com/user-attachments/assets/adea3a47-bbdb-4752-9b9d-e65c6b7ad5a8)

## Maintenance Checklist

- [x] Update the Go runtime directive to Go 1.26.5.
- [x] Add Makefile targets for vulnerability checks and arbitrary commands.
- [x] Run `make test`.
- [x] Run `make verify-dependency-security`.
- [x] Run `make vulncheck`; no vulnerabilities were found.
