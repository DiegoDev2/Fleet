package main

// Gnutls - GNU Transport Layer Security (TLS) Library
// Homepage: https://gnutls.org/

import (
	"fmt"
	
	"os/exec"
)

func installGnutls() {
	// Método 1: Descargar y extraer .tar.gz
	gnutls_tar_url := "https://www.gnupg.org/ftp/gcrypt/gnutls/v3.8/gnutls-3.8.4.tar.xz"
	gnutls_cmd_tar := exec.Command("curl", "-L", gnutls_tar_url, "-o", "package.tar.gz")
	err := gnutls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnutls_zip_url := "https://www.gnupg.org/ftp/gcrypt/gnutls/v3.8/gnutls-3.8.4.tar.xz"
	gnutls_cmd_zip := exec.Command("curl", "-L", gnutls_zip_url, "-o", "package.zip")
	err = gnutls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnutls_bin_url := "https://www.gnupg.org/ftp/gcrypt/gnutls/v3.8/gnutls-3.8.4.tar.xz"
	gnutls_cmd_bin := exec.Command("curl", "-L", gnutls_bin_url, "-o", "binary.bin")
	err = gnutls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnutls_src_url := "https://www.gnupg.org/ftp/gcrypt/gnutls/v3.8/gnutls-3.8.4.tar.xz"
	gnutls_cmd_src := exec.Command("curl", "-L", gnutls_src_url, "-o", "source.tar.gz")
	err = gnutls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnutls_cmd_direct := exec.Command("./binary")
	err = gnutls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: ca-certificates")
	exec.Command("latte", "install", "ca-certificates").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: libidn2")
	exec.Command("latte", "install", "libidn2").Run()
	fmt.Println("Instalando dependencia: libtasn1")
	exec.Command("latte", "install", "libtasn1").Run()
	fmt.Println("Instalando dependencia: libunistring")
	exec.Command("latte", "install", "libunistring").Run()
	fmt.Println("Instalando dependencia: nettle")
	exec.Command("latte", "install", "nettle").Run()
	fmt.Println("Instalando dependencia: p11-kit")
	exec.Command("latte", "install", "p11-kit").Run()
	fmt.Println("Instalando dependencia: unbound")
	exec.Command("latte", "install", "unbound").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
