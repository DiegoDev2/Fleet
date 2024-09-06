package main

// Fabio - Zero-conf load balancing HTTP(S) router
// Homepage: https://github.com/fabiolb/fabio

import (
	"fmt"
	
	"os/exec"
)

func installFabio() {
	// Método 1: Descargar y extraer .tar.gz
	fabio_tar_url := "https://github.com/fabiolb/fabio/archive/refs/tags/v1.6.3.tar.gz"
	fabio_cmd_tar := exec.Command("curl", "-L", fabio_tar_url, "-o", "package.tar.gz")
	err := fabio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fabio_zip_url := "https://github.com/fabiolb/fabio/archive/refs/tags/v1.6.3.zip"
	fabio_cmd_zip := exec.Command("curl", "-L", fabio_zip_url, "-o", "package.zip")
	err = fabio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fabio_bin_url := "https://github.com/fabiolb/fabio/archive/refs/tags/v1.6.3.bin"
	fabio_cmd_bin := exec.Command("curl", "-L", fabio_bin_url, "-o", "binary.bin")
	err = fabio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fabio_src_url := "https://github.com/fabiolb/fabio/archive/refs/tags/v1.6.3.src.tar.gz"
	fabio_cmd_src := exec.Command("curl", "-L", fabio_src_url, "-o", "source.tar.gz")
	err = fabio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fabio_cmd_direct := exec.Command("./binary")
	err = fabio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: consul")
exec.Command("latte", "install", "consul")
}
