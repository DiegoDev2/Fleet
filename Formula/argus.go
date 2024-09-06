package main

// Argus - Audit Record Generation and Utilization System server
// Homepage: https://openargus.org

import (
	"fmt"
	
	"os/exec"
)

func installArgus() {
	// Método 1: Descargar y extraer .tar.gz
	argus_tar_url := "https://github.com/openargus/argus/archive/refs/tags/v5.0.0.tar.gz"
	argus_cmd_tar := exec.Command("curl", "-L", argus_tar_url, "-o", "package.tar.gz")
	err := argus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	argus_zip_url := "https://github.com/openargus/argus/archive/refs/tags/v5.0.0.zip"
	argus_cmd_zip := exec.Command("curl", "-L", argus_zip_url, "-o", "package.zip")
	err = argus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	argus_bin_url := "https://github.com/openargus/argus/archive/refs/tags/v5.0.0.bin"
	argus_cmd_bin := exec.Command("curl", "-L", argus_bin_url, "-o", "binary.bin")
	err = argus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	argus_src_url := "https://github.com/openargus/argus/archive/refs/tags/v5.0.0.src.tar.gz"
	argus_cmd_src := exec.Command("curl", "-L", argus_src_url, "-o", "source.tar.gz")
	err = argus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	argus_cmd_direct := exec.Command("./binary")
	err = argus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtirpc")
	exec.Command("latte", "install", "libtirpc").Run()
}
