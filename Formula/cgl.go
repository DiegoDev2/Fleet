package main

// Cgl - Cut Generation Library
// Homepage: https://github.com/coin-or/Cgl

import (
	"fmt"
	
	"os/exec"
)

func installCgl() {
	// Método 1: Descargar y extraer .tar.gz
	cgl_tar_url := "https://github.com/coin-or/Cgl/archive/refs/tags/releases/0.60.9.tar.gz"
	cgl_cmd_tar := exec.Command("curl", "-L", cgl_tar_url, "-o", "package.tar.gz")
	err := cgl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cgl_zip_url := "https://github.com/coin-or/Cgl/archive/refs/tags/releases/0.60.9.zip"
	cgl_cmd_zip := exec.Command("curl", "-L", cgl_zip_url, "-o", "package.zip")
	err = cgl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cgl_bin_url := "https://github.com/coin-or/Cgl/archive/refs/tags/releases/0.60.9.bin"
	cgl_cmd_bin := exec.Command("curl", "-L", cgl_bin_url, "-o", "binary.bin")
	err = cgl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cgl_src_url := "https://github.com/coin-or/Cgl/archive/refs/tags/releases/0.60.9.src.tar.gz"
	cgl_cmd_src := exec.Command("curl", "-L", cgl_src_url, "-o", "source.tar.gz")
	err = cgl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cgl_cmd_direct := exec.Command("./binary")
	err = cgl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: clp")
	exec.Command("latte", "install", "clp").Run()
	fmt.Println("Instalando dependencia: coinutils")
	exec.Command("latte", "install", "coinutils").Run()
	fmt.Println("Instalando dependencia: osi")
	exec.Command("latte", "install", "osi").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
