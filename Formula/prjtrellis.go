package main

// Prjtrellis - Documenting the Lattice ECP5 bit-stream format
// Homepage: https://github.com/YosysHQ/prjtrellis

import (
	"fmt"
	
	"os/exec"
)

func installPrjtrellis() {
	// Método 1: Descargar y extraer .tar.gz
	prjtrellis_tar_url := "https://github.com/YosysHQ/prjtrellis/archive/refs/tags/1.4.tar.gz"
	prjtrellis_cmd_tar := exec.Command("curl", "-L", prjtrellis_tar_url, "-o", "package.tar.gz")
	err := prjtrellis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	prjtrellis_zip_url := "https://github.com/YosysHQ/prjtrellis/archive/refs/tags/1.4.zip"
	prjtrellis_cmd_zip := exec.Command("curl", "-L", prjtrellis_zip_url, "-o", "package.zip")
	err = prjtrellis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	prjtrellis_bin_url := "https://github.com/YosysHQ/prjtrellis/archive/refs/tags/1.4.bin"
	prjtrellis_cmd_bin := exec.Command("curl", "-L", prjtrellis_bin_url, "-o", "binary.bin")
	err = prjtrellis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	prjtrellis_src_url := "https://github.com/YosysHQ/prjtrellis/archive/refs/tags/1.4.src.tar.gz"
	prjtrellis_cmd_src := exec.Command("curl", "-L", prjtrellis_src_url, "-o", "source.tar.gz")
	err = prjtrellis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	prjtrellis_cmd_direct := exec.Command("./binary")
	err = prjtrellis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: boost-python3")
	exec.Command("latte", "install", "boost-python3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
