package main

// Unbound - Validating, recursive, caching DNS resolver
// Homepage: https://www.unbound.net

import (
	"fmt"
	
	"os/exec"
)

func installUnbound() {
	// Método 1: Descargar y extraer .tar.gz
	unbound_tar_url := "https://nlnetlabs.nl/downloads/unbound/unbound-1.21.0.tar.gz"
	unbound_cmd_tar := exec.Command("curl", "-L", unbound_tar_url, "-o", "package.tar.gz")
	err := unbound_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unbound_zip_url := "https://nlnetlabs.nl/downloads/unbound/unbound-1.21.0.zip"
	unbound_cmd_zip := exec.Command("curl", "-L", unbound_zip_url, "-o", "package.zip")
	err = unbound_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unbound_bin_url := "https://nlnetlabs.nl/downloads/unbound/unbound-1.21.0.bin"
	unbound_cmd_bin := exec.Command("curl", "-L", unbound_bin_url, "-o", "binary.bin")
	err = unbound_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unbound_src_url := "https://nlnetlabs.nl/downloads/unbound/unbound-1.21.0.src.tar.gz"
	unbound_cmd_src := exec.Command("curl", "-L", unbound_src_url, "-o", "source.tar.gz")
	err = unbound_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unbound_cmd_direct := exec.Command("./binary")
	err = unbound_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: libnghttp2")
exec.Command("latte", "install", "libnghttp2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
