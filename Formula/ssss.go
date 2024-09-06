package main

// Ssss - Shamir's secret sharing scheme implementation
// Homepage: http://point-at-infinity.org/ssss/

import (
	"fmt"
	
	"os/exec"
)

func installSsss() {
	// Método 1: Descargar y extraer .tar.gz
	ssss_tar_url := "http://point-at-infinity.org/ssss/ssss-0.5.tar.gz"
	ssss_cmd_tar := exec.Command("curl", "-L", ssss_tar_url, "-o", "package.tar.gz")
	err := ssss_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ssss_zip_url := "http://point-at-infinity.org/ssss/ssss-0.5.zip"
	ssss_cmd_zip := exec.Command("curl", "-L", ssss_zip_url, "-o", "package.zip")
	err = ssss_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ssss_bin_url := "http://point-at-infinity.org/ssss/ssss-0.5.bin"
	ssss_cmd_bin := exec.Command("curl", "-L", ssss_bin_url, "-o", "binary.bin")
	err = ssss_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ssss_src_url := "http://point-at-infinity.org/ssss/ssss-0.5.src.tar.gz"
	ssss_cmd_src := exec.Command("curl", "-L", ssss_src_url, "-o", "source.tar.gz")
	err = ssss_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ssss_cmd_direct := exec.Command("./binary")
	err = ssss_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: xmltoman")
	exec.Command("latte", "install", "xmltoman").Run()
}
