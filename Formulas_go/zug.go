package main

// Zug - C++ library providing transducers
// Homepage: https://sinusoid.es/zug/

import (
	"fmt"
	
	"os/exec"
)

func installZug() {
	// Método 1: Descargar y extraer .tar.gz
	zug_tar_url := "https://github.com/arximboldi/zug/archive/refs/tags/v0.1.1.tar.gz"
	zug_cmd_tar := exec.Command("curl", "-L", zug_tar_url, "-o", "package.tar.gz")
	err := zug_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zug_zip_url := "https://github.com/arximboldi/zug/archive/refs/tags/v0.1.1.zip"
	zug_cmd_zip := exec.Command("curl", "-L", zug_zip_url, "-o", "package.zip")
	err = zug_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zug_bin_url := "https://github.com/arximboldi/zug/archive/refs/tags/v0.1.1.bin"
	zug_cmd_bin := exec.Command("curl", "-L", zug_bin_url, "-o", "binary.bin")
	err = zug_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zug_src_url := "https://github.com/arximboldi/zug/archive/refs/tags/v0.1.1.src.tar.gz"
	zug_cmd_src := exec.Command("curl", "-L", zug_src_url, "-o", "source.tar.gz")
	err = zug_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zug_cmd_direct := exec.Command("./binary")
	err = zug_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
