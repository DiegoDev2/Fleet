package main

// Pbzx - Parser for pbzx stream
// Homepage: https://github.com/NiklasRosenstein/pbzx/

import (
	"fmt"
	
	"os/exec"
)

func installPbzx() {
	// Método 1: Descargar y extraer .tar.gz
	pbzx_tar_url := "https://github.com/NiklasRosenstein/pbzx/archive/refs/tags/v1.0.2.tar.gz"
	pbzx_cmd_tar := exec.Command("curl", "-L", pbzx_tar_url, "-o", "package.tar.gz")
	err := pbzx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pbzx_zip_url := "https://github.com/NiklasRosenstein/pbzx/archive/refs/tags/v1.0.2.zip"
	pbzx_cmd_zip := exec.Command("curl", "-L", pbzx_zip_url, "-o", "package.zip")
	err = pbzx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pbzx_bin_url := "https://github.com/NiklasRosenstein/pbzx/archive/refs/tags/v1.0.2.bin"
	pbzx_cmd_bin := exec.Command("curl", "-L", pbzx_bin_url, "-o", "binary.bin")
	err = pbzx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pbzx_src_url := "https://github.com/NiklasRosenstein/pbzx/archive/refs/tags/v1.0.2.src.tar.gz"
	pbzx_cmd_src := exec.Command("curl", "-L", pbzx_src_url, "-o", "source.tar.gz")
	err = pbzx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pbzx_cmd_direct := exec.Command("./binary")
	err = pbzx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
}
