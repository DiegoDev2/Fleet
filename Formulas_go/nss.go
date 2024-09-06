package main

// Nss - Libraries for security-enabled client and server applications
// Homepage: https://firefox-source-docs.mozilla.org/security/nss/index.html

import (
	"fmt"
	
	"os/exec"
)

func installNss() {
	// Método 1: Descargar y extraer .tar.gz
	nss_tar_url := "https://ftp.mozilla.org/pub/security/nss/releases/NSS_3_103_RTM/src/nss-3.103.tar.gz"
	nss_cmd_tar := exec.Command("curl", "-L", nss_tar_url, "-o", "package.tar.gz")
	err := nss_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nss_zip_url := "https://ftp.mozilla.org/pub/security/nss/releases/NSS_3_103_RTM/src/nss-3.103.zip"
	nss_cmd_zip := exec.Command("curl", "-L", nss_zip_url, "-o", "package.zip")
	err = nss_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nss_bin_url := "https://ftp.mozilla.org/pub/security/nss/releases/NSS_3_103_RTM/src/nss-3.103.bin"
	nss_cmd_bin := exec.Command("curl", "-L", nss_bin_url, "-o", "binary.bin")
	err = nss_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nss_src_url := "https://ftp.mozilla.org/pub/security/nss/releases/NSS_3_103_RTM/src/nss-3.103.src.tar.gz"
	nss_cmd_src := exec.Command("curl", "-L", nss_src_url, "-o", "source.tar.gz")
	err = nss_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nss_cmd_direct := exec.Command("./binary")
	err = nss_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: nspr")
exec.Command("latte", "install", "nspr")
}
