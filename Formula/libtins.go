package main

// Libtins - C++ network packet sniffing and crafting library
// Homepage: https://libtins.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installLibtins() {
	// Método 1: Descargar y extraer .tar.gz
	libtins_tar_url := "https://github.com/mfontanini/libtins/archive/refs/tags/v4.5.tar.gz"
	libtins_cmd_tar := exec.Command("curl", "-L", libtins_tar_url, "-o", "package.tar.gz")
	err := libtins_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtins_zip_url := "https://github.com/mfontanini/libtins/archive/refs/tags/v4.5.zip"
	libtins_cmd_zip := exec.Command("curl", "-L", libtins_zip_url, "-o", "package.zip")
	err = libtins_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtins_bin_url := "https://github.com/mfontanini/libtins/archive/refs/tags/v4.5.bin"
	libtins_cmd_bin := exec.Command("curl", "-L", libtins_bin_url, "-o", "binary.bin")
	err = libtins_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtins_src_url := "https://github.com/mfontanini/libtins/archive/refs/tags/v4.5.src.tar.gz"
	libtins_cmd_src := exec.Command("curl", "-L", libtins_src_url, "-o", "source.tar.gz")
	err = libtins_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtins_cmd_direct := exec.Command("./binary")
	err = libtins_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
