package main

// RpkiClient - OpenBSD portable rpki-client
// Homepage: https://www.rpki-client.org/

import (
	"fmt"
	
	"os/exec"
)

func installRpkiClient() {
	// Método 1: Descargar y extraer .tar.gz
	rpkiclient_tar_url := "https://ftp.openbsd.org/pub/OpenBSD/rpki-client/rpki-client-9.2.tar.gz"
	rpkiclient_cmd_tar := exec.Command("curl", "-L", rpkiclient_tar_url, "-o", "package.tar.gz")
	err := rpkiclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rpkiclient_zip_url := "https://ftp.openbsd.org/pub/OpenBSD/rpki-client/rpki-client-9.2.zip"
	rpkiclient_cmd_zip := exec.Command("curl", "-L", rpkiclient_zip_url, "-o", "package.zip")
	err = rpkiclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rpkiclient_bin_url := "https://ftp.openbsd.org/pub/OpenBSD/rpki-client/rpki-client-9.2.bin"
	rpkiclient_cmd_bin := exec.Command("curl", "-L", rpkiclient_bin_url, "-o", "binary.bin")
	err = rpkiclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rpkiclient_src_url := "https://ftp.openbsd.org/pub/OpenBSD/rpki-client/rpki-client-9.2.src.tar.gz"
	rpkiclient_cmd_src := exec.Command("curl", "-L", rpkiclient_src_url, "-o", "source.tar.gz")
	err = rpkiclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rpkiclient_cmd_direct := exec.Command("./binary")
	err = rpkiclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libretls")
	exec.Command("latte", "install", "libretls").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: rsync")
	exec.Command("latte", "install", "rsync").Run()
}
