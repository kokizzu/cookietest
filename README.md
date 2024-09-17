
# test httpOnly secure cookie

```shell
caddy run
# if blocked by browser
# firefox > Advanced > click the certificate, download the PEM (not chain)
# Settings > Advanced > Certificates > Import ^

# install autorecompile
go install github.com/air-verse/air

cd sts
air
```

open the https://console.local.benalu.dev afterward and click the button, inspect element

![image](https://github.com/user-attachments/assets/adea3a47-bbdb-4752-9b9d-e65c6b7ad5a8)
