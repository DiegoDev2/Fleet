package main

// Mailutils - Swiss Army knife of email handling
// Homepage: https://mailutils.org/

import (
	"fmt"
	
	"os/exec"
)

func installMailutils() {
	// Método 1: Descargar y extraer .tar.gz
	mailutils_tar_url := "https://ftp.gnu.org/gnu/mailutils/mailutils-3.17.tar.gz"
	mailutils_cmd_tar := exec.Command("curl", "-L", mailutils_tar_url, "-o", "package.tar.gz")
	err := mailutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mailutils_zip_url := "https://ftp.gnu.org/gnu/mailutils/mailutils-3.17.zip"
	mailutils_cmd_zip := exec.Command("curl", "-L", mailutils_zip_url, "-o", "package.zip")
	err = mailutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mailutils_bin_url := "https://ftp.gnu.org/gnu/mailutils/mailutils-3.17.bin"
	mailutils_cmd_bin := exec.Command("curl", "-L", mailutils_bin_url, "-o", "binary.bin")
	err = mailutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mailutils_src_url := "https://ftp.gnu.org/gnu/mailutils/mailutils-3.17.src.tar.gz"
	mailutils_cmd_src := exec.Command("curl", "-L", mailutils_src_url, "-o", "source.tar.gz")
	err = mailutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mailutils_cmd_direct := exec.Command("./binary")
	err = mailutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: gsasl")
exec.Command("latte", "install", "gsasl")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: libunistring")
exec.Command("latte", "install", "libunistring")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
