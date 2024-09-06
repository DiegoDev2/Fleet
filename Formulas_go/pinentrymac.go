package main

// PinentryMac - Pinentry for GPG on Mac
// Homepage: https://github.com/GPGTools/pinentry

import (
	"fmt"
	
	"os/exec"
)

func installPinentryMac() {
	// Método 1: Descargar y extraer .tar.gz
	pinentrymac_tar_url := "https://github.com/GPGTools/pinentry/archive/refs/tags/v1.1.1.1.tar.gz"
	pinentrymac_cmd_tar := exec.Command("curl", "-L", pinentrymac_tar_url, "-o", "package.tar.gz")
	err := pinentrymac_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pinentrymac_zip_url := "https://github.com/GPGTools/pinentry/archive/refs/tags/v1.1.1.1.zip"
	pinentrymac_cmd_zip := exec.Command("curl", "-L", pinentrymac_zip_url, "-o", "package.zip")
	err = pinentrymac_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pinentrymac_bin_url := "https://github.com/GPGTools/pinentry/archive/refs/tags/v1.1.1.1.bin"
	pinentrymac_cmd_bin := exec.Command("curl", "-L", pinentrymac_bin_url, "-o", "binary.bin")
	err = pinentrymac_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pinentrymac_src_url := "https://github.com/GPGTools/pinentry/archive/refs/tags/v1.1.1.1.src.tar.gz"
	pinentrymac_cmd_src := exec.Command("curl", "-L", pinentrymac_src_url, "-o", "source.tar.gz")
	err = pinentrymac_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pinentrymac_cmd_direct := exec.Command("./binary")
	err = pinentrymac_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libassuan@2")
exec.Command("latte", "install", "libassuan@2")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
}
