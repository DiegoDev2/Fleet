package main

// Macosvpn - Create Mac OS VPNs programmatically
// Homepage: https://github.com/halo/macosvpn

import (
	"fmt"
	
	"os/exec"
)

func installMacosvpn() {
	// Método 1: Descargar y extraer .tar.gz
	macosvpn_tar_url := "https://github.com/halo/macosvpn/archive/refs/tags/2.0.0.tar.gz"
	macosvpn_cmd_tar := exec.Command("curl", "-L", macosvpn_tar_url, "-o", "package.tar.gz")
	err := macosvpn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	macosvpn_zip_url := "https://github.com/halo/macosvpn/archive/refs/tags/2.0.0.zip"
	macosvpn_cmd_zip := exec.Command("curl", "-L", macosvpn_zip_url, "-o", "package.zip")
	err = macosvpn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	macosvpn_bin_url := "https://github.com/halo/macosvpn/archive/refs/tags/2.0.0.bin"
	macosvpn_cmd_bin := exec.Command("curl", "-L", macosvpn_bin_url, "-o", "binary.bin")
	err = macosvpn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	macosvpn_src_url := "https://github.com/halo/macosvpn/archive/refs/tags/2.0.0.src.tar.gz"
	macosvpn_cmd_src := exec.Command("curl", "-L", macosvpn_src_url, "-o", "source.tar.gz")
	err = macosvpn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	macosvpn_cmd_direct := exec.Command("./binary")
	err = macosvpn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
