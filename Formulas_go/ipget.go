package main

// Ipget - Retrieve files over IPFS and save them locally
// Homepage: https://github.com/ipfs/ipget/

import (
	"fmt"
	
	"os/exec"
)

func installIpget() {
	// Método 1: Descargar y extraer .tar.gz
	ipget_tar_url := "https://github.com/ipfs/ipget/archive/refs/tags/v0.10.0.tar.gz"
	ipget_cmd_tar := exec.Command("curl", "-L", ipget_tar_url, "-o", "package.tar.gz")
	err := ipget_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ipget_zip_url := "https://github.com/ipfs/ipget/archive/refs/tags/v0.10.0.zip"
	ipget_cmd_zip := exec.Command("curl", "-L", ipget_zip_url, "-o", "package.zip")
	err = ipget_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ipget_bin_url := "https://github.com/ipfs/ipget/archive/refs/tags/v0.10.0.bin"
	ipget_cmd_bin := exec.Command("curl", "-L", ipget_bin_url, "-o", "binary.bin")
	err = ipget_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ipget_src_url := "https://github.com/ipfs/ipget/archive/refs/tags/v0.10.0.src.tar.gz"
	ipget_cmd_src := exec.Command("curl", "-L", ipget_src_url, "-o", "source.tar.gz")
	err = ipget_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ipget_cmd_direct := exec.Command("./binary")
	err = ipget_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
