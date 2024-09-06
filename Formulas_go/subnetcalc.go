package main

// Subnetcalc - IPv4/IPv6 subnet calculator
// Homepage: https://www.nntb.no/~dreibh/subnetcalc/index.html

import (
	"fmt"
	
	"os/exec"
)

func installSubnetcalc() {
	// Método 1: Descargar y extraer .tar.gz
	subnetcalc_tar_url := "https://github.com/dreibh/subnetcalc/archive/refs/tags/subnetcalc-2.5.1.tar.gz"
	subnetcalc_cmd_tar := exec.Command("curl", "-L", subnetcalc_tar_url, "-o", "package.tar.gz")
	err := subnetcalc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	subnetcalc_zip_url := "https://github.com/dreibh/subnetcalc/archive/refs/tags/subnetcalc-2.5.1.zip"
	subnetcalc_cmd_zip := exec.Command("curl", "-L", subnetcalc_zip_url, "-o", "package.zip")
	err = subnetcalc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	subnetcalc_bin_url := "https://github.com/dreibh/subnetcalc/archive/refs/tags/subnetcalc-2.5.1.bin"
	subnetcalc_cmd_bin := exec.Command("curl", "-L", subnetcalc_bin_url, "-o", "binary.bin")
	err = subnetcalc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	subnetcalc_src_url := "https://github.com/dreibh/subnetcalc/archive/refs/tags/subnetcalc-2.5.1.src.tar.gz"
	subnetcalc_cmd_src := exec.Command("curl", "-L", subnetcalc_src_url, "-o", "source.tar.gz")
	err = subnetcalc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	subnetcalc_cmd_direct := exec.Command("./binary")
	err = subnetcalc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
