package main

// Gwenhywfar - Utility library required by aqbanking and related software
// Homepage: https://www.aquamaniac.de/rdm/projects/gwenhywfar

import (
	"fmt"
	
	"os/exec"
)

func installGwenhywfar() {
	// Método 1: Descargar y extraer .tar.gz
	gwenhywfar_tar_url := "https://www.aquamaniac.de/rdm/attachments/download/501/gwenhywfar-5.10.2.tar.gz"
	gwenhywfar_cmd_tar := exec.Command("curl", "-L", gwenhywfar_tar_url, "-o", "package.tar.gz")
	err := gwenhywfar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gwenhywfar_zip_url := "https://www.aquamaniac.de/rdm/attachments/download/501/gwenhywfar-5.10.2.zip"
	gwenhywfar_cmd_zip := exec.Command("curl", "-L", gwenhywfar_zip_url, "-o", "package.zip")
	err = gwenhywfar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gwenhywfar_bin_url := "https://www.aquamaniac.de/rdm/attachments/download/501/gwenhywfar-5.10.2.bin"
	gwenhywfar_cmd_bin := exec.Command("curl", "-L", gwenhywfar_bin_url, "-o", "binary.bin")
	err = gwenhywfar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gwenhywfar_src_url := "https://www.aquamaniac.de/rdm/attachments/download/501/gwenhywfar-5.10.2.src.tar.gz"
	gwenhywfar_cmd_src := exec.Command("curl", "-L", gwenhywfar_src_url, "-o", "source.tar.gz")
	err = gwenhywfar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gwenhywfar_cmd_direct := exec.Command("./binary")
	err = gwenhywfar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
