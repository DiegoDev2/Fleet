package main

// Atomicparsley - MPEG-4 command-line tool
// Homepage: https://github.com/wez/atomicparsley

import (
	"fmt"
	
	"os/exec"
)

func installAtomicparsley() {
	// Método 1: Descargar y extraer .tar.gz
	atomicparsley_tar_url := "https://github.com/wez/atomicparsley/archive/refs/tags/20240608.083822.1ed9031.tar.gz"
	atomicparsley_cmd_tar := exec.Command("curl", "-L", atomicparsley_tar_url, "-o", "package.tar.gz")
	err := atomicparsley_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	atomicparsley_zip_url := "https://github.com/wez/atomicparsley/archive/refs/tags/20240608.083822.1ed9031.zip"
	atomicparsley_cmd_zip := exec.Command("curl", "-L", atomicparsley_zip_url, "-o", "package.zip")
	err = atomicparsley_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	atomicparsley_bin_url := "https://github.com/wez/atomicparsley/archive/refs/tags/20240608.083822.1ed9031.bin"
	atomicparsley_cmd_bin := exec.Command("curl", "-L", atomicparsley_bin_url, "-o", "binary.bin")
	err = atomicparsley_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	atomicparsley_src_url := "https://github.com/wez/atomicparsley/archive/refs/tags/20240608.083822.1ed9031.src.tar.gz"
	atomicparsley_cmd_src := exec.Command("curl", "-L", atomicparsley_src_url, "-o", "source.tar.gz")
	err = atomicparsley_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	atomicparsley_cmd_direct := exec.Command("./binary")
	err = atomicparsley_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
