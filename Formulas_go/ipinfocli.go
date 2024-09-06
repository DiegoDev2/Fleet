package main

// IpinfoCli - Official CLI for the IPinfo IP Address API
// Homepage: https://ipinfo.io/

import (
	"fmt"
	
	"os/exec"
)

func installIpinfoCli() {
	// Método 1: Descargar y extraer .tar.gz
	ipinfocli_tar_url := "https://github.com/ipinfo/cli/archive/refs/tags/ipinfo-3.3.1.tar.gz"
	ipinfocli_cmd_tar := exec.Command("curl", "-L", ipinfocli_tar_url, "-o", "package.tar.gz")
	err := ipinfocli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ipinfocli_zip_url := "https://github.com/ipinfo/cli/archive/refs/tags/ipinfo-3.3.1.zip"
	ipinfocli_cmd_zip := exec.Command("curl", "-L", ipinfocli_zip_url, "-o", "package.zip")
	err = ipinfocli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ipinfocli_bin_url := "https://github.com/ipinfo/cli/archive/refs/tags/ipinfo-3.3.1.bin"
	ipinfocli_cmd_bin := exec.Command("curl", "-L", ipinfocli_bin_url, "-o", "binary.bin")
	err = ipinfocli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ipinfocli_src_url := "https://github.com/ipinfo/cli/archive/refs/tags/ipinfo-3.3.1.src.tar.gz"
	ipinfocli_cmd_src := exec.Command("curl", "-L", ipinfocli_src_url, "-o", "source.tar.gz")
	err = ipinfocli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ipinfocli_cmd_direct := exec.Command("./binary")
	err = ipinfocli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
