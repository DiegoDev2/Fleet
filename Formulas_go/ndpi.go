package main

// Ndpi - Deep Packet Inspection (DPI) library
// Homepage: https://www.ntop.org/products/deep-packet-inspection/ndpi/

import (
	"fmt"
	
	"os/exec"
)

func installNdpi() {
	// Método 1: Descargar y extraer .tar.gz
	ndpi_tar_url := "https://github.com/ntop/nDPI/archive/refs/tags/4.4.tar.gz"
	ndpi_cmd_tar := exec.Command("curl", "-L", ndpi_tar_url, "-o", "package.tar.gz")
	err := ndpi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ndpi_zip_url := "https://github.com/ntop/nDPI/archive/refs/tags/4.4.zip"
	ndpi_cmd_zip := exec.Command("curl", "-L", ndpi_zip_url, "-o", "package.zip")
	err = ndpi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ndpi_bin_url := "https://github.com/ntop/nDPI/archive/refs/tags/4.4.bin"
	ndpi_cmd_bin := exec.Command("curl", "-L", ndpi_bin_url, "-o", "binary.bin")
	err = ndpi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ndpi_src_url := "https://github.com/ntop/nDPI/archive/refs/tags/4.4.src.tar.gz"
	ndpi_cmd_src := exec.Command("curl", "-L", ndpi_src_url, "-o", "source.tar.gz")
	err = ndpi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ndpi_cmd_direct := exec.Command("./binary")
	err = ndpi_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: json-c")
exec.Command("latte", "install", "json-c")
}
