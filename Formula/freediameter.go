package main

// Freediameter - Open source Diameter (Authentication) protocol implementation
// Homepage: http://www.freediameter.net

import (
	"fmt"
	
	"os/exec"
)

func installFreediameter() {
	// Método 1: Descargar y extraer .tar.gz
	freediameter_tar_url := "http://www.freediameter.net/hg/freeDiameter/archive/1.5.0.tar.gz"
	freediameter_cmd_tar := exec.Command("curl", "-L", freediameter_tar_url, "-o", "package.tar.gz")
	err := freediameter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	freediameter_zip_url := "http://www.freediameter.net/hg/freeDiameter/archive/1.5.0.zip"
	freediameter_cmd_zip := exec.Command("curl", "-L", freediameter_zip_url, "-o", "package.zip")
	err = freediameter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	freediameter_bin_url := "http://www.freediameter.net/hg/freeDiameter/archive/1.5.0.bin"
	freediameter_cmd_bin := exec.Command("curl", "-L", freediameter_bin_url, "-o", "binary.bin")
	err = freediameter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	freediameter_src_url := "http://www.freediameter.net/hg/freeDiameter/archive/1.5.0.src.tar.gz"
	freediameter_cmd_src := exec.Command("curl", "-L", freediameter_src_url, "-o", "source.tar.gz")
	err = freediameter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	freediameter_cmd_direct := exec.Command("./binary")
	err = freediameter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: libidn2")
	exec.Command("latte", "install", "libidn2").Run()
}
