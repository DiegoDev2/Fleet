package main

// Bgpq3 - BGP filtering automation for Cisco, Juniper, BIRD and OpenBGPD routers
// Homepage: http://snar.spb.ru/prog/bgpq3/

import (
	"fmt"
	
	"os/exec"
)

func installBgpq3() {
	// Método 1: Descargar y extraer .tar.gz
	bgpq3_tar_url := "https://github.com/snar/bgpq3/archive/refs/tags/v0.1.36.1.tar.gz"
	bgpq3_cmd_tar := exec.Command("curl", "-L", bgpq3_tar_url, "-o", "package.tar.gz")
	err := bgpq3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bgpq3_zip_url := "https://github.com/snar/bgpq3/archive/refs/tags/v0.1.36.1.zip"
	bgpq3_cmd_zip := exec.Command("curl", "-L", bgpq3_zip_url, "-o", "package.zip")
	err = bgpq3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bgpq3_bin_url := "https://github.com/snar/bgpq3/archive/refs/tags/v0.1.36.1.bin"
	bgpq3_cmd_bin := exec.Command("curl", "-L", bgpq3_bin_url, "-o", "binary.bin")
	err = bgpq3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bgpq3_src_url := "https://github.com/snar/bgpq3/archive/refs/tags/v0.1.36.1.src.tar.gz"
	bgpq3_cmd_src := exec.Command("curl", "-L", bgpq3_src_url, "-o", "source.tar.gz")
	err = bgpq3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bgpq3_cmd_direct := exec.Command("./binary")
	err = bgpq3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
