package main

// Libaec - Adaptive Entropy Coding implementing Golomb-Rice algorithm
// Homepage: https://gitlab.dkrz.de/k202009/libaec

import (
	"fmt"
	
	"os/exec"
)

func installLibaec() {
	// Método 1: Descargar y extraer .tar.gz
	libaec_tar_url := "https://gitlab.dkrz.de/k202009/libaec/-/archive/v1.1.3/libaec-v1.1.3.tar.bz2"
	libaec_cmd_tar := exec.Command("curl", "-L", libaec_tar_url, "-o", "package.tar.gz")
	err := libaec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libaec_zip_url := "https://gitlab.dkrz.de/k202009/libaec/-/archive/v1.1.3/libaec-v1.1.3.tar.bz2"
	libaec_cmd_zip := exec.Command("curl", "-L", libaec_zip_url, "-o", "package.zip")
	err = libaec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libaec_bin_url := "https://gitlab.dkrz.de/k202009/libaec/-/archive/v1.1.3/libaec-v1.1.3.tar.bz2"
	libaec_cmd_bin := exec.Command("curl", "-L", libaec_bin_url, "-o", "binary.bin")
	err = libaec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libaec_src_url := "https://gitlab.dkrz.de/k202009/libaec/-/archive/v1.1.3/libaec-v1.1.3.tar.bz2"
	libaec_cmd_src := exec.Command("curl", "-L", libaec_src_url, "-o", "source.tar.gz")
	err = libaec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libaec_cmd_direct := exec.Command("./binary")
	err = libaec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
