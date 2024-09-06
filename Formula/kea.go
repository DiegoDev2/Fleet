package main

// Kea - DHCP server
// Homepage: https://www.isc.org/kea/

import (
	"fmt"
	
	"os/exec"
)

func installKea() {
	// Método 1: Descargar y extraer .tar.gz
	kea_tar_url := "https://ftp.isc.org/isc/kea/2.6.1/kea-2.6.1.tar.gz"
	kea_cmd_tar := exec.Command("curl", "-L", kea_tar_url, "-o", "package.tar.gz")
	err := kea_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kea_zip_url := "https://ftp.isc.org/isc/kea/2.6.1/kea-2.6.1.zip"
	kea_cmd_zip := exec.Command("curl", "-L", kea_zip_url, "-o", "package.zip")
	err = kea_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kea_bin_url := "https://ftp.isc.org/isc/kea/2.6.1/kea-2.6.1.bin"
	kea_cmd_bin := exec.Command("curl", "-L", kea_bin_url, "-o", "binary.bin")
	err = kea_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kea_src_url := "https://ftp.isc.org/isc/kea/2.6.1/kea-2.6.1.src.tar.gz"
	kea_cmd_src := exec.Command("curl", "-L", kea_src_url, "-o", "source.tar.gz")
	err = kea_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kea_cmd_direct := exec.Command("./binary")
	err = kea_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: log4cplus")
	exec.Command("latte", "install", "log4cplus").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
