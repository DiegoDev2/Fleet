package main

// Swtpm - Software TPM Emulator based on libtpms
// Homepage: https://github.com/stefanberger/swtpm

import (
	"fmt"
	
	"os/exec"
)

func installSwtpm() {
	// Método 1: Descargar y extraer .tar.gz
	swtpm_tar_url := "https://github.com/stefanberger/swtpm/archive/refs/tags/v0.9.0.tar.gz"
	swtpm_cmd_tar := exec.Command("curl", "-L", swtpm_tar_url, "-o", "package.tar.gz")
	err := swtpm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swtpm_zip_url := "https://github.com/stefanberger/swtpm/archive/refs/tags/v0.9.0.zip"
	swtpm_cmd_zip := exec.Command("curl", "-L", swtpm_zip_url, "-o", "package.zip")
	err = swtpm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swtpm_bin_url := "https://github.com/stefanberger/swtpm/archive/refs/tags/v0.9.0.bin"
	swtpm_cmd_bin := exec.Command("curl", "-L", swtpm_bin_url, "-o", "binary.bin")
	err = swtpm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swtpm_src_url := "https://github.com/stefanberger/swtpm/archive/refs/tags/v0.9.0.src.tar.gz"
	swtpm_cmd_src := exec.Command("curl", "-L", swtpm_src_url, "-o", "source.tar.gz")
	err = swtpm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swtpm_cmd_direct := exec.Command("./binary")
	err = swtpm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gawk")
	exec.Command("latte", "install", "gawk").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: socat")
	exec.Command("latte", "install", "socat").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: json-glib")
	exec.Command("latte", "install", "json-glib").Run()
	fmt.Println("Instalando dependencia: libtasn1")
	exec.Command("latte", "install", "libtasn1").Run()
	fmt.Println("Instalando dependencia: libtpms")
	exec.Command("latte", "install", "libtpms").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libseccomp")
	exec.Command("latte", "install", "libseccomp").Run()
	fmt.Println("Instalando dependencia: net-tools")
	exec.Command("latte", "install", "net-tools").Run()
}
